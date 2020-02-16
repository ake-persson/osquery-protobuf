// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/known_host.proto

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

// A line-delimited known_hosts table.
type KnownHost struct {
	// parsed authorized keys line
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// Path to known_hosts file
	KeyFile              string   `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KnownHost) Reset()         { *m = KnownHost{} }
func (m *KnownHost) String() string { return proto.CompactTextString(m) }
func (*KnownHost) ProtoMessage()    {}
func (*KnownHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_7862cd453b30539b, []int{0}
}

func (m *KnownHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnownHost.Unmarshal(m, b)
}
func (m *KnownHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnownHost.Marshal(b, m, deterministic)
}
func (m *KnownHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnownHost.Merge(m, src)
}
func (m *KnownHost) XXX_Size() int {
	return xxx_messageInfo_KnownHost.Size(m)
}
func (m *KnownHost) XXX_DiscardUnknown() {
	xxx_messageInfo_KnownHost.DiscardUnknown(m)
}

var xxx_messageInfo_KnownHost proto.InternalMessageInfo

func (m *KnownHost) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KnownHost) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func init() {
	proto.RegisterType((*KnownHost)(nil), "schemas.KnownHost")
}

func init() { proto.RegisterFile("pb/known_host.proto", fileDescriptor_7862cd453b30539b) }

var fileDescriptor_7862cd453b30539b = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xd2, 0xcf,
	0xce, 0xcb, 0x2f, 0xcf, 0x8b, 0xcf, 0xc8, 0x2f, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0xb2, 0xe0, 0xe2, 0xf4, 0x06, 0x49, 0x7a,
	0xe4, 0x17, 0x97, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0x98, 0x42, 0x92, 0x5c, 0x1c, 0xd9, 0xa9, 0x95, 0xf1, 0x69, 0x99, 0x39, 0xa9, 0x12,
	0x4c, 0x60, 0x61, 0xf6, 0xec, 0xd4, 0x4a, 0xb7, 0xcc, 0x9c, 0x54, 0x27, 0xa3, 0x28, 0x83, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xdc, 0xcc, 0xe4, 0xec, 0xd4, 0x02,
	0x73, 0x33, 0xfd, 0xfc, 0xe2, 0xc2, 0xd2, 0xd4, 0xa2, 0x4a, 0x5d, 0xb0, 0x3d, 0x49, 0xa5, 0x69,
	0xfa, 0x05, 0xd9, 0xe9, 0xd6, 0x50, 0xdb, 0x92, 0xd8, 0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0x44, 0x77, 0x95, 0x94, 0x00, 0x00, 0x00,
}