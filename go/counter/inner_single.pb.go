// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_single.proto

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

// Single key.
type SingleKey struct {
	Product              *string            `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string            `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateId             *string            `protobuf:"bytes,3,req,name=relate_id,json=relateId" json:"relate_id,omitempty"`
	Value                *int64             `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	Expire               *duration.Duration `protobuf:"bytes,5,opt,name=expire" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SingleKey) Reset()         { *m = SingleKey{} }
func (m *SingleKey) String() string { return proto.CompactTextString(m) }
func (*SingleKey) ProtoMessage()    {}
func (*SingleKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4acfde5c641b5c3, []int{0}
}

func (m *SingleKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleKey.Unmarshal(m, b)
}
func (m *SingleKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleKey.Marshal(b, m, deterministic)
}
func (m *SingleKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleKey.Merge(m, src)
}
func (m *SingleKey) XXX_Size() int {
	return xxx_messageInfo_SingleKey.Size(m)
}
func (m *SingleKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleKey.DiscardUnknown(m)
}

var xxx_messageInfo_SingleKey proto.InternalMessageInfo

func (m *SingleKey) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *SingleKey) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *SingleKey) GetRelateId() string {
	if m != nil && m.RelateId != nil {
		return *m.RelateId
	}
	return ""
}

func (m *SingleKey) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *SingleKey) GetExpire() *duration.Duration {
	if m != nil {
		return m.Expire
	}
	return nil
}

func init() {
	proto.RegisterType((*SingleKey)(nil), "appootb.counter.SingleKey")
}

func init() { proto.RegisterFile("inner_single.proto", fileDescriptor_a4acfde5c641b5c3) }

