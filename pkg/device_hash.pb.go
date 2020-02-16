// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/device_hash.proto

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

// Similar to the hash table, but use TSK and allow block address access.
type DeviceHash struct {
	// Absolute file path to device node
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device"`
	// A partition number
	Partition string `protobuf:"bytes,2,opt,name=partition,proto3" json:"partition"`
	// Filesystem inode number
	Inode int64 `protobuf:"varint,3,opt,name=inode,proto3" json:"inode"`
	// MD5 hash of provided inode data
	Md5 string `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5"`
	// SHA1 hash of provided inode data
	Sha1 string `protobuf:"bytes,5,opt,name=sha1,proto3" json:"sha1"`
	// SHA256 hash of provided inode data
	Sha256               string   `protobuf:"bytes,6,opt,name=sha256,proto3" json:"sha256"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceHash) Reset()         { *m = DeviceHash{} }
func (m *DeviceHash) String() string { return proto.CompactTextString(m) }
func (*DeviceHash) ProtoMessage()    {}
func (*DeviceHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1978d61c69d025e, []int{0}
}

func (m *DeviceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceHash.Unmarshal(m, b)
}
func (m *DeviceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceHash.Marshal(b, m, deterministic)
}
func (m *DeviceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceHash.Merge(m, src)
}
func (m *DeviceHash) XXX_Size() int {
	return xxx_messageInfo_DeviceHash.Size(m)
}
func (m *DeviceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceHash.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceHash proto.InternalMessageInfo

func (m *DeviceHash) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DeviceHash) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *DeviceHash) GetInode() int64 {
	if m != nil {
		return m.Inode
	}
	return 0
}

func (m *DeviceHash) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *DeviceHash) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *DeviceHash) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceHash)(nil), "schemas.DeviceHash")
}

func init() { proto.RegisterFile("pb/device_hash.proto", fileDescriptor_b1978d61c69d025e) }

var fileDescriptor_b1978d61c69d025e = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0x4a, 0xc6, 0x30,
	0x1c, 0xc4, 0x89, 0xfd, 0xbe, 0x4a, 0xff, 0x93, 0x84, 0x22, 0x19, 0x1c, 0x8a, 0x53, 0x17, 0x1b,
	0xad, 0xb4, 0x0e, 0x6e, 0xe2, 0xe0, 0xdc, 0xd1, 0x45, 0x92, 0x34, 0x36, 0xa1, 0xa4, 0x89, 0x4d,
	0x2a, 0xf8, 0x22, 0x3e, 0xaf, 0x34, 0x29, 0xb8, 0xdd, 0xfd, 0xee, 0xe0, 0x38, 0x28, 0x1d, 0xa7,
	0xa3, 0xfc, 0xd6, 0x42, 0x7e, 0x28, 0xe6, 0x55, 0xe3, 0x56, 0x1b, 0x2c, 0xbe, 0xf4, 0x42, 0x49,
	0xc3, 0xfc, 0xed, 0x2f, 0x02, 0x78, 0x8d, 0xf1, 0x1b, 0xf3, 0x0a, 0x5f, 0x43, 0x9e, 0xca, 0x04,
	0x55, 0xa8, 0x2e, 0x86, 0xc3, 0xe1, 0x1b, 0x28, 0x1c, 0x5b, 0x83, 0x0e, 0xda, 0x2e, 0xe4, 0x22,
	0x46, 0xff, 0x00, 0x97, 0x70, 0xd6, 0x8b, 0x1d, 0x25, 0xc9, 0x2a, 0x54, 0x67, 0x43, 0x32, 0xf8,
	0x0a, 0x32, 0x33, 0x76, 0xe4, 0x14, 0xdb, 0xbb, 0xc4, 0x18, 0x4e, 0x5e, 0xb1, 0x07, 0x72, 0x8e,
	0x28, 0xea, 0x7d, 0xd1, 0x2b, 0xd6, 0x76, 0x3d, 0xc9, 0xd3, 0x62, 0x72, 0x2f, 0xed, 0xfb, 0xfd,
	0xa4, 0x83, 0xda, 0x78, 0x23, 0xac, 0xa1, 0x46, 0x8b, 0x59, 0xba, 0xa7, 0x9e, 0x5a, 0xff, 0xb5,
	0xc9, 0xf5, 0xe7, 0x2e, 0xde, 0xe0, 0xdb, 0x27, 0x75, 0xf3, 0xf4, 0x7c, 0x9c, 0xe1, 0x79, 0xa4,
	0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0xb7, 0x5d, 0x73, 0xf4, 0x00, 0x00, 0x00,
}
