// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_hash_fields.proto

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

// Hash fields key.
type HashFieldsKey struct {
	Product              *string          `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string          `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateId             *string          `protobuf:"bytes,3,req,name=relate_id,json=relateId" json:"relate_id,omitempty"`
	Fields               []string         `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	Values               map[string]int64 `protobuf:"bytes,5,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HashFieldsKey) Reset()         { *m = HashFieldsKey{} }
func (m *HashFieldsKey) String() string { return proto.CompactTextString(m) }
func (*HashFieldsKey) ProtoMessage()    {}
func (*HashFieldsKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_af168d4c90cbcd5c, []int{0}
}

func (m *HashFieldsKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashFieldsKey.Unmarshal(m, b)
}
func (m *HashFieldsKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashFieldsKey.Marshal(b, m, deterministic)
}
func (m *HashFieldsKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashFieldsKey.Merge(m, src)
}
func (m *HashFieldsKey) XXX_Size() int {
	return xxx_messageInfo_HashFieldsKey.Size(m)
}
func (m *HashFieldsKey) XXX_DiscardUnknown() {
	xxx_messageInfo_HashFieldsKey.DiscardUnknown(m)
}

var xxx_messageInfo_HashFieldsKey proto.InternalMessageInfo

func (m *HashFieldsKey) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *HashFieldsKey) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *HashFieldsKey) GetRelateId() string {
	if m != nil && m.RelateId != nil {
		return *m.RelateId
	}
	return ""
}

func (m *HashFieldsKey) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *HashFieldsKey) GetValues() map[string]int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

// Hash keys.
type HashKeys struct {
	Product              *string  `protobuf:"bytes,1,req,name=product" json:"product,omitempty"`
	Type                 *string  `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	RelateIds            []string `protobuf:"bytes,3,rep,name=relate_ids,json=relateIds" json:"relate_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashKeys) Reset()         { *m = HashKeys{} }
func (m *HashKeys) String() string { return proto.CompactTextString(m) }
func (*HashKeys) ProtoMessage()    {}
func (*HashKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_af168d4c90cbcd5c, []int{1}
}

func (m *HashKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashKeys.Unmarshal(m, b)
}
func (m *HashKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashKeys.Marshal(b, m, deterministic)
}
func (m *HashKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashKeys.Merge(m, src)
}
func (m *HashKeys) XXX_Size() int {
	return xxx_messageInfo_HashKeys.Size(m)
}
func (m *HashKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_HashKeys.DiscardUnknown(m)
}

var xxx_messageInfo_HashKeys proto.InternalMessageInfo

func (m *HashKeys) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *HashKeys) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *HashKeys) GetRelateIds() []string {
	if m != nil {
		return m.RelateIds
	}
	return nil
}

