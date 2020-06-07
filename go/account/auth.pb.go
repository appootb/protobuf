// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: auth.proto

package account

import (
	context "context"
	captcha "github.com/appootb/protobuf/go/captcha"
	_ "github.com/appootb/protobuf/go/permission"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Login request.
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelateId *string `protobuf:"bytes,1,req,name=relate_id,json=relateId" json:"relate_id,omitempty"` // Relate ID of the account, mobile or email
	// Types that are assignable to Secure:
	//	*LoginRequest_Code
	//	*LoginRequest_Password
	Secure isLoginRequest_Secure `protobuf_oneof:"secure"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetRelateId() string {
	if x != nil && x.RelateId != nil {
		return *x.RelateId
	}
	return ""
}

func (m *LoginRequest) GetSecure() isLoginRequest_Secure {
	if m != nil {
		return m.Secure
	}
	return nil
}

func (x *LoginRequest) GetCode() string {
	if x, ok := x.GetSecure().(*LoginRequest_Code); ok {
		return x.Code
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x, ok := x.GetSecure().(*LoginRequest_Password); ok {
		return x.Password
	}
	return ""
}

type isLoginRequest_Secure interface {
	isLoginRequest_Secure()
}

type LoginRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,oneof"` // Random verification code
}

type LoginRequest_Password struct {
	Password string `protobuf:"bytes,3,opt,name=password,oneof"` // Account password
}

func (*LoginRequest_Code) isLoginRequest_Secure() {}

func (*LoginRequest_Password) isLoginRequest_Secure() {}

// OAuth request.
type OAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        *AuthType `protobuf:"varint,1,req,name=type,enum=appootb.account.AuthType" json:"type,omitempty"`   // Auth type
	OpenId      *string   `protobuf:"bytes,2,req,name=open_id,json=openId" json:"open_id,omitempty"`                // Authentication ID
	AccessToken *string   `protobuf:"bytes,3,req,name=access_token,json=accessToken" json:"access_token,omitempty"` // Authorization token
}

func (x *OAuthRequest) Reset() {
	*x = OAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthRequest) ProtoMessage() {}

func (x *OAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthRequest.ProtoReflect.Descriptor instead.
func (*OAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthRequest) GetType() AuthType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return AuthType_AUTH_TYPE_UNSPECIFIED
}

func (x *OAuthRequest) GetOpenId() string {
	if x != nil && x.OpenId != nil {
		return *x.OpenId
	}
	return ""
}

func (x *OAuthRequest) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

