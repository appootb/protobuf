// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_multi.proto

package counter

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
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
type MultiKeys struct {
	Product              *string          `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string          `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateIds            []string         `protobuf:"bytes,3,rep,name=relate_ids,json=relateIds" json:"relate_ids,omitempty"`
	Values               map[string]int64 `protobuf:"bytes,4,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MultiKeys) Reset()         { *m = MultiKeys{} }
func (m *MultiKeys) String() string { return proto.CompactTextString(m) }
func (*MultiKeys) ProtoMessage()    {}
func (*MultiKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_319822df4b7c645a, []int{0}
}

func (m *MultiKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiKeys.Unmarshal(m, b)
}
func (m *MultiKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiKeys.Marshal(b, m, deterministic)
}
func (m *MultiKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiKeys.Merge(m, src)
}
func (m *MultiKeys) XXX_Size() int {
	return xxx_messageInfo_MultiKeys.Size(m)
}
func (m *MultiKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MultiKeys proto.InternalMessageInfo

func (m *MultiKeys) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *MultiKeys) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *MultiKeys) GetRelateIds() []string {
	if m != nil {
		return m.RelateIds
	}
	return nil
}

func (m *MultiKeys) GetValues() map[string]int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

// Mix key.
type MixKeys struct {
	Product              *string  `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Types                []string `protobuf:"bytes,2,rep,name=types" json:"types,omitempty"`
	RelateIds            []string `protobuf:"bytes,3,rep,name=relate_ids,json=relateIds" json:"relate_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MixKeys) Reset()         { *m = MixKeys{} }
func (m *MixKeys) String() string { return proto.CompactTextString(m) }
func (*MixKeys) ProtoMessage()    {}
func (*MixKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_319822df4b7c645a, []int{1}
}

func (m *MixKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MixKeys.Unmarshal(m, b)
}
func (m *MixKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MixKeys.Marshal(b, m, deterministic)
}
func (m *MixKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MixKeys.Merge(m, src)
}
func (m *MixKeys) XXX_Size() int {
	return xxx_messageInfo_MixKeys.Size(m)
}
func (m *MixKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MixKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MixKeys proto.InternalMessageInfo

func (m *MixKeys) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *MixKeys) GetTypes() []string {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *MixKeys) GetRelateIds() []string {
	if m != nil {
		return m.RelateIds
	}
	return nil
}

// Mix value.
type MixValue struct {
	Values               map[string]int64 `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MixValue) Reset()         { *m = MixValue{} }
func (m *MixValue) String() string { return proto.CompactTextString(m) }
func (*MixValue) ProtoMessage()    {}
func (*MixValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_319822df4b7c645a, []int{2}
}

func (m *MixValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MixValue.Unmarshal(m, b)
}
func (m *MixValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MixValue.Marshal(b, m, deterministic)
}
func (m *MixValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MixValue.Merge(m, src)
}
func (m *MixValue) XXX_Size() int {
	return xxx_messageInfo_MixValue.Size(m)
}
func (m *MixValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MixValue.DiscardUnknown(m)
}

var xxx_messageInfo_MixValue proto.InternalMessageInfo

func (m *MixValue) GetValues() map[string]int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