// Hash all fields values.
type HashValue struct {
	FieldValues          map[string]int64 `protobuf:"bytes,1,rep,name=field_values,json=fieldValues" json:"field_values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HashValue) Reset()         { *m = HashValue{} }
func (m *HashValue) String() string { return proto.CompactTextString(m) }
func (*HashValue) ProtoMessage()    {}
func (*HashValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_af168d4c90cbcd5c, []int{2}
}

func (m *HashValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashValue.Unmarshal(m, b)
}
func (m *HashValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashValue.Marshal(b, m, deterministic)
}
func (m *HashValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashValue.Merge(m, src)
}
func (m *HashValue) XXX_Size() int {
	return xxx_messageInfo_HashValue.Size(m)
}
func (m *HashValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HashValue.DiscardUnknown(m)
}

var xxx_messageInfo_HashValue proto.InternalMessageInfo

func (m *HashValue) GetFieldValues() map[string]int64 {
	if m != nil {
		return m.FieldValues
	}
	return nil
}

// Hash counter key values.
type HashValues struct {
	Values               map[string]*HashValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HashValues) Reset()         { *m = HashValues{} }
func (m *HashValues) String() string { return proto.CompactTextString(m) }
func (*HashValues) ProtoMessage()    {}
func (*HashValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_af168d4c90cbcd5c, []int{3}
}

func (m *HashValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashValues.Unmarshal(m, b)
}
func (m *HashValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashValues.Marshal(b, m, deterministic)
}
func (m *HashValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashValues.Merge(m, src)
}
func (m *HashValues) XXX_Size() int {
	return xxx_messageInfo_HashValues.Size(m)
}
func (m *HashValues) XXX_DiscardUnknown() {
	xxx_messageInfo_HashValues.DiscardUnknown(m)
}

var xxx_messageInfo_HashValues proto.InternalMessageInfo

func (m *HashValues) GetValues() map[string]*HashValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*HashFieldsKey)(nil), "appootb.counter.HashFieldsKey")
	proto.RegisterMapType((map[string]int64)(nil), "appootb.counter.HashFieldsKey.ValuesEntry")
	proto.RegisterType((*HashKeys)(nil), "appootb.counter.HashKeys")
	proto.RegisterType((*HashValue)(nil), "appootb.counter.HashValue")
	proto.RegisterMapType((map[string]int64)(nil), "appootb.counter.HashValue.FieldValuesEntry")
	proto.RegisterType((*HashValues)(nil), "appootb.counter.HashValues")
	proto.RegisterMapType((map[string]*HashValue)(nil), "appootb.counter.HashValues.ValuesEntry")
}

func init() { proto.RegisterFile("inner_hash_fields.proto", fileDescriptor_af168d4c90cbcd5c) }

var fileDescriptor_af168d4c90cbcd5c = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6b, 0x13, 0x4f,
	0x18, 0x67, 0x76, 0xd3, 0x36, 0xfb, 0xe4, 0xff, 0xa7, 0x65, 0xac, 0x75, 0xdd, 0xa8, 0x2c, 0x7b,
	0xd0, 0x50, 0x61, 0x47, 0xe3, 0x41, 0x5b, 0xc4, 0x97, 0x2d, 0xb5, 0x09, 0x05, 0x09, 0x11, 0x15,
	0xa4, 0x18, 0xb6, 0xc9, 0x24, 0x59, 0x4c, 0x76, 0x96, 0x9d, 0x49, 0x61, 0x09, 0x01, 0xed, 0x07,
	0xd0, 0x83, 0x17, 0xc1, 0x9b, 0x78, 0xf2, 0x1b, 0xf8, 0x15, 0xbc, 0xfa, 0x15, 0xfc, 0x02, 0xde,
	0x3c, 0xca, 0xce, 0xee, 0xa6, 0x31, 0xa4, 0xa9, 0xb5, 0xb7, 0x79, 0xde, 0x7f, 0xbf, 0x79, 0x5e,
	0xe0, 0x82, 0xe7, 0xfb, 0x34, 0x6c, 0x74, 0x5d, 0xde, 0x6d, 0xb4, 0x3d, 0xda, 0x6b, 0x71, 0x3b,
	0x08, 0x99, 0x60, 0x78, 0xd9, 0x0d, 0x02, 0xc6, 0xc4, 0xbe, 0xdd, 0x64, 0x03, 0x5f, 0xd0, 0xd0,
	0x30, 0x53, 0x05, 0x09, 0x68, 0xd8, 0xf7, 0x38, 0xf7, 0x98, 0x4f, 0x38, 0x0d, 0x0f, 0xbc, 0x26,
	0x4d, 0x42, 0x8c, 0x4b, 0x1d, 0xc6, 0x3a, 0x3d, 0x4a, 0xdc, 0xc0, 0x23, 0xae, 0xef, 0x33, 0xe1,
	0x0a, 0x8f, 0xf9, 0x69, 0x42, 0xa3, 0x98, 0x5a, 0xa5, 0xb4, 0x3f, 0x68, 0x13, 0xda, 0x0f, 0x44,
	0x94, 0x18, 0xad, 0x9f, 0x08, 0xfe, 0xaf, 0xb8, 0xbc, 0xfb, 0x48, 0x42, 0xd8, 0xa5, 0x11, 0xd6,
	0x61, 0x29, 0x08, 0x59, 0x6b, 0xd0, 0x14, 0x3a, 0x32, 0x95, 0x92, 0x56, 0xcf, 0x44, 0x8c, 0x21,
	0x27, 0xa2, 0x80, 0xea, 0x8a, 0x54, 0xcb, 0x37, 0x2e, 0x82, 0x16, 0xd2, 0x9e, 0x2b, 0x68, 0xc3,
	0x6b, 0xe9, 0xaa, 0x34, 0xe4, 0x13, 0x45, 0xb5, 0x85, 0xd7, 0x60, 0x31, 0xa1, 0xa6, 0xe7, 0x4c,
	0xb5, 0xa4, 0xd5, 0x53, 0x09, 0x3b, 0xb0, 0x78, 0xe0, 0xf6, 0x06, 0x94, 0xeb, 0x0b, 0xa6, 0x5a,
	0x2a, 0x94, 0xd7, 0xed, 0x29, 0xce, 0xf6, 0x1f, 0x90, 0xec, 0x67, 0xd2, 0x79, 0xdb, 0x17, 0x61,
	0x54, 0x4f, 0x23, 0x8d, 0x0d, 0x28, 0x4c, 0xa8, 0xf1, 0x0a, 0xa8, 0xaf, 0x68, 0xa4, 0x23, 0x13,
	0x95, 0xb4, 0x7a, 0xfc, 0xc4, 0xab, 0xb0, 0x20, 0x5d, 0x75, 0xc5, 0x44, 0x25, 0xb5, 0x9e, 0x08,
	0x9b, 0xca, 0x1d, 0x64, 0x3d, 0x87, 0x7c, 0x9c, 0x7f, 0x97, 0x46, 0xfc, 0x94, 0x6c, 0x2f, 0x03,
	0x8c, 0xd9, 0x72, 0x5d, 0x95, 0xa4, 0xb4, 0x8c, 0x2e, 0xb7, 0x3e, 0x22, 0xd0, 0xe2, 0xcc, 0x12,
	0x18, 0x7e, 0x0c, 0xff, 0x49, 0xbe, 0x8d, 0x94, 0x2b, 0x92, 0x5c, 0xaf, 0xcf, 0xe4, 0x2a, 0x23,
	0x6c, 0xc9, 0x78, 0x92, 0x6c, 0xa1, 0x7d, 0xa4, 0x31, 0xee, 0xc1, 0xca, 0xb4, 0xc3, 0xa9, 0x68,
	0x7f, 0x46, 0x00, 0xe3, 0x5a, 0x1c, 0xdf, 0x1f, 0x37, 0x21, 0x01, 0x76, 0xed, 0x78, 0x60, 0x7c,
	0x66, 0x07, 0x9e, 0x9e, 0xd4, 0x81, 0x1b, 0x93, 0x50, 0x0a, 0x65, 0xe3, 0xf8, 0x02, 0x13, 0x30,
	0xcb, 0x5f, 0x73, 0xb0, 0x5c, 0x8d, 0x77, 0xe3, 0x68, 0x06, 0xf0, 0x5b, 0x04, 0xf9, 0xaa, 0xdf,
	0x0c, 0xa9, 0xcb, 0x29, 0xbe, 0x32, 0x7f, 0x5a, 0x8c, 0x13, 0xec, 0x96, 0x73, 0xf8, 0xfd, 0xc7,
	0x7b, 0xe5, 0xae, 0x71, 0x9b, 0xa4, 0x76, 0x22, 0x57, 0x91, 0x4c, 0xac, 0x22, 0x19, 0xa6, 0x53,
	0x30, 0x22, 0xc3, 0xb8, 0xf1, 0x23, 0x32, 0x1c, 0xf7, 0x7d, 0xb4, 0x89, 0xd6, 0xf1, 0x3b, 0x04,
	0xea, 0x0e, 0x15, 0x67, 0xc6, 0x52, 0x91, 0x58, 0x1c, 0xfc, 0xe0, 0x1f, 0xb1, 0x90, 0x61, 0xe2,
	0x34, 0xc2, 0x6f, 0x10, 0xa8, 0x4f, 0xfe, 0x02, 0xd1, 0x9a, 0x9d, 0x9c, 0x03, 0x3b, 0x3b, 0x07,
	0xf6, 0x76, 0x7c, 0x0e, 0xb2, 0x5f, 0xb1, 0xce, 0xf2, 0x2b, 0x11, 0xe4, 0x76, 0xa8, 0xe0, 0xf8,
	0xe2, 0x4c, 0x0c, 0xf1, 0xbe, 0x19, 0xc5, 0x39, 0x53, 0x66, 0x6d, 0x48, 0x0c, 0xb7, 0xf0, 0xcd,
	0x19, 0x18, 0xe6, 0x15, 0xe7, 0x23, 0x23, 0xf7, 0xe1, 0xe5, 0x6b, 0xc5, 0x39, 0x44, 0x70, 0xae,
	0xc9, 0xfa, 0xd3, 0x35, 0x9c, 0xd5, 0xa9, 0x81, 0xaa, 0xc5, 0xdc, 0x2b, 0xa8, 0x86, 0x5e, 0x5c,
	0xed, 0x78, 0xa2, 0x3b, 0x88, 0x3d, 0xfb, 0x64, 0x7c, 0x67, 0xb3, 0x43, 0xd9, 0x61, 0x19, 0x90,
	0x5f, 0x08, 0x7d, 0x52, 0xd4, 0xad, 0x9a, 0xf3, 0x45, 0x59, 0xda, 0x4a, 0x54, 0xdf, 0x94, 0xf3,
	0x0f, 0x93, 0x80, 0x3d, 0x99, 0x6e, 0x2f, 0xd5, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x94,
	0x97, 0x92, 0xda, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InnerHashFieldsClient is the client API for InnerHashFields service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerHashFieldsClient interface {
	// Increase hash fields values.
	Increase(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*HashFieldsKey, error)
	// Get hash fields values.
	Get(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*HashFieldsKey, error)
	// Set hash fields values.
	Set(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get hash keys.
	Gets(ctx context.Context, in *HashKeys, opts ...grpc.CallOption) (*HashValues, error)
}

type innerHashFieldsClient struct {
	cc *grpc.ClientConn
}

func NewInnerHashFieldsClient(cc *grpc.ClientConn) InnerHashFieldsClient {
	return &innerHashFieldsClient{cc}
}

func (c *innerHashFieldsClient) Increase(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*HashFieldsKey, error) {
	out := new(HashFieldsKey)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerHashFields/Increase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerHashFieldsClient) Get(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*HashFieldsKey, error) {
	out := new(HashFieldsKey)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerHashFields/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerHashFieldsClient) Set(ctx context.Context, in *HashFieldsKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerHashFields/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerHashFieldsClient) Gets(ctx context.Context, in *HashKeys, opts ...grpc.CallOption) (*HashValues, error) {
	out := new(HashValues)
	err := c.cc.Invoke(ctx, "/appootb.counter.InnerHashFields/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerHashFieldsServer is the server API for InnerHashFields service.
type InnerHashFieldsServer interface {
	// Increase hash fields values.
	Increase(context.Context, *HashFieldsKey) (*HashFieldsKey, error)
	// Get hash fields values.
	Get(context.Context, *HashFieldsKey) (*HashFieldsKey, error)
	// Set hash fields values.
	Set(context.Context, *HashFieldsKey) (*empty.Empty, error)
	// Get hash keys.
	Gets(context.Context, *HashKeys) (*HashValues, error)
}

// UnimplementedInnerHashFieldsServer can be embedded to have forward compatible implementations.
type UnimplementedInnerHashFieldsServer struct {
}

func (*UnimplementedInnerHashFieldsServer) Increase(ctx context.Context, req *HashFieldsKey) (*HashFieldsKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increase not implemented")
}
func (*UnimplementedInnerHashFieldsServer) Get(ctx context.Context, req *HashFieldsKey) (*HashFieldsKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInnerHashFieldsServer) Set(ctx context.Context, req *HashFieldsKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedInnerHashFieldsServer) Gets(ctx context.Context, req *HashKeys) (*HashValues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}

func RegisterInnerHashFieldsServer(s *grpc.Server, srv InnerHashFieldsServer) {
	s.RegisterService(&_InnerHashFields_serviceDesc, srv)
}

func _InnerHashFields_Increase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashFieldsKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerHashFieldsServer).Increase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerHashFields/Increase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerHashFieldsServer).Increase(ctx, req.(*HashFieldsKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerHashFields_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashFieldsKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerHashFieldsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerHashFields/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerHashFieldsServer).Get(ctx, req.(*HashFieldsKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerHashFields_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashFieldsKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerHashFieldsServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerHashFields/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerHashFieldsServer).Set(ctx, req.(*HashFieldsKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _InnerHashFields_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerHashFieldsServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.counter.InnerHashFields/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerHashFieldsServer).Gets(ctx, req.(*HashKeys))
	}
	return interceptor(ctx, in, info, handler)
}

var _InnerHashFields_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.counter.InnerHashFields",
	HandlerType: (*InnerHashFieldsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increase",
			Handler:    _InnerHashFields_Increase_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InnerHashFields_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _InnerHashFields_Set_Handler,
		},
		{
			MethodName: "Gets",
			Handler:    _InnerHashFields_Gets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner_hash_fields.proto",
}
