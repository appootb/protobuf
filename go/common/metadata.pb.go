// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

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

// Request metadata.
type Metadata struct {
	Account              *uint64   `protobuf:"varint,1,req,name=account" json:"account,omitempty"`
	Token                *string   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Platform             *Platform `protobuf:"varint,3,req,name=platform,enum=appootb.common.Platform" json:"platform,omitempty"`
	Network              *Network  `protobuf:"varint,4,opt,name=network,enum=appootb.common.Network" json:"network,omitempty"`
	Package              *string   `protobuf:"bytes,5,opt,name=package" json:"package,omitempty"`
	Version              *string   `protobuf:"bytes,6,req,name=version" json:"version,omitempty"`
	OsVersion            *string   `protobuf:"bytes,7,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	Brand                *string   `protobuf:"bytes,8,opt,name=brand" json:"brand,omitempty"`
	Model                *string   `protobuf:"bytes,9,opt,name=model" json:"model,omitempty"`
	DeviceId             *string   `protobuf:"bytes,10,req,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Timestamp            *int64    `protobuf:"varint,11,req,name=timestamp" json:"timestamp,omitempty"`
	IsEmulator           *bool     `protobuf:"varint,12,opt,name=is_emulator,json=isEmulator,def=1" json:"is_emulator,omitempty"`
	Latitude             *string   `protobuf:"bytes,13,opt,name=latitude" json:"latitude,omitempty"`
	Longitude            *string   `protobuf:"bytes,14,opt,name=longitude" json:"longitude,omitempty"`
	Locale               *string   `protobuf:"bytes,15,opt,name=locale" json:"locale,omitempty"`
	ClientIp             *string   `protobuf:"bytes,16,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	Channel              *string   `protobuf:"bytes,17,opt,name=channel" json:"channel,omitempty"`
	Product              *string   `protobuf:"bytes,18,opt,name=product" json:"product,omitempty"`
	TraceId              *string   `protobuf:"bytes,19,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

const Default_Metadata_IsEmulator bool = true

func (m *Metadata) GetAccount() uint64 {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return 0
}

func (m *Metadata) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Metadata) GetPlatform() Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return Platform_PLATFORM_UNSPECIFIED
}

func (m *Metadata) GetNetwork() Network {
	if m != nil && m.Network != nil {
		return *m.Network
	}
	return Network_NETWORK_UNSPECIFIED
}

func (m *Metadata) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *Metadata) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *Metadata) GetOsVersion() string {
	if m != nil && m.OsVersion != nil {
		return *m.OsVersion
	}
	return ""
}

func (m *Metadata) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *Metadata) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Metadata) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Metadata) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Metadata) GetIsEmulator() bool {
	if m != nil && m.IsEmulator != nil {
		return *m.IsEmulator
	}
	return Default_Metadata_IsEmulator
}

func (m *Metadata) GetLatitude() string {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return ""
}

func (m *Metadata) GetLongitude() string {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return ""
}

func (m *Metadata) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *Metadata) GetClientIp() string {
	if m != nil && m.ClientIp != nil {
		return *m.ClientIp
	}
	return ""
}

func (m *Metadata) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *Metadata) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *Metadata) GetTraceId() string {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "appootb.common.Metadata")
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor_56d9f74966f40d04) }

