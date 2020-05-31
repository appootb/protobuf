// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: method.proto

package permission

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Method token level.
type TokenLevel int32

const (
	TokenLevel_NONE_TOKEN   TokenLevel = 0   // Public access
	TokenLevel_GUEST_TOKEN  TokenLevel = 100 // Guest token
	TokenLevel_LOW_TOKEN    TokenLevel = 200 // Lower token, for h5, web
	TokenLevel_MIDDLE_TOKEN TokenLevel = 300 // Middle token, for pc
	TokenLevel_HIGH_TOKEN   TokenLevel = 400 // High token, for native app
	TokenLevel_INNER_TOKEN  TokenLevel = 999 // For server internal
)

// Enum value maps for TokenLevel.
var (
	TokenLevel_name = map[int32]string{
		0:   "NONE_TOKEN",
		100: "GUEST_TOKEN",
		200: "LOW_TOKEN",
		300: "MIDDLE_TOKEN",
		400: "HIGH_TOKEN",
		999: "INNER_TOKEN",
	}
	TokenLevel_value = map[string]int32{
		"NONE_TOKEN":   0,
		"GUEST_TOKEN":  100,
		"LOW_TOKEN":    200,
		"MIDDLE_TOKEN": 300,
		"HIGH_TOKEN":   400,
		"INNER_TOKEN":  999,
	}
)

func (x TokenLevel) Enum() *TokenLevel {
	p := new(TokenLevel)
	*p = x
	return p
}

func (x TokenLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_method_proto_enumTypes[0].Descriptor()
}

func (TokenLevel) Type() protoreflect.EnumType {
	return &file_method_proto_enumTypes[0]
}

func (x TokenLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenLevel.Descriptor instead.
func (TokenLevel) EnumDescriptor() ([]byte, []int) {
	return file_method_proto_rawDescGZIP(), []int{0}
}

// Method token.
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required TokenLevel `protobuf:"varint,1,opt,name=required,proto3,enum=appootb.permission.method.TokenLevel" json:"required,omitempty"` // Token level
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_method_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetRequired() TokenLevel {
	if x != nil {
		return x.Required
	}
	return TokenLevel_NONE_TOKEN
}

var file_method_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Token)(nil),
		Field:         2507,
		Name:          "appootb.permission.method.token",
		Tag:           "bytes,2507,opt,name=token",
		Filename:      "method.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for grpc-gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional appootb.permission.method.Token token = 2507;
	E_Token = &file_method_proto_extTypes[0] // Method option
)

var File_method_proto protoreflect.FileDescriptor

var file_method_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2a, 0x73, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x64, 0x12, 0x0e, 0x0a, 0x09, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x10, 0xc8, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x4d, 0x49, 0x44, 0x44, 0x4c, 0x45,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xac, 0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x48, 0x49, 0x47,
	0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x90, 0x03, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x4e,
	0x4e, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xe7, 0x07, 0x3a, 0x57, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcb, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x7d, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01, 0x50, 0x01,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x70,
	0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x43, 0x50, 0x42, 0xaa, 0x02, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0xca, 0x02, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_method_proto_rawDescOnce sync.Once
	file_method_proto_rawDescData = file_method_proto_rawDesc
)

func file_method_proto_rawDescGZIP() []byte {
	file_method_proto_rawDescOnce.Do(func() {
		file_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_method_proto_rawDescData)
	})
	return file_method_proto_rawDescData
}

var file_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_method_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_method_proto_goTypes = []interface{}{
	(TokenLevel)(0),                  // 0: appootb.permission.method.TokenLevel
	(*Token)(nil),                    // 1: appootb.permission.method.Token
	(*descriptor.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_method_proto_depIdxs = []int32{
	0, // 0: appootb.permission.method.Token.required:type_name -> appootb.permission.method.TokenLevel
	2, // 1: appootb.permission.method.token:extendee -> google.protobuf.MethodOptions
	1, // 2: appootb.permission.method.token:type_name -> appootb.permission.method.Token
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_method_proto_init() }
func file_method_proto_init() {
	if File_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_method_proto_goTypes,
		DependencyIndexes: file_method_proto_depIdxs,
		EnumInfos:         file_method_proto_enumTypes,
		MessageInfos:      file_method_proto_msgTypes,
		ExtensionInfos:    file_method_proto_extTypes,
	}.Build()
	File_method_proto = out.File
	file_method_proto_rawDesc = nil
	file_method_proto_goTypes = nil
	file_method_proto_depIdxs = nil
}