var fileDescriptor_a4acfde5c641b5c3 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x49, 0x3b, 0xeb, 0xee, 0x64, 0x0f, 0x4a, 0xfc, 0x43, 0xed, 0x8a, 0x94, 0x1e, 0x64,
	0xf0, 0xd0, 0xe0, 0x20, 0x08, 0x2b, 0x1e, 0x76, 0x56, 0xd1, 0xc1, 0xcb, 0x30, 0x7b, 0x93, 0xc5,
	0xa5, 0xdb, 0xbe, 0xd6, 0x40, 0x9b, 0x84, 0x34, 0x5d, 0xa6, 0x0c, 0x23, 0x22, 0x78, 0x17, 0xbc,
	0x78, 0x13, 0x3c, 0xfa, 0x51, 0xbc, 0xfa, 0x15, 0xfc, 0x10, 0x1e, 0xa5, 0x49, 0x3a, 0xc8, 0x08,
	0x2e, 0xcc, 0xad, 0x79, 0x9e, 0x37, 0x7d, 0x7f, 0xcf, 0xfb, 0x06, 0x13, 0xc6, 0x39, 0xa8, 0xb3,
	0x9a, 0xf1, 0xa2, 0x84, 0x44, 0x2a, 0xa1, 0x05, 0xb9, 0x9a, 0x4a, 0x29, 0x84, 0x3e, 0x4f, 0x32,
	0xd1, 0x70, 0x0d, 0x2a, 0x8c, 0x9c, 0x40, 0x25, 0xa8, 0x8a, 0xd5, 0x35, 0x13, 0x9c, 0xd6, 0xa0,
	0x2e, 0x58, 0xe6, 0xae, 0x84, 0x77, 0x0a, 0x21, 0x8a, 0x12, 0x68, 0x2a, 0x19, 0x4d, 0x39, 0x17,
	0x3a, 0xd5, 0x4c, 0xf0, 0xda, 0xb9, 0x77, 0x9d, 0x6b, 0x4e, 0xe7, 0xcd, 0x1b, 0x9a, 0x37, 0xca,
	0x14, 0x38, 0xff, 0x60, 0xd3, 0x87, 0x4a, 0xea, 0xd6, 0x9a, 0xf1, 0x57, 0x84, 0x87, 0x27, 0x06,
	0xef, 0x25, 0xb4, 0x24, 0xc0, 0xbb, 0x52, 0x89, 0xbc, 0xc9, 0x74, 0x80, 0x22, 0x6f, 0x34, 0x9c,
	0xf7, 0x47, 0x42, 0xf0, 0x40, 0xb7, 0x12, 0x02, 0xcf, 0xc8, 0xe6, 0x9b, 0x1c, 0xe0, 0xa1, 0x82,
	0x32, 0xd5, 0x70, 0xc6, 0xf2, 0xc0, 0x37, 0xc6, 0x9e, 0x15, 0xa6, 0x39, 0xb9, 0x81, 0x77, 0x2e,
	0xd2, 0xb2, 0x81, 0x60, 0x10, 0xa1, 0x91, 0x3f, 0xb7, 0x07, 0xf2, 0x00, 0x5f, 0x81, 0x85, 0x64,
	0x0a, 0x82, 0x9d, 0x08, 0x8d, 0xf6, 0xc7, 0xb7, 0x13, 0x0b, 0x97, 0xf4, 0x70, 0xc9, 0x53, 0x07,
	0x3f, 0x77, 0x85, 0xe3, 0x4f, 0x3e, 0xde, 0x9f, 0x76, 0x63, 0xb4, 0x98, 0xe4, 0x23, 0xc2, 0x7b,
	0x53, 0x9e, 0x29, 0x48, 0x6b, 0x20, 0x61, 0xb2, 0x31, 0xcd, 0x64, 0x1d, 0x26, 0xfc, 0x8f, 0x17,
	0x1f, 0x7d, 0xf8, 0xf9, 0xeb, 0xb3, 0xf7, 0x38, 0x1c, 0x53, 0xe7, 0x51, 0xb3, 0x29, 0x6a, 0x37,
	0x45, 0x97, 0x2e, 0xf6, 0x8a, 0x2e, 0xbb, 0xa4, 0x2b, 0xba, 0x5c, 0x07, 0x5d, 0x1d, 0xba, 0x28,
	0x2d, 0xf6, 0x9f, 0x83, 0xde, 0x9a, 0xe0, 0xd0, 0x10, 0x3c, 0x24, 0x5b, 0x10, 0x90, 0x05, 0xf6,
	0x4f, 0x2e, 0x69, 0x7d, 0xeb, 0x9f, 0xc1, 0x3e, 0xeb, 0xb6, 0x1e, 0x3f, 0x31, 0x6d, 0x1f, 0xc5,
	0xdb, 0x04, 0x47, 0xf7, 0xc3, 0xc1, 0x97, 0xd7, 0xef, 0xbd, 0xc9, 0x3b, 0x7c, 0x3d, 0x13, 0xd5,
	0x66, 0xf7, 0xc9, 0xb5, 0xbf, 0xd6, 0x34, 0xeb, 0xda, 0xbe, 0x40, 0x33, 0xf4, 0xea, 0x5e, 0xc1,
	0xf4, 0xdb, 0xa6, 0xab, 0xaa, 0xe8, 0xfa, 0xa5, 0xf7, 0x4f, 0xb1, 0x10, 0x3d, 0xc7, 0x6f, 0x84,
	0xbe, 0x79, 0xfe, 0xf1, 0x6c, 0xf2, 0xdd, 0xdb, 0x3d, 0xb6, 0xd2, 0x0f, 0xef, 0xe6, 0x91, 0xbd,
	0x70, 0x6a, 0x7e, 0x77, 0xea, 0xf4, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0xfe, 0x65, 0xe6,
	0x57, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InnerSingleClient is the client API for InnerSingle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerSingleClient interface {
	// Increase counter value.
	Increase(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*SingleKey, error)
	// Get counter value.
	Get(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*SingleKey, error)
	// Set counter value.
	Set(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type innerSingleClient struct {
	cc *grpc.ClientConn
}

func NewInnerSingleClient(cc *grpc.ClientConn) InnerSingleClient {
	return &innerSingleClient{cc}
}

func (c *innerSingleClient) Increase(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*SingleKey, error) {
	out := new(SingleKey)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerSingle/Increase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerSingleClient) Get(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*SingleKey, error) {
	out := new(SingleKey)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerSingle/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerSingleClient) Set(ctx context.Context, in *SingleKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerSingle/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerSingleServer is the server API for InnerSingle service.
type InnerSingleServer interface {
	// Increase counter value.
	Increase(context.Context, *SingleKey) (*SingleKey, error)
	// Get counter value.
	Get(context.Context, *SingleKey) (*SingleKey, error)
	// Set counter value.
	Set(context.Context, *SingleKey) (*empty.Empty, error)
}

// UnimplementedInnerSingleServer can be embedded to have forward compatible implementations.
type UnimplementedInnerSingleServer struct {
}

func (*UnimplementedInnerSingleServer) Increase(ctx context.Context, req *SingleKey) (*SingleKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increase not implemented")
}
func (*UnimplementedInnerSingleServer) Get(ctx context.Context, req *SingleKey) (*SingleKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInnerSingleServer) Set(ctx context.Context, req *SingleKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterInnerSingleServer(s *grpc.Server, srv InnerSingleServer) {
	s.RegisterService(&_InnerSingle_serviceDesc, srv)
}

func _InnerSingle_Increase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerSingleServer).Increase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerSingle/Increase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerSingleServer).Increase(ctx, req.(*SingleKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerSingle_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerSingleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerSingle/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerSingleServer).Get(ctx, req.(*SingleKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerSingle_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerSingleServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerSingle/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerSingleServer).Set(ctx, req.(*SingleKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _InnerSingle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.counter.InnerSingle",
	HandlerType: (*InnerSingleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increase",
			Handler:    _InnerSingle_Increase_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InnerSingle_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _InnerSingle_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner_single.proto",
}
