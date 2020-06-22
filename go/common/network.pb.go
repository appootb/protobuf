// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Network enums.
type Network int32

const (
	Network_NETWORK_UNSPECIFIED Network = 0
	Network_NETWORK_ETHERNET    Network = 1
	Network_NETWORK_WIFI        Network = 2
	Network_NETWORK_CELLULAR    Network = 3
)

var Network_name = map[int32]string{
	0: "NETWORK_UNSPECIFIED",
	1: "NETWORK_ETHERNET",
	2: "NETWORK_WIFI",
	3: "NETWORK_CELLULAR",
}

var Network_value = map[string]int32{
	"NETWORK_UNSPECIFIED": 0,
	"NETWORK_ETHERNET":    1,
	"NETWORK_WIFI":        2,
	"NETWORK_CELLULAR":    3,
}

func (x Network) Enum() *Network {
	p := new(Network)
	*p = x
	return p
}

func (x Network) String() string {
	return proto.EnumName(Network_name, int32(x))
}

func (x *Network) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Network_value, data, "Network")
	if err != nil {
		return err
	}
	*x = Network(value)
	return nil
}

func (Network) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func init() {
	proto.RegisterEnum("appootb.common.Network", Network_name, Network_value)
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x2c, 0x28, 0xc8, 0xcf,
	0x2f, 0x49, 0xd2, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x4a, 0xe0, 0x62, 0xf7, 0x83, 0x28,
	0x10, 0x12, 0xe7, 0x12, 0xf6, 0x73, 0x0d, 0x09, 0xf7, 0x0f, 0xf2, 0x8e, 0x0f, 0xf5, 0x0b, 0x0e,
	0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0x12, 0xe1, 0x12, 0x80, 0x49, 0xb8,
	0x86, 0x78, 0xb8, 0x06, 0xf9, 0xb9, 0x86, 0x08, 0x30, 0x0a, 0x09, 0x70, 0xf1, 0xc0, 0x44, 0xc3,
	0x3d, 0xdd, 0x3c, 0x05, 0x98, 0x90, 0xd5, 0x39, 0xbb, 0xfa, 0xf8, 0x84, 0xfa, 0x38, 0x06, 0x09,
	0x30, 0x3b, 0x95, 0x71, 0x09, 0x25, 0xe7, 0xe7, 0xea, 0xa1, 0xda, 0xeb, 0xc4, 0x03, 0xb5, 0x35,
	0x00, 0xe4, 0x2a, 0x0f, 0xc6, 0x00, 0xc6, 0x28, 0xd5, 0xf4, 0xcc, 0x92, 0x8c, 0x52, 0xb0, 0x02,
	0x7d, 0xa8, 0x62, 0x7d, 0xb0, 0x9b, 0x93, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xf5, 0x21, 0x1a, 0x7f,
	0x30, 0x32, 0x2e, 0x62, 0x62, 0x76, 0x0e, 0x70, 0x5a, 0xc5, 0xc4, 0xe6, 0x0c, 0x16, 0x39, 0xc5,
	0x24, 0xe2, 0x08, 0x51, 0x1e, 0x03, 0x36, 0x2c, 0x06, 0x22, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x6d, 0xa1, 0x98, 0xf9, 0x00, 0x00, 0x00,
}
