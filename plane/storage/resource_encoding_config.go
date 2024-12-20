/*
Copyright 2024 The olive Authors

This program is offered under a commercial and under the AGPL license.
For AGPL licensing, see below.

AGPL licensing:
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package storage

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apimachineryversion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/component-base/version"
)

type ResourceEncodingConfig interface {
	// StorageEncoding returns the serialization format for the resource.
	// TODO this should actually return a GroupVersionKind since you can logically have multiple "matching" Kinds
	// For now, it returns just the GroupVersion for consistency with old behavior
	StorageEncodingFor(schema.GroupResource) (schema.GroupVersion, error)

	// InMemoryEncodingFor returns the groupVersion for the in memory representation the storage should convert to.
	InMemoryEncodingFor(schema.GroupResource) (schema.GroupVersion, error)
}

type CompatibilityResourceEncodingConfig interface {
	BackwardCompatibileStorageEncodingFor(schema.GroupResource, runtime.Object) (schema.GroupVersion, error)
}

type DefaultResourceEncodingConfig struct {
	// resources records the overriding encoding configs for individual resources.
	resources        map[schema.GroupResource]*OverridingResourceEncoding
	scheme           *runtime.Scheme
	effectiveVersion version.EffectiveVersion
}

type OverridingResourceEncoding struct {
	ExternalResourceEncoding schema.GroupVersion
	InternalResourceEncoding schema.GroupVersion
}

var _ ResourceEncodingConfig = &DefaultResourceEncodingConfig{}

func NewDefaultResourceEncodingConfig(scheme *runtime.Scheme) *DefaultResourceEncodingConfig {
	return &DefaultResourceEncodingConfig{resources: map[schema.GroupResource]*OverridingResourceEncoding{}, scheme: scheme, effectiveVersion: version.DefaultKubeEffectiveVersion()}
}

func (o *DefaultResourceEncodingConfig) SetResourceEncoding(resourceBeingStored schema.GroupResource, externalEncodingVersion, internalVersion schema.GroupVersion) {
	o.resources[resourceBeingStored] = &OverridingResourceEncoding{
		ExternalResourceEncoding: externalEncodingVersion,
		InternalResourceEncoding: internalVersion,
	}
}

func (o *DefaultResourceEncodingConfig) SetEffectiveVersion(effectiveVersion version.EffectiveVersion) {
	o.effectiveVersion = effectiveVersion
}

func (o *DefaultResourceEncodingConfig) StorageEncodingFor(resource schema.GroupResource) (schema.GroupVersion, error) {
	if !o.scheme.IsGroupRegistered(resource.Group) {
		return schema.GroupVersion{}, fmt.Errorf("group %q is not registered in scheme", resource.Group)
	}

	resourceOverride, resourceExists := o.resources[resource]
	if resourceExists {
		return resourceOverride.ExternalResourceEncoding, nil
	}

	// return the most preferred external version for the group
	return o.scheme.PrioritizedVersionsForGroup(resource.Group)[0], nil
}

func (o *DefaultResourceEncodingConfig) BackwardCompatibileStorageEncodingFor(resource schema.GroupResource, example runtime.Object) (schema.GroupVersion, error) {
	if !o.scheme.IsGroupRegistered(resource.Group) {
		return schema.GroupVersion{}, fmt.Errorf("group %q is not registered in scheme", resource.Group)
	}

	// Always respect overrides
	resourceOverride, resourceExists := o.resources[resource]
	if resourceExists {
		return resourceOverride.ExternalResourceEncoding, nil
	}

	return emulatedStorageVersion(
		o.scheme.PrioritizedVersionsForGroup(resource.Group)[0],
		example,
		o.effectiveVersion,
		o.scheme)
}

func (o *DefaultResourceEncodingConfig) InMemoryEncodingFor(resource schema.GroupResource) (schema.GroupVersion, error) {
	if !o.scheme.IsGroupRegistered(resource.Group) {
		return schema.GroupVersion{}, fmt.Errorf("group %q is not registered in scheme", resource.Group)
	}

	resourceOverride, resourceExists := o.resources[resource]
	if resourceExists {
		return resourceOverride.InternalResourceEncoding, nil
	}
	return schema.GroupVersion{Group: resource.Group, Version: runtime.APIVersionInternal}, nil
}

// Object interface generated from "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
type introducedInterface interface {
	APILifecycleIntroduced() (major, minor int)
}

type replacementInterface interface {
	APILifecycleReplacement() schema.GroupVersionKind
}

func emulatedStorageVersion(binaryVersionOfResource schema.GroupVersion, example runtime.Object, effectiveVersion version.EffectiveVersion, scheme *runtime.Scheme) (schema.GroupVersion, error) {
	if example == nil || effectiveVersion == nil {
		return binaryVersionOfResource, nil
	}

	// Look up example in scheme to find all objects of the same Group-Kind
	// Use the highest priority version for that group-kind whose lifecycle window
	// includes the current emulation version.
	// If no version is found, use the binary version
	// (in this case the API should be disabled anyway)
	gvks, _, err := scheme.ObjectKinds(example)
	if err != nil {
		return schema.GroupVersion{}, err
	}

	var gvk schema.GroupVersionKind
	for _, item := range gvks {
		if item.Group != binaryVersionOfResource.Group {
			continue
		}

		gvk = item
		break
	}

	if len(gvk.Kind) == 0 {
		return schema.GroupVersion{}, fmt.Errorf("object %T has no GVKs registered in scheme", example)
	}

	// VersionsForGroupKind returns versions in priority order
	versions := scheme.VersionsForGroupKind(schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind})

	compatibilityVersion := effectiveVersion.MinCompatibilityVersion()

	for _, gv := range versions {
		if gv.Version == runtime.APIVersionInternal {
			continue
		}

		gvk := schema.GroupVersionKind{
			Group:   gv.Group,
			Version: gv.Version,
			Kind:    gvk.Kind,
		}

		exampleOfGVK, err := scheme.New(gvk)
		if err != nil {
			return schema.GroupVersion{}, err
		}

		// If it was introduced after current compatibility version, don't use it
		// skip the introduced check for test when currentVersion is 0.0 to test all apis
		if introduced, hasIntroduced := exampleOfGVK.(introducedInterface); hasIntroduced && (compatibilityVersion.Major() > 0 || compatibilityVersion.Minor() > 0) {

			// Skip versions that have a replacement.
			// This can be used to override this storage version selection by
			// marking a storage version has having a replacement and preventing a
			// that storage version from being selected.
			if _, hasReplacement := exampleOfGVK.(replacementInterface); hasReplacement {
				continue
			}
			// API resource lifecycles should be relative to k8s api version
			majorIntroduced, minorIntroduced := introduced.APILifecycleIntroduced()
			introducedVer := apimachineryversion.MajorMinor(uint(majorIntroduced), uint(minorIntroduced))
			if introducedVer.GreaterThan(compatibilityVersion) {
				continue
			}
		}

		// versions is returned in priority order, so just use first result
		return gvk.GroupVersion(), nil
	}

	// Getting here means we're serving a version that is unknown to the
	// min-compatibility-version server.
	//
	// This is only expected to happen when serving an alpha API type due
	// to missing pre-release lifecycle information
	// (which doesn't happen by default), or when emulation-version and
	// min-compatibility-version are several versions apart so a beta or GA API
	// was being served which didn't exist at all in min-compatibility-version.
	//
	// In the alpha case - we do not support compatibility versioning of
	// alpha types and recommend users do not mix the two.
	// In the skip-level case - The version of apiserver we are retaining
	// compatibility with has no knowledge of the type,
	// so storing it in another type is no issue.
	return binaryVersionOfResource, nil
}
