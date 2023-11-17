// Code generated by proto-gen-gogo. DO NOT EDIT.
// source: github.com/olive-io/olive/api/raftpb/raft.proto

package raftpb

import (
	ebinary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = ebinary.BigEndian

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

type ConfChangeType int32

const (
	ConfChangeType_ConfChangeAddNode        ConfChangeType = 0
	ConfChangeType_ConfChangeRemoveNode     ConfChangeType = 1
	ConfChangeType_ConfChangeUpdateNode     ConfChangeType = 2
	ConfChangeType_ConfChangeAddLearnerNode ConfChangeType = 3
)

var ConfChangeType_name = map[int32]string{
	0: "ConfChangeAddNode",
	1: "ConfChangeRemoveNode",
	2: "ConfChangeUpdateNode",
	3: "ConfChangeAddLearnerNode",
}

var ConfChangeType_value = map[string]int32{
	"ConfChangeAddNode":        0,
	"ConfChangeRemoveNode":     1,
	"ConfChangeUpdateNode":     2,
	"ConfChangeAddLearnerNode": 3,
}

func (x ConfChangeType) Enum() *ConfChangeType {
	p := new(ConfChangeType)
	*p = x
	return p
}

func (x ConfChangeType) String() string {
	return proto.EnumName(ConfChangeType_name, int32(x))
}

func (x *ConfChangeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConfChangeType_value, data, "ConfChangeType")
	if err != nil {
		return err
	}
	*x = ConfChangeType(value)
	return nil
}

func (ConfChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c7ed0c999701337a, []int{0}
}

type ConfChange struct {
	Type    ConfChangeType `protobuf:"varint,2,opt,name=type,enum=raftpb.ConfChangeType" json:"type"`
	NodeID  uint64         `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id"`
	Context []byte         `protobuf:"bytes,4,opt,name=context" json:"context"`
	// NB: this is used only by olive to thread through a unique identifier.
	// Ideally it should really use the Context instead. No counterpart to
	// this field exists in ConfChangeV2.
	ID uint64 `protobuf:"varint,1,opt,name=id" json:"id"`
}

func (m *ConfChange) Reset()         { *m = ConfChange{} }
func (m *ConfChange) String() string { return proto.CompactTextString(m) }
func (*ConfChange) ProtoMessage()    {}
func (*ConfChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7ed0c999701337a, []int{0}
}
func (m *ConfChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfChange.Merge(m, src)
}
func (m *ConfChange) XXX_Size() int {
	return m.XSize()
}
func (m *ConfChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfChange.DiscardUnknown(m)
}

var xxx_messageInfo_ConfChange proto.InternalMessageInfo

// ConfChangeSingle is an individual configuration change operation. Multiple
// such operations can be carried out atomically via a ConfChangeV2.
type ConfChangeSingle struct {
	Type   ConfChangeType `protobuf:"varint,1,opt,name=type,enum=raftpb.ConfChangeType" json:"type"`
	NodeID uint64         `protobuf:"varint,2,opt,name=node_id,json=nodeId" json:"node_id"`
}

func (m *ConfChangeSingle) Reset()         { *m = ConfChangeSingle{} }
func (m *ConfChangeSingle) String() string { return proto.CompactTextString(m) }
func (*ConfChangeSingle) ProtoMessage()    {}
func (*ConfChangeSingle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7ed0c999701337a, []int{1}
}
func (m *ConfChangeSingle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfChangeSingle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfChangeSingle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfChangeSingle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfChangeSingle.Merge(m, src)
}
func (m *ConfChangeSingle) XXX_Size() int {
	return m.XSize()
}
func (m *ConfChangeSingle) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfChangeSingle.DiscardUnknown(m)
}

var xxx_messageInfo_ConfChangeSingle proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("raftpb.ConfChangeType", ConfChangeType_name, ConfChangeType_value)
	proto.RegisterType((*ConfChange)(nil), "raftpb.ConfChange")
	proto.RegisterType((*ConfChangeSingle)(nil), "raftpb.ConfChangeSingle")
}

func init() {
	proto.RegisterFile("github.com/olive-io/olive/api/raftpb/raft.proto", fileDescriptor_c7ed0c999701337a)
}