var fileDescriptor_56d9f74966f40d04 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x6c, 0x3f, 0x12, 0x2f, 0x0d, 0x60, 0x2a, 0x30, 0x05, 0xa4, 0x08, 0x69, 0xa5,
	0x9c, 0x5a, 0xb1, 0xe2, 0xc4, 0x8d, 0xae, 0x90, 0xe8, 0x01, 0x54, 0xe5, 0xc0, 0x01, 0xad, 0x54,
	0xb9, 0x8e, 0xb7, 0x6b, 0x35, 0xf1, 0x44, 0xc9, 0x64, 0xf7, 0x7d, 0x38, 0xf2, 0x18, 0x1c, 0x79,
	0x22, 0x8e, 0x28, 0x1e, 0xa7, 0x88, 0x3d, 0xfe, 0xfe, 0xff, 0xf9, 0xf2, 0x78, 0x58, 0x52, 0x69,
	0x94, 0x85, 0x44, 0xb9, 0xac, 0x1b, 0x40, 0xe0, 0x89, 0xac, 0x6b, 0x00, 0xdc, 0x2f, 0x15, 0x54,
	0x15, 0xd8, 0xc5, 0xcc, 0x6a, 0xbc, 0x87, 0xe6, 0x48, 0xf6, 0x22, 0xa9, 0x4b, 0x89, 0x37, 0xd0,
	0x54, 0xc4, 0x6f, 0x7f, 0x8d, 0x58, 0xf4, 0xc5, 0x57, 0xe0, 0x82, 0x4d, 0xa5, 0x52, 0xd0, 0x59,
	0x14, 0x41, 0x1a, 0x66, 0xa3, 0x7c, 0x40, 0x3e, 0x67, 0x63, 0x84, 0xa3, 0xb6, 0x22, 0x4c, 0x83,
	0x2c, 0xce, 0x09, 0xf8, 0x7b, 0x16, 0x0d, 0xe5, 0xc4, 0x59, 0x1a, 0x66, 0xc9, 0xa5, 0x58, 0xfe,
	0xdf, 0x7e, 0xb9, 0xf5, 0x7e, 0x7e, 0x8a, 0xe4, 0xef, 0xd8, 0xd4, 0xcf, 0x24, 0x46, 0x69, 0x90,
	0x25, 0x97, 0x2f, 0x1e, 0x26, 0x7d, 0x25, 0x3b, 0x1f, 0xe2, 0xfa, 0xc1, 0x6a, 0xa9, 0x8e, 0xf2,
	0xa0, 0xc5, 0xd8, 0x0d, 0x30, 0x60, 0xef, 0xdc, 0xe9, 0xa6, 0x35, 0x60, 0xc5, 0x24, 0x0d, 0x7b,
	0xc7, 0x23, 0x7f, 0xc3, 0x18, 0xb4, 0xbb, 0xc1, 0x9c, 0xba, 0xb4, 0x18, 0xda, 0x6f, 0xde, 0x9e,
	0xb3, 0xf1, 0xbe, 0x91, 0xb6, 0x10, 0x11, 0xbd, 0xc8, 0x41, 0xaf, 0x56, 0x50, 0xe8, 0x52, 0xc4,
	0xa4, 0x3a, 0xe0, 0xaf, 0x58, 0x5c, 0xe8, 0x3b, 0xa3, 0xf4, 0xce, 0x14, 0x82, 0xb9, 0x36, 0x11,
	0x09, 0x9b, 0x82, 0xbf, 0x66, 0x31, 0x9a, 0x4a, 0xb7, 0x28, 0xab, 0x5a, 0x9c, 0xa7, 0x61, 0x76,
	0x96, 0xff, 0x13, 0xf8, 0x05, 0x3b, 0x37, 0xed, 0x4e, 0x57, 0x5d, 0x29, 0x11, 0x1a, 0xf1, 0x28,
	0x0d, 0xb2, 0xe8, 0xc3, 0x08, 0x9b, 0x4e, 0xe7, 0xcc, 0xb4, 0x9f, 0xbc, 0xce, 0x17, 0x2c, 0x2a,
	0x25, 0x1a, 0xec, 0x0a, 0x2d, 0x66, 0xae, 0xf5, 0x89, 0xfb, 0x06, 0x25, 0xd8, 0x03, 0x99, 0x09,
	0xbd, 0xe3, 0x24, 0xf0, 0xe7, 0x6c, 0x52, 0x82, 0x92, 0xa5, 0x16, 0x8f, 0x9d, 0xe5, 0xa9, 0x9f,
	0x59, 0x95, 0x46, 0x5b, 0xdc, 0x99, 0x5a, 0x3c, 0xa1, 0x92, 0x24, 0x6c, 0xea, 0x7e, 0x6b, 0xea,
	0x56, 0x5a, 0xab, 0x4b, 0xf1, 0x94, 0xf6, 0xe9, 0xd1, 0x6d, 0xba, 0x81, 0xa2, 0x53, 0x28, 0xb8,
	0xdf, 0x34, 0x21, 0x7f, 0xc9, 0x22, 0x6c, 0x24, 0xed, 0xe0, 0x19, 0x59, 0x8e, 0x37, 0xc5, 0xfa,
	0x9e, 0x71, 0x05, 0xd5, 0x83, 0x5f, 0x5c, 0xcf, 0x86, 0xbb, 0xda, 0xf6, 0x97, 0xf6, 0x39, 0xd8,
	0x06, 0xdf, 0x2f, 0x0e, 0x06, 0x6f, 0x3b, 0x17, 0xb1, 0xf2, 0xd1, 0x2b, 0x77, 0x87, 0xfb, 0xee,
	0x66, 0x75, 0x80, 0x15, 0x65, 0xfe, 0x09, 0x82, 0x1f, 0xe1, 0xd9, 0xd5, 0x76, 0xfd, 0x33, 0x9c,
	0x5c, 0x39, 0xe5, 0x77, 0x38, 0xff, 0x48, 0xe1, 0xd7, 0xae, 0xd8, 0x35, 0xc9, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0xfa, 0x07, 0x48, 0xfd, 0x02, 0x00, 0x00,
}
