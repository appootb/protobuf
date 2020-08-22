// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner_session.proto

package message

import (
	context "context"
	fmt "fmt"
	common "github.com/appootb/protobuf/go/common"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

// Account session type.
type SessionType int32

const (
	SessionType_SESSION_TYPE_UNSPECIFIED SessionType = 0
	SessionType_SESSION_TYPE_CHAT        SessionType = 1
	SessionType_SESSION_TYPE_ROOM        SessionType = 2
)

var SessionType_name = map[int32]string{
	0: "SESSION_TYPE_UNSPECIFIED",
	1: "SESSION_TYPE_CHAT",
	2: "SESSION_TYPE_ROOM",
}

var SessionType_value = map[string]int32{
	"SESSION_TYPE_UNSPECIFIED": 0,
	"SESSION_TYPE_CHAT":        1,
	"SESSION_TYPE_ROOM":        2,
}

func (x SessionType) Enum() *SessionType {
	p := new(SessionType)
	*p = x
	return p
}

func (x SessionType) String() string {
	return proto.EnumName(SessionType_name, int32(x))
}

func (x *SessionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionType_value, data, "SessionType")
	if err != nil {
		return err
	}
	*x = SessionType(value)
	return nil
}

func (SessionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10dd3a7c9f4ef2c3, []int{0}
}

