package buckets

import (
	"bytes"

	"github.com/olive-io/olive/server/mvcc/backend"
)

var (
	keyBucketName   = []byte("key")
	metaBucketName  = []byte("meta")
	leaseBucketName = []byte("lease")

	clusterBucketName = []byte("cluster")

	membersBucketName        = []byte("members")
	membersRemovedBucketName = []byte("members_removed")

	authBucketName      = []byte("auth")
	authUsersBucketName = []byte("authUsers")
	authRolesBucketName = []byte("authRoles")

	testBucketName = []byte("test")
)

var (
	Key   = backend.IBucket(bucket{id: 1, name: keyBucketName, safeRangeBucket: true})
	Meta  = backend.IBucket(bucket{id: 2, name: metaBucketName, safeRangeBucket: false})
	Lease = backend.IBucket(bucket{id: 3, name: leaseBucketName, safeRangeBucket: false})

	Cluster = backend.IBucket(bucket{id: 5, name: clusterBucketName, safeRangeBucket: false})

	Members        = backend.IBucket(bucket{id: 10, name: membersBucketName, safeRangeBucket: false})
	MembersRemoved = backend.IBucket(bucket{id: 11, name: membersRemovedBucketName, safeRangeBucket: false})

	Auth      = backend.IBucket(bucket{id: 20, name: authBucketName, safeRangeBucket: false})
	AuthUsers = backend.IBucket(bucket{id: 21, name: authUsersBucketName, safeRangeBucket: false})
	AuthRoles = backend.IBucket(bucket{id: 22, name: authRolesBucketName, safeRangeBucket: false})

	Test = backend.IBucket(bucket{id: 100, name: testBucketName, safeRangeBucket: false})
)

type bucket struct {
	id              backend.BucketID
	name            []byte
	safeRangeBucket bool
}

func (b bucket) ID() backend.BucketID    { return b.id }
func (b bucket) Name() []byte            { return b.name }
func (b bucket) String() string          { return string(b.Name()) }
func (b bucket) IsSafeRangeBucket() bool { return b.safeRangeBucket }

var (
	MetaConsistentIndexKeyName = []byte("consistent_index")
	MetaTermKeyName            = []byte("term")
)

// DefaultIgnores defines buckets & keys to ignore in hash checking.
func DefaultIgnores(bucket, key []byte) bool {
	// consistent index & term might be changed due to v2 internal sync, which
	// is not controllable by the user.
	return bytes.Compare(bucket, Meta.Name()) == 0 &&
		(bytes.Compare(key, MetaTermKeyName) == 0 || bytes.Compare(key, MetaConsistentIndexKeyName) == 0)
}
