// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

package account

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Account property.
type Property struct {
	Name                 *string         `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value                *_struct.Struct `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Property) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

// Account properties.
type Properties struct {
	Kvs                  map[string]*_struct.Struct `protobuf:"bytes,1,rep,name=kvs" json:"kvs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Properties) Reset()         { *m = Properties{} }
func (m *Properties) String() string { return proto.CompactTextString(m) }
func (*Properties) ProtoMessage()    {}
func (*Properties) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *Properties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Properties.Unmarshal(m, b)
}
func (m *Properties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Properties.Marshal(b, m, deterministic)
}
func (m *Properties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Properties.Merge(m, src)
}
func (m *Properties) XXX_Size() int {
	return xxx_messageInfo_Properties.Size(m)
}
func (m *Properties) XXX_DiscardUnknown() {
	xxx_messageInfo_Properties.DiscardUnknown(m)
}

var xxx_messageInfo_Properties proto.InternalMessageInfo

func (m *Properties) GetKvs() map[string]*_struct.Struct {
	if m != nil {
		return m.Kvs
	}
	return nil
}

func init() {
	proto.RegisterType((*Property)(nil), "appootb.account.Property")
	proto.RegisterType((*Properties)(nil), "appootb.account.Properties")
	proto.RegisterMapType((map[string]*_struct.Struct)(nil), "appootb.account.Properties.KvsEntry")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x99, 0xc9, 0x4a, 0xdb, 0xa9, 0xa2, 0x4e, 0xd1, 0x6e, 0xd3, 0xa2, 0x21, 0x14, 0x0d,
	0x82, 0x33, 0xb0, 0x07, 0x91, 0xde, 0xba, 0x52, 0x2a, 0x88, 0x18, 0xd2, 0x9b, 0xf6, 0x32, 0x1b,
	0xa7, 0x69, 0xe8, 0x26, 0x6f, 0xc8, 0x4c, 0x16, 0xa2, 0x78, 0xf1, 0x2b, 0x88, 0x37, 0x4f, 0x1e,
	0xc5, 0x0f, 0x22, 0x5e, 0xbd, 0xf8, 0x01, 0xfc, 0x10, 0x1e, 0x25, 0xc9, 0xa4, 0x0b, 0x91, 0x2c,
	0xf4, 0x36, 0x79, 0xff, 0xf9, 0xff, 0xfe, 0x2f, 0x6f, 0x1e, 0xb9, 0xa1, 0x0a, 0x38, 0x4b, 0xe7,
	0x92, 0xa9, 0x02, 0x0c, 0xd0, 0x9b, 0x42, 0x29, 0x00, 0x33, 0x63, 0x22, 0x8e, 0xa1, 0xcc, 0x8d,
	0x7b, 0xdf, 0x16, 0xb8, 0x92, 0x45, 0x96, 0x6a, 0x9d, 0x42, 0xce, 0x33, 0x69, 0xce, 0xe1, 0x6d,
	0xeb, 0x70, 0xf7, 0x12, 0x80, 0x64, 0x2e, 0xb9, 0x50, 0x29, 0x17, 0x79, 0x0e, 0x46, 0x98, 0x14,
	0x72, 0x6d, 0xd5, 0x5d, 0xab, 0x36, 0x5f, 0xb3, 0xf2, 0x8c, 0xcb, 0x4c, 0x99, 0xaa, 0x67, 0xbd,
	0x14, 0xb5, 0x29, 0xca, 0xd8, 0xb4, 0xaa, 0xff, 0x92, 0xac, 0x87, 0x05, 0x28, 0x59, 0x98, 0x8a,
	0x52, 0x32, 0xca, 0x45, 0x26, 0xc7, 0xc8, 0xc3, 0xc1, 0x46, 0xd4, 0x9c, 0xe9, 0x63, 0x72, 0x6d,
	0x21, 0xe6, 0xa5, 0x1c, 0x63, 0x0f, 0x05, 0x9b, 0x93, 0x6d, 0xd6, 0xd2, 0x58, 0x47, 0x63, 0x27,
	0x0d, 0x2d, 0x6a, 0x6f, 0xf9, 0x9f, 0x11, 0x21, 0x96, 0x97, 0x4a, 0x4d, 0x9f, 0x10, 0xe7, 0x62,
	0xa1, 0xc7, 0xc8, 0x73, 0x82, 0xcd, 0xc9, 0x3e, 0xeb, 0xfd, 0x36, 0x5b, 0xde, 0x64, 0x2f, 0x16,
	0xfa, 0x28, 0x37, 0x45, 0x15, 0xd5, 0x06, 0xf7, 0x15, 0x59, 0xef, 0x0a, 0xf4, 0x16, 0x71, 0x2e,
	0x64, 0x35, 0x46, 0x1e, 0x0a, 0x36, 0xa2, 0xfa, 0x78, 0xc5, 0x9e, 0x0e, 0xf0, 0x53, 0x34, 0xf9,
	0x8e, 0xc9, 0x5a, 0xd8, 0xbe, 0x01, 0x15, 0xc4, 0x39, 0x91, 0x86, 0xee, 0x0c, 0xb5, 0x53, 0xb9,
	0x77, 0xff, 0x23, 0x1e, 0xd5, 0x03, 0xf5, 0x1f, 0xfe, 0xfe, 0x82, 0x7e, 0x8c, 0x3e, 0xfe, 0xfa,
	0xf3, 0x09, 0xef, 0xb9, 0xdb, 0xdc, 0xfa, 0xb8, 0x7d, 0x5c, 0xfe, 0xbe, 0x9e, 0xd8, 0x87, 0x03,
	0xf4, 0xa8, 0x8e, 0x38, 0x5e, 0x1d, 0x31, 0x2c, 0xf9, 0xfb, 0xcb, 0x94, 0x1d, 0x3a, 0x94, 0x42,
	0xdf, 0x90, 0xd1, 0xb1, 0x34, 0x9a, 0x0e, 0xf4, 0xea, 0xee, 0xae, 0x98, 0xb6, 0x7f, 0x6f, 0x19,
	0xb1, 0x45, 0x6f, 0xf7, 0x23, 0xf4, 0xf4, 0x1d, 0xd9, 0x8a, 0x21, 0xeb, 0x13, 0xa6, 0xd7, 0xed,
	0x08, 0xc3, 0x3a, 0xea, 0x39, 0x0a, 0xd1, 0xeb, 0x07, 0x49, 0x6a, 0xce, 0xcb, 0x19, 0x8b, 0x21,
	0xe3, 0x97, 0x3b, 0xdc, 0x2d, 0x5a, 0x02, 0x1d, 0xf9, 0x2f, 0x42, 0x5f, 0xb1, 0xf3, 0x2c, 0x9c,
	0x7e, 0xc3, 0x6b, 0x87, 0x6d, 0xe9, 0x27, 0xbe, 0x73, 0xd8, 0x1a, 0x4e, 0x1b, 0xdc, 0xa9, 0xad,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x55, 0xfc, 0x54, 0x21, 0x2c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileClient interface {
	// Set or update account property.
	Set(ctx context.Context, in *Property, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get account property.
	Get(ctx context.Context, in *Property, opts ...grpc.CallOption) (*Property, error)
	// Get account all properties.
	Gets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Properties, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Set(ctx context.Context, in *Property, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.account.Profile/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Get(ctx context.Context, in *Property, opts ...grpc.CallOption) (*Property, error) {
	out := new(Property)
	err := c.cc.Invoke(ctx, "/appootb.account.Profile/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Gets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Properties, error) {
	out := new(Properties)
	err := c.cc.Invoke(ctx, "/appootb.account.Profile/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
type ProfileServer interface {
	// Set or update account property.
	Set(context.Context, *Property) (*empty.Empty, error)
	// Get account property.
	Get(context.Context, *Property) (*Property, error)
	// Get account all properties.
	Gets(context.Context, *empty.Empty) (*Properties, error)
}

// UnimplementedProfileServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (*UnimplementedProfileServer) Set(ctx context.Context, req *Property) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedProfileServer) Get(ctx context.Context, req *Property) (*Property, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProfileServer) Gets(ctx context.Context, req *empty.Empty) (*Properties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Property)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Profile/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Set(ctx, req.(*Property))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Property)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Profile/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Get(ctx, req.(*Property))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Profile/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Gets(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.account.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Profile_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Profile_Get_Handler,
		},
		{
			MethodName: "Gets",
			Handler:    _Profile_Gets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
