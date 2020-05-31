// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: password.proto

package account

import (
	context "context"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Account password request.
type PasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password    *string `protobuf:"bytes,1,req,name=password" json:"password,omitempty"`                          // Account password
	OldPassword *string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword" json:"old_password,omitempty"` // Account old password
	VerifyCode  *string `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode" json:"verify_code,omitempty"`    // Verify code
}

func (x *PasswordRequest) Reset() {
	*x = PasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRequest) ProtoMessage() {}

func (x *PasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRequest.ProtoReflect.Descriptor instead.
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *PasswordRequest) GetOldPassword() string {
	if x != nil && x.OldPassword != nil {
		return *x.OldPassword
	}
	return ""
}

func (x *PasswordRequest) GetVerifyCode() string {
	if x != nil && x.VerifyCode != nil {
		return *x.VerifyCode
	}
	return ""
}

var File_password_proto protoreflect.FileDescriptor

var file_password_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71,
	0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0xc4, 0x02, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x65,
	0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74,
	0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x23, 0xda, 0x9c, 0x01, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x23, 0xda, 0x9c, 0x01, 0x03,
	0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x32, 0x11, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x01, 0x2a, 0x12,
	0x67, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f,
	0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x23, 0xda, 0x9c, 0x01, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x1a, 0x11, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x73, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01,
	0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x43, 0x50, 0x42, 0xaa, 0x02, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xca, 0x02, 0x0d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
}

var (
	file_password_proto_rawDescOnce sync.Once
	file_password_proto_rawDescData = file_password_proto_rawDesc
)

func file_password_proto_rawDescGZIP() []byte {
	file_password_proto_rawDescOnce.Do(func() {
		file_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_password_proto_rawDescData)
	})
	return file_password_proto_rawDescData
}

var file_password_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_password_proto_goTypes = []interface{}{
	(*PasswordRequest)(nil), // 0: appootb.account.PasswordRequest
	(*Secret)(nil),          // 1: appootb.account.Secret
}
var file_password_proto_depIdxs = []int32{
	0, // 0: appootb.account.Password.Set:input_type -> appootb.account.PasswordRequest
	0, // 1: appootb.account.Password.Update:input_type -> appootb.account.PasswordRequest
	0, // 2: appootb.account.Password.Reset:input_type -> appootb.account.PasswordRequest
	1, // 3: appootb.account.Password.Set:output_type -> appootb.account.Secret
	1, // 4: appootb.account.Password.Update:output_type -> appootb.account.Secret
	1, // 5: appootb.account.Password.Reset:output_type -> appootb.account.Secret
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_password_proto_init() }
func file_password_proto_init() {
	if File_password_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_password_proto_goTypes,
		DependencyIndexes: file_password_proto_depIdxs,
		MessageInfos:      file_password_proto_msgTypes,
	}.Build()
	File_password_proto = out.File
	file_password_proto_rawDesc = nil
	file_password_proto_goTypes = nil
	file_password_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

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
	cc grpc.ClientConnInterface
}

func NewPasswordClient(cc grpc.ClientConnInterface) PasswordClient {
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

func (*UnimplementedPasswordServer) Set(context.Context, *PasswordRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedPasswordServer) Update(context.Context, *PasswordRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPasswordServer) Reset(context.Context, *PasswordRequest) (*Secret, error) {
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
