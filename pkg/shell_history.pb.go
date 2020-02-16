// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/shell_history.proto

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

// A line-delimited (command) table of per-user .*_history data.
type ShellHistory struct {
	// Shell history owner
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// Entry timestamp. It could be absent
	Time int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time"`
	// Unparsed date/line/command history line
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command"`
	// Path to the .*_history for this user
	HistoryFile          string   `protobuf:"bytes,4,opt,name=history_file,json=historyFile,proto3" json:"history_file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShellHistory) Reset()         { *m = ShellHistory{} }
func (m *ShellHistory) String() string { return proto.CompactTextString(m) }
func (*ShellHistory) ProtoMessage()    {}
func (*ShellHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_34d20134e6ce98bb, []int{0}
}

func (m *ShellHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShellHistory.Unmarshal(m, b)
}
func (m *ShellHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShellHistory.Marshal(b, m, deterministic)
}
func (m *ShellHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShellHistory.Merge(m, src)
}
func (m *ShellHistory) XXX_Size() int {
	return xxx_messageInfo_ShellHistory.Size(m)
}
func (m *ShellHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ShellHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ShellHistory proto.InternalMessageInfo

func (m *ShellHistory) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ShellHistory) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ShellHistory) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ShellHistory) GetHistoryFile() string {
	if m != nil {
		return m.HistoryFile
	}
	return ""
}

func init() {
	proto.RegisterType((*ShellHistory)(nil), "schemas.ShellHistory")
}

func init() { proto.RegisterFile("pb/shell_history.proto", fileDescriptor_34d20134e6ce98bb) }

var fileDescriptor_34d20134e6ce98bb = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x2f,
	0xce, 0x48, 0xcd, 0xc9, 0x89, 0xcf, 0xc8, 0x2c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0x2a, 0xe4, 0xe2, 0x09,
	0x06, 0xc9, 0x7b, 0x40, 0xa4, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0x40, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xcc, 0xdc, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f, 0x37, 0x37, 0x31, 0x2f,
	0x45, 0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x52, 0xe4, 0xe2, 0x81, 0xda, 0x14,
	0x9f, 0x96, 0x99, 0x93, 0x2a, 0xc1, 0x02, 0x96, 0xe6, 0x86, 0x8a, 0xb9, 0x65, 0xe6, 0xa4, 0x3a,
	0x19, 0x45, 0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x66,
	0x26, 0x67, 0xa7, 0x16, 0x98, 0x9b, 0xe9, 0xe7, 0x17, 0x17, 0x96, 0xa6, 0x16, 0x55, 0xea, 0x82,
	0x1d, 0x98, 0x54, 0x9a, 0xa6, 0x5f, 0x90, 0x9d, 0x6e, 0x0d, 0x75, 0x66, 0x12, 0x1b, 0x58, 0xd4,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x76, 0xc6, 0x70, 0x08, 0xd0, 0x00, 0x00, 0x00,
}
