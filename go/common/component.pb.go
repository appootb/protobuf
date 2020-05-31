// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: component.proto

package common

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

// Component enums.
type Component int32

const (
	Component_COMPONENT_UNSPECIFIED Component = 0 // Unspecified
	Component_COMPONENT_CAPTCHA     Component = 1 // Captcha
	Component_COMPONENT_ACCOUNT     Component = 2 // Account
	Component_COMPONENT_STRAINER    Component = 3 // Strainer
)

// Enum value maps for Component.
var (
	Component_name = map[int32]string{
		0: "COMPONENT_UNSPECIFIED",
		1: "COMPONENT_CAPTCHA",
		2: "COMPONENT_ACCOUNT",
		3: "COMPONENT_STRAINER",
	}
	Component_value = map[string]int32{
		"COMPONENT_UNSPECIFIED": 0,
		"COMPONENT_CAPTCHA":     1,
		"COMPONENT_ACCOUNT":     2,
		"COMPONENT_STRAINER":    3,
	}
)

func (x Component) Enum() *Component {
	p := new(Component)
	*p = x
	return p
}

func (x Component) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Component) Descriptor() protoreflect.EnumDescriptor {
	return file_component_proto_enumTypes[0].Descriptor()
}

func (Component) Type() protoreflect.EnumType {
	return &file_component_proto_enumTypes[0]
}

func (x Component) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Component) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Component(num)
	return nil
}

// Deprecated: Use Component.Descriptor instead.
func (Component) EnumDescriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{0}
}

var File_component_proto protoreflect.FileDescriptor

var file_component_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2a, 0x6c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4d,
	0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x43, 0x48, 0x41, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x50, 0x4f,
	0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x42,
	0x53, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x6f, 0x74, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01, 0x50, 0x01, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x42, 0xaa, 0x02, 0x06, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e,
}

var (
	file_component_proto_rawDescOnce sync.Once
	file_component_proto_rawDescData = file_component_proto_rawDesc
)

func file_component_proto_rawDescGZIP() []byte {
	file_component_proto_rawDescOnce.Do(func() {
		file_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_component_proto_rawDescData)
	})
	return file_component_proto_rawDescData
}

var file_component_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_component_proto_goTypes = []interface{}{
	(Component)(0), // 0: appootb.common.Component
}
var file_component_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_component_proto_init() }
func file_component_proto_init() {
	if File_component_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_component_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_component_proto_goTypes,
		DependencyIndexes: file_component_proto_depIdxs,
		EnumInfos:         file_component_proto_enumTypes,
	}.Build()
	File_component_proto = out.File
	file_component_proto_rawDesc = nil
	file_component_proto_goTypes = nil
	file_component_proto_depIdxs = nil
}
