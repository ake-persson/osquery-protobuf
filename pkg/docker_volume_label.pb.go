// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_volume_label.proto

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

// Docker volume labels.
type DockerVolumeLabel struct {
	// Volume name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Label key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key"`
	// Optional label value
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerVolumeLabel) Reset()         { *m = DockerVolumeLabel{} }
func (m *DockerVolumeLabel) String() string { return proto.CompactTextString(m) }
func (*DockerVolumeLabel) ProtoMessage()    {}
func (*DockerVolumeLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b837dd88bbd9d7, []int{0}
}

func (m *DockerVolumeLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerVolumeLabel.Unmarshal(m, b)
}
func (m *DockerVolumeLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerVolumeLabel.Marshal(b, m, deterministic)
}
func (m *DockerVolumeLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerVolumeLabel.Merge(m, src)
}
func (m *DockerVolumeLabel) XXX_Size() int {
	return xxx_messageInfo_DockerVolumeLabel.Size(m)
}
func (m *DockerVolumeLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerVolumeLabel.DiscardUnknown(m)
}

var xxx_messageInfo_DockerVolumeLabel proto.InternalMessageInfo

func (m *DockerVolumeLabel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerVolumeLabel) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DockerVolumeLabel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerVolumeLabel)(nil), "schemas.DockerVolumeLabel")
}

func init() { proto.RegisterFile("pb/docker_volume_label.proto", fileDescriptor_f9b837dd88bbd9d7) }

var fileDescriptor_f9b837dd88bbd9d7 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x48, 0xd2, 0x4f,
	0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0x8a, 0x2f, 0xcb, 0xcf, 0x29, 0xcd, 0x4d, 0x8d, 0xcf, 0x49, 0x4c,
	0x4a, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d,
	0x2c, 0x56, 0xf2, 0xe7, 0x12, 0x74, 0x01, 0xab, 0x0a, 0x03, 0x2b, 0xf2, 0x01, 0xa9, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85,
	0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0xc0, 0x42, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b,
	0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x33, 0x58, 0x0c, 0xc2, 0x71, 0x32, 0x8a, 0x32, 0x48, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0xce, 0x4e, 0x2d, 0x30,
	0x37, 0xd3, 0xcf, 0x2f, 0x2e, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x05, 0x5b, 0x9f, 0x54, 0x9a, 0xa6,
	0x5f, 0x90, 0x9d, 0x6e, 0x0d, 0x75, 0x44, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0x22, 0xf4, 0x46, 0xb4, 0x00, 0x00, 0x00,
}