// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/rest/group/group.proto

package group

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateReq struct {
	// 群主
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// 成员
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// 群名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 群ID，如果不传，zim会生成一个群ID
	GroupId string `protobuf:"bytes,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Notice  string `protobuf:"bytes,5,opt,name=notice,proto3" json:"notice,omitempty"`
	Intro   string `protobuf:"bytes,6,opt,name=intro,proto3" json:"intro,omitempty"`
	Avatar  string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_663754992404ef89, []int{0}
}
func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (*CreateReq) XXX_MessageName() string {
	return "github.com.zchat.team.zim.api.rest.group.CreateReq"
}

type CreateRsp struct {
	// 群ID
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (m *CreateRsp) Reset()         { *m = CreateRsp{} }
func (m *CreateRsp) String() string { return proto.CompactTextString(m) }
func (*CreateRsp) ProtoMessage()    {}
func (*CreateRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_663754992404ef89, []int{1}
}
func (m *CreateRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRsp.Merge(m, src)
}
func (m *CreateRsp) XXX_Size() int {
	return m.Size()
}
func (m *CreateRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRsp proto.InternalMessageInfo

func (*CreateRsp) XXX_MessageName() string {
	return "github.com.zchat.team.zim.api.rest.group.CreateRsp"
}
func init() {
	proto.RegisterType((*CreateReq)(nil), "github.com.zchat.team.zim.api.rest.group.CreateReq")
	proto.RegisterType((*CreateRsp)(nil), "github.com.zchat.team.zim.api.rest.group.CreateRsp")
}

func init() { proto.RegisterFile("api/rest/group/group.proto", fileDescriptor_663754992404ef89) }

var fileDescriptor_663754992404ef89 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0x8e, 0xef, 0x9a, 0x1c, 0x67, 0x16, 0x64, 0x2a, 0x94, 0x46, 0x60, 0xaa, 0x0c, 0xa8, 0x42,
	0x9c, 0x23, 0xb5, 0x1b, 0x5b, 0xcb, 0x80, 0x58, 0x33, 0xb2, 0x54, 0x4e, 0xce, 0xe4, 0x2c, 0x35,
	0x7e, 0xc6, 0x71, 0x0f, 0xe9, 0x46, 0xc4, 0xc0, 0x88, 0x60, 0x41, 0x62, 0x61, 0xe8, 0xc4, 0x4f,
	0x68, 0xa5, 0x8a, 0xf1, 0xc6, 0x4a, 0x2c, 0x8c, 0xe5, 0x42, 0x75, 0x7f, 0x03, 0xc5, 0x6e, 0xab,
	0x9e, 0xc4, 0x00, 0x4b, 0xf2, 0xbe, 0xcf, 0xdf, 0xf7, 0xbd, 0x97, 0xf8, 0xe1, 0x84, 0x6b, 0x99,
	0x19, 0xd1, 0xd8, 0xac, 0x32, 0x70, 0xa8, 0xfd, 0x93, 0x69, 0x03, 0x16, 0xc8, 0x56, 0x25, 0xed,
	0xe4, 0xb0, 0x60, 0x25, 0xd4, 0x6c, 0x56, 0x4e, 0xb8, 0x65, 0x56, 0xf0, 0x9a, 0xcd, 0x64, 0xcd,
	0xb8, 0x96, 0xac, 0x73, 0x31, 0xa7, 0x4f, 0xee, 0x57, 0x00, 0xd5, 0x81, 0xc8, 0xba, 0x30, 0xae,
	0x14, 0x58, 0x6e, 0x25, 0xa8, 0xc6, 0xe7, 0x24, 0xeb, 0x15, 0x54, 0xe0, 0xca, 0xac, 0xab, 0x2e,
	0x59, 0xff, 0x2a, 0x47, 0x95, 0x50, 0x23, 0xd0, 0x42, 0x71, 0x2d, 0xa7, 0xdb, 0x19, 0x68, 0xe7,
	0xfc, 0x4b, 0xca, 0xc3, 0x1b, 0x3d, 0x5e, 0x49, 0x71, 0x30, 0xde, 0x2f, 0xc4, 0x84, 0x4f, 0x25,
	0x18, 0x2f, 0x48, 0x8f, 0x11, 0x1e, 0x3e, 0x33, 0x82, 0x5b, 0x91, 0x8b, 0xd7, 0x64, 0x03, 0x87,
	0xf0, 0x46, 0x09, 0x13, 0xa3, 0x4d, 0xb4, 0x35, 0xdc, 0xeb, 0x9f, 0xef, 0xf6, 0x72, 0xcf, 0x90,
	0x07, 0x78, 0x50, 0x8b, 0xba, 0x10, 0xa6, 0x89, 0x7b, 0x9b, 0xfd, 0xab, 0xc3, 0x2b, 0x8e, 0x10,
	0xbc, 0xa6, 0x78, 0x2d, 0xe2, 0x7e, 0x67, 0xcc, 0x5d, 0x4d, 0x36, 0xf0, 0x2d, 0xf7, 0xa5, 0xfb,
	0x72, 0x1c, 0xaf, 0x39, 0x7e, 0xe0, 0xf0, 0x8b, 0x31, 0xb9, 0x87, 0x23, 0x05, 0x56, 0x96, 0x22,
	0x0e, 0xdd, 0xc1, 0x25, 0x22, 0xeb, 0x38, 0x94, 0xca, 0x1a, 0x88, 0x23, 0x47, 0x7b, 0xd0, 0xa9,
	0xf9, 0x94, 0x5b, 0x6e, 0xe2, 0x81, 0x57, 0x7b, 0x94, 0x3e, 0xba, 0x9e, 0xbd, 0xd1, 0x2b, 0xdd,
	0xd0, 0x4a, 0xb7, 0xed, 0x6f, 0x08, 0x87, 0xcf, 0xbb, 0x9a, 0x7c, 0x45, 0x38, 0xf2, 0x16, 0xb2,
	0xc3, 0xfe, 0xf5, 0xa6, 0xd8, 0xf5, 0x0f, 0x4a, 0xfe, 0xdf, 0xd4, 0xe8, 0xf4, 0xc9, 0xc7, 0xdd,
	0xbb, 0x38, 0x5a, 0x9e, 0x1e, 0x2d, 0x4f, 0xde, 0x91, 0xe1, 0xc5, 0xfb, 0x2f, 0x17, 0x27, 0xc7,
	0xcb, 0xd3, 0xa3, 0xb7, 0x3f, 0x7e, 0x7f, 0xea, 0xdd, 0x49, 0x6f, 0x67, 0x33, 0x59, 0xfb, 0x15,
	0x6a, 0x9e, 0xa2, 0xc7, 0x7b, 0xf9, 0xfc, 0x17, 0x0d, 0xe6, 0x0b, 0x8a, 0xce, 0x16, 0x14, 0x9d,
	0x2f, 0x28, 0xfa, 0xd0, 0xd2, 0xe0, 0x73, 0x4b, 0x83, 0xef, 0x2d, 0x45, 0xf3, 0x96, 0xa2, 0xb3,
	0x96, 0x06, 0x3f, 0x5b, 0x1a, 0xbc, 0xbc, 0xb1, 0x6a, 0x99, 0x9b, 0x65, 0xd4, 0xcd, 0xe2, 0xe2,
	0x56, 0x17, 0xb4, 0x88, 0xdc, 0x65, 0xef, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xff, 0x5d, 0xea,
	0x0a, 0xb9, 0x02, 0x00, 0x00,
}

func (m *CreateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Avatar) > 0 {
		i -= len(m.Avatar)
		copy(dAtA[i:], m.Avatar)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Avatar)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Intro) > 0 {
		i -= len(m.Intro)
		copy(dAtA[i:], m.Intro)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Intro)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Notice) > 0 {
		i -= len(m.Notice)
		copy(dAtA[i:], m.Notice)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Notice)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintGroup(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Notice)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Intro)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	return n
}

func (m *CreateRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: CreateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intro", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Intro = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *CreateRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: CreateRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
