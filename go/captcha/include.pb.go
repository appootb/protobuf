// Code generated by protoc-gen-go. DO NOT EDIT.
// source: include.proto

package captcha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Captcha code channel enums.
type Channel int32

const (
	Channel_CHANNEL_UNSPECIFIED Channel = 0
	Channel_CHANNEL_EMAIL       Channel = 1
	Channel_CHANNEL_SMS         Channel = 2
	Channel_CHANNEL_PHONE       Channel = 3
	Channel_CHANNEL_OTP         Channel = 4
)

var Channel_name = map[int32]string{
	0: "CHANNEL_UNSPECIFIED",
	1: "CHANNEL_EMAIL",
	2: "CHANNEL_SMS",
	3: "CHANNEL_PHONE",
	4: "CHANNEL_OTP",
}

var Channel_value = map[string]int32{
	"CHANNEL_UNSPECIFIED": 0,
	"CHANNEL_EMAIL":       1,
	"CHANNEL_SMS":         2,
	"CHANNEL_PHONE":       3,
	"CHANNEL_OTP":         4,
}

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return proto.EnumName(Channel_name, int32(x))
}

func (x *Channel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Channel_value, data, "Channel")
	if err != nil {
		return err
	}
	*x = Channel(value)
	return nil
}

func (Channel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1eef7b97e24a48e, []int{0}
}

// Captcha code category enums.
type Category int32

const (
	Category_CATEGORY_UNSPECIFIED Category = 0
	Category_CATEGORY_REGISTER    Category = 1
	Category_CATEGORY_LOGIN       Category = 2
	Category_CATEGORY_RESET_PWD   Category = 3
	Category_CATEGORY_BIND        Category = 4
	Category_CATEGORY_UNBIND      Category = 5
)

var Category_name = map[int32]string{
	0: "CATEGORY_UNSPECIFIED",
	1: "CATEGORY_REGISTER",
	2: "CATEGORY_LOGIN",
	3: "CATEGORY_RESET_PWD",
	4: "CATEGORY_BIND",
	5: "CATEGORY_UNBIND",
}

var Category_value = map[string]int32{
	"CATEGORY_UNSPECIFIED": 0,
	"CATEGORY_REGISTER":    1,
	"CATEGORY_LOGIN":       2,
	"CATEGORY_RESET_PWD":   3,
	"CATEGORY_BIND":        4,
	"CATEGORY_UNBIND":      5,
}

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}

func (x *Category) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Category_value, data, "Category")
	if err != nil {
		return err
	}
	*x = Category(value)
	return nil
}

func (Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1eef7b97e24a48e, []int{1}
}

func init() {
	proto.RegisterEnum("appootb.captcha.Channel", Channel_name, Channel_value)
	proto.RegisterEnum("appootb.captcha.Category", Category_name, Category_value)
}

func init() { proto.RegisterFile("include.proto", fileDescriptor_f1eef7b97e24a48e) }

var fileDescriptor_f1eef7b97e24a48e = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4e, 0xc2, 0x40,
	0x18, 0x85, 0x9d, 0x82, 0xc1, 0xfc, 0x04, 0x19, 0x06, 0x51, 0x4f, 0xe0, 0x82, 0x45, 0x7b, 0x86,
	0x76, 0x18, 0x61, 0x12, 0x98, 0x4e, 0xda, 0x1a, 0xa3, 0x21, 0x21, 0x6d, 0xad, 0x2d, 0x86, 0x76,
	0x1a, 0xd2, 0x2e, 0x38, 0x85, 0x77, 0x70, 0xe9, 0x51, 0x3c, 0x91, 0x4b, 0x63, 0x69, 0x49, 0xc3,
	0xf6, 0xfb, 0xbf, 0x37, 0x6f, 0xf2, 0x60, 0xb0, 0xcd, 0xc2, 0x5d, 0xf9, 0x16, 0xe9, 0xf9, 0x5e,
	0x15, 0x8a, 0x0c, 0xfd, 0x3c, 0x57, 0xaa, 0x08, 0xf4, 0xd0, 0xcf, 0x8b, 0x30, 0xf1, 0xa7, 0x1f,
	0xd0, 0xa3, 0x89, 0x9f, 0x65, 0xd1, 0x8e, 0xdc, 0xc1, 0x98, 0x2e, 0x4c, 0x21, 0xd8, 0x72, 0xf3,
	0x24, 0x5c, 0xc9, 0x28, 0x7f, 0xe4, 0x6c, 0x86, 0x2f, 0xc8, 0x08, 0x06, 0xcd, 0x81, 0xad, 0x4c,
	0xbe, 0xc4, 0x88, 0x0c, 0xa1, 0xdf, 0x20, 0x77, 0xe5, 0x62, 0xad, 0xed, 0xc8, 0x85, 0x2d, 0x18,
	0xee, 0xb4, 0x1d, 0xdb, 0x93, 0xb8, 0x3b, 0xfd, 0x44, 0x70, 0x45, 0xfd, 0x22, 0x8a, 0xd5, 0xfe,
	0x40, 0xee, 0xe1, 0x86, 0x9a, 0x1e, 0x9b, 0xdb, 0xce, 0xcb, 0x59, 0xdd, 0x04, 0x46, 0xa7, 0x8b,
	0xc3, 0xe6, 0xdc, 0xf5, 0x98, 0x83, 0x11, 0x21, 0x70, 0x7d, 0xc2, 0x4b, 0x7b, 0xce, 0x05, 0xd6,
	0xc8, 0x2d, 0x90, 0x96, 0xea, 0x32, 0x6f, 0x23, 0x9f, 0x67, 0xb8, 0x53, 0xfd, 0xa6, 0xe1, 0x16,
	0x17, 0x33, 0xdc, 0x25, 0x63, 0x18, 0xb6, 0xfa, 0x2a, 0x78, 0x69, 0x1d, 0x60, 0x1c, 0xaa, 0x54,
	0x3f, 0x1b, 0xc5, 0xea, 0x53, 0x95, 0xa6, 0x2a, 0x93, 0xff, 0x93, 0x2d, 0x90, 0x44, 0xaf, 0x0f,
	0xf1, 0xb6, 0x48, 0xca, 0x40, 0x0f, 0x55, 0x6a, 0xd4, 0xb2, 0x51, 0x0d, 0x1a, 0x94, 0xef, 0x46,
	0xac, 0x8c, 0x3a, 0xf8, 0x8b, 0xd0, 0x97, 0xd6, 0xa1, 0xd2, 0xfa, 0xd6, 0x7a, 0xf4, 0x88, 0x7e,
	0xb4, 0x89, 0x79, 0x0c, 0xac, 0xab, 0xe7, 0xd6, 0x35, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x94, 0x97, 0x3d, 0x99, 0x01, 0x00, 0x00,
}
