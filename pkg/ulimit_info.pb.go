// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/ulimit_info.proto

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

// System resource usage limits.
type UlimitInfo struct {
	// System resource to be limited
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// Current limit value
	SoftLimit string `protobuf:"bytes,2,opt,name=soft_limit,json=softLimit,proto3" json:"soft_limit"`
	// Maximum limit value
	HardLimit            string   `protobuf:"bytes,3,opt,name=hard_limit,json=hardLimit,proto3" json:"hard_limit"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UlimitInfo) Reset()         { *m = UlimitInfo{} }
func (m *UlimitInfo) String() string { return proto.CompactTextString(m) }
func (*UlimitInfo) ProtoMessage()    {}
func (*UlimitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_08bbabe09f7de7a9, []int{0}
}

func (m *UlimitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UlimitInfo.Unmarshal(m, b)
}
func (m *UlimitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UlimitInfo.Marshal(b, m, deterministic)
}
func (m *UlimitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UlimitInfo.Merge(m, src)
}
func (m *UlimitInfo) XXX_Size() int {
	return xxx_messageInfo_UlimitInfo.Size(m)
}
func (m *UlimitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UlimitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UlimitInfo proto.InternalMessageInfo

func (m *UlimitInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UlimitInfo) GetSoftLimit() string {
	if m != nil {
		return m.SoftLimit
	}
	return ""
}

func (m *UlimitInfo) GetHardLimit() string {
	if m != nil {
		return m.HardLimit
	}
	return ""
}

func init() {
	proto.RegisterType((*UlimitInfo)(nil), "schemas.UlimitInfo")
}

func init() { proto.RegisterFile("pb/ulimit_info.proto", fileDescriptor_08bbabe09f7de7a9) }

var fileDescriptor_08bbabe09f7de7a9 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x48, 0xd2, 0x2f,
	0xcd, 0xc9, 0xcc, 0xcd, 0x2c, 0x89, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0x8a, 0xe3, 0xe2, 0x0a, 0x05, 0xcb,
	0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0xb2, 0x5c, 0x5c, 0xc5, 0xf9, 0x69, 0x25, 0xf1, 0x60, 0x55,
	0x12, 0x4c, 0x60, 0x19, 0x4e, 0x90, 0x88, 0x0f, 0x48, 0x00, 0x24, 0x9d, 0x91, 0x58, 0x94, 0x02,
	0x95, 0x66, 0x86, 0x48, 0x83, 0x44, 0xc0, 0xd2, 0x4e, 0x46, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0x99, 0xc9, 0xd9, 0xa9, 0x05, 0xe6, 0x66, 0xfa,
	0xf9, 0xc5, 0x85, 0xa5, 0xa9, 0x45, 0x95, 0xba, 0x60, 0xd7, 0x24, 0x95, 0xa6, 0xe9, 0x17, 0x64,
	0xa7, 0x5b, 0x43, 0xdd, 0x94, 0xc4, 0x06, 0x16, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06,
	0xc6, 0x17, 0xc6, 0xbb, 0x00, 0x00, 0x00,
}
