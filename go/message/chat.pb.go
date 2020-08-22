// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package message

import (
	context "context"
	fmt "fmt"
	_ "github.com/appootb/protobuf/go/api"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
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
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x2c, 0x28, 0xc8, 0xcf, 0x2f, 0x49, 0xd2,
	0xcb, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x92, 0x86, 0x0a, 0xe8, 0x27, 0x16, 0x64, 0xea,
	0x97, 0xa7, 0x26, 0x15, 0xe7, 0x27, 0x67, 0xa7, 0x42, 0x55, 0x4b, 0xc9, 0xc3, 0x24, 0x0b, 0x52,
	0x8b, 0x72, 0x33, 0x8b, 0x8b, 0x33, 0xf3, 0xf3, 0xf4, 0x73, 0x53, 0x4b, 0x32, 0xf2, 0x53, 0xa0,
	0x0a, 0x78, 0x33, 0xf3, 0x92, 0x73, 0x4a, 0x53, 0x52, 0x21, 0x5c, 0xa3, 0x24, 0x2e, 0x16, 0xe7,
	0x8c, 0xc4, 0x12, 0xa1, 0x28, 0x2e, 0x0e, 0xcf, 0xbc, 0x92, 0xd4, 0xa2, 0xc4, 0xe4, 0x12, 0x21,
	0x71, 0x3d, 0x34, 0x2b, 0xf5, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0xa5, 0x70, 0x49, 0x28, 0x49,
	0xdc, 0x98, 0xc3, 0x78, 0x92, 0x65, 0xd6, 0x6d, 0x46, 0x7e, 0x2e, 0x5e, 0x7d, 0xa8, 0x9c, 0x3e,
	0xc8, 0x0f, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0xe5, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0xe8, 0x7a, 0x9d,
	0x38, 0x41, 0x16, 0x07, 0x80, 0x5c, 0xe1, 0xc1, 0x18, 0xc0, 0x18, 0xa5, 0x96, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf7, 0x04, 0x48, 0x36, 0xa9, 0x34, 0x4d, 0x3f,
	0x3d, 0x1f, 0x66, 0xec, 0x0f, 0x46, 0xc6, 0x45, 0x4c, 0xcc, 0xce, 0x01, 0x4e, 0xab, 0x98, 0xd8,
	0x7d, 0x21, 0x42, 0xa7, 0x98, 0x44, 0x1d, 0x21, 0x1a, 0x62, 0xc0, 0xc6, 0xc5, 0x40, 0xc5, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x83, 0xe3, 0xcd, 0x47, 0x01, 0x00, 0x00,
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

// ChatServer is the server API for Chat service.
type ChatServer interface {
	// Chat streaming interaction.
	Interact(Chat_InteractServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Interact(srv Chat_InteractServer) error {
	return status.Errorf(codes.Unimplemented, "method Interact not implemented")
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

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.message.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
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
