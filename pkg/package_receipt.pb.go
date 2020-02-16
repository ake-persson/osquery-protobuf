// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/package_receipt.proto

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

// OS X package receipt details.
type PackageReceipt struct {
	// Package domain identifier
	PackageId string `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id"`
	// Installed package version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version"`
	// Optional relative install path on volume
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location"`
	// Timestamp of install time
	InstallTime float64 `protobuf:"fixed64,4,opt,name=install_time,json=installTime,proto3" json:"install_time"`
	// Name of installer process
	InstallerName string `protobuf:"bytes,5,opt,name=installer_name,json=installerName,proto3" json:"installer_name"`
	// Path of receipt plist
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageReceipt) Reset()         { *m = PackageReceipt{} }
func (m *PackageReceipt) String() string { return proto.CompactTextString(m) }
func (*PackageReceipt) ProtoMessage()    {}
func (*PackageReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_98eddbefe2841fd0, []int{0}
}

func (m *PackageReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageReceipt.Unmarshal(m, b)
}
func (m *PackageReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageReceipt.Marshal(b, m, deterministic)
}
func (m *PackageReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageReceipt.Merge(m, src)
}
func (m *PackageReceipt) XXX_Size() int {
	return xxx_messageInfo_PackageReceipt.Size(m)
}
func (m *PackageReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_PackageReceipt proto.InternalMessageInfo

func (m *PackageReceipt) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *PackageReceipt) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PackageReceipt) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *PackageReceipt) GetInstallTime() float64 {
	if m != nil {
		return m.InstallTime
	}
	return 0
}

func (m *PackageReceipt) GetInstallerName() string {
	if m != nil {
		return m.InstallerName
	}
	return ""
}

func (m *PackageReceipt) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*PackageReceipt)(nil), "schemas.PackageReceipt")
}

func init() { proto.RegisterFile("pb/package_receipt.proto", fileDescriptor_98eddbefe2841fd0) }

var fileDescriptor_98eddbefe2841fd0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0xd6, 0xd6, 0x8e, 0xda, 0x43, 0x4e, 0x41, 0x10, 0xaa, 0x20, 0xf4, 0x62, 0x57,
	0x14, 0xf4, 0xe0, 0xcd, 0x9b, 0x17, 0x91, 0xc5, 0x93, 0x97, 0x65, 0x36, 0x1d, 0x77, 0x87, 0xdd,
	0x6c, 0x62, 0x92, 0x15, 0xfc, 0x7d, 0xfe, 0x31, 0x31, 0x4d, 0x7b, 0xcb, 0xfb, 0xbe, 0x97, 0x81,
	0x07, 0xca, 0xd5, 0x85, 0x43, 0xdd, 0x61, 0x43, 0x95, 0x27, 0x4d, 0xec, 0xe2, 0xda, 0x79, 0x1b,
	0xad, 0x9c, 0x05, 0xdd, 0x92, 0xc1, 0x70, 0xf5, 0x2b, 0x60, 0xf1, 0xb6, 0xad, 0x94, 0xdb, 0x86,
	0xbc, 0x00, 0xd8, 0x7d, 0xe2, 0x8d, 0x12, 0x4b, 0xb1, 0x9a, 0x97, 0xf3, 0x4c, 0x5e, 0x36, 0x52,
	0xc1, 0xec, 0x9b, 0x7c, 0x60, 0x3b, 0xa8, 0x83, 0xe4, 0x76, 0x51, 0x9e, 0xc3, 0x71, 0x6f, 0x35,
	0xc6, 0x7f, 0x75, 0x98, 0xd4, 0x3e, 0xcb, 0x4b, 0x38, 0xe5, 0x21, 0x44, 0xec, 0xfb, 0x2a, 0xb2,
	0x21, 0x35, 0x59, 0x8a, 0x95, 0x28, 0x4f, 0x32, 0x7b, 0x67, 0x43, 0xf2, 0x1a, 0x16, 0x39, 0x92,
	0xaf, 0x06, 0x34, 0xa4, 0x8e, 0xd2, 0x91, 0xb3, 0x3d, 0x7d, 0x45, 0x43, 0x52, 0xc2, 0xc4, 0x61,
	0x6c, 0xd5, 0x34, 0xc9, 0xf4, 0x7e, 0xbe, 0xfb, 0xb8, 0x6d, 0x38, 0xb6, 0x63, 0xbd, 0xd6, 0xd6,
	0x14, 0x86, 0x75, 0x47, 0xee, 0xf1, 0xa1, 0xb0, 0xe1, 0x6b, 0x24, 0xff, 0x73, 0x93, 0x36, 0xd7,
	0xe3, 0x67, 0xe1, 0xba, 0xe6, 0x29, 0x2f, 0xaf, 0xa7, 0x89, 0xde, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0xa0, 0xed, 0xb8, 0x25, 0x01, 0x00, 0x00,
}
