// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group.proto

package relation

import (
	context "context"
	fmt "fmt"
	common "github.com/appootb/protobuf/go/common"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
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

// Group info message type.
type GroupInfo struct {
	Id                   *uint64              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 *string              `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Notice               *string              `protobuf:"bytes,3,opt,name=notice" json:"notice,omitempty"`
	Members              []uint64             `protobuf:"varint,4,rep,name=members" json:"members,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{0}
}

func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupInfo.Unmarshal(m, b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
}
func (m *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(m, src)
}
func (m *GroupInfo) XXX_Size() int {
	return xxx_messageInfo_GroupInfo.Size(m)
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GroupInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GroupInfo) GetNotice() string {
	if m != nil && m.Notice != nil {
		return *m.Notice
	}
	return ""
}

func (m *GroupInfo) GetMembers() []uint64 {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GroupInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GroupInfo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupInfo)(nil), "appootb.relation.GroupInfo")
}

func init() { proto.RegisterFile("group.proto", fileDescriptor_e10f4c9b19ad8eee) }

var fileDescriptor_e10f4c9b19ad8eee = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x3f, 0x8f, 0xd3, 0x30,
	0x14, 0x97, 0x93, 0x50, 0xe8, 0xab, 0x04, 0xc8, 0xaa, 0x4a, 0x2e, 0x9c, 0xb8, 0x52, 0x21, 0x5d,
	0x75, 0x43, 0x2c, 0x75, 0x83, 0xad, 0xbd, 0xe1, 0xb8, 0xad, 0x8a, 0xb8, 0x05, 0x9d, 0x54, 0xa5,
	0x89, 0x9b, 0x5a, 0x3a, 0xfb, 0x85, 0xc4, 0x61, 0x01, 0x16, 0x06, 0xbe, 0x00, 0x2b, 0x13, 0x23,
	0x9f, 0x04, 0xb1, 0x32, 0xc0, 0x07, 0xe0, 0x43, 0x30, 0xa2, 0xb8, 0x71, 0x8b, 0x0a, 0x2d, 0x12,
	0xb7, 0xc5, 0xef, 0xbd, 0xdf, 0xbf, 0x3c, 0x1b, 0x3a, 0x59, 0x81, 0x55, 0x1e, 0xe6, 0x05, 0x6a,
	0xa4, 0x77, 0xe3, 0x3c, 0x47, 0xd4, 0xf3, 0xb0, 0xe0, 0x57, 0xb1, 0x16, 0xa8, 0x82, 0x07, 0x4d,
	0x85, 0x25, 0x28, 0x25, 0x2a, 0x56, 0x29, 0xf1, 0xa2, 0xe2, 0x33, 0x91, 0xae, 0x10, 0xc1, 0x91,
	0xed, 0xe7, 0xbc, 0x90, 0xa2, 0x2c, 0x05, 0x2a, 0x26, 0xb9, 0x5e, 0xa2, 0x1d, 0x38, 0xcc, 0x10,
	0xb3, 0x2b, 0xce, 0xe2, 0x5c, 0xb0, 0x58, 0x29, 0xd4, 0x86, 0xb7, 0xb4, 0xf0, 0xa6, 0x6b, 0x4e,
	0xf3, 0x6a, 0xc1, 0xb4, 0x90, 0xbc, 0xd4, 0xb1, 0x6c, 0x1c, 0x0d, 0xbe, 0x11, 0x68, 0x9f, 0xd5,
	0x0e, 0xcf, 0xd5, 0x02, 0xe9, 0x6d, 0x70, 0x44, 0xea, 0x93, 0x3e, 0x19, 0x7a, 0x91, 0x23, 0x52,
	0x4a, 0xc1, 0x53, 0xb1, 0xe4, 0xbe, 0xd3, 0x27, 0xc3, 0x76, 0x64, 0xbe, 0x69, 0x0f, 0x5a, 0x0a,
	0xb5, 0x48, 0xb8, 0xef, 0x9a, 0x6a, 0x73, 0xa2, 0x3e, 0xdc, 0x94, 0x5c, 0xce, 0x79, 0x51, 0xfa,
	0x5e, 0xdf, 0x1d, 0x7a, 0x91, 0x3d, 0xd2, 0xc7, 0x00, 0x49, 0xc1, 0x63, 0xcd, 0xd3, 0x59, 0xac,
	0x7d, 0xe8, 0x93, 0x61, 0x67, 0x14, 0x84, 0x2b, 0x67, 0xa1, 0x75, 0x16, 0x3e, 0xb3, 0xce, 0xa2,
	0x76, 0x33, 0x3d, 0xd6, 0x35, 0xb4, 0xca, 0x53, 0x0b, 0xed, 0xfc, 0x1b, 0xda, 0x4c, 0x8f, 0xf5,
	0xe8, 0x9d, 0x07, 0x37, 0x4c, 0x32, 0x9a, 0x40, 0xeb, 0xd4, 0x30, 0xd2, 0xfb, 0xe1, 0xf6, 0x02,
	0xc2, 0x75, 0xf8, 0x60, 0x5f, 0x73, 0x70, 0xf4, 0xfd, 0x03, 0xf9, 0xec, 0xbd, 0xfd, 0xfa, 0xe3,
	0xbd, 0xd3, 0x1d, 0xdc, 0x61, 0x76, 0x82, 0x99, 0xed, 0x3e, 0x21, 0x27, 0x74, 0x06, 0xee, 0x19,
	0xd7, 0xd4, 0x5f, 0x93, 0xac, 0x16, 0x1a, 0x5e, 0x98, 0x85, 0x9e, 0xa7, 0xfb, 0xe9, 0x1f, 0x6e,
	0xe8, 0x7b, 0xb4, 0xbb, 0x45, 0xcf, 0x5e, 0x89, 0xf4, 0x0d, 0x5d, 0x42, 0xeb, 0xc2, 0x84, 0xbb,
	0x46, 0x8a, 0x47, 0x1b, 0x99, 0x83, 0xd1, 0x5f, 0x65, 0xea, 0x28, 0x0b, 0x70, 0xc7, 0x69, 0x4a,
	0x0f, 0x76, 0x45, 0x29, 0xf7, 0x8b, 0x1c, 0x6f, 0x44, 0x0e, 0x83, 0x7b, 0xdb, 0x22, 0xcd, 0xa5,
	0xa8, 0x75, 0x38, 0xb4, 0x22, 0x2e, 0xf1, 0x25, 0xff, 0xdf, 0xbf, 0xf6, 0x7b, 0x9c, 0x93, 0x5d,
	0x4a, 0x93, 0xd7, 0xd0, 0x4d, 0x50, 0xfe, 0xc1, 0x33, 0x01, 0x43, 0x34, 0xad, 0x2f, 0xd1, 0x53,
	0x32, 0x25, 0xcf, 0x8f, 0x33, 0xa1, 0x97, 0x95, 0xb1, 0xc0, 0xd6, 0xaf, 0xce, 0xbe, 0x9b, 0x0c,
	0xd7, 0xf4, 0x3f, 0x09, 0xf9, 0xe8, 0xb8, 0xa7, 0xd3, 0xc9, 0x27, 0xe7, 0x56, 0xd4, 0xd4, 0xbe,
	0x38, 0xbd, 0xf1, 0x0a, 0x72, 0x69, 0x08, 0x2f, 0x6d, 0xe3, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x08, 0x33, 0xc9, 0x00, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupClient interface {
	// Create group.
	Create(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*GroupInfo, error)
	// Get group info.
	Get(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*GroupInfo, error)
	// Update group info.
	Update(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*GroupInfo, error)
	// Add group members.
	Add(ctx context.Context, in *common.UniqueIds, opts ...grpc.CallOption) (*GroupInfo, error)
	// Kick group member.
	Remove(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*GroupInfo, error)
}

