// Code generated by protoc-gen-go. DO NOT EDIT.
// source: channel.proto

package relation

import (
	context "context"
	fmt "fmt"
	common "github.com/appootb/protobuf/go/common"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Channel info.
type ChannelInfo struct {
	Id                   *uint64              `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string              `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Description          *string              `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,req,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,req,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f385724121f37b, []int{0}
}

func (m *ChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInfo.Unmarshal(m, b)
}
func (m *ChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInfo.Marshal(b, m, deterministic)
}
func (m *ChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInfo.Merge(m, src)
}
func (m *ChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ChannelInfo.Size(m)
}
func (m *ChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInfo proto.InternalMessageInfo

func (m *ChannelInfo) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ChannelInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChannelInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ChannelInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ChannelInfo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelInfo)(nil), "appootb.relation.ChannelInfo")
}

func init() { proto.RegisterFile("channel.proto", fileDescriptor_c8f385724121f37b) }

var fileDescriptor_c8f385724121f37b = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0x96, 0xdd, 0x0a, 0xa8, 0xcb, 0x21, 0x64, 0xa1, 0x2a, 0x0a, 0x70, 0x17, 0x95, 0x81, 0xaa,
	0x43, 0x2c, 0xdd, 0x06, 0x5b, 0x5a, 0x21, 0x71, 0x5b, 0x15, 0xc1, 0x82, 0x4e, 0xaa, 0x7c, 0xf1,
	0x7b, 0xa9, 0xa5, 0xf8, 0xe3, 0x12, 0x67, 0x38, 0xa1, 0x63, 0xe0, 0x2f, 0xb0, 0x32, 0x31, 0xf2,
	0x4b, 0x80, 0x95, 0x85, 0x81, 0x91, 0x1f, 0xc1, 0x88, 0x9a, 0xd8, 0xe5, 0x74, 0x45, 0x62, 0x8b,
	0xfd, 0x7c, 0xbc, 0xcf, 0x1b, 0x3f, 0xe4, 0xa0, 0xd8, 0x70, 0xad, 0xa1, 0x4a, 0x6d, 0x6d, 0x9c,
	0xa1, 0xf7, 0xb9, 0xb5, 0xc6, 0xb8, 0xb3, 0xb4, 0x86, 0x8a, 0x3b, 0x69, 0x74, 0x7c, 0xe8, 0x6f,
	0x58, 0x61, 0x94, 0x32, 0x9a, 0xb5, 0x5a, 0x5e, 0xb4, 0xb0, 0x96, 0xa2, 0x57, 0xc4, 0x47, 0x01,
	0xb7, 0x50, 0x2b, 0xd9, 0x34, 0xd2, 0x68, 0xa6, 0xc0, 0x6d, 0x4c, 0x20, 0x3c, 0x2a, 0x8d, 0x29,
	0x2b, 0x60, 0xdc, 0x4a, 0xc6, 0xb5, 0x36, 0xae, 0xf3, 0x6d, 0x3c, 0xfa, 0xd0, 0xa3, 0xdd, 0xe9,
	0xac, 0x3d, 0x67, 0xa0, 0xac, 0xbb, 0x0c, 0xde, 0x37, 0x41, 0x27, 0x15, 0x34, 0x8e, 0x2b, 0xeb,
	0x09, 0x07, 0x52, 0x17, 0x55, 0x2b, 0xa0, 0x3f, 0x4e, 0xbf, 0x22, 0x32, 0x5e, 0xf6, 0xfb, 0x9c,
	0xe8, 0x73, 0x43, 0xef, 0x11, 0x2c, 0x45, 0x84, 0x12, 0x3c, 0x1b, 0xe6, 0x58, 0x0a, 0x4a, 0xc9,
	0x50, 0x73, 0x05, 0x11, 0x4e, 0xf0, 0x6c, 0x94, 0x77, 0xdf, 0x34, 0x21, 0x63, 0x01, 0x4d, 0x51,
	0x4b, 0xbb, 0x8d, 0x15, 0x0d, 0x12, 0x34, 0x1b, 0xe5, 0xd7, 0xaf, 0xe8, 0x33, 0x42, 0x8a, 0x1a,
	0xb8, 0x03, 0xb1, 0xe6, 0x2e, 0x22, 0x09, 0x9e, 0x8d, 0x8f, 0xe3, 0xb4, 0x8f, 0x96, 0x86, 0x68,
	0xe9, 0xab, 0x10, 0x2d, 0x1f, 0x79, 0x76, 0xe6, 0xb6, 0xd2, 0xd6, 0x8a, 0x20, 0x1d, 0xff, 0x5f,
	0xea, 0xd9, 0x99, 0x3b, 0xfe, 0x89, 0xc8, 0x6d, 0xbf, 0x0b, 0xb5, 0x64, 0x90, 0x09, 0x41, 0x0f,
	0xd3, 0x9b, 0xaf, 0x93, 0x66, 0xd6, 0x56, 0x97, 0x39, 0x5c, 0xb4, 0xd0, 0xb8, 0xf8, 0xf1, 0x3e,
	0x7e, 0xed, 0x6f, 0x4c, 0xe7, 0x3f, 0x3e, 0xa2, 0x2f, 0xc3, 0xf7, 0xdf, 0x7f, 0x7d, 0xc0, 0x47,
	0xd3, 0x98, 0x05, 0x0e, 0xf3, 0x0d, 0x60, 0x6f, 0x1d, 0xaf, 0x4b, 0x70, 0x57, 0xcf, 0xd1, 0x9c,
	0xae, 0xc9, 0xad, 0x25, 0xd7, 0x05, 0x54, 0x34, 0xda, 0x99, 0xf6, 0x05, 0x48, 0x5f, 0x77, 0x05,
	0x38, 0x11, 0xf1, 0x64, 0x6f, 0x91, 0x17, 0xdb, 0xb7, 0x9b, 0x3e, 0xf9, 0x3b, 0x27, 0x9a, 0x4f,
	0xfe, 0x31, 0x47, 0x8a, 0xab, 0xc5, 0x3b, 0xf2, 0xa0, 0x30, 0x6a, 0x2f, 0xf0, 0xe2, 0xae, 0x4f,
	0xbc, 0xda, 0x7a, 0xbe, 0x44, 0x2b, 0xf4, 0xe6, 0x69, 0x29, 0xdd, 0xa6, 0xed, 0xc6, 0xb3, 0x5d,
	0xd7, 0x42, 0x21, 0x4a, 0xb3, 0xf3, 0xff, 0x8d, 0xd0, 0x27, 0x3c, 0x58, 0xae, 0x16, 0x9f, 0xf1,
	0x9d, 0xdc, 0xdf, 0x7d, 0xc3, 0x93, 0xac, 0x97, 0x9c, 0x76, 0x86, 0xa7, 0x01, 0xf8, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xe5, 0xb8, 0x99, 0xf8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChannelClient interface {
	// Follow channel.
	Add(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ChannelInfo, error)
	// Unfollow channel.
	Cancel(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*empty.Empty, error)
}

type channelClient struct {
	cc *grpc.ClientConn
}

func NewChannelClient(cc *grpc.ClientConn) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) Add(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ChannelInfo, error) {
	out := new(ChannelInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Channel/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) Cancel(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.relation.Channel/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
type ChannelServer interface {
	// Follow channel.
	Add(context.Context, *ApplyRequest) (*ChannelInfo, error)
	// Unfollow channel.
	Cancel(context.Context, *common.UniqueId) (*empty.Empty, error)
}

// UnimplementedChannelServer can be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (*UnimplementedChannelServer) Add(ctx context.Context, req *ApplyRequest) (*ChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedChannelServer) Cancel(ctx context.Context, req *common.UniqueId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterChannelServer(s *grpc.Server, srv ChannelServer) {
	s.RegisterService(&_Channel_serviceDesc, srv)
}

func _Channel_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Channel/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).Add(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UniqueId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Channel/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).Cancel(ctx, req.(*common.UniqueId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Channel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.relation.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Channel_Add_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Channel_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel.proto",
}
