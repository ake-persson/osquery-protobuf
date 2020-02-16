// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/winbaseobj.proto

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

// Lists named Windows objects in the default object directories, across all terminal services sessions.
type Winbaseobj struct {
	// Terminal Services Session Id
	SessionId int32 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id"`
	// Object Name
	ObjectName string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name"`
	// Object Type
	ObjectType           string   `protobuf:"bytes,3,opt,name=object_type,json=objectType,proto3" json:"object_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Winbaseobj) Reset()         { *m = Winbaseobj{} }
func (m *Winbaseobj) String() string { return proto.CompactTextString(m) }
func (*Winbaseobj) ProtoMessage()    {}
func (*Winbaseobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd5e458a5182be0, []int{0}
}

func (m *Winbaseobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Winbaseobj.Unmarshal(m, b)
}
func (m *Winbaseobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Winbaseobj.Marshal(b, m, deterministic)
}
func (m *Winbaseobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Winbaseobj.Merge(m, src)
}
func (m *Winbaseobj) XXX_Size() int {
	return xxx_messageInfo_Winbaseobj.Size(m)
}
func (m *Winbaseobj) XXX_DiscardUnknown() {
	xxx_messageInfo_Winbaseobj.DiscardUnknown(m)
}

var xxx_messageInfo_Winbaseobj proto.InternalMessageInfo

func (m *Winbaseobj) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Winbaseobj) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *Winbaseobj) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func init() {
	proto.RegisterType((*Winbaseobj)(nil), "schemas.Winbaseobj")
}

func init() { proto.RegisterFile("pb/winbaseobj.proto", fileDescriptor_7fd5e458a5182be0) }

var fileDescriptor_7fd5e458a5182be0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xd2, 0x2f,
	0xcf, 0xcc, 0x4b, 0x4a, 0x2c, 0x4e, 0xcd, 0x4f, 0xca, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0xca, 0xe5, 0xe2, 0x0a, 0x87, 0x4b, 0x0a,
	0xc9, 0x72, 0x71, 0x15, 0xa7, 0x16, 0x17, 0x67, 0xe6, 0xe7, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x06, 0x71, 0x42, 0x45, 0x3c, 0x53, 0x84, 0xe4, 0xb9, 0xb8, 0xf3, 0x93, 0xb2,
	0x52, 0x93, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xb8,
	0x20, 0x42, 0x7e, 0x89, 0xb9, 0xa9, 0x48, 0x0a, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x91, 0x15,
	0x84, 0x54, 0x16, 0xa4, 0x3a, 0x19, 0x45, 0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xe7, 0x66, 0x26, 0x67, 0xa7, 0x16, 0x98, 0x9b, 0xe9, 0xe7, 0x17, 0x17, 0x96,
	0xa6, 0x16, 0x55, 0xea, 0x82, 0x1d, 0x97, 0x54, 0x9a, 0xa6, 0x5f, 0x90, 0x9d, 0x6e, 0x0d, 0x75,
	0x62, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x2a, 0x06, 0x7f, 0xc9,
	0x00, 0x00, 0x00,
}