type groupClient struct {
	cc *grpc.ClientConn
}

func NewGroupClient(cc *grpc.ClientConn) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) Create(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Group/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Get(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Group/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Update(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Group/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Add(ctx context.Context, in *common.UniqueIds, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Group/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Remove(ctx context.Context, in *common.UniqueId, opts ...grpc.CallOption) (*GroupInfo, error) {
	out := new(GroupInfo)
	err := c.cc.Invoke(ctx, "/appootb.relation.Group/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
type GroupServer interface {
	// Create group.
	Create(context.Context, *GroupInfo) (*GroupInfo, error)
	// Get group info.
	Get(context.Context, *common.UniqueId) (*GroupInfo, error)
	// Update group info.
	Update(context.Context, *GroupInfo) (*GroupInfo, error)
	// Add group members.
	Add(context.Context, *common.UniqueIds) (*GroupInfo, error)
	// Kick group member.
	Remove(context.Context, *common.UniqueId) (*GroupInfo, error)
}

// UnimplementedGroupServer can be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (*UnimplementedGroupServer) Create(ctx context.Context, req *GroupInfo) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGroupServer) Get(ctx context.Context, req *common.UniqueId) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGroupServer) Update(ctx context.Context, req *GroupInfo) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGroupServer) Add(ctx context.Context, req *common.UniqueIds) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedGroupServer) Remove(ctx context.Context, req *common.UniqueId) (*GroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterGroupServer(s *grpc.Server, srv GroupServer) {
	s.RegisterService(&_Group_serviceDesc, srv)
}

func _Group_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Group/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Create(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UniqueId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Group/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Get(ctx, req.(*common.UniqueId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Group/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Update(ctx, req.(*GroupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UniqueIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Group/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Add(ctx, req.(*common.UniqueIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UniqueId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.relation.Group/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Remove(ctx, req.(*common.UniqueId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Group_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.relation.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Group_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Group_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Group_Update_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Group_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Group_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
