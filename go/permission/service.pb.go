// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Service visible scope.
type VisibleScope int32

const (
	VisibleScope_DEFAULT_SCOPE VisibleScope = 0
	VisibleScope_INNER_SCOPE   VisibleScope = 100
	VisibleScope_ALL_SCOPES    VisibleScope = 999
)

var VisibleScope_name = map[int32]string{
	0:   "DEFAULT_SCOPE",
	100: "INNER_SCOPE",
	999: "ALL_SCOPES",
}

var VisibleScope_value = map[string]int32{
	"DEFAULT_SCOPE": 0,
	"INNER_SCOPE":   100,
	"ALL_SCOPES":    999,
}

func (x VisibleScope) String() string {
	return proto.EnumName(VisibleScope_name, int32(x))
}

func (VisibleScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

// Service visible.
type ServiceVisible struct {
	Scope                VisibleScope `protobuf:"varint,1,opt,name=scope,proto3,enum=appootb.permission.service.VisibleScope" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ServiceVisible) Reset()         { *m = ServiceVisible{} }
func (m *ServiceVisible) String() string { return proto.CompactTextString(m) }
func (*ServiceVisible) ProtoMessage()    {}
func (*ServiceVisible) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ServiceVisible) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVisible.Unmarshal(m, b)
}
func (m *ServiceVisible) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVisible.Marshal(b, m, deterministic)
}
func (m *ServiceVisible) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVisible.Merge(m, src)
}
func (m *ServiceVisible) XXX_Size() int {
	return xxx_messageInfo_ServiceVisible.Size(m)
}
func (m *ServiceVisible) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVisible.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVisible proto.InternalMessageInfo

func (m *ServiceVisible) GetScope() VisibleScope {
	if m != nil {
		return m.Scope
	}
	return VisibleScope_DEFAULT_SCOPE
}

var E_Visible = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ServiceVisible)(nil),
	Field:         1507,
	Name:          "appootb.permission.service.visible",
	Tag:           "bytes,1507,opt,name=visible",
	Filename:      "service.proto",
}

func init() {
	proto.RegisterEnum("appootb.permission.service.VisibleScope", VisibleScope_name, VisibleScope_value)
	proto.RegisterType((*ServiceVisible)(nil), "appootb.permission.service.ServiceVisible")
	proto.RegisterExtension(E_Visible)
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0x2c, 0x28, 0xc8, 0xcf,
	0x2f, 0x49, 0xd2, 0x2b, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x83, 0xaa,
	0x90, 0x52, 0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xab, 0x4c, 0x2a, 0x4d, 0xd3, 0x4f,
	0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x82, 0xe8, 0x56, 0x0a, 0xe0, 0xe2, 0x0b,
	0x86, 0x28, 0x0e, 0xcb, 0x2c, 0xce, 0x4c, 0xca, 0x49, 0x15, 0xb2, 0xe3, 0x62, 0x2d, 0x4e, 0xce,
	0x2f, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd0, 0xc3, 0x6d, 0xbe, 0x1e, 0x54,
	0x4f, 0x30, 0x48, 0x7d, 0x10, 0x44, 0x9b, 0x96, 0x33, 0x17, 0x0f, 0xb2, 0xb0, 0x90, 0x20, 0x17,
	0xaf, 0x8b, 0xab, 0x9b, 0x63, 0xa8, 0x4f, 0x48, 0x7c, 0xb0, 0xb3, 0x7f, 0x80, 0xab, 0x00, 0x83,
	0x10, 0x3f, 0x17, 0xb7, 0xa7, 0x9f, 0x9f, 0x6b, 0x10, 0x54, 0x20, 0x45, 0x88, 0x9f, 0x8b, 0xcb,
	0xd1, 0xc7, 0x07, 0xc2, 0x0d, 0x16, 0x78, 0xce, 0x6e, 0x95, 0xc6, 0xc5, 0x5e, 0x06, 0x75, 0x8f,
	0xbc, 0x1e, 0xc4, 0x13, 0x7a, 0x30, 0x4f, 0xe8, 0x41, 0x1d, 0xec, 0x5f, 0x50, 0x92, 0x99, 0x9f,
	0x57, 0x2c, 0xf1, 0x98, 0x5b, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x0b, 0x9f, 0x43, 0x51, 0xfd, 0x18,
	0x04, 0x33, 0xdc, 0x29, 0x91, 0x4b, 0x2c, 0x39, 0x3f, 0x17, 0x8b, 0x4e, 0x27, 0x1e, 0xa8, 0x96,
	0x00, 0x90, 0xad, 0x1e, 0x8c, 0x01, 0x8c, 0x51, 0x3c, 0x7a, 0xd6, 0x08, 0xf9, 0x1f, 0x8c, 0x8c,
	0x8b, 0x98, 0x98, 0x9d, 0x03, 0x9c, 0x56, 0x31, 0x71, 0x05, 0xc0, 0x45, 0x4f, 0x31, 0x09, 0x80,
	0xd5, 0xc7, 0x20, 0x84, 0x92, 0xd8, 0xc0, 0xee, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x86,
	0xc6, 0xd0, 0x70, 0xb7, 0x01, 0x00, 0x00,
}
