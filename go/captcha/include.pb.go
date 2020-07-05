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
)

var Category_name = map[int32]string{
	0: "CATEGORY_UNSPECIFIED",
	1: "CATEGORY_REGISTER",
	2: "CATEGORY_LOGIN",
	3: "CATEGORY_RESET_PWD",
}

var Category_value = map[string]int32{
	"CATEGORY_UNSPECIFIED": 0,
	"CATEGORY_REGISTER":    1,
	"CATEGORY_LOGIN":       2,
	"CATEGORY_RESET_PWD":   3,
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

// Captcha code request.
type CodeRequest struct {
	Channel              *Channel  `protobuf:"varint,1,req,name=channel,enum=appootb.captcha.Channel" json:"channel,omitempty"`
	Category             *Category `protobuf:"varint,2,req,name=category,enum=appootb.captcha.Category" json:"category,omitempty"`
	Target               *string   `protobuf:"bytes,3,req,name=target" json:"target,omitempty"`
	Value                *string   `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CodeRequest) Reset()         { *m = CodeRequest{} }
func (m *CodeRequest) String() string { return proto.CompactTextString(m) }
func (*CodeRequest) ProtoMessage()    {}
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1eef7b97e24a48e, []int{0}
}

func (m *CodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeRequest.Unmarshal(m, b)
}
func (m *CodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeRequest.Marshal(b, m, deterministic)
}
func (m *CodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeRequest.Merge(m, src)
}
func (m *CodeRequest) XXX_Size() int {
	return xxx_messageInfo_CodeRequest.Size(m)
}
func (m *CodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeRequest proto.InternalMessageInfo

func (m *CodeRequest) GetChannel() Channel {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return Channel_CHANNEL_UNSPECIFIED
}

func (m *CodeRequest) GetCategory() Category {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return Category_CATEGORY_UNSPECIFIED
}

func (m *CodeRequest) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

func (m *CodeRequest) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("appootb.captcha.Channel", Channel_name, Channel_value)
	proto.RegisterEnum("appootb.captcha.Category", Category_name, Category_value)
	proto.RegisterType((*CodeRequest)(nil), "appootb.captcha.CodeRequest")
}

func init() { proto.RegisterFile("include.proto", fileDescriptor_f1eef7b97e24a48e) }

var fileDescriptor_f1eef7b97e24a48e = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x9d, 0xe9, 0x2a, 0xbb, 0x87, 0xec, 0xee, 0xec, 0xec, 0x82, 0xf5, 0x8e, 0x78, 0x61,
	0x08, 0x17, 0x6d, 0x42, 0xe2, 0x03, 0x94, 0x61, 0x84, 0x26, 0xd0, 0x36, 0xd3, 0x1a, 0xa3, 0x21,
	0x21, 0xa5, 0x8c, 0x05, 0x43, 0x3b, 0x15, 0xa7, 0x26, 0xbc, 0x8e, 0x17, 0x5e, 0xf8, 0x28, 0x3e,
	0x91, 0x97, 0x46, 0xda, 0x12, 0x82, 0x97, 0xe7, 0xff, 0xbf, 0xaf, 0xa7, 0x99, 0x03, 0xb7, 0xdb,
	0x3c, 0xd9, 0x95, 0x6b, 0x69, 0x15, 0x7b, 0xa5, 0x15, 0xbd, 0x8f, 0x8b, 0x42, 0x29, 0xbd, 0xb2,
	0x92, 0xb8, 0xd0, 0xc9, 0x26, 0x7e, 0xfd, 0x13, 0x41, 0x9b, 0xa9, 0xb5, 0x14, 0xf2, 0x6b, 0x29,
	0xbf, 0x69, 0x3a, 0x84, 0x56, 0xb2, 0x89, 0xf3, 0x5c, 0xee, 0x4c, 0xd4, 0xc3, 0xfd, 0xbb, 0xa1,
	0x69, 0x5d, 0x28, 0x16, 0xab, 0x7a, 0xd1, 0x80, 0xf4, 0x2d, 0x5c, 0x27, 0xb1, 0x96, 0xa9, 0xda,
	0x1f, 0x4c, 0x7c, 0x94, 0x5e, 0xfd, 0x2f, 0xd5, 0x80, 0x38, 0xa1, 0xb4, 0x0b, 0x2f, 0x74, 0xbc,
	0x4f, 0xa5, 0x36, 0x8d, 0x1e, 0xee, 0xdf, 0x88, 0x7a, 0xa2, 0x4f, 0xf0, 0xfc, 0x7b, 0xbc, 0x2b,
	0xa5, 0x79, 0xd5, 0x43, 0xfd, 0x1b, 0x51, 0x0d, 0x83, 0x2f, 0xd0, 0xaa, 0x17, 0xd3, 0x97, 0xf0,
	0xc8, 0xa6, 0x8e, 0xe7, 0xf1, 0xd9, 0xf2, 0xbd, 0x17, 0x06, 0x9c, 0xb9, 0xef, 0x5c, 0x3e, 0x26,
	0xcf, 0xe8, 0x03, 0xdc, 0x36, 0x05, 0x9f, 0x3b, 0xee, 0x8c, 0x20, 0x7a, 0x0f, 0xed, 0x26, 0x0a,
	0xe7, 0x21, 0xc1, 0xe7, 0x4c, 0x30, 0xf5, 0x3d, 0x4e, 0x8c, 0x73, 0xc6, 0x8f, 0x02, 0x72, 0x35,
	0x48, 0xe1, 0xba, 0xf9, 0x5f, 0x6a, 0xc2, 0x13, 0x73, 0x22, 0x3e, 0xf1, 0xc5, 0xc7, 0x8b, 0x6d,
	0x1d, 0x78, 0x38, 0x35, 0x82, 0x4f, 0xdc, 0x30, 0xe2, 0x82, 0x20, 0x4a, 0xe1, 0xee, 0x14, 0xcf,
	0xfc, 0x89, 0xeb, 0x11, 0x4c, 0xbb, 0x40, 0xcf, 0xd0, 0x90, 0x47, 0xcb, 0xe0, 0xc3, 0x98, 0x18,
	0xa3, 0x03, 0x3c, 0x26, 0x2a, 0xbb, 0x7c, 0xac, 0x51, 0x9b, 0xa9, 0x2c, 0x53, 0x79, 0xf0, 0xef,
	0x64, 0x53, 0x14, 0xa0, 0x4f, 0x6f, 0xd2, 0xad, 0xde, 0x94, 0x2b, 0x2b, 0x51, 0x99, 0x5d, 0xc3,
	0xf6, 0xf1, 0xa0, 0xab, 0xf2, 0xb3, 0x9d, 0x2a, 0xbb, 0x16, 0xff, 0x20, 0xf4, 0x03, 0x1b, 0x2c,
	0x18, 0xfd, 0xc2, 0x2d, 0x56, 0x45, 0xbf, 0x71, 0xc7, 0xa9, 0x84, 0xc5, 0xf1, 0x73, 0x8b, 0x3a,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xba, 0xbb, 0xda, 0xdc, 0x19, 0x02, 0x00, 0x00,
}