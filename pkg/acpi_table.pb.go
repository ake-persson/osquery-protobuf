// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/acpi_table.proto

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

// Firmware ACPI functional table common metadata and content.
type AcpiTable struct {
	// ACPI table name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Size of compiled table data
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size"`
	// MD5 hash of table content
	Md5                  string   `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcpiTable) Reset()         { *m = AcpiTable{} }
func (m *AcpiTable) String() string { return proto.CompactTextString(m) }
func (*AcpiTable) ProtoMessage()    {}
func (*AcpiTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bf045d391bff6d, []int{0}
}

func (m *AcpiTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcpiTable.Unmarshal(m, b)
}
func (m *AcpiTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcpiTable.Marshal(b, m, deterministic)
}
func (m *AcpiTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcpiTable.Merge(m, src)
}
func (m *AcpiTable) XXX_Size() int {
	return xxx_messageInfo_AcpiTable.Size(m)
}
func (m *AcpiTable) XXX_DiscardUnknown() {
	xxx_messageInfo_AcpiTable.DiscardUnknown(m)
}

var xxx_messageInfo_AcpiTable proto.InternalMessageInfo

func (m *AcpiTable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AcpiTable) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AcpiTable) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func init() {
	proto.RegisterType((*AcpiTable)(nil), "schemas.AcpiTable")
}

func init() { proto.RegisterFile("pb/acpi_table.proto", fileDescriptor_d7bf045d391bff6d) }

var fileDescriptor_d7bf045d391bff6d = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xd2, 0x4f,
	0x4c, 0x2e, 0xc8, 0x8c, 0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0x72, 0xe5, 0xe2, 0x74, 0x4c, 0x2e, 0xc8,
	0x0c, 0x01, 0xc9, 0x09, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0xe2, 0xcc, 0xaa, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0x37, 0xc5, 0x54, 0x82, 0x19, 0xac, 0x0c, 0xc4, 0x74,
	0x32, 0x8a, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd,
	0x4c, 0xce, 0x4e, 0x2d, 0x30, 0x37, 0xd3, 0xcf, 0x2f, 0x2e, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x05,
	0x5b, 0x9a, 0x54, 0x9a, 0xa6, 0x5f, 0x90, 0x9d, 0x6e, 0x0d, 0xb5, 0x3a, 0x89, 0x0d, 0x2c, 0x6a,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0xbf, 0xa3, 0x8c, 0xa1, 0x00, 0x00, 0x00,
}