// Account session.
type Session struct {
	Type                 *SessionType         `protobuf:"varint,1,req,name=type,enum=appootb.message.SessionType" json:"type,omitempty"`
	Entry                *string              `protobuf:"bytes,2,req,name=entry" json:"entry,omitempty"`
	Metadata             *common.Metadata     `protobuf:"bytes,3,req,name=metadata" json:"metadata,omitempty"`
	AddTime              *timestamp.Timestamp `protobuf:"bytes,4,req,name=add_time,json=addTime" json:"add_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_10dd3a7c9f4ef2c3, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetType() SessionType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SessionType_SESSION_TYPE_UNSPECIFIED
}

func (m *Session) GetEntry() string {
	if m != nil && m.Entry != nil {
		return *m.Entry
	}
	return ""
}

func (m *Session) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Session) GetAddTime() *timestamp.Timestamp {
	if m != nil {
		return m.AddTime
	}
	return nil
}

// Account sessions.
type Sessions struct {
	Sessions             []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Sessions) Reset()         { *m = Sessions{} }
func (m *Sessions) String() string { return proto.CompactTextString(m) }
func (*Sessions) ProtoMessage()    {}
func (*Sessions) Descriptor() ([]byte, []int) {
	return fileDescriptor_10dd3a7c9f4ef2c3, []int{1}
}

func (m *Sessions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sessions.Unmarshal(m, b)
}
func (m *Sessions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sessions.Marshal(b, m, deterministic)
}
func (m *Sessions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sessions.Merge(m, src)
}
func (m *Sessions) XXX_Size() int {
	return xxx_messageInfo_Sessions.Size(m)
}
func (m *Sessions) XXX_DiscardUnknown() {
	xxx_messageInfo_Sessions.DiscardUnknown(m)
}

var xxx_messageInfo_Sessions proto.InternalMessageInfo

func (m *Sessions) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func init() {
	proto.RegisterEnum("appootb.message.SessionType", SessionType_name, SessionType_value)
	proto.RegisterType((*Session)(nil), "appootb.message.Session")
	proto.RegisterType((*Sessions)(nil), "appootb.message.Sessions")
}

func init() { proto.RegisterFile("inner_session.proto", fileDescriptor_10dd3a7c9f4ef2c3) }

var fileDescriptor_10dd3a7c9f4ef2c3 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x9d, 0x01, 0x03, 0xce, 0xf5, 0x0f, 0xb7, 0xf7, 0xde, 0xa4, 0x69, 0x30, 0x36, 0x2c,
	0x0c, 0x71, 0x31, 0x35, 0x44, 0xf7, 0x02, 0xd6, 0xd0, 0x05, 0xd0, 0xb4, 0x75, 0x81, 0x21, 0x36,
	0x43, 0x3b, 0xd6, 0x26, 0x4c, 0xa7, 0xe9, 0x0c, 0x26, 0xac, 0xf4, 0x31, 0x5c, 0xbb, 0xf4, 0x21,
	0x7c, 0x00, 0x9f, 0xc8, 0xa5, 0x69, 0x3b, 0x6d, 0x80, 0xbb, 0x3c, 0xe7, 0xfc, 0xce, 0x37, 0xdf,
	0x37, 0x07, 0xdd, 0xa4, 0x59, 0x46, 0x8b, 0x50, 0x50, 0x21, 0x52, 0x9e, 0xe1, 0xbc, 0xe0, 0x92,
	0x6b, 0xcf, 0x48, 0x9e, 0x73, 0x2e, 0x77, 0x98, 0x51, 0x21, 0x48, 0x42, 0x8d, 0xe7, 0xaa, 0x61,
	0x45, 0x9c, 0x31, 0x9e, 0x59, 0x8c, 0x4a, 0x12, 0x13, 0x49, 0x6a, 0xde, 0x30, 0x9b, 0x71, 0x4e,
	0x0b, 0x96, 0x56, 0x4a, 0x96, 0xa0, 0xc5, 0xb7, 0x34, 0xa2, 0x8a, 0x78, 0x91, 0x70, 0x9e, 0xec,
	0xa9, 0x55, 0x55, 0xbb, 0xc3, 0x17, 0x4b, 0xa6, 0x8c, 0x0a, 0x49, 0x58, 0xae, 0x80, 0x27, 0x69,
	0x16, 0xed, 0x0f, 0xb1, 0xe2, 0x47, 0x7f, 0x00, 0xea, 0xf9, 0xb5, 0x27, 0xed, 0x35, 0xea, 0xca,
	0x63, 0x4e, 0x75, 0x60, 0xc2, 0xf1, 0xd3, 0xc9, 0x10, 0x5f, 0x98, 0xc3, 0x8a, 0x0b, 0x8e, 0x39,
	0xf5, 0x2a, 0x52, 0xbb, 0x45, 0x0f, 0x69, 0x26, 0x8b, 0xa3, 0x0e, 0x4d, 0x38, 0x7e, 0xe4, 0xd5,
	0x85, 0xf6, 0x06, 0xf5, 0x1b, 0xdf, 0x7a, 0xc7, 0x84, 0xe3, 0xab, 0x89, 0xde, 0x6a, 0xd5, 0xb9,
	0xf0, 0x52, 0xcd, 0xbd, 0x96, 0xd4, 0xde, 0xa2, 0x3e, 0x89, 0xe3, 0xb0, 0xf4, 0xab, 0x77, 0xab,
	0x2d, 0x03, 0xd7, 0x61, 0x70, 0x13, 0x06, 0x07, 0x4d, 0x18, 0xaf, 0x47, 0xe2, 0xb8, 0xac, 0x46,
	0xef, 0x50, 0x5f, 0xf9, 0x12, 0xe5, 0xc3, 0xea, 0x7f, 0x85, 0x0e, 0xcc, 0xce, 0xd9, 0xc3, 0x17,
	0x21, 0xbc, 0x96, 0x7c, 0xb5, 0x41, 0x57, 0x27, 0xc9, 0xb4, 0x21, 0xd2, 0x7d, 0xdb, 0xf7, 0x9d,
	0xf5, 0x2a, 0x0c, 0x36, 0xae, 0x1d, 0x7e, 0x5c, 0xf9, 0xae, 0x3d, 0x77, 0x3e, 0x38, 0xf6, 0xfb,
	0xc1, 0x03, 0xed, 0x0e, 0x5d, 0x9f, 0x4d, 0xe7, 0x8b, 0x69, 0x30, 0x00, 0xf7, 0xda, 0xde, 0x7a,
	0xbd, 0x1c, 0xc0, 0xc9, 0x2d, 0x7a, 0xec, 0x94, 0x67, 0x57, 0xfa, 0x46, 0xf7, 0xe7, 0xe7, 0x1f,
	0x70, 0xf6, 0x1d, 0xdd, 0x44, 0x9c, 0x5d, 0x3a, 0x9b, 0x5d, 0x9f, 0xa2, 0x6e, 0x19, 0x79, 0x01,
	0x5c, 0xf0, 0xe9, 0x65, 0x92, 0xca, 0xaf, 0x87, 0xea, 0xe7, 0xac, 0xf6, 0xfc, 0xcd, 0x75, 0x13,
	0x6e, 0xa9, 0xf5, 0x7f, 0x00, 0xfc, 0x82, 0x9d, 0xb9, 0x3b, 0xfb, 0x0d, 0x7b, 0xcb, 0xba, 0xf5,
	0x17, 0xde, 0x4d, 0xeb, 0x85, 0x6d, 0x25, 0xb7, 0x55, 0xfd, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x30, 0xdb, 0x57, 0x9c, 0x8c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InnerSessionClient is the client API for InnerSession service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerSessionClient interface {
}

type innerSessionClient struct {
	cc *grpc.ClientConn
}

func NewInnerSessionClient(cc *grpc.ClientConn) InnerSessionClient {
	return &innerSessionClient{cc}
}

// InnerSessionServer is the server API for InnerSession service.
type InnerSessionServer interface {
}

// UnimplementedInnerSessionServer can be embedded to have forward compatible implementations.
type UnimplementedInnerSessionServer struct {
}

func RegisterInnerSessionServer(s *grpc.Server, srv InnerSessionServer) {
	s.RegisterService(&_InnerSession_serviceDesc, srv)
}

var _InnerSession_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.message.InnerSession",
	HandlerType: (*InnerSessionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "inner_session.proto",
}