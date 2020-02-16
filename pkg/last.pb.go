// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/last.proto

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

// System logins and logouts.
type Last struct {
	// Entry username
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	// Entry terminal
	Tty string `protobuf:"bytes,2,opt,name=tty,proto3" json:"tty"`
	// Process (or thread) ID
	Pid int32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid"`
	// Entry type
	Type int32 `protobuf:"varint,4,opt,name=type,proto3" json:"type"`
	// Entry timestamp
	Time int32 `protobuf:"varint,5,opt,name=time,proto3" json:"time"`
	// Entry hostname
	Host                 string   `protobuf:"bytes,6,opt,name=host,proto3" json:"host"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Last) Reset()         { *m = Last{} }
func (m *Last) String() string { return proto.CompactTextString(m) }
func (*Last) ProtoMessage()    {}
func (*Last) Descriptor() ([]byte, []int) {
	return fileDescriptor_049cd416fff4d298, []int{0}
}

func (m *Last) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Last.Unmarshal(m, b)
}
func (m *Last) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Last.Marshal(b, m, deterministic)
}
func (m *Last) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Last.Merge(m, src)
}
func (m *Last) XXX_Size() int {
	return xxx_messageInfo_Last.Size(m)
}
func (m *Last) XXX_DiscardUnknown() {
	xxx_messageInfo_Last.DiscardUnknown(m)
}

var xxx_messageInfo_Last proto.InternalMessageInfo

func (m *Last) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Last) GetTty() string {
	if m != nil {
		return m.Tty
	}
	return ""
}

func (m *Last) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Last) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Last) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Last) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*Last)(nil), "schemas.Last")
}

func init() { proto.RegisterFile("pb/last.proto", fileDescriptor_049cd416fff4d298) }

var fileDescriptor_049cd416fff4d298 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0x3d, 0xab, 0xc2, 0x30,
	0x14, 0xc6, 0x71, 0x72, 0xfb, 0x72, 0xef, 0x0d, 0x08, 0x92, 0x29, 0x38, 0x15, 0xa7, 0x2e, 0x36,
	0xa2, 0xa0, 0x83, 0x9b, 0xb3, 0x53, 0x47, 0xb7, 0xa4, 0xc6, 0x36, 0xd4, 0x98, 0xd8, 0x9c, 0x0c,
	0x5d, 0xfd, 0xe4, 0xd2, 0x43, 0x71, 0xfb, 0x9f, 0xdf, 0x70, 0x78, 0xe8, 0xc2, 0x2b, 0xf1, 0x90,
	0x01, 0x2a, 0x3f, 0x38, 0x70, 0xec, 0x37, 0x34, 0x9d, 0xb6, 0x32, 0xac, 0xdf, 0x84, 0xa6, 0x17,
	0x19, 0x80, 0xad, 0xe8, 0x5f, 0x0c, 0x7a, 0x78, 0x4a, 0xab, 0x39, 0x29, 0x48, 0xf9, 0x5f, 0x7f,
	0x6f, 0xb6, 0xa4, 0x09, 0xc0, 0xc8, 0x7f, 0x90, 0xa7, 0x9c, 0xc4, 0x9b, 0x1b, 0x4f, 0x0a, 0x52,
	0x66, 0xf5, 0x94, 0x8c, 0xd1, 0x14, 0x46, 0xaf, 0x79, 0x8a, 0x84, 0x8d, 0x66, 0xac, 0xe6, 0xd9,
	0x6c, 0xc6, 0xa2, 0x75, 0x2e, 0x00, 0xcf, 0xf1, 0x19, 0xf6, 0x79, 0x77, 0xdd, 0xb6, 0x06, 0xba,
	0xa8, 0xaa, 0xc6, 0x59, 0x61, 0x4d, 0xd3, 0x6b, 0x7f, 0x3c, 0x08, 0x17, 0x5e, 0x51, 0x0f, 0xe3,
	0x06, 0x27, 0xab, 0x78, 0x17, 0xbe, 0x6f, 0x4f, 0xf3, 0x70, 0x95, 0xa3, 0xee, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x90, 0xff, 0x91, 0x1e, 0xd9, 0x00, 0x00, 0x00,
}