var fileDescriptor_c7ed0c999701337a = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x69, 0x68, 0x61, 0x90, 0x52, 0x97, 0x2a, 0xa1, 0xc8, 0xb6, 0xf4, 0x62, 0x51,
	0x4c, 0xc4, 0xab, 0x27, 0xdb, 0x5e, 0x02, 0xe2, 0x21, 0xea, 0xc5, 0x8b, 0xa4, 0xdd, 0x6d, 0xba,
	0xd0, 0x66, 0x42, 0x58, 0x8b, 0x7d, 0x0b, 0xdf, 0xc1, 0x97, 0xe9, 0xb1, 0x47, 0x4f, 0x45, 0xd3,
	0x17, 0x91, 0xec, 0x56, 0x42, 0x2f, 0x22, 0x9e, 0x76, 0xf8, 0xbf, 0x99, 0x8f, 0x1f, 0x16, 0xfc,
	0x58, 0xaa, 0xe9, 0xcb, 0xc8, 0x1b, 0xe3, 0xdc, 0xc7, 0x99, 0x5c, 0x88, 0x0b, 0x89, 0x66, 0xf0,
	0xa3, 0x54, 0xfa, 0x59, 0x34, 0x51, 0xe9, 0x48, 0x3f, 0x5e, 0x9a, 0xa1, 0x42, 0x5a, 0x35, 0x51,
	0xab, 0x19, 0x63, 0x8c, 0x3a, 0xf2, 0x8b, 0xc9, 0xd0, 0xee, 0x3b, 0x01, 0x18, 0x60, 0x32, 0x19,
	0x4c, 0xa3, 0x24, 0x16, 0xf4, 0x12, 0x1c, 0xb5, 0x4c, 0x85, 0x6b, 0x77, 0x48, 0xaf, 0x7e, 0x75,
	0xec, 0x99, 0x5b, 0xaf, 0xdc, 0x78, 0x58, 0xa6, 0xa2, 0xef, 0xac, 0x36, 0x6d, 0x2b, 0xd4, 0x9b,
	0xf4, 0x14, 0x6a, 0x09, 0x72, 0xf1, 0x2c, 0xb9, 0x5b, 0xe9, 0x90, 0x9e, 0xd3, 0xaf, 0x17, 0x30,
	0xdf, 0xb4, 0xab, 0x77, 0xc8, 0x45, 0x30, 0x0c, 0xab, 0x05, 0x0e, 0x38, 0x65, 0x50, 0x1b, 0x63,
	0xa2, 0xc4, 0xab, 0x72, 0x9d, 0x0e, 0xe9, 0x1d, 0xec, 0x2c, 0x3f, 0x21, 0x6d, 0x81, 0x2d, 0xb9,
	0x4b, 0xb4, 0x03, 0x76, 0x0e, 0x3b, 0x18, 0x86, 0xb6, 0xe4, 0xdd, 0x39, 0x34, 0xca, 0x0a, 0xf7,
	0x32, 0x89, 0x67, 0x65, 0x55, 0xf2, 0x9f, 0xaa, 0xf6, 0x6f, 0x55, 0xcf, 0x96, 0x50, 0xdf, 0xd7,
	0xd0, 0x23, 0x38, 0x2c, 0x93, 0x1b, 0xce, 0x8b, 0x83, 0x86, 0x45, 0x5d, 0x68, 0x96, 0x71, 0x28,
	0xe6, 0xb8, 0x10, 0x9a, 0x90, 0x7d, 0xf2, 0x98, 0xf2, 0x48, 0x19, 0x62, 0xd3, 0x13, 0x70, 0xf7,
	0x54, 0xb7, 0x22, 0xca, 0x12, 0x91, 0x69, 0x5a, 0xe9, 0x07, 0xab, 0x2f, 0x66, 0xad, 0x72, 0x46,
	0xd6, 0x39, 0x23, 0x9f, 0x39, 0x23, 0x6f, 0x5b, 0x66, 0xad, 0xb7, 0xcc, 0xfa, 0xd8, 0x32, 0xeb,
	0xe9, 0xfc, 0x2f, 0x9f, 0x7f, 0x6d, 0x9e, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x07, 0xb3,
	0x07, 0x2a, 0x02, 0x00, 0x00,
}

func (m *ConfChange) XSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.NodeID))
	if m.Context != nil {
		l = len(m.Context)
		n += 1 + l + sovRaft(uint64(l))
	}
	n += 1 + sovRaft(uint64(m.ID))
	return n
}

func (m *ConfChangeSingle) XSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.NodeID))
	return n
}

func sovRaft(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfChange) Marshal() (dAtA []byte, err error) {
	size := m.XSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.XSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		i -= len(m.Context)
		copy(dAtA[i:], m.Context)
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Context)))
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintRaft(dAtA, i, uint64(m.NodeID))
	i--
	dAtA[i] = 0x18
	i = encodeVarintRaft(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x10
	i = encodeVarintRaft(dAtA, i, uint64(m.ID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *ConfChangeSingle) Marshal() (dAtA []byte, err error) {
	size := m.XSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfChangeSingle) MarshalTo(dAtA []byte) (int, error) {
	size := m.XSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfChangeSingle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintRaft(dAtA, i, uint64(m.NodeID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintRaft(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintRaft(dAtA []byte, offset int, v uint64) int {
	offset -= sovRaft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ConfChangeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
			if m.Context == nil {
				m.Context = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfChangeSingle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfChangeSingle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChangeSingle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ConfChangeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRaft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRaft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRaft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRaft = fmt.Errorf("proto: unexpected end of group")
)
