// Code generated by protoc-gen-go. DO NOT EDIT.
// source: password.proto

package account

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
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

// Account password request.
type PasswordRequest struct {
	Password             *string  `protobuf:"bytes,1,req,name=password" json:"password,omitempty"`
	OldPassword          *string  `protobuf:"bytes,2,opt,name=old_password,json=oldPassword" json:"old_password,omitempty"`
	VerifyCode           *string  `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode" json:"verify_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordRequest) Reset()         { *m = PasswordRequest{} }
func (m *PasswordRequest) String() string { return proto.CompactTextString(m) }
func (*PasswordRequest) ProtoMessage()    {}
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3819adb3c6759fd8, []int{0}
}

func (m *PasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordRequest.Unmarshal(m, b)
}
func (m *PasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordRequest.Marshal(b, m, deterministic)
}
func (m *PasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordRequest.Merge(m, src)
}
func (m *PasswordRequest) XXX_Size() int {
	return xxx_messageInfo_PasswordRequest.Size(m)
}
func (m *PasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordRequest proto.InternalMessageInfo

func (m *PasswordRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PasswordRequest) GetOldPassword() string {
	if m != nil && m.OldPassword != nil {
		return *m.OldPassword
	}
	return ""
}

func (m *PasswordRequest) GetVerifyCode() string {
	if m != nil && m.VerifyCode != nil {
		return *m.VerifyCode
	}
	return ""
}

func init() {
	proto.RegisterType((*PasswordRequest)(nil), "appootb.account.PasswordRequest")
}

func init() { proto.RegisterFile("password.proto", fileDescriptor_3819adb3c6759fd8) }

var fileDescriptor_3819adb3c6759fd8 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x4f, 0x2a, 0x41,
	0x14, 0xcd, 0xcc, 0xe6, 0xbd, 0x07, 0x03, 0xef, 0x91, 0xb7, 0x46, 0x25, 0x1b, 0x13, 0x10, 0x13,
	0x43, 0x2c, 0x76, 0x12, 0x4a, 0x3b, 0xa0, 0xb1, 0xdc, 0x2c, 0xb1, 0x31, 0x24, 0x64, 0x99, 0xbd,
	0x2c, 0x9b, 0xb0, 0x7b, 0x87, 0x9d, 0x59, 0x8d, 0xb1, 0xf3, 0x2f, 0xd8, 0x5a, 0x59, 0x1a, 0x7f,
	0x86, 0x95, 0xad, 0x9d, 0xb5, 0x3f, 0xc2, 0xd2, 0xc0, 0x7e, 0x14, 0x48, 0x49, 0x79, 0xcf, 0x39,
	0xf7, 0x9c, 0x73, 0x33, 0xc3, 0xfe, 0x49, 0x4f, 0xa9, 0x1b, 0x4c, 0x7c, 0x5b, 0x26, 0xa8, 0xd1,
	0x6c, 0x78, 0x52, 0x22, 0xea, 0xa9, 0xed, 0x09, 0x81, 0x69, 0xac, 0xad, 0x56, 0x0e, 0x70, 0x09,
	0x49, 0x14, 0x2a, 0x15, 0x62, 0xcc, 0x23, 0xd0, 0x73, 0xcc, 0x37, 0xac, 0xa3, 0x00, 0x31, 0x58,
	0x00, 0xf7, 0x64, 0xc8, 0xbd, 0x38, 0x46, 0xed, 0xe9, 0x10, 0x63, 0x95, 0xb3, 0x75, 0x81, 0x51,
	0x84, 0x71, 0x36, 0x75, 0x96, 0xac, 0xe1, 0xe4, 0x79, 0x2e, 0x2c, 0x53, 0x50, 0xda, 0xb4, 0x58,
	0xa5, 0xa8, 0xd0, 0x24, 0x6d, 0xda, 0xad, 0xba, 0xe5, 0x6c, 0x1e, 0xb3, 0x3a, 0x2e, 0xfc, 0x49,
	0xc9, 0xd3, 0x36, 0xe9, 0x56, 0xdd, 0x1a, 0x2e, 0xfc, 0xc2, 0xc5, 0x6c, 0xb1, 0xda, 0x35, 0x24,
	0xe1, 0xec, 0x76, 0x22, 0xd0, 0x87, 0xa6, 0xb1, 0x56, 0xb0, 0x0c, 0x1a, 0xa2, 0x0f, 0xbd, 0x57,
	0xca, 0x2a, 0xa5, 0x1a, 0x98, 0x31, 0x02, 0x6d, 0xb6, 0xed, 0x8d, 0x2b, 0xed, 0x8d, 0x56, 0xd6,
	0xe1, 0x0f, 0xc5, 0x08, 0x44, 0x02, 0xba, 0x73, 0xf2, 0xf1, 0x48, 0x8c, 0xca, 0x0b, 0xbd, 0x7f,
	0xff, 0x7c, 0xa0, 0x07, 0x9d, 0xff, 0x3c, 0xe7, 0x79, 0x51, 0xf2, 0x9c, 0x9c, 0x99, 0x73, 0xf6,
	0xfb, 0x52, 0xfa, 0x9e, 0x86, 0xdd, 0x25, 0xf5, 0xb6, 0x27, 0x05, 0xec, 0x97, 0x0b, 0x6a, 0x97,
	0x27, 0x59, 0x5b, 0x83, 0x06, 0x77, 0x6c, 0x4f, 0x60, 0xb4, 0x69, 0x31, 0xf8, 0x5b, 0xa4, 0x38,
	0xab, 0xf7, 0xbd, 0x20, 0x0e, 0xb9, 0x3a, 0x0d, 0x42, 0x3d, 0x4f, 0xa7, 0xb6, 0xc0, 0x88, 0x97,
	0x7f, 0x67, 0xc5, 0x4e, 0xd3, 0x19, 0x0f, 0xb0, 0x70, 0xff, 0x22, 0xe4, 0x89, 0x1a, 0x43, 0x67,
	0xf0, 0x4c, 0xff, 0xf4, 0x33, 0xe8, 0x8d, 0xee, 0xf7, 0xb3, 0x85, 0xf1, 0xda, 0x6e, 0x9c, 0xe3,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x1d, 0x78, 0xdb, 0xa5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PasswordClient is the client API for Password service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PasswordClient interface {
	// Set account password.
	Set(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error)
	// Update account password.
	Update(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error)
	// Reset account password.
	Reset(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error)
}

type passwordClient struct {
	cc *grpc.ClientConn
}

func NewPasswordClient(cc *grpc.ClientConn) PasswordClient {
	return &passwordClient{cc}
}

func (c *passwordClient) Set(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/appootb.account.Password/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordClient) Update(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/appootb.account.Password/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordClient) Reset(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/appootb.account.Password/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordServer is the server API for Password service.
type PasswordServer interface {
	// Set account password.
	Set(context.Context, *PasswordRequest) (*Secret, error)
	// Update account password.
	Update(context.Context, *PasswordRequest) (*Secret, error)
	// Reset account password.
	Reset(context.Context, *PasswordRequest) (*Secret, error)
}

// UnimplementedPasswordServer can be embedded to have forward compatible implementations.
type UnimplementedPasswordServer struct {
}

func (*UnimplementedPasswordServer) Set(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedPasswordServer) Update(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPasswordServer) Reset(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}

func RegisterPasswordServer(s *grpc.Server, srv PasswordServer) {
	s.RegisterService(&_Password_serviceDesc, srv)
}

func _Password_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Password/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordServer).Set(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Password_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Password/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordServer).Update(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Password_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Password/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordServer).Reset(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Password_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.account.Password",
	HandlerType: (*PasswordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Password_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Password_Update_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Password_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "password.proto",
}
