// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/http/rest/swagger/swagger.proto

package swagger

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ResponseBody struct {
	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message  string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail   string            `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ResponseBody) Reset()         { *m = ResponseBody{} }
func (m *ResponseBody) String() string { return proto.CompactTextString(m) }
func (*ResponseBody) ProtoMessage()    {}
func (*ResponseBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c31151fad61cfa, []int{0}
}
func (m *ResponseBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBody.Merge(m, src)
}
func (m *ResponseBody) XXX_Size() int {
	return m.Size()
}
func (m *ResponseBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBody.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBody proto.InternalMessageInfo

func (*ResponseBody) XXX_MessageName() string {
	return "github.com.zchat.team.zim.proto.http.rest.swagger.ResponseBody"
}
func init() {
	proto.RegisterType((*ResponseBody)(nil), "github.com.zchat.team.zim.proto.http.rest.swagger.ResponseBody")
	proto.RegisterMapType((map[string]string)(nil), "github.com.zchat.team.zim.proto.http.rest.swagger.ResponseBody.MetadataEntry")
}

func init() {
	proto.RegisterFile("proto/http/rest/swagger/swagger.proto", fileDescriptor_b8c31151fad61cfa)
}

var fileDescriptor_b8c31151fad61cfa = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xce, 0x26, 0xe9, 0x9f, 0xa1, 0x55, 0x65, 0x0a, 0x84, 0x1c, 0x56, 0xa6, 0x08, 0x14, 0x2a,
	0xc5, 0x4e, 0xd2, 0x56, 0x6a, 0xc3, 0xa9, 0x45, 0x55, 0x25, 0xa0, 0x1c, 0xf6, 0xd8, 0x9b, 0xb3,
	0x3b, 0xda, 0x35, 0xcd, 0xda, 0x96, 0xed, 0x4d, 0x49, 0x9e, 0x02, 0x38, 0xf1, 0x38, 0x1c, 0x7b,
	0xec, 0x91, 0x23, 0x34, 0x2f, 0xc0, 0x03, 0x70, 0x40, 0xbb, 0x69, 0x68, 0xab, 0x42, 0x10, 0xa7,
	0x7c, 0xe3, 0x6f, 0xf2, 0x7d, 0x33, 0xb3, 0x33, 0xe8, 0xa9, 0x36, 0xca, 0x29, 0x96, 0x38, 0xa7,
	0x99, 0x01, 0xeb, 0x98, 0x3d, 0xe5, 0x71, 0x0c, 0x66, 0xfa, 0x4b, 0x0b, 0x1e, 0xb7, 0x63, 0xe1,
	0x92, 0xac, 0x47, 0x43, 0x95, 0xd2, 0x51, 0x98, 0x70, 0x47, 0x1d, 0xf0, 0x94, 0x8e, 0x44, 0x3a,
	0x49, 0xa0, 0xb9, 0x00, 0xcd, 0x05, 0xe8, 0xe5, 0x1f, 0xeb, 0x13, 0x22, 0x6c, 0xc6, 0x20, 0x9b,
	0x4a, 0x83, 0xe4, 0x5a, 0x0c, 0x3a, 0x4c, 0x69, 0x27, 0x94, 0xb4, 0x8c, 0x4b, 0xa9, 0x1c, 0x2f,
	0xf0, 0x24, 0xb1, 0xbe, 0x16, 0xab, 0x58, 0x4d, 0xaa, 0xc9, 0xd1, 0xe4, 0x75, 0xfd, 0xa7, 0x87,
	0xee, 0x06, 0x60, 0xb5, 0x92, 0x16, 0xf6, 0x55, 0x34, 0xc4, 0x18, 0x55, 0x43, 0x15, 0x41, 0xcd,
	0x23, 0x5e, 0x63, 0x2e, 0x28, 0x30, 0xae, 0xa1, 0x85, 0x14, 0xac, 0xe5, 0x31, 0xd4, 0xca, 0xc4,
	0x6b, 0x2c, 0x05, 0xd3, 0x10, 0x3f, 0x40, 0xf3, 0x11, 0x38, 0x2e, 0xfa, 0xb5, 0x4a, 0x41, 0x5c,
	0x46, 0x58, 0xa0, 0xc5, 0x14, 0x1c, 0x8f, 0xb8, 0xe3, 0xb5, 0x2a, 0xa9, 0x34, 0xee, 0x74, 0x8e,
	0xe8, 0x7f, 0xb7, 0x48, 0xaf, 0x17, 0x46, 0x8f, 0x2e, 0xf5, 0x0e, 0xa4, 0x33, 0xc3, 0xe0, 0xb7,
	0x7c, 0xfd, 0x05, 0x5a, 0xbe, 0x41, 0xe1, 0x55, 0x54, 0x39, 0x81, 0x61, 0xd1, 0xc0, 0x52, 0x90,
	0x43, 0xbc, 0x86, 0xe6, 0x06, 0xbc, 0x9f, 0x4d, 0xab, 0x9f, 0x04, 0xdd, 0xf2, 0x8e, 0xb7, 0xff,
	0xa3, 0xfa, 0x69, 0xef, 0x63, 0x15, 0xaf, 0xa0, 0x6a, 0x6e, 0xdc, 0x99, 0x1f, 0xb4, 0x68, 0x9b,
	0xb6, 0xea, 0x2b, 0x7d, 0x15, 0xf2, 0x7e, 0xa2, 0xac, 0xeb, 0x6e, 0xb7, 0x76, 0x5a, 0xeb, 0x0b,
	0x8c, 0x6b, 0xc1, 0x06, 0xed, 0x8d, 0xb2, 0x57, 0xee, 0xac, 0x72, 0xad, 0xfb, 0x22, 0x2c, 0x86,
	0xcb, 0xde, 0x59, 0x25, 0xbb, 0xb7, 0x5e, 0x82, 0x57, 0xa8, 0xb2, 0xd5, 0x6a, 0xe3, 0x97, 0xe8,
	0x59, 0x00, 0x2e, 0x33, 0x12, 0x22, 0x72, 0x9a, 0x80, 0x24, 0x2e, 0x01, 0x92, 0x59, 0x30, 0x24,
	0x52, 0x60, 0x89, 0x54, 0x8e, 0xf4, 0x55, 0x2c, 0x24, 0xc5, 0x8f, 0xd0, 0xc3, 0xfa, 0xfd, 0x3f,
	0x36, 0x1e, 0x9c, 0xe4, 0x5a, 0x9b, 0x38, 0x42, 0x87, 0xff, 0xd2, 0x4a, 0xf8, 0x00, 0x88, 0x06,
	0x93, 0x0a, 0x6b, 0x85, 0x92, 0xc4, 0x29, 0xc2, 0xc3, 0x10, 0xac, 0x2d, 0x72, 0x0d, 0x58, 0x95,
	0x99, 0x10, 0x66, 0x9a, 0xbd, 0xcd, 0xcd, 0xb6, 0xf0, 0x21, 0xda, 0xb8, 0x6d, 0x36, 0x15, 0xb8,
	0x32, 0x84, 0xf7, 0xc2, 0xba, 0x99, 0x7a, 0x6f, 0x50, 0x65, 0x7b, 0x77, 0x17, 0x1f, 0xa0, 0xc6,
	0x6d, 0xbd, 0x5e, 0x66, 0x85, 0xcc, 0xab, 0x03, 0x63, 0x94, 0x21, 0x09, 0xd7, 0x1a, 0x66, 0x8e,
	0xe2, 0xf8, 0x09, 0x7a, 0x8c, 0xd0, 0x9e, 0x16, 0xaf, 0x61, 0xb8, 0x97, 0xb9, 0x04, 0xdf, 0x5b,
	0x2c, 0xd7, 0x97, 0x73, 0xa4, 0x8c, 0x18, 0x15, 0xe3, 0x27, 0xe5, 0xde, 0x2a, 0x5a, 0xb9, 0x91,
	0x54, 0x32, 0x2d, 0x54, 0x19, 0x89, 0x14, 0x3f, 0xcf, 0x37, 0xcc, 0x76, 0x19, 0xbb, 0xda, 0xc4,
	0xe9, 0x19, 0x36, 0xf3, 0xef, 0x3b, 0xc5, 0x99, 0x38, 0xfb, 0xee, 0x97, 0xce, 0x2e, 0x7c, 0xef,
	0xfc, 0xc2, 0xf7, 0xbe, 0x5d, 0xf8, 0xde, 0x87, 0xb1, 0x5f, 0xfa, 0x3c, 0xf6, 0x4b, 0x5f, 0xc6,
	0xbe, 0x77, 0x36, 0xf6, 0xbd, 0xf3, 0xb1, 0x5f, 0xfa, 0x3a, 0xf6, 0x4b, 0xc7, 0xd7, 0xee, 0x96,
	0x15, 0x4b, 0xdd, 0xcc, 0x97, 0x9a, 0x8d, 0x44, 0xca, 0xfe, 0x72, 0xf8, 0xbd, 0xf9, 0x82, 0xd8,
	0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0xda, 0x5e, 0x12, 0xe8, 0x1a, 0x04, 0x00, 0x00,
}

func (m *ResponseBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintSwagger(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSwagger(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSwagger(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintSwagger(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintSwagger(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintSwagger(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwagger(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwagger(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResponseBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovSwagger(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovSwagger(uint64(l))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovSwagger(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovSwagger(uint64(len(k))) + 1 + len(v) + sovSwagger(uint64(len(v)))
			n += mapEntrySize + 1 + sovSwagger(uint64(mapEntrySize))
		}
	}
	return n
}

func sovSwagger(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwagger(x uint64) (n int) {
	return sovSwagger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResponseBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwagger
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
			return fmt.Errorf("proto: ResponseBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwagger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwagger
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
				return ErrInvalidLengthSwagger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwagger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwagger
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
				return ErrInvalidLengthSwagger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwagger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwagger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSwagger
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwagger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSwagger
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSwagger
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthSwagger
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSwagger
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSwagger
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthSwagger
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthSwagger
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSwagger(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthSwagger
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwagger(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwagger
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
func skipSwagger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwagger
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
					return 0, ErrIntOverflowSwagger
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
					return 0, ErrIntOverflowSwagger
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
				return 0, ErrInvalidLengthSwagger
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwagger
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwagger
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwagger        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwagger          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwagger = fmt.Errorf("proto: unexpected end of group")
)