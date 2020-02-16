// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/package_install_history.proto

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

// OS X package install history.
type PackageInstallHistory struct {
	// Label packageIdentifiers
	PackageId string `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id"`
	// Label date as UNIX timestamp
	Time int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time"`
	// Package display name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// Package display version
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version"`
	// Install source: usually the installer process name
	Source string `protobuf:"bytes,5,opt,name=source,proto3" json:"source"`
	// Package content_type (optional)
	ContentType          string   `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageInstallHistory) Reset()         { *m = PackageInstallHistory{} }
func (m *PackageInstallHistory) String() string { return proto.CompactTextString(m) }
func (*PackageInstallHistory) ProtoMessage()    {}
func (*PackageInstallHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d4d93055c9463a, []int{0}
}

func (m *PackageInstallHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageInstallHistory.Unmarshal(m, b)
}
func (m *PackageInstallHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageInstallHistory.Marshal(b, m, deterministic)
}
func (m *PackageInstallHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageInstallHistory.Merge(m, src)
}
func (m *PackageInstallHistory) XXX_Size() int {
	return xxx_messageInfo_PackageInstallHistory.Size(m)
}
func (m *PackageInstallHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageInstallHistory.DiscardUnknown(m)
}

var xxx_messageInfo_PackageInstallHistory proto.InternalMessageInfo

func (m *PackageInstallHistory) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *PackageInstallHistory) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PackageInstallHistory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageInstallHistory) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PackageInstallHistory) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *PackageInstallHistory) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func init() {
	proto.RegisterType((*PackageInstallHistory)(nil), "schemas.PackageInstallHistory")
}

func init() { proto.RegisterFile("pb/package_install_history.proto", fileDescriptor_73d4d93055c9463a) }

var fileDescriptor_73d4d93055c9463a = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x68, 0x53, 0xd5, 0x30, 0x59, 0x02, 0x79, 0x41, 0x0a, 0x4c, 0x5d, 0x68, 0x10,
	0x48, 0x30, 0xb0, 0x31, 0xd1, 0x0d, 0x55, 0x4c, 0x2c, 0x91, 0xe3, 0x1e, 0x89, 0x95, 0xd8, 0x67,
	0xec, 0x0b, 0x52, 0x7e, 0x17, 0x7f, 0x10, 0x71, 0x09, 0xdd, 0xee, 0x7d, 0xef, 0xd3, 0x49, 0x4f,
	0x14, 0xa1, 0x2e, 0x83, 0x36, 0x9d, 0x6e, 0xa0, 0xb2, 0x3e, 0x91, 0xee, 0xfb, 0xaa, 0xb5, 0x89,
	0x30, 0x8e, 0xdb, 0x10, 0x91, 0x50, 0xae, 0x92, 0x69, 0xc1, 0xe9, 0x74, 0xf3, 0x93, 0x89, 0x8b,
	0xb7, 0x49, 0xdd, 0x4d, 0xe6, 0xeb, 0x24, 0xca, 0x2b, 0x21, 0x8e, 0x3f, 0x0e, 0x2a, 0x2b, 0xb2,
	0xcd, 0x7a, 0xbf, 0x9e, 0xc9, 0xee, 0x20, 0xa5, 0x58, 0x90, 0x75, 0xa0, 0x4e, 0x8a, 0x6c, 0xb3,
	0xdc, 0xf3, 0xfd, 0xc7, 0xbc, 0x76, 0xa0, 0x4e, 0x59, 0xe6, 0x5b, 0x2a, 0xb1, 0xfa, 0x86, 0x98,
	0x2c, 0x7a, 0xb5, 0x60, 0xfc, 0x1f, 0xe5, 0xa5, 0xc8, 0x13, 0x0e, 0xd1, 0x80, 0x5a, 0x72, 0x31,
	0x27, 0x79, 0x2d, 0xce, 0x0d, 0x7a, 0x02, 0x4f, 0x15, 0x8d, 0x01, 0x54, 0xce, 0xed, 0xd9, 0xcc,
	0xde, 0xc7, 0x00, 0x2f, 0xf7, 0x1f, 0x77, 0x8d, 0xa5, 0x76, 0xa8, 0xb7, 0x06, 0x5d, 0xe9, 0xac,
	0xe9, 0x20, 0x3c, 0x3d, 0x96, 0x98, 0xbe, 0x06, 0x88, 0xe3, 0x2d, 0x6f, 0xac, 0x87, 0xcf, 0x32,
	0x74, 0xcd, 0xf3, 0xbc, 0xb4, 0xce, 0x99, 0x3e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x1d,
	0x70, 0xb8, 0x1d, 0x01, 0x00, 0x00,
}
