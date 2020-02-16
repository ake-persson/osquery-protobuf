// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/interface_address.proto

package schemas

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

// Network interfaces and relevant metadata.
type InterfaceAddress struct {
	// Interface name
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface"`
	// Specific address for interface
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	// Interface netmask
	Mask string `protobuf:"bytes,3,opt,name=mask,proto3" json:"mask"`
	// Broadcast address for the interface
	Broadcast string `protobuf:"bytes,4,opt,name=broadcast,proto3" json:"broadcast"`
	// PtP address for the interface
	PointToPoint string `protobuf:"bytes,5,opt,name=point_to_point,json=pointToPoint,proto3" json:"point_to_point"`
	// Type of address. One of dhcp
	Type string `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	// The friendly display name of the interface.
	FriendlyName         string   `protobuf:"bytes,7,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceAddress) Reset()         { *m = InterfaceAddress{} }
func (m *InterfaceAddress) String() string { return proto.CompactTextString(m) }
func (*InterfaceAddress) ProtoMessage()    {}
func (*InterfaceAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_089f50b4763f328a, []int{0}
}

func (m *InterfaceAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceAddress.Unmarshal(m, b)
}
func (m *InterfaceAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceAddress.Marshal(b, m, deterministic)
}
func (m *InterfaceAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceAddress.Merge(m, src)
}
func (m *InterfaceAddress) XXX_Size() int {
	return xxx_messageInfo_InterfaceAddress.Size(m)
}
func (m *InterfaceAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceAddress.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceAddress proto.InternalMessageInfo

func (m *InterfaceAddress) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *InterfaceAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InterfaceAddress) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *InterfaceAddress) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *InterfaceAddress) GetPointToPoint() string {
	if m != nil {
		return m.PointToPoint
	}
	return ""
}

func (m *InterfaceAddress) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InterfaceAddress) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func init() {
	proto.RegisterType((*InterfaceAddress)(nil), "schemas.InterfaceAddress")
}

func init() { proto.RegisterFile("pb/interface_address.proto", fileDescriptor_089f50b4763f328a) }

var fileDescriptor_089f50b4763f328a = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x28, 0x8d, 0x6a, 0x01, 0x42, 0x9e, 0x2c, 0xc4, 0x80, 0x80, 0x81, 0x85, 0x06,
	0x81, 0x04, 0x03, 0x13, 0x6c, 0x2c, 0x08, 0x21, 0x26, 0x96, 0xe8, 0xec, 0x5c, 0x5a, 0x2b, 0x38,
	0x67, 0x6c, 0x67, 0xc8, 0xb3, 0xf2, 0x32, 0x28, 0x97, 0x26, 0x4c, 0xfe, 0xef, 0xf3, 0x7f, 0xdf,
	0x70, 0xe2, 0xd4, 0xeb, 0xc2, 0xb6, 0x09, 0x43, 0x0d, 0x06, 0x4b, 0xa8, 0xaa, 0x80, 0x31, 0xae,
	0x7d, 0xa0, 0x44, 0x32, 0x8f, 0x66, 0x8b, 0x0e, 0xe2, 0xc5, 0x6f, 0x26, 0x4e, 0x5e, 0xa7, 0xd2,
	0xf3, 0xd8, 0x91, 0x67, 0x62, 0x35, 0x2f, 0xaa, 0xec, 0x3c, 0xbb, 0x5e, 0x7d, 0xfc, 0x03, 0xa9,
	0x44, 0xbe, 0x93, 0xa9, 0x3d, 0xfe, 0x9b, 0x46, 0x29, 0xc5, 0xc2, 0x41, 0x6c, 0xd4, 0x3e, 0x63,
	0xce, 0x83, 0x4b, 0x07, 0x82, 0xca, 0x40, 0x4c, 0x6a, 0x31, 0xba, 0x66, 0x20, 0xaf, 0xc4, 0xb1,
	0x27, 0xdb, 0xa6, 0x32, 0x51, 0xc9, 0x41, 0x1d, 0x70, 0xe5, 0x90, 0x87, 0x4f, 0x7a, 0x1f, 0x9e,
	0xc1, 0x9b, 0x7a, 0x8f, 0x6a, 0x39, 0x7a, 0x87, 0x2c, 0x2f, 0xc5, 0x51, 0x1d, 0x2c, 0xb6, 0xd5,
	0x77, 0x5f, 0xb6, 0xe0, 0x50, 0xe5, 0xe3, 0xe2, 0x04, 0xdf, 0xc0, 0xe1, 0xcb, 0xdd, 0xd7, 0xed,
	0xc6, 0xa6, 0x6d, 0xa7, 0xd7, 0x86, 0x5c, 0xe1, 0xac, 0x69, 0xd0, 0x3f, 0x3e, 0x14, 0x14, 0x7f,
	0x3a, 0x0c, 0xfd, 0x0d, 0xdf, 0x42, 0x77, 0x75, 0xe1, 0x9b, 0xcd, 0xd3, 0xee, 0x22, 0x7a, 0xc9,
	0xf4, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x76, 0xb4, 0x95, 0x3f, 0x01, 0x00, 0x00,
}
