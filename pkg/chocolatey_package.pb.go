// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/chocolatey_package.proto

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

// Chocolatey packages installed in a system.
type ChocolateyPackage struct {
	// Package display name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Package-supplied version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version"`
	// Package-supplied summary
	Summary string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary"`
	// Optional package author
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author"`
	// License under which package is launched
	License string `protobuf:"bytes,5,opt,name=license,proto3" json:"license"`
	// Path at which this package resides
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChocolateyPackage) Reset()         { *m = ChocolateyPackage{} }
func (m *ChocolateyPackage) String() string { return proto.CompactTextString(m) }
func (*ChocolateyPackage) ProtoMessage()    {}
func (*ChocolateyPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_84943a17ac88c95f, []int{0}
}

func (m *ChocolateyPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChocolateyPackage.Unmarshal(m, b)
}
func (m *ChocolateyPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChocolateyPackage.Marshal(b, m, deterministic)
}
func (m *ChocolateyPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChocolateyPackage.Merge(m, src)
}
func (m *ChocolateyPackage) XXX_Size() int {
	return xxx_messageInfo_ChocolateyPackage.Size(m)
}
func (m *ChocolateyPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChocolateyPackage.DiscardUnknown(m)
}

var xxx_messageInfo_ChocolateyPackage proto.InternalMessageInfo

func (m *ChocolateyPackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChocolateyPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChocolateyPackage) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ChocolateyPackage) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ChocolateyPackage) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *ChocolateyPackage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*ChocolateyPackage)(nil), "schemas.ChocolateyPackage")
}

func init() { proto.RegisterFile("pb/chocolatey_package.proto", fileDescriptor_84943a17ac88c95f) }

var fileDescriptor_84943a17ac88c95f = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xbd, 0x4a, 0x05, 0x31,
	0x10, 0x85, 0x59, 0xbd, 0xee, 0xc5, 0x74, 0xa6, 0x90, 0x80, 0x8d, 0x58, 0xd9, 0x78, 0x23, 0x0a,
	0x5a, 0xd8, 0xe9, 0x0b, 0x88, 0xa5, 0x8d, 0x4c, 0xc2, 0xb8, 0x09, 0xbb, 0xd9, 0x89, 0xf9, 0x11,
	0xf6, 0x71, 0x7c, 0x53, 0xc9, 0x8f, 0x76, 0xf3, 0xcd, 0x39, 0x87, 0xc3, 0x61, 0x17, 0x5e, 0x49,
	0x6d, 0x48, 0xd3, 0x02, 0x09, 0xb7, 0x0f, 0x0f, 0x7a, 0x86, 0x09, 0x0f, 0x3e, 0x50, 0x22, 0xbe,
	0x8f, 0xda, 0xa0, 0x83, 0x78, 0xf5, 0x33, 0xb0, 0xb3, 0x97, 0x7f, 0xd7, 0x6b, 0x33, 0x71, 0xce,
	0x76, 0x2b, 0x38, 0x14, 0xc3, 0xe5, 0x70, 0x7d, 0xfa, 0x56, 0x6f, 0x2e, 0xd8, 0xfe, 0x1b, 0x43,
	0xb4, 0xb4, 0x8a, 0xa3, 0xfa, 0xfe, 0xc3, 0xa2, 0xc4, 0xec, 0x1c, 0x84, 0x4d, 0x1c, 0x37, 0xa5,
	0x23, 0x3f, 0x67, 0x23, 0xe4, 0x64, 0x28, 0x88, 0x5d, 0x15, 0x3a, 0x95, 0xc4, 0x62, 0x35, 0xae,
	0x11, 0xc5, 0x49, 0x4b, 0x74, 0x2c, 0xcd, 0x1e, 0x92, 0x11, 0x63, 0x6b, 0x2e, 0xf7, 0xf3, 0xdd,
	0xfb, 0xed, 0x64, 0x93, 0xc9, 0xea, 0xa0, 0xc9, 0x49, 0x67, 0xf5, 0x8c, 0xfe, 0xf1, 0x41, 0x52,
	0xfc, 0xca, 0x18, 0xb6, 0x9b, 0xba, 0x48, 0xe5, 0x4f, 0xe9, 0xe7, 0xe9, 0xa9, 0xef, 0x52, 0x63,
	0xfd, 0xde, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xce, 0xb3, 0x78, 0x06, 0x01, 0x00, 0x00,
}