// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_key.proto

package counter

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Counter key.
type Key struct {
	Product              *string            `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string            `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateId             *string            `protobuf:"bytes,3,req,name=relate_id,json=relateId" json:"relate_id,omitempty"`
	Value                *int64             `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	Expire               *duration.Duration `protobuf:"bytes,5,opt,name=expire" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc24b3d33c9123b, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *Key) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Key) GetRelateId() string {
	if m != nil && m.RelateId != nil {
		return *m.RelateId
	}
	return ""
}

func (m *Key) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Key) GetExpire() *duration.Duration {
	if m != nil {
		return m.Expire
	}
	return nil
}

// Multi keys.
type Keys struct {
	Product              *string          `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string          `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateIds            []string         `protobuf:"bytes,3,rep,name=relate_ids,json=relateIds" json:"relate_ids,omitempty"`
	Values               map[string]int64 `protobuf:"bytes,4,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc24b3d33c9123b, []int{1}
}

func (m *Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keys.Unmarshal(m, b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return xxx_messageInfo_Keys.Size(m)
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *Keys) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Keys) GetRelateIds() []string {
	if m != nil {
		return m.RelateIds
	}
	return nil
}

func (m *Keys) GetValues() map[string]int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "appootb.counter.Key")
	proto.RegisterType((*Keys)(nil), "appootb.counter.Keys")
	proto.RegisterMapType((map[string]int64)(nil), "appootb.counter.Keys.ValuesEntry")
}

func init() { proto.RegisterFile("inner_key.proto", fileDescriptor_1dc24b3d33c9123b) }

var fileDescriptor_1dc24b3d33c9123b = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xe5, 0x24, 0xdd, 0x56, 0x57, 0xd3, 0xfe, 0xf2, 0x7f, 0x43, 0x21, 0x03, 0x14, 0x22,
	0x84, 0x42, 0x0f, 0xb1, 0x5a, 0x10, 0xac, 0x15, 0x42, 0xa2, 0x63, 0x1a, 0xd3, 0x84, 0x54, 0x65,
	0x12, 0x07, 0x34, 0x31, 0x65, 0xad, 0x57, 0x42, 0xdb, 0x38, 0xb2, 0x9d, 0x8a, 0xa8, 0xaa, 0x84,
	0xf8, 0x0a, 0x5c, 0xe0, 0xca, 0x91, 0xaf, 0xc0, 0x27, 0x80, 0x2b, 0x5f, 0x81, 0x0f, 0xc1, 0x11,
	0xc5, 0x71, 0xca, 0xd4, 0x75, 0x13, 0xb9, 0xd9, 0xcf, 0x6b, 0xf7, 0xf7, 0x3c, 0xaf, 0xdf, 0x06,
	0x6e, 0x84, 0x51, 0x44, 0xd8, 0xc9, 0x90, 0xa4, 0x5e, 0xcc, 0xa8, 0xa0, 0x68, 0x23, 0x88, 0x63,
	0x4a, 0xc5, 0xa9, 0xd7, 0xa3, 0x49, 0x24, 0x08, 0xb3, 0x6c, 0x25, 0xe0, 0x98, 0xb0, 0x71, 0xc8,
	0x79, 0x48, 0x23, 0xcc, 0x09, 0x9b, 0x84, 0x3d, 0x92, 0x5f, 0xb1, 0x6e, 0x0c, 0x28, 0x1d, 0x8c,
	0x08, 0x0e, 0xe2, 0x10, 0x07, 0x51, 0x44, 0x45, 0x20, 0x42, 0x1a, 0x71, 0x55, 0xbd, 0xa5, 0xaa,
	0x72, 0x77, 0x9a, 0x9c, 0xe1, 0x7e, 0xc2, 0xe4, 0x01, 0x55, 0xdf, 0x5e, 0xac, 0x93, 0x71, 0x2c,
	0x94, 0x1b, 0xe7, 0x33, 0x80, 0xfa, 0x21, 0x49, 0x91, 0x09, 0x57, 0x63, 0x46, 0xfb, 0x49, 0x4f,
	0x98, 0xc0, 0xd6, 0xdc, 0xaa, 0x5f, 0x6c, 0x11, 0x82, 0x86, 0x48, 0x63, 0x62, 0x6a, 0x52, 0x96,
	0x6b, 0xb4, 0x0d, 0xab, 0x8c, 0x8c, 0x02, 0x41, 0x4e, 0xc2, 0xbe, 0xa9, 0xcb, 0xc2, 0x5a, 0x2e,
	0x1c, 0xf4, 0xd1, 0x26, 0xac, 0x4c, 0x82, 0x51, 0x42, 0x4c, 0xc3, 0x06, 0xae, 0xee, 0xe7, 0x1b,
	0xd4, 0x80, 0x2b, 0xe4, 0x5d, 0x1c, 0x32, 0x62, 0x56, 0x6c, 0xe0, 0xd6, 0x9a, 0xd7, 0xbd, 0xdc,
	0x96, 0x57, 0xd8, 0xf2, 0x9e, 0x29, 0xdb, 0xbe, 0x3a, 0xe8, 0x7c, 0x07, 0xd0, 0x38, 0x24, 0x29,
	0x2f, 0x69, 0xee, 0x26, 0x84, 0x73, 0x73, 0xdc, 0xd4, 0x6d, 0xdd, 0xad, 0xfa, 0xd5, 0xc2, 0x1d,
	0x47, 0x2d, 0xb8, 0x22, 0x1d, 0x71, 0xd3, 0xb0, 0x75, 0xb7, 0xd6, 0xbc, 0xed, 0x2d, 0x3c, 0x88,
	0x97, 0x31, 0xbd, 0x97, 0xf2, 0xcc, 0x5e, 0x24, 0x58, 0xea, 0xab, 0x0b, 0x56, 0x0b, 0xd6, 0xce,
	0xc9, 0xe8, 0x3f, 0xa8, 0x0f, 0x49, 0x6a, 0x02, 0x1b, 0xb8, 0x55, 0x3f, 0x5b, 0xfe, 0x8d, 0xae,
	0x9d, 0x8b, 0xde, 0xd6, 0x76, 0x40, 0xf3, 0x5b, 0x05, 0xae, 0x1d, 0x64, 0x93, 0x90, 0x35, 0x7b,
	0x92, 0xad, 0x7b, 0x8c, 0x04, 0x9c, 0xa0, 0xcd, 0x65, 0x78, 0x6b, 0xa9, 0xea, 0x3c, 0xf9, 0xf0,
	0xf3, 0xd7, 0x47, 0x6d, 0xc7, 0xc2, 0x58, 0xa9, 0x58, 0x0e, 0x17, 0x1e, 0x92, 0x14, 0x4f, 0x55,
	0x4b, 0x66, 0x78, 0x9a, 0x75, 0x61, 0x86, 0xa7, 0xf3, 0x26, 0xcc, 0xda, 0xea, 0x0d, 0xa6, 0x70,
	0xfd, 0x45, 0x32, 0x12, 0xe1, 0x1c, 0xbe, 0xb5, 0x34, 0xbb, 0xb5, 0x5c, 0x76, 0x1e, 0x4b, 0xfc,
	0x43, 0xab, 0x71, 0x11, 0xcf, 0xaf, 0xe2, 0xf3, 0x59, 0x1b, 0xd4, 0xd1, 0x5b, 0xa8, 0xef, 0x13,
	0x51, 0x2a, 0xef, 0x23, 0x09, 0x6c, 0xa0, 0xb2, 0x79, 0x11, 0x85, 0xc6, 0x3e, 0x11, 0xbc, 0x64,
	0xbe, 0x96, 0xc4, 0xdd, 0x47, 0xe5, 0xf3, 0x21, 0x0a, 0xf5, 0xa3, 0x4b, 0xc3, 0x5d, 0xbb, 0x30,
	0xea, 0x7b, 0xd9, 0x3f, 0xd0, 0x69, 0x4b, 0xde, 0x03, 0xa7, 0xf4, 0x73, 0x82, 0x3a, 0x3a, 0x83,
	0xc6, 0xd1, 0x15, 0x09, 0x2f, 0x43, 0x62, 0x89, 0xbc, 0xe7, 0xdc, 0xf9, 0x97, 0x88, 0x6d, 0x50,
	0xb7, 0x8c, 0x4f, 0xaf, 0xdf, 0x6b, 0x9d, 0x29, 0xfc, 0xbf, 0x47, 0xc7, 0x8b, 0xa8, 0xce, 0x7a,
	0x31, 0xd1, 0xdd, 0x8c, 0xf2, 0x1c, 0x74, 0xc1, 0xab, 0xbb, 0x83, 0x50, 0xbc, 0x49, 0xb2, 0x23,
	0x63, 0x3c, 0xff, 0xae, 0x15, 0x1f, 0x9e, 0x01, 0x2d, 0xb0, 0xbf, 0x01, 0xf8, 0xa2, 0xe9, 0xbb,
	0xdd, 0xce, 0x57, 0x6d, 0x75, 0x37, 0x97, 0x7e, 0x68, 0x5b, 0x4f, 0xf3, 0x0b, 0xc7, 0xf2, 0xe7,
	0x8e, 0x95, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xf2, 0x4f, 0xbd, 0x42, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InnerKeyClient is the client API for InnerKey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerKeyClient interface {
	// Increase counter value.
	Increase(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error)
	// Multiple increase counter values.
	MultiIncrease(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Keys, error)
	// Get counter value.
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error)
	// Multi get counter values.
	Gets(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Keys, error)
	// Set counter value.
	Set(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error)
	// Multi set counter values.
	Sets(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*empty.Empty, error)
}

type innerKeyClient struct {
	cc *grpc.ClientConn
}

func NewInnerKeyClient(cc *grpc.ClientConn) InnerKeyClient {
	return &innerKeyClient{cc}
}

func (c *innerKeyClient) Increase(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/Increase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerKeyClient) MultiIncrease(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Keys, error) {
	out := new(Keys)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/MultiIncrease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerKeyClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerKeyClient) Gets(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Keys, error) {
	out := new(Keys)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerKeyClient) Set(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerKeyClient) Sets(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerKey/Sets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerKeyServer is the server API for InnerKey service.
type InnerKeyServer interface {
	// Increase counter value.
	Increase(context.Context, *Key) (*Key, error)
	// Multiple increase counter values.
	MultiIncrease(context.Context, *Keys) (*Keys, error)
	// Get counter value.
	Get(context.Context, *Key) (*Key, error)
	// Multi get counter values.
	Gets(context.Context, *Keys) (*Keys, error)
	// Set counter value.
	Set(context.Context, *Key) (*empty.Empty, error)
	// Multi set counter values.
	Sets(context.Context, *Keys) (*empty.Empty, error)
}

// UnimplementedInnerKeyServer can be embedded to have forward compatible implementations.
type UnimplementedInnerKeyServer struct {
}

func (*UnimplementedInnerKeyServer) Increase(ctx context.Context, req *Key) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increase not implemented")
}
func (*UnimplementedInnerKeyServer) MultiIncrease(ctx context.Context, req *Keys) (*Keys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiIncrease not implemented")
}
func (*UnimplementedInnerKeyServer) Get(ctx context.Context, req *Key) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInnerKeyServer) Gets(ctx context.Context, req *Keys) (*Keys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}
func (*UnimplementedInnerKeyServer) Set(ctx context.Context, req *Key) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedInnerKeyServer) Sets(ctx context.Context, req *Keys) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sets not implemented")
}

func RegisterInnerKeyServer(s *grpc.Server, srv InnerKeyServer) {
	s.RegisterService(&_InnerKey_serviceDesc, srv)
}

func _InnerKey_Increase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).Increase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/Increase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).Increase(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerKey_MultiIncrease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).MultiIncrease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/MultiIncrease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).MultiIncrease(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerKey_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerKey_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).Gets(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerKey_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).Set(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerKey_Sets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerKeyServer).Sets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerKey/Sets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerKeyServer).Sets(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

var _InnerKey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.counter.InnerKey",
	HandlerType: (*InnerKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increase",
			Handler:    _InnerKey_Increase_Handler,
		},
		{
			MethodName: "MultiIncrease",
			Handler:    _InnerKey_MultiIncrease_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InnerKey_Get_Handler,
		},
		{
			MethodName: "Gets",
			Handler:    _InnerKey_Gets_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _InnerKey_Set_Handler,
		},
		{
			MethodName: "Sets",
			Handler:    _InnerKey_Sets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner_key.proto",
}
