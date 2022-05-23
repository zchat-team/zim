// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errno/errno.proto

package errno

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/zmicro-team/zmicro/core/errors"
	math "math"
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

type ErrorReason int32

const (
	ErrorReason_internal_server ErrorReason = 0
	ErrorReason_bad_request     ErrorReason = 1
	ErrorReason_custom          ErrorReason = 100
	ErrorReason_invalid_token   ErrorReason = 101
	ErrorReason_token_expired   ErrorReason = 102
	ErrorReason_token_revoked   ErrorReason = 103
	ErrorReason_login_conflict  ErrorReason = 104
)

var ErrorReason_name = map[int32]string{
	0:   "internal_server",
	1:   "bad_request",
	100: "custom",
	101: "invalid_token",
	102: "token_expired",
	103: "token_revoked",
	104: "login_conflict",
}

var ErrorReason_value = map[string]int32{
	"internal_server": 0,
	"bad_request":     1,
	"custom":          100,
	"invalid_token":   101,
	"token_expired":   102,
	"token_revoked":   103,
	"login_conflict":  104,
}

func (x ErrorReason) String() string {
	return proto.EnumName(ErrorReason_name, int32(x))
}

func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_123bf0aeec5a6d70, []int{0}
}

func init() {
	proto.RegisterEnum("github.com.zmicro.team.zim.errno.ErrorReason", ErrorReason_name, ErrorReason_value)
}

func init() { proto.RegisterFile("errno/errno.proto", fileDescriptor_123bf0aeec5a6d70) }

var fileDescriptor_123bf0aeec5a6d70 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2d, 0x2a, 0xca,
	0xcb, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x7a, 0x55, 0xb9, 0x99, 0xc9, 0x45, 0xf9, 0x7a, 0x25, 0xa9,
	0x89, 0xb9, 0x7a, 0x55, 0x99, 0xb9, 0x7a, 0x60, 0x75, 0x52, 0x66, 0x08, 0x15, 0xfa, 0x10, 0x15,
	0xba, 0x20, 0x15, 0x50, 0xb6, 0x7e, 0x72, 0x7e, 0x51, 0x2a, 0xc8, 0xc4, 0xfc, 0xa2, 0x62, 0x28,
	0x05, 0x31, 0x59, 0x6b, 0x17, 0x13, 0x17, 0xb7, 0x2b, 0x48, 0x20, 0x28, 0x35, 0xb1, 0x38, 0x3f,
	0x4f, 0x48, 0x9b, 0x8b, 0x3f, 0x33, 0xaf, 0x24, 0xb5, 0x28, 0x2f, 0x31, 0x27, 0xbe, 0x38, 0xb5,
	0xa8, 0x2c, 0xb5, 0x48, 0x80, 0x41, 0x4a, 0xec, 0x84, 0xdd, 0x17, 0xe6, 0x4b, 0x76, 0xfc, 0xcf,
	0xe6, 0xf4, 0x3e, 0xed, 0x5a, 0xf8, 0x74, 0xe6, 0x8a, 0x97, 0x53, 0x66, 0xbe, 0x58, 0xbf, 0x5e,
	0x48, 0x8b, 0x8b, 0x3b, 0x29, 0x31, 0x25, 0xbe, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x80,
	0x51, 0x4a, 0xf2, 0x84, 0xdd, 0x04, 0xe6, 0x4b, 0x76, 0x42, 0x2f, 0xd6, 0x6f, 0x7f, 0xb6, 0xb1,
	0xe9, 0x69, 0x7f, 0xd3, 0xb3, 0xa9, 0x1b, 0xa0, 0x6a, 0x95, 0xb8, 0xd8, 0x92, 0x4b, 0x8b, 0x4b,
	0xf2, 0x73, 0x05, 0x52, 0x40, 0xe6, 0xbd, 0x60, 0xbf, 0x64, 0xc7, 0xff, 0xa2, 0x7d, 0xd5, 0xd3,
	0x75, 0xb3, 0x9e, 0xec, 0xec, 0x84, 0xaa, 0xd1, 0xe4, 0xe2, 0xcd, 0xcc, 0x2b, 0x4b, 0xcc, 0xc9,
	0x4c, 0x89, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x48, 0x05, 0x29, 0x7d, 0x09, 0x52, 0xfa, 0x6c,
	0xfa, 0x82, 0x67, 0x53, 0x3b, 0x9e, 0xcf, 0x6a, 0x79, 0xb2, 0x7b, 0xc9, 0xf3, 0xce, 0x1e, 0x21,
	0x35, 0x2e, 0x5e, 0xb0, 0x92, 0xf8, 0xd4, 0x8a, 0x82, 0xcc, 0xa2, 0xd4, 0x14, 0x81, 0x34, 0x29,
	0xe1, 0x13, 0x76, 0xaf, 0xd8, 0x2f, 0xd9, 0xf1, 0x40, 0x54, 0xbc, 0xd8, 0xdf, 0xfe, 0x6c, 0xce,
	0x7c, 0x21, 0x1d, 0x98, 0xba, 0xa2, 0xd4, 0xb2, 0xfc, 0xec, 0xd4, 0x14, 0x81, 0x74, 0x90, 0x23,
	0x5f, 0xb3, 0x5f, 0xb2, 0x13, 0x82, 0xa8, 0x7b, 0xba, 0x7d, 0xd3, 0x8b, 0x45, 0xab, 0x9f, 0x4e,
	0xe8, 0x7a, 0x39, 0xa5, 0x41, 0x48, 0x9d, 0x8b, 0x2f, 0x27, 0x3f, 0x3d, 0x33, 0x2f, 0x3e, 0x39,
	0x3f, 0x2f, 0x2d, 0x27, 0x33, 0xb9, 0x44, 0x20, 0x03, 0x64, 0xec, 0x1b, 0x90, 0xb1, 0xcf, 0x67,
	0xee, 0x7e, 0xba, 0x77, 0xea, 0xd3, 0xb6, 0x4d, 0xcf, 0x57, 0x35, 0x4a, 0xb1, 0x1c, 0xb0, 0xfb,
	0xc2, 0xec, 0x64, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xce, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x80, 0x2b, 0x1a, 0x32, 0x73, 0x21, 0x11, 0x9a, 0xc4, 0x06,
	0x0e, 0x77, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x72, 0xc1, 0x9f, 0xe6, 0x01, 0x00,
	0x00,
}