// Region.
type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`   // Region name
	Local *string `protobuf:"bytes,2,req,name=local" json:"local,omitempty"` // Local name
	Abbr  *string `protobuf:"bytes,3,req,name=abbr" json:"abbr,omitempty"`   // Abbreviation of region name
	Code  *string `protobuf:"bytes,4,req,name=code" json:"code,omitempty"`   // Region code
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Region) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Region) GetLocal() string {
	if x != nil && x.Local != nil {
		return *x.Local
	}
	return ""
}

func (x *Region) GetAbbr() string {
	if x != nil && x.Abbr != nil {
		return *x.Abbr
	}
	return ""
}

func (x *Region) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

// Regions.
type Regions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []*Region `protobuf:"bytes,1,rep,name=regions" json:"regions,omitempty"` // Regions
}

func (x *Regions) Reset() {
	*x = Regions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regions) ProtoMessage() {}

func (x *Regions) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regions.ProtoReflect.Descriptor instead.
func (*Regions) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Regions) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70,
	0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x61,
	0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70,
	0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x22, 0x79, 0x0a, 0x0c, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x62, 0x62, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x62, 0x62, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3c, 0x0a, 0x07, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xf8, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x5b,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x6f, 0x74, 0x62, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x64, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01,
	0x2a, 0x12, 0x64, 0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22,
	0x13, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e,
	0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12,
	0x15, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x68, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x6f, 0x74, 0x62, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0xda, 0x9c, 0x01, 0x03, 0x08, 0xc8, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x3a, 0x01, 0x2a,
	0x42, 0x77, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x01, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x42, 0xaa, 0x02, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0xca, 0x02, 0x15, 0x41, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x5c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),        // 0: appootb.account.LoginRequest
	(*OAuthRequest)(nil),        // 1: appootb.account.OAuthRequest
	(*Region)(nil),              // 2: appootb.account.Region
	(*Regions)(nil),             // 3: appootb.account.Regions
	(AuthType)(0),               // 4: appootb.account.AuthType
	(*captcha.CodeRequest)(nil), // 5: appootb.captcha.CodeRequest
	(*empty.Empty)(nil),         // 6: google.protobuf.Empty
	(*AccountInfo)(nil),         // 7: appootb.account.AccountInfo
}
var file_auth_proto_depIdxs = []int32{
	4, // 0: appootb.account.OAuthRequest.type:type_name -> appootb.account.AuthType
	2, // 1: appootb.account.Regions.regions:type_name -> appootb.account.Region
	5, // 2: appootb.account.Auth.GetCode:input_type -> appootb.captcha.CodeRequest
	0, // 3: appootb.account.Auth.Login:input_type -> appootb.account.LoginRequest
	1, // 4: appootb.account.Auth.OAuth:input_type -> appootb.account.OAuthRequest
	6, // 5: appootb.account.Auth.GetRegions:input_type -> google.protobuf.Empty
	6, // 6: appootb.account.Auth.Refresh:input_type -> google.protobuf.Empty
	6, // 7: appootb.account.Auth.GetCode:output_type -> google.protobuf.Empty
	7, // 8: appootb.account.Auth.Login:output_type -> appootb.account.AccountInfo
	7, // 9: appootb.account.Auth.OAuth:output_type -> appootb.account.AccountInfo
	3, // 10: appootb.account.Auth.GetRegions:output_type -> appootb.account.Regions
	7, // 11: appootb.account.Auth.Refresh:output_type -> appootb.account.AccountInfo
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthRequest); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regions); i {
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
	file_auth_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LoginRequest_Code)(nil),
		(*LoginRequest_Password)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// Get auth code.
	GetCode(ctx context.Context, in *captcha.CodeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Login.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccountInfo, error)
	// OAuth.
	OAuth(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AccountInfo, error)
	// Get regions.
	GetRegions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Regions, error)
	// Refresh account token.
	Refresh(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountInfo, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetCode(ctx context.Context, in *captcha.CodeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appootb.account.Auth/GetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/appootb.account.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) OAuth(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/appootb.account.Auth/OAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRegions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Regions, error) {
	out := new(Regions)
	err := c.cc.Invoke(ctx, "/appootb.account.Auth/GetRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/appootb.account.Auth/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// Get auth code.
	GetCode(context.Context, *captcha.CodeRequest) (*empty.Empty, error)
	// Login.
	Login(context.Context, *LoginRequest) (*AccountInfo, error)
	// OAuth.
	OAuth(context.Context, *OAuthRequest) (*AccountInfo, error)
	// Get regions.
	GetRegions(context.Context, *empty.Empty) (*Regions, error)
	// Refresh account token.
	Refresh(context.Context, *empty.Empty) (*AccountInfo, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) GetCode(context.Context, *captcha.CodeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (*UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) OAuth(context.Context, *OAuthRequest) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuth not implemented")
}
func (*UnimplementedAuthServer) GetRegions(context.Context, *empty.Empty) (*Regions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegions not implemented")
}
func (*UnimplementedAuthServer) Refresh(context.Context, *empty.Empty) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(captcha.CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Auth/GetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetCode(ctx, req.(*captcha.CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_OAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).OAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Auth/OAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).OAuth(ctx, req.(*OAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Auth/GetRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetRegions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appootb.account.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appootb.account.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCode",
			Handler:    _Auth_GetCode_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "OAuth",
			Handler:    _Auth_OAuth_Handler,
		},
		{
			MethodName: "GetRegions",
			Handler:    _Auth_GetRegions_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
