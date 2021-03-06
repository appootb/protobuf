// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package message

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/api"
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

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0xb9, 0xfc, 0x0b, 0x7f, 0x3d, 0x28, 0x85, 0x48, 0xa9, 0x46, 0x41, 0xc8, 0x20, 0xe2,
	0x90, 0x13, 0x47, 0x37, 0xdb, 0x45, 0x07, 0x21, 0xd0, 0xad, 0x74, 0xb9, 0x5c, 0x5e, 0x93, 0xa3,
	0xbd, 0x7b, 0x43, 0xee, 0x2d, 0xdd, 0xfd, 0x0a, 0x6e, 0xe2, 0xe4, 0xe8, 0xa7, 0x70, 0x14, 0x57,
	0x17, 0xc1, 0xd5, 0x0f, 0xe1, 0x28, 0x69, 0xaf, 0x82, 0x55, 0xc7, 0x3c, 0xbf, 0xf7, 0xf9, 0x91,
	0xe7, 0x38, 0x57, 0xa5, 0xa4, 0xa4, 0xaa, 0x91, 0x30, 0xec, 0xc8, 0xaa, 0x42, 0xa4, 0x2c, 0x31,
	0xe0, 0x9c, 0x2c, 0x20, 0xda, 0xf5, 0x81, 0x90, 0x95, 0x16, 0x73, 0xc8, 0x1c, 0xaa, 0x09, 0xf8,
	0xeb, 0x68, 0x7f, 0x05, 0x2b, 0xa8, 0x8d, 0x76, 0x4e, 0xa3, 0x15, 0x06, 0xa8, 0xc4, 0xdc, 0x1f,
	0xec, 0x15, 0x88, 0xc5, 0x14, 0x16, 0x65, 0x69, 0x2d, 0x92, 0x24, 0x8d, 0xd6, 0x79, 0xda, 0xd6,
	0x56, 0x4d, 0x67, 0x39, 0x2c, 0x3f, 0x4f, 0x1e, 0x19, 0x6f, 0x0d, 0x4a, 0x49, 0xe1, 0x88, 0x6f,
	0x5c, 0x58, 0x82, 0x5a, 0x2a, 0x0a, 0x7b, 0xc9, 0xda, 0x1f, 0x25, 0x43, 0xaa, 0x41, 0x9a, 0xe8,
	0x2f, 0x10, 0x6f, 0xbf, 0xde, 0xb1, 0xa7, 0xd6, 0xed, 0x1b, 0xeb, 0xf0, 0xb6, 0xf0, 0x4c, 0x34,
	0x13, 0x0f, 0xd9, 0x31, 0x0b, 0xc7, 0xbc, 0x35, 0x04, 0x9b, 0x87, 0xdd, 0x1f, 0xf5, 0x14, 0x1d,
	0x45, 0x3b, 0xbf, 0xc6, 0x46, 0xd6, 0x93, 0x38, 0x5e, 0x78, 0xaf, 0x5f, 0xde, 0x6f, 0x82, 0x5e,
	0x1c, 0x7e, 0x33, 0x0b, 0x07, 0x36, 0x3f, 0x65, 0x47, 0xfd, 0x39, 0xdf, 0x52, 0x68, 0xd6, 0x1d,
	0xfd, 0xcd, 0x66, 0x56, 0xda, 0x8c, 0x3c, 0x67, 0x29, 0x1b, 0x1d, 0x14, 0x9a, 0xca, 0x59, 0x96,
	0x28, 0x34, 0xe2, 0xeb, 0x05, 0x1b, 0x9a, 0xcd, 0xae, 0x44, 0x81, 0x2b, 0xf5, 0x07, 0x63, 0xf7,
	0xc1, 0xbf, 0x41, 0xda, 0x7f, 0x08, 0xfe, 0x5f, 0x2e, 0xa3, 0xe7, 0xa0, 0x7b, 0xb6, 0x2c, 0x8c,
	0x17, 0xba, 0xb1, 0xcf, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x3b, 0x67, 0x91, 0xc4, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	// Chat streaming interaction.
	Interact(ctx context.Context, opts ...grpc.CallOption) (Chat_InteractClient, error)
	// Send chat message.
	Send(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Postmark, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Interact(ctx context.Context, opts ...grpc.CallOption) (Chat_InteractClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/appootb.message.Chat/Interact", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatInteractClient{stream}
	return x, nil
}

type Chat_InteractClient interface {
	Send(*Stream) error
	Recv() (*Stream, error)
	grpc.ClientStream
}

type chatInteractClient struct {
	grpc.ClientStream
}

func (x *chatInteractClient) Send(m *Stream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatInteractClient) Recv() (*Stream, error) {
	m := new(Stream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) Send(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Postmark, error) {
	out := new(Postmark)
	err := c.cc.Invoke(ctx, "/appootb.message.Chat/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	// Chat streaming interaction.
	Interact(Chat_InteractServer) error
	// Send chat message.
	Send(context.Context, *Post) (*Postmark, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Interact(srv Chat_InteractServer) error {
	return status.Errorf(codes.Unimplemented, "method Interact not implemented")
}
func (*UnimplementedChatServer) Send(ctx context.Context, req *Post) (*Postmark, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Interact_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Interact(&chatInteractServer{stream})
}

type Chat_InteractServer interface {
	Send(*Stream) error
	Recv() (*Stream, error)
	grpc.ServerStream
}

type chatInteractServer struct {
	grpc.ServerStream
}

func (x *chatInteractServer) Send(m *Stream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatInteractServer) Recv() (*Stream, error) {
	m := new(Stream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.message.Chat/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Send(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.message.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Chat_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Interact",
			Handler:       _Chat_Interact_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
