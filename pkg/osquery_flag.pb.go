// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/osquery_flag.proto

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

// Configurable flags that modify osquery's behavior.
type OsqueryFlag struct {
	// Flag name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Flag type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// Flag description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// Flag default value
	DefaultValue string `protobuf:"bytes,4,opt,name=default_value,json=defaultValue,proto3" json:"default_value"`
	// Flag value
	Value string `protobuf:"bytes,5,opt,name=value,proto3" json:"value"`
	// Is the flag shell only?
	ShellOnly            int32    `protobuf:"varint,6,opt,name=shell_only,json=shellOnly,proto3" json:"shell_only"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OsqueryFlag) Reset()         { *m = OsqueryFlag{} }
func (m *OsqueryFlag) String() string { return proto.CompactTextString(m) }
func (*OsqueryFlag) ProtoMessage()    {}
func (*OsqueryFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef3724eaeec4be1, []int{0}
}

func (m *OsqueryFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OsqueryFlag.Unmarshal(m, b)
}
func (m *OsqueryFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OsqueryFlag.Marshal(b, m, deterministic)
}
func (m *OsqueryFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsqueryFlag.Merge(m, src)
}
func (m *OsqueryFlag) XXX_Size() int {
	return xxx_messageInfo_OsqueryFlag.Size(m)
}
func (m *OsqueryFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_OsqueryFlag.DiscardUnknown(m)
}

var xxx_messageInfo_OsqueryFlag proto.InternalMessageInfo

func (m *OsqueryFlag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OsqueryFlag) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OsqueryFlag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OsqueryFlag) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *OsqueryFlag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *OsqueryFlag) GetShellOnly() int32 {
	if m != nil {
		return m.ShellOnly
	}
	return 0
}

func init() {
	proto.RegisterType((*OsqueryFlag)(nil), "schemas.OsqueryFlag")
}

func init() { proto.RegisterFile("pb/osquery_flag.proto", fileDescriptor_7ef3724eaeec4be1) }

var fileDescriptor_7ef3724eaeec4be1 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0xc6, 0x71, 0x05, 0x9a, 0xa2, 0x5e, 0x61, 0xb1, 0x40, 0xf2, 0x82, 0x14, 0xc1, 0xd2, 0x85,
	0x06, 0x81, 0x04, 0x03, 0x1b, 0x03, 0x6b, 0xa5, 0x0e, 0x0c, 0x2c, 0x91, 0x93, 0x5e, 0x5e, 0xd4,
	0xf3, 0x0b, 0xb1, 0x8d, 0xe4, 0x8f, 0xc5, 0x37, 0x44, 0xb1, 0x83, 0xd4, 0xed, 0xfc, 0xfb, 0x7b,
	0x79, 0xe0, 0xc6, 0xd4, 0xa5, 0xb6, 0xdf, 0x1e, 0xc7, 0x50, 0xb5, 0x24, 0xba, 0xad, 0x19, 0xb5,
	0xd3, 0xec, 0xc2, 0x36, 0x3d, 0x4a, 0x61, 0xef, 0x7e, 0x33, 0x58, 0xef, 0x52, 0xff, 0x20, 0xd1,
	0x31, 0x06, 0x0b, 0x25, 0x24, 0xf2, 0xac, 0xc8, 0x36, 0xab, 0x7d, 0xbc, 0x27, 0x73, 0xc1, 0x20,
	0x3f, 0x4b, 0x36, 0xdd, 0xac, 0x80, 0xf5, 0x01, 0x6d, 0x33, 0x0e, 0xc6, 0x0d, 0x5a, 0xf1, 0xf3,
	0x98, 0x4e, 0x89, 0xdd, 0xc3, 0xd5, 0x01, 0x5b, 0xe1, 0xc9, 0x55, 0x3f, 0x82, 0x3c, 0xf2, 0x45,
	0xfc, 0x73, 0x39, 0xe3, 0xe7, 0x64, 0xec, 0x1a, 0xf2, 0x14, 0xf3, 0x18, 0xd3, 0x83, 0xdd, 0x02,
	0xd8, 0x1e, 0x89, 0x2a, 0xad, 0x28, 0xf0, 0x65, 0x91, 0x6d, 0xf2, 0xfd, 0x2a, 0xca, 0x4e, 0x51,
	0x78, 0x7f, 0xfa, 0x7a, 0xec, 0x06, 0xd7, 0xfb, 0x7a, 0xdb, 0x68, 0x59, 0xca, 0xa1, 0x39, 0xa2,
	0x79, 0x7d, 0xf9, 0x9f, 0xf9, 0x10, 0x17, 0xd6, 0xbe, 0x2d, 0xcd, 0xb1, 0x7b, 0x9b, 0x77, 0xd6,
	0xcb, 0xa8, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x29, 0xef, 0x2b, 0x10, 0x01, 0x00,
	0x00,
}
