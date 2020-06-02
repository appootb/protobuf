// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: common.proto

package captcha

import (
	proto "github.com/golang/protobuf/proto"
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

// Captcha code channel enums.
type Channel int32

const (
	Channel_CHANNEL_UNSPECIFIED Channel = 0 // Unspecified
	Channel_CHANNEL_EMAIL       Channel = 1 // Email
	Channel_CHANNEL_SMS         Channel = 2 // Mobile SMS
	Channel_CHANNEL_PHONE       Channel = 3 // Phone call
	Channel_CHANNEL_OTP         Channel = 4 // One-Time password
)

// Enum value maps for Channel.
var (
	Channel_name = map[int32]string{
		0: "CHANNEL_UNSPECIFIED",
		1: "CHANNEL_EMAIL",
		2: "CHANNEL_SMS",
		3: "CHANNEL_PHONE",
		4: "CHANNEL_OTP",
	}
	Channel_value = map[string]int32{
		"CHANNEL_UNSPECIFIED": 0,
		"CHANNEL_EMAIL":       1,
		"CHANNEL_SMS":         2,
		"CHANNEL_PHONE":       3,
		"CHANNEL_OTP":         4,
	}
)

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Channel) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Channel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Channel(num)
	return nil
}

// Deprecated: Use Channel.Descriptor instead.
func (Channel) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Captcha code category enums.
type Category int32

const (
	Category_CATEGORY_UNSPECIFIED Category = 0 // Unspecified
	Category_CATEGORY_REGISTER    Category = 1 // Register
	Category_CATEGORY_LOGIN       Category = 2 // Login
	Category_CATEGORY_RESET_PWD   Category = 3 // Reset password
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "CATEGORY_UNSPECIFIED",
		1: "CATEGORY_REGISTER",
		2: "CATEGORY_LOGIN",
		3: "CATEGORY_RESET_PWD",
	}
	Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED": 0,
		"CATEGORY_REGISTER":    1,
		"CATEGORY_LOGIN":       2,
		"CATEGORY_RESET_PWD":   3,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Category) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Category(num)
	return nil
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

// Captcha code request.
type CodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel  *Channel  `protobuf:"varint,1,req,name=channel,enum=appootb.captcha.Channel" json:"channel,omitempty"`    // Code channel
	Category *Category `protobuf:"varint,2,req,name=category,enum=appootb.captcha.Category" json:"category,omitempty"` // Code category
	Target   *string   `protobuf:"bytes,3,req,name=target" json:"target,omitempty"`                                    // Target, e.g. email address, phone number
	Value    *string   `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`                                      // Code value for verification
}

func (x *CodeRequest) Reset() {
	*x = CodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRequest) ProtoMessage() {}

func (x *CodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRequest.ProtoReflect.Descriptor instead.
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *CodeRequest) GetChannel() Channel {
	if x != nil && x.Channel != nil {
		return *x.Channel
	}
	return Channel_CHANNEL_UNSPECIFIED
}

func (x *CodeRequest) GetCategory() Category {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return Category_CATEGORY_UNSPECIFIED
}

func (x *CodeRequest) GetTarget() string {
	if x != nil && x.Target != nil {
		return *x.Target
	}
	return ""
}

func (x *CodeRequest) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22,
	0xa6, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x6a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x4d, 0x53, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x50, 0x48, 0x4f, 0x4e,
	0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4f,
	0x54, 0x50, 0x10, 0x04, 0x2a, 0x67, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x57, 0x44, 0x10, 0x03, 0x42, 0x71, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x48, 0x01, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x43, 0x50, 0x42, 0xaa, 0x02, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0xca, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_proto_goTypes = []interface{}{
	(Channel)(0),        // 0: appootb.captcha.Channel
	(Category)(0),       // 1: appootb.captcha.Category
	(*CodeRequest)(nil), // 2: appootb.captcha.CodeRequest
}
var file_common_proto_depIdxs = []int32{
	0, // 0: appootb.captcha.CodeRequest.channel:type_name -> appootb.captcha.Channel
	1, // 1: appootb.captcha.CodeRequest.category:type_name -> appootb.captcha.Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRequest); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}