// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/carf.proto

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

// Forensic Carves.
type Carf struct {
	// Time at which the carve was kicked off
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	// A SHA256 sum of the carved archive
	Sha256 string `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256"`
	// Size of the carved archive
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// The path of the requested carve
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path"`
	// Status of the carve
	Status string `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	// Identifying value of the carve session
	CarveGuid string `protobuf:"bytes,6,opt,name=carve_guid,json=carveGuid,proto3" json:"carve_guid"`
	// Set this value to '1' to start a file carve
	Carve                int32    `protobuf:"varint,7,opt,name=carve,proto3" json:"carve"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Carf) Reset()         { *m = Carf{} }
func (m *Carf) String() string { return proto.CompactTextString(m) }
func (*Carf) ProtoMessage()    {}
func (*Carf) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f79b0172cb63854, []int{0}
}

func (m *Carf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Carf.Unmarshal(m, b)
}
func (m *Carf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Carf.Marshal(b, m, deterministic)
}
func (m *Carf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Carf.Merge(m, src)
}
func (m *Carf) XXX_Size() int {
	return xxx_messageInfo_Carf.Size(m)
}
func (m *Carf) XXX_DiscardUnknown() {
	xxx_messageInfo_Carf.DiscardUnknown(m)
}

var xxx_messageInfo_Carf proto.InternalMessageInfo

func (m *Carf) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Carf) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Carf) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Carf) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Carf) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Carf) GetCarveGuid() string {
	if m != nil {
		return m.CarveGuid
	}
	return ""
}

func (m *Carf) GetCarve() int32 {
	if m != nil {
		return m.Carve
	}
	return 0
}

func init() {
	proto.RegisterType((*Carf)(nil), "schemas.Carf")
}

func init() { proto.RegisterFile("pb/carf.proto", fileDescriptor_4f79b0172cb63854) }

var fileDescriptor_4f79b0172cb63854 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0x3f, 0x4b, 0xc0, 0x30,
	0x14, 0xc4, 0x89, 0xfd, 0x47, 0x03, 0x2e, 0x41, 0x24, 0x8b, 0x50, 0x9c, 0xba, 0xd8, 0x48, 0xc5,
	0x3a, 0xb8, 0xe9, 0xe0, 0xde, 0xd1, 0x45, 0x92, 0x34, 0x6d, 0x42, 0x09, 0x89, 0xf9, 0x23, 0xe8,
	0x17, 0xf2, 0x6b, 0x4a, 0x9f, 0xdd, 0xee, 0x77, 0xef, 0xee, 0xc1, 0xe1, 0x4b, 0x2f, 0x98, 0xe4,
	0x61, 0x1d, 0x7c, 0x70, 0xc9, 0x91, 0x26, 0x4a, 0xad, 0x2c, 0x8f, 0xb7, 0xbf, 0x08, 0x97, 0xaf,
	0x3c, 0xac, 0x84, 0xe0, 0x32, 0x19, 0xab, 0x28, 0xea, 0x50, 0x5f, 0xcc, 0xa0, 0xc9, 0x35, 0xae,
	0xa3, 0xe6, 0xe3, 0xe3, 0x44, 0x2f, 0x3a, 0xd4, 0xb7, 0xf3, 0x49, 0x47, 0x36, 0x9a, 0x1f, 0x45,
	0x8b, 0x0e, 0xf5, 0xd5, 0x0c, 0xfa, 0xf0, 0x3c, 0x4f, 0x9a, 0x96, 0x90, 0x04, 0x0d, 0xfd, 0xc4,
	0x53, 0x8e, 0xb4, 0x3a, 0xfb, 0x40, 0xe4, 0x06, 0x63, 0xc9, 0xc3, 0x97, 0xfa, 0xd8, 0xb2, 0x59,
	0x68, 0x0d, 0xb7, 0x16, 0x9c, 0xb7, 0x6c, 0x16, 0x72, 0x85, 0x2b, 0x00, 0xda, 0xc0, 0xff, 0x7f,
	0x78, 0x19, 0xdf, 0xef, 0x37, 0x93, 0x74, 0x16, 0x83, 0x74, 0x96, 0x59, 0x23, 0x77, 0xe5, 0x9f,
	0x26, 0xe6, 0xe2, 0x67, 0x56, 0xe1, 0xfb, 0x0e, 0x76, 0x89, 0xbc, 0x32, 0xbf, 0x6f, 0xcf, 0xe7,
	0x3a, 0x51, 0x83, 0xfb, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0x13, 0x2c, 0x69, 0xc4, 0xfe, 0x00,
	0x00, 0x00,
}
