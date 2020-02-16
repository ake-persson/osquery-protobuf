// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/xprotect_metum.proto

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

// Database of the machine's XProtect browser-related signatures.
type XprotectMetum struct {
	// Browser plugin or extension identifier
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier"`
	// Either plugin or extension
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// Developer identity (SHA1) of extension
	DeveloperId string `protobuf:"bytes,3,opt,name=developer_id,json=developerId,proto3" json:"developer_id"`
	// The minimum allowed plugin version.
	MinVersion           string   `protobuf:"bytes,4,opt,name=min_version,json=minVersion,proto3" json:"min_version"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XprotectMetum) Reset()         { *m = XprotectMetum{} }
func (m *XprotectMetum) String() string { return proto.CompactTextString(m) }
func (*XprotectMetum) ProtoMessage()    {}
func (*XprotectMetum) Descriptor() ([]byte, []int) {
	return fileDescriptor_728522e3519474eb, []int{0}
}

func (m *XprotectMetum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XprotectMetum.Unmarshal(m, b)
}
func (m *XprotectMetum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XprotectMetum.Marshal(b, m, deterministic)
}
func (m *XprotectMetum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XprotectMetum.Merge(m, src)
}
func (m *XprotectMetum) XXX_Size() int {
	return xxx_messageInfo_XprotectMetum.Size(m)
}
func (m *XprotectMetum) XXX_DiscardUnknown() {
	xxx_messageInfo_XprotectMetum.DiscardUnknown(m)
}

var xxx_messageInfo_XprotectMetum proto.InternalMessageInfo

func (m *XprotectMetum) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *XprotectMetum) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *XprotectMetum) GetDeveloperId() string {
	if m != nil {
		return m.DeveloperId
	}
	return ""
}

func (m *XprotectMetum) GetMinVersion() string {
	if m != nil {
		return m.MinVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*XprotectMetum)(nil), "schemas.XprotectMetum")
}

func init() { proto.RegisterFile("pb/xprotect_metum.proto", fileDescriptor_728522e3519474eb) }

var fileDescriptor_728522e3519474eb = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x48, 0xd2, 0xaf,
	0x28, 0x28, 0xca, 0x2f, 0x49, 0x4d, 0x2e, 0x89, 0xcf, 0x4d, 0x2d, 0x29, 0xcd, 0xd5, 0x03, 0xf1,
	0xf2, 0x85, 0xd8, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0x8b, 0x95, 0xda, 0x19, 0xb9, 0x78, 0x23,
	0xa0, 0x2a, 0x7c, 0x41, 0x0a, 0x84, 0xe4, 0xb8, 0xb8, 0x32, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0xd3,
	0x32, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x44, 0x84, 0x84, 0xb8, 0x58,
	0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0xc0, 0x32, 0x60, 0xb6, 0x90, 0x22, 0x17, 0x4f, 0x4a, 0x6a,
	0x59, 0x6a, 0x4e, 0x7e, 0x41, 0x6a, 0x51, 0x7c, 0x66, 0x8a, 0x04, 0x33, 0x58, 0x8e, 0x1b, 0x2e,
	0xe6, 0x99, 0x22, 0x24, 0xcf, 0xc5, 0x9d, 0x9b, 0x99, 0x17, 0x5f, 0x96, 0x5a, 0x54, 0x9c, 0x99,
	0x9f, 0x27, 0xc1, 0x02, 0x31, 0x37, 0x37, 0x33, 0x2f, 0x0c, 0x22, 0xe2, 0x64, 0x14, 0x65, 0x90,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9b, 0x99, 0x9c, 0x9d, 0x5a,
	0x60, 0x6e, 0xa6, 0x9f, 0x5f, 0x5c, 0x58, 0x9a, 0x5a, 0x54, 0xa9, 0x0b, 0x76, 0x77, 0x52, 0x69,
	0x9a, 0x7e, 0x41, 0x76, 0xba, 0x35, 0xd4, 0xf5, 0x49, 0x6c, 0x60, 0x51, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0xc1, 0xbd, 0x89, 0xe8, 0x00, 0x00, 0x00,
}