// Mix values.
type MixValues struct {
	Types                map[string]*MixValue `protobuf:"bytes,1,rep,name=types" json:"types,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MixValues) Reset()         { *m = MixValues{} }
func (m *MixValues) String() string { return proto.CompactTextString(m) }
func (*MixValues) ProtoMessage()    {}
func (*MixValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_319822df4b7c645a, []int{3}
}

func (m *MixValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MixValues.Unmarshal(m, b)
}
func (m *MixValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MixValues.Marshal(b, m, deterministic)
}
func (m *MixValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MixValues.Merge(m, src)
}
func (m *MixValues) XXX_Size() int {
	return xxx_messageInfo_MixValues.Size(m)
}
func (m *MixValues) XXX_DiscardUnknown() {
	xxx_messageInfo_MixValues.DiscardUnknown(m)
}

var xxx_messageInfo_MixValues proto.InternalMessageInfo

func (m *MixValues) GetTypes() map[string]*MixValue {
	if m != nil {
		return m.Types
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiKeys)(nil), "appootb.counter.MultiKeys")
	proto.RegisterMapType((map[string]int64)(nil), "appootb.counter.MultiKeys.ValuesEntry")
	proto.RegisterType((*MixKeys)(nil), "appootb.counter.MixKeys")
	proto.RegisterType((*MixValue)(nil), "appootb.counter.MixValue")
	proto.RegisterMapType((map[string]int64)(nil), "appootb.counter.MixValue.ValuesEntry")
	proto.RegisterType((*MixValues)(nil), "appootb.counter.MixValues")
	proto.RegisterMapType((map[string]*MixValue)(nil), "appootb.counter.MixValues.TypesEntry")
}

func init() { proto.RegisterFile("inner_multi.proto", fileDescriptor_319822df4b7c645a) }

var fileDescriptor_319822df4b7c645a = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x66, 0x66, 0x53, 0x63, 0x5e, 0x0e, 0xd5, 0xb1, 0xca, 0xba, 0x2a, 0x2c, 0x0b, 0x2d, 0xa1,
	0x87, 0x1d, 0x89, 0x0a, 0x9a, 0xa2, 0x60, 0x4a, 0xd1, 0x22, 0x81, 0x90, 0x8a, 0x88, 0x14, 0xcb,
	0x36, 0x19, 0xe3, 0xd0, 0x64, 0x67, 0xd9, 0x99, 0x2d, 0x59, 0x42, 0x40, 0xc4, 0x7f, 0xe0, 0xc5,
	0x8b, 0x17, 0x8f, 0xfe, 0x14, 0xaf, 0xfd, 0x0b, 0xfe, 0x04, 0x0f, 0x1e, 0x65, 0x67, 0x76, 0x93,
	0x58, 0x1b, 0xd3, 0x43, 0x6f, 0xfb, 0xbe, 0xf7, 0xde, 0x7c, 0xdf, 0xb7, 0x6f, 0xde, 0xc0, 0x55,
	0x1e, 0x86, 0x2c, 0x3e, 0x18, 0x26, 0x03, 0xc5, 0xfd, 0x28, 0x16, 0x4a, 0x90, 0xd5, 0x20, 0x8a,
	0x84, 0x50, 0x87, 0x7e, 0x57, 0x24, 0xa1, 0x62, 0xb1, 0xe3, 0xe6, 0x00, 0x8d, 0x58, 0x3c, 0xe4,
	0x52, 0x72, 0x11, 0x52, 0xc9, 0xe2, 0x63, 0xde, 0x65, 0xa6, 0xc5, 0xb9, 0xdd, 0x17, 0xa2, 0x3f,
	0x60, 0x34, 0x88, 0x38, 0x0d, 0xc2, 0x50, 0xa8, 0x40, 0x71, 0x11, 0xca, 0x3c, 0x7b, 0x2b, 0xcf,
	0xea, 0xe8, 0x30, 0x79, 0x47, 0xd9, 0x30, 0x52, 0xa9, 0x49, 0x7a, 0x27, 0x08, 0x2a, 0xad, 0x8c,
	0xfd, 0x05, 0x4b, 0x25, 0xb1, 0xa1, 0x1c, 0xc5, 0xa2, 0x97, 0x74, 0x95, 0x8d, 0x5c, 0x5c, 0xab,
	0x74, 0x8a, 0x90, 0x10, 0x28, 0xa9, 0x34, 0x62, 0x36, 0xd6, 0xb0, 0xfe, 0x26, 0x77, 0x00, 0x62,
	0x36, 0x08, 0x14, 0x3b, 0xe0, 0x3d, 0x69, 0x5b, 0xae, 0x55, 0xab, 0x74, 0x2a, 0x06, 0xd9, 0xed,
	0x49, 0xf2, 0x04, 0x2e, 0x1d, 0x07, 0x83, 0x84, 0x49, 0xbb, 0xe4, 0x5a, 0xb5, 0x6a, 0x7d, 0xc3,
	0x3f, 0xe5, 0xcc, 0x9f, 0x12, 0xfb, 0xaf, 0x74, 0xe1, 0x4e, 0xa8, 0xe2, 0xb4, 0x93, 0x77, 0x39,
	0x8f, 0xa0, 0x3a, 0x07, 0x93, 0x2b, 0x60, 0x1d, 0xb1, 0xd4, 0x46, 0x2e, 0xaa, 0x55, 0x3a, 0xd9,
	0x27, 0x59, 0x83, 0x15, 0x5d, 0x6a, 0x63, 0x17, 0xd5, 0xac, 0x8e, 0x09, 0x1a, 0xf8, 0x21, 0xf2,
	0x5e, 0x43, 0xb9, 0xc5, 0x47, 0x4b, 0x2c, 0xad, 0xc1, 0x4a, 0x66, 0x43, 0xda, 0x58, 0x2b, 0x37,
	0xc1, 0x12, 0x53, 0xde, 0x27, 0x04, 0x97, 0x5b, 0x7c, 0xa4, 0x85, 0x91, 0xc7, 0x53, 0x87, 0x48,
	0x3b, 0x5c, 0xff, 0xd7, 0x61, 0x5e, 0x7a, 0xd1, 0x06, 0xbf, 0x66, 0x63, 0xcb, 0xcf, 0x96, 0x64,
	0xab, 0x70, 0xb2, 0x4c, 0x86, 0xf4, 0x5f, 0x66, 0x75, 0x46, 0x86, 0xe9, 0x71, 0xf6, 0x00, 0x66,
	0xe0, 0x19, 0x22, 0xe8, 0xbc, 0x88, 0x6a, 0xfd, 0xe6, 0xc2, 0xc3, 0xe7, 0xf4, 0xd5, 0x7f, 0x61,
	0x80, 0xdd, 0xec, 0x6a, 0xeb, 0x11, 0x93, 0x14, 0xac, 0x67, 0x4c, 0x11, 0x67, 0xf1, 0x0d, 0x70,
	0xfe, 0x93, 0xf3, 0x1a, 0x1f, 0x4f, 0x7e, 0x7e, 0xc6, 0xf7, 0x49, 0x9d, 0xe6, 0x39, 0xaa, 0xd7,
	0x86, 0xea, 0xb5, 0xa1, 0xe3, 0x7c, 0xa0, 0x13, 0x3a, 0xce, 0x1c, 0x4d, 0xe8, 0x78, 0x36, 0xc0,
	0x09, 0x39, 0x02, 0x6b, 0x6f, 0x09, 0xf5, 0x0d, 0xdf, 0x6c, 0x88, 0x5f, 0x6c, 0x88, 0xbf, 0x93,
	0x6d, 0x88, 0x77, 0x57, 0xd3, 0x6e, 0x7a, 0xeb, 0xe7, 0xa2, 0x6d, 0xa0, 0x4d, 0x32, 0x02, 0xab,
	0xc5, 0x47, 0xc4, 0x3e, 0xeb, 0x1f, 0x2d, 0x72, 0x59, 0x8c, 0xc6, 0xdb, 0xd2, 0x74, 0x0f, 0xc8,
	0xbd, 0xf3, 0xd0, 0xc9, 0xbf, 0x6d, 0x3a, 0xa5, 0x2f, 0x6f, 0x3f, 0xe0, 0xe6, 0x04, 0xae, 0x75,
	0xc5, 0xf0, 0x34, 0x47, 0x73, 0x75, 0x36, 0x8a, 0x76, 0x66, 0xf1, 0x39, 0x6a, 0xa3, 0x37, 0x1b,
	0x7d, 0xae, 0xde, 0x27, 0x59, 0xd1, 0x90, 0x4e, 0x5f, 0x98, 0xe2, 0x89, 0xe8, 0x8b, 0x42, 0xc4,
	0x6f, 0x84, 0xbe, 0x61, 0x6b, 0xbb, 0xdd, 0xfc, 0x8e, 0xcb, 0xdb, 0x06, 0xfa, 0x81, 0xaf, 0x3f,
	0x35, 0x0d, 0xfb, 0xfa, 0xb8, 0xfd, 0x1c, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x12, 0xe2,
	0x68, 0xce, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InnerMultiClient is the client API for InnerMulti service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerMultiClient interface {
	// Multi get counter values.
	Get(ctx context.Context, in *MultiKeys, opts ...grpc.CallOption) (*MultiKeys, error)
	// Multi set counter values.
	Set(ctx context.Context, in *MultiKeys, opts ...grpc.CallOption) (*empty.Empty, error)
	// Mix gets counter values.
	Mix(ctx context.Context, in *MixKeys, opts ...grpc.CallOption) (*MixValues, error)
}

type innerMultiClient struct {
	cc *grpc.ClientConn
}

func NewInnerMultiClient(cc *grpc.ClientConn) InnerMultiClient {
	return &innerMultiClient{cc}
}

func (c *innerMultiClient) Get(ctx context.Context, in *MultiKeys, opts ...grpc.CallOption) (*MultiKeys, error) {
	out := new(MultiKeys)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerMulti/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerMultiClient) Set(ctx context.Context, in *MultiKeys, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerMulti/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerMultiClient) Mix(ctx context.Context, in *MixKeys, opts ...grpc.CallOption) (*MixValues, error) {
	out := new(MixValues)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerMulti/Mix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerMultiServer is the server API for InnerMulti service.
type InnerMultiServer interface {
	// Multi get counter values.
	Get(context.Context, *MultiKeys) (*MultiKeys, error)
	// Multi set counter values.
	Set(context.Context, *MultiKeys) (*empty.Empty, error)
	// Mix gets counter values.
	Mix(context.Context, *MixKeys) (*MixValues, error)
}

// UnimplementedInnerMultiServer can be embedded to have forward compatible implementations.
type UnimplementedInnerMultiServer struct {
}

func (*UnimplementedInnerMultiServer) Get(ctx context.Context, req *MultiKeys) (*MultiKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInnerMultiServer) Set(ctx context.Context, req *MultiKeys) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedInnerMultiServer) Mix(ctx context.Context, req *MixKeys) (*MixValues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mix not implemented")
}

func RegisterInnerMultiServer(s *grpc.Server, srv InnerMultiServer) {
	s.RegisterService(&_InnerMulti_serviceDesc, srv)
}

func _InnerMulti_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerMultiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerMulti/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerMultiServer).Get(ctx, req.(*MultiKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerMulti_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerMultiServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerMulti/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerMultiServer).Set(ctx, req.(*MultiKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerMulti_Mix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MixKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerMultiServer).Mix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerMulti/Mix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerMultiServer).Mix(ctx, req.(*MixKeys))
	}
	return interceptor(ctx, in, info, handler)
}

var _InnerMulti_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.counter.InnerMulti",
	HandlerType: (*InnerMultiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InnerMulti_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _InnerMulti_Set_Handler,
		},
		{
			MethodName: "Mix",
			Handler:    _InnerMulti_Mix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner_multi.proto",
}