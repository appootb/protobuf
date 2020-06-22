// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bind.proto

package account

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

// Account bind information.
type BindInfo struct {
	SourceType           *AuthType `protobuf:"varint,1,req,name=source_type,json=sourceType,enum=appootb.account.AuthType" json:"source_type,omitempty"`
	SourceId             *string   `protobuf:"bytes,2,req,name=source_id,json=sourceId" json:"source_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BindInfo) Reset()         { *m = BindInfo{} }
func (m *BindInfo) String() string { return proto.CompactTextString(m) }
func (*BindInfo) ProtoMessage()    {}
func (*BindInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ba67facb4874d5, []int{0}
}

func (m *BindInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindInfo.Unmarshal(m, b)
}
func (m *BindInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindInfo.Marshal(b, m, deterministic)
}
func (m *BindInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindInfo.Merge(m, src)
}
func (m *BindInfo) XXX_Size() int {
	return xxx_messageInfo_BindInfo.Size(m)
}
func (m *BindInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BindInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BindInfo proto.InternalMessageInfo

func (m *BindInfo) GetSourceType() AuthType {
	if m != nil && m.SourceType != nil {
		return *m.SourceType
	}
	return AuthType_AUTH_TYPE_UNSPECIFIED
}

func (m *BindInfo) GetSourceId() string {
	if m != nil && m.SourceId != nil {
		return *m.SourceId
	}
	return ""
}

// Account bind request.
type BindRequest struct {
	Bind                 *BindInfo `protobuf:"bytes,1,req,name=bind" json:"bind,omitempty"`
	Code                 *string   `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BindRequest) Reset()         { *m = BindRequest{} }
func (m *BindRequest) String() string { return proto.CompactTextString(m) }
func (*BindRequest) ProtoMessage()    {}
func (*BindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ba67facb4874d5, []int{1}
}

