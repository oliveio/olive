//
//Copyright 2023 The olive Authors
//
//This program is offered under a commercial and under the AGPL license.
//For AGPL licensing, see below.
//
//AGPL licensing:
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU Affero General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU Affero General Public License for more details.
//
//You should have received a copy of the GNU Affero General Public License
//along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: github.com/olive-io/olive/api/discoverypb/discovery.proto

package discoverypb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BoxType int32

const (
	BoxType_string  BoxType = 0
	BoxType_integer BoxType = 1
	BoxType_float   BoxType = 2
	BoxType_boolean BoxType = 3
	BoxType_array   BoxType = 4
	BoxType_object  BoxType = 5
	BoxType_map     BoxType = 6
)

// Enum value maps for BoxType.
var (
	BoxType_name = map[int32]string{
		0: "string",
		1: "integer",
		2: "float",
		3: "boolean",
		4: "array",
		5: "object",
		6: "map",
	}
	BoxType_value = map[string]int32{
		"string":  0,
		"integer": 1,
		"float":   2,
		"boolean": 3,
		"array":   4,
		"object":  5,
		"map":     6,
	}
)

func (x BoxType) Enum() *BoxType {
	p := new(BoxType)
	*p = x
	return p
}

func (x BoxType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoxType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_enumTypes[0].Descriptor()
}

func (BoxType) Type() protoreflect.EnumType {
	return &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_enumTypes[0]
}

func (x BoxType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoxType.Descriptor instead.
func (BoxType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{0}
}

// Box is an opaque value for a request or response
type Box struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BoxType `protobuf:"varint,1,opt,name=type,proto3,enum=discoverypb.BoxType" json:"type,omitempty"`
	// Box Value by json.Marshal
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// the reference, points to OpenAPI Component when type is Object
	Ref        string          `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Parameters map[string]*Box `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Box) Reset() {
	*x = Box{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Box) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Box) ProtoMessage() {}

func (x *Box) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Box.ProtoReflect.Descriptor instead.
func (*Box) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *Box) GetType() BoxType {
	if x != nil {
		return x.Type
	}
	return BoxType_string
}

func (x *Box) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Box) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Box) GetParameters() map[string]*Box {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// Endpoint is a endpoint provided by a service
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Request  *Box              `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response *Box              `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *Endpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Endpoint) GetRequest() *Box {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Endpoint) GetResponse() *Box {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Endpoint) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Event is registry event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of event
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service entry
	Service *Service `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

// Node represents the node the service is on
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port     int64             `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Node) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Result is returns by the watcher
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    string    `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Service   *Service  `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint  *Endpoint `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Timestamp int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Result) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Result) GetEndpoint() *Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *Result) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Service represents a olive service
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Namespace string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Endpoints []*Endpoint       `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Nodes     []*Node           `protobuf:"bytes,6,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Ttl       int64             `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{5}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Service) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Service) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Service) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Service) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Hook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Hook) Reset() {
	*x = Hook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hook) ProtoMessage() {}

func (x *Hook) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hook.ProtoReflect.Descriptor instead.
func (*Hook) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{6}
}

func (x *Hook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GlobalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GlobalConfig) Reset() {
	*x = GlobalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalConfig) ProtoMessage() {}

func (x *GlobalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalConfig.ProtoReflect.Descriptor instead.
func (*GlobalConfig) Descriptor() ([]byte, []int) {
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP(), []int{7}
}

var File_github_com_olive_io_olive_api_discoverypb_discovery_proto protoreflect.FileDescriptor

var file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x69,
	0x76, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x6f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x22, 0xe8, 0x01, 0x0a, 0x03, 0x42, 0x6f, 0x78,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x78,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66,
	0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x70, 0x62, 0x2e, 0x42, 0x6f, 0x78, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x4f, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x78, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xf6, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x78, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62,
	0x2e, 0x42, 0x6f, 0x78, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x79, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3b,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc2, 0x02, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x74, 0x74, 0x6c, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2c, 0x0a, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a,
	0x5a, 0x0a, 0x07, 0x42, 0x6f, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x10, 0x06, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x69, 0x76, 0x65, 0x2d,
	0x69, 0x6f, 0x2f, 0x6f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescOnce sync.Once
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescData = file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDesc
)

func file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescGZIP() []byte {
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescOnce.Do(func() {
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescData)
	})
	return file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDescData
}

var file_github_com_olive_io_olive_api_discoverypb_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_github_com_olive_io_olive_api_discoverypb_discovery_proto_goTypes = []interface{}{
	(BoxType)(0),         // 0: discoverypb.BoxType
	(*Box)(nil),          // 1: discoverypb.Box
	(*Endpoint)(nil),     // 2: discoverypb.Endpoint
	(*Event)(nil),        // 3: discoverypb.Event
	(*Node)(nil),         // 4: discoverypb.Node
	(*Result)(nil),       // 5: discoverypb.Result
	(*Service)(nil),      // 6: discoverypb.Service
	(*Hook)(nil),         // 7: discoverypb.Hook
	(*GlobalConfig)(nil), // 8: discoverypb.GlobalConfig
	nil,                  // 9: discoverypb.Box.ParametersEntry
	nil,                  // 10: discoverypb.Endpoint.MetadataEntry
	nil,                  // 11: discoverypb.Node.MetadataEntry
	nil,                  // 12: discoverypb.Service.MetadataEntry
}
var file_github_com_olive_io_olive_api_discoverypb_discovery_proto_depIdxs = []int32{
	0,  // 0: discoverypb.Box.type:type_name -> discoverypb.BoxType
	9,  // 1: discoverypb.Box.parameters:type_name -> discoverypb.Box.ParametersEntry
	1,  // 2: discoverypb.Endpoint.request:type_name -> discoverypb.Box
	1,  // 3: discoverypb.Endpoint.response:type_name -> discoverypb.Box
	10, // 4: discoverypb.Endpoint.metadata:type_name -> discoverypb.Endpoint.MetadataEntry
	6,  // 5: discoverypb.Event.service:type_name -> discoverypb.Service
	11, // 6: discoverypb.Node.metadata:type_name -> discoverypb.Node.MetadataEntry
	6,  // 7: discoverypb.Result.service:type_name -> discoverypb.Service
	2,  // 8: discoverypb.Result.endpoint:type_name -> discoverypb.Endpoint
	12, // 9: discoverypb.Service.metadata:type_name -> discoverypb.Service.MetadataEntry
	2,  // 10: discoverypb.Service.endpoints:type_name -> discoverypb.Endpoint
	4,  // 11: discoverypb.Service.nodes:type_name -> discoverypb.Node
	1,  // 12: discoverypb.Box.ParametersEntry.value:type_name -> discoverypb.Box
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_github_com_olive_io_olive_api_discoverypb_discovery_proto_init() }
func file_github_com_olive_io_olive_api_discoverypb_discovery_proto_init() {
	if File_github_com_olive_io_olive_api_discoverypb_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Box); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hook); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_olive_io_olive_api_discoverypb_discovery_proto_goTypes,
		DependencyIndexes: file_github_com_olive_io_olive_api_discoverypb_discovery_proto_depIdxs,
		EnumInfos:         file_github_com_olive_io_olive_api_discoverypb_discovery_proto_enumTypes,
		MessageInfos:      file_github_com_olive_io_olive_api_discoverypb_discovery_proto_msgTypes,
	}.Build()
	File_github_com_olive_io_olive_api_discoverypb_discovery_proto = out.File
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_rawDesc = nil
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_goTypes = nil
	file_github_com_olive_io_olive_api_discoverypb_discovery_proto_depIdxs = nil
}