func (m *BindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindRequest.Unmarshal(m, b)
}
func (m *BindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindRequest.Marshal(b, m, deterministic)
}
func (m *BindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindRequest.Merge(m, src)
}
func (m *BindRequest) XXX_Size() int {
	return xxx_messageInfo_BindRequest.Size(m)
}
func (m *BindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindRequest proto.InternalMessageInfo

func (m *BindRequest) GetBind() *BindInfo {
	if m != nil {
		return m.Bind
	}
	return nil
}

func (m *BindRequest) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

// Account binds information.
type AccountBinds struct {
	Binds                []*BindInfo `protobuf:"bytes,1,rep,name=binds" json:"binds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AccountBinds) Reset()         { *m = AccountBinds{} }
func (m *AccountBinds) String() string { return proto.CompactTextString(m) }
func (*AccountBinds) ProtoMessage()    {}
func (*AccountBinds) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ba67facb4874d5, []int{2}
}

func (m *AccountBinds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBinds.Unmarshal(m, b)
}
func (m *AccountBinds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBinds.Marshal(b, m, deterministic)
}
func (m *AccountBinds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBinds.Merge(m, src)
}
func (m *AccountBinds) XXX_Size() int {
	return xxx_messageInfo_AccountBinds.Size(m)
}
func (m *AccountBinds) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBinds.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBinds proto.InternalMessageInfo

func (m *AccountBinds) GetBinds() []*BindInfo {
	if m != nil {
		return m.Binds
	}
	return nil
}

func init() {
	proto.RegisterType((*BindInfo)(nil), "appootb.account.BindInfo")
	proto.RegisterType((*BindRequest)(nil), "appootb.account.BindRequest")
	proto.RegisterType((*AccountBinds)(nil), "appootb.account.AccountBinds")
}

func init() { proto.RegisterFile("bind.proto", fileDescriptor_93ba67facb4874d5) }

var fileDescriptor_93ba67facb4874d5 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0x65, 0x27, 0x85, 0xe4, 0xa5, 0xb4, 0xc8, 0x08, 0x14, 0xd2, 0x54, 0x8d, 0x6e, 0x40,
	0x51, 0x25, 0x6c, 0xa9, 0x63, 0x17, 0x94, 0x54, 0x08, 0xba, 0x45, 0x27, 0x26, 0x14, 0x81, 0x2e,
	0x3e, 0x37, 0xb1, 0x94, 0xf3, 0x33, 0xb1, 0x4f, 0x28, 0x2b, 0x5f, 0x81, 0x95, 0xa9, 0x23, 0xe2,
	0x93, 0xb0, 0xb2, 0x31, 0xf3, 0x21, 0x18, 0x91, 0xef, 0x1c, 0x94, 0x96, 0xc0, 0xc0, 0x66, 0xbf,
	0xf7, 0xff, 0xff, 0x7f, 0xef, 0xde, 0x19, 0x60, 0xa6, 0x4d, 0xce, 0xed, 0x0a, 0x3d, 0xb2, 0xc3,
	0xcc, 0x5a, 0x44, 0x3f, 0xe3, 0x99, 0x94, 0x58, 0x1a, 0xdf, 0x3b, 0x89, 0x05, 0x61, 0xd5, 0xaa,
	0xd0, 0xce, 0x69, 0x34, 0xa2, 0x50, 0x7e, 0x81, 0xd1, 0xd1, 0xeb, 0xcf, 0x11, 0xe7, 0x4b, 0x25,
	0x32, 0xab, 0x45, 0x66, 0x0c, 0xfa, 0xcc, 0x6b, 0x34, 0x2e, 0x76, 0x8f, 0x62, 0xb7, 0xba, 0xcd,
	0xca, 0x2b, 0xa1, 0x0a, 0xeb, 0xd7, 0xb1, 0xb9, 0x2f, 0xb1, 0x28, 0xd0, 0xd4, 0xb7, 0x44, 0x42,
	0x6b, 0xac, 0x4d, 0x7e, 0x69, 0xae, 0x90, 0x9d, 0x43, 0xc7, 0x61, 0xb9, 0x92, 0xea, 0xad, 0x5f,
	0x5b, 0xd5, 0x25, 0x03, 0x3a, 0x3c, 0x38, 0x7b, 0xcc, 0x6f, 0x0d, 0xc7, 0x47, 0xa5, 0x5f, 0xbc,
	0x5a, 0x5b, 0x95, 0x42, 0xad, 0x0e, 0x67, 0x76, 0x04, 0xed, 0xe8, 0xd5, 0x79, 0x97, 0x0e, 0xe8,
	0xb0, 0x9d, 0xb6, 0xea, 0xc2, 0x65, 0x9e, 0x4c, 0xa0, 0x13, 0x20, 0xa9, 0x7a, 0x57, 0x2a, 0xe7,
	0xd9, 0x53, 0x68, 0x86, 0x8f, 0xaf, 0x00, 0x9d, 0x1d, 0x80, 0xcd, 0x40, 0x69, 0x25, 0x63, 0x0c,
	0x9a, 0x12, 0x73, 0xd5, 0xa5, 0x03, 0x32, 0x6c, 0xa7, 0xd5, 0x39, 0x79, 0x06, 0xfb, 0xa3, 0x5a,
	0x1d, 0xc4, 0x8e, 0x09, 0xd8, 0x0b, 0x5a, 0xd7, 0x25, 0x83, 0xc6, 0xbf, 0x33, 0x6b, 0xdd, 0xd9,
	0x35, 0x85, 0x66, 0xa8, 0xb1, 0x37, 0xb0, 0x37, 0xb2, 0x76, 0xb9, 0x66, 0xfd, 0x9d, 0x9e, 0x38,
	0x73, 0xef, 0x11, 0xaf, 0x77, 0xca, 0x37, 0x3b, 0xe5, 0xcf, 0xc3, 0x4e, 0x93, 0x93, 0xef, 0x9f,
	0x48, 0xa3, 0xf5, 0x85, 0x7e, 0xf8, 0xf6, 0xe3, 0x23, 0x65, 0xc9, 0x3d, 0x11, 0xcd, 0x22, 0x60,
	0xce, 0xc9, 0x29, 0x9b, 0xc2, 0x9d, 0x8b, 0xcc, 0x48, 0xb5, 0xfc, 0x4f, 0x40, 0x7f, 0x1b, 0x70,
	0x78, 0x7a, 0x13, 0xc0, 0xa6, 0xd0, 0x7c, 0xa1, 0xbc, 0x63, 0x7f, 0x71, 0xf7, 0x8e, 0xff, 0xfc,
	0x7b, 0x5b, 0x6b, 0x4b, 0x8e, 0xb7, 0xc3, 0xef, 0xb3, 0x83, 0x1b, 0xe1, 0x6e, 0xfc, 0x1e, 0x1e,
	0x48, 0x2c, 0x6e, 0x47, 0x8c, 0xdb, 0xc1, 0x3c, 0x09, 0xa0, 0x97, 0x64, 0x42, 0x5e, 0x3f, 0x99,
	0x6b, 0xbf, 0x28, 0x67, 0x5c, 0x62, 0x21, 0x7e, 0xbf, 0xdb, 0xcd, 0xcb, 0x9b, 0xe3, 0x26, 0xf3,
	0x27, 0x21, 0xd7, 0xb4, 0x71, 0x31, 0x19, 0x7f, 0xa6, 0x77, 0xe3, 0x10, 0x5f, 0xe9, 0xc3, 0x51,
	0x6d, 0x98, 0x56, 0x71, 0xd3, 0x58, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x61, 0x2e, 0xf1, 0x1b,
	0x1d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BindClient is the client API for Bind service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BindClient interface {
	// Apply account bind.
	Apply(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Cancel account bind.
	Cancel(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get all bind info.
	Gets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountBinds, error)
}

type bindClient struct {
	cc *grpc.ClientConn
}

func NewBindClient(cc *grpc.ClientConn) BindClient {
	return &bindClient{cc}
}

func (c *bindClient) Apply(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.account.Bind/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindClient) Cancel(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.account.Bind/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindClient) Gets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountBinds, error) {
	out := new(AccountBinds)
	err := c.cc.Invoke(ctx, "/appootb.account.Bind/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BindServer is the server API for Bind service.
type BindServer interface {
	// Apply account bind.
	Apply(context.Context, *BindRequest) (*empty.Empty, error)
	// Cancel account bind.
	Cancel(context.Context, *BindRequest) (*empty.Empty, error)
	// Get all bind info.
	Gets(context.Context, *empty.Empty) (*AccountBinds, error)
}

// UnimplementedBindServer can be embedded to have forward compatible implementations.
type UnimplementedBindServer struct {
}

func (*UnimplementedBindServer) Apply(ctx context.Context, req *BindRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (*UnimplementedBindServer) Cancel(ctx context.Context, req *BindRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedBindServer) Gets(ctx context.Context, req *empty.Empty) (*AccountBinds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}

func RegisterBindServer(s *grpc.Server, srv BindServer) {
	s.RegisterService(&_Bind_serviceDesc, srv)
}

func _Bind_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Bind/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindServer).Apply(ctx, req.(*BindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bind_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Bind/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindServer).Cancel(ctx, req.(*BindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bind_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Bind/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindServer).Gets(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bind_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.account.Bind",
	HandlerType: (*BindServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _Bind_Apply_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Bind_Cancel_Handler,
		},
		{
			MethodName: "Gets",
			Handler:    _Bind_Gets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bind.proto",
}
