// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/crash.proto

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

// Application, System, and Mobile App crash logs.
type Crash struct {
	// Type of crash log
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// Process (or thread) ID of the crashed process
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid"`
	// Path to the crashed process
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path"`
	// Location of log file
	CrashPath string `protobuf:"bytes,4,opt,name=crash_path,json=crashPath,proto3" json:"crash_path"`
	// Identifier of the crashed process
	Identifier string `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier"`
	// Version info of the crashed process
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version"`
	// Parent PID of the crashed process
	Parent int64 `protobuf:"varint,7,opt,name=parent,proto3" json:"parent"`
	// Process responsible for the crashed process
	Responsible string `protobuf:"bytes,8,opt,name=responsible,proto3" json:"responsible"`
	// User ID of the crashed process
	Uid int32 `protobuf:"varint,9,opt,name=uid,proto3" json:"uid"`
	// Date/Time at which the crash occurred
	Datetime string `protobuf:"bytes,10,opt,name=datetime,proto3" json:"datetime"`
	// Thread ID which crashed
	CrashedThread int64 `protobuf:"varint,11,opt,name=crashed_thread,json=crashedThread,proto3" json:"crashed_thread"`
	// Most recent frame from the stack trace
	StackTrace string `protobuf:"bytes,12,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace"`
	// Exception type of the crash
	ExceptionType string `protobuf:"bytes,13,opt,name=exception_type,json=exceptionType,proto3" json:"exception_type"`
	// Exception codes from the crash
	ExceptionCodes string `protobuf:"bytes,14,opt,name=exception_codes,json=exceptionCodes,proto3" json:"exception_codes"`
	// Exception notes from the crash
	ExceptionNotes string `protobuf:"bytes,15,opt,name=exception_notes,json=exceptionNotes,proto3" json:"exception_notes"`
	// The value of the system registers
	Registers            string   `protobuf:"bytes,16,opt,name=registers,proto3" json:"registers"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crash) Reset()         { *m = Crash{} }
func (m *Crash) String() string { return proto.CompactTextString(m) }
func (*Crash) ProtoMessage()    {}
func (*Crash) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd62cd1ba883daf, []int{0}
}

func (m *Crash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crash.Unmarshal(m, b)
}
func (m *Crash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crash.Marshal(b, m, deterministic)
}
func (m *Crash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crash.Merge(m, src)
}
func (m *Crash) XXX_Size() int {
	return xxx_messageInfo_Crash.Size(m)
}
func (m *Crash) XXX_DiscardUnknown() {
	xxx_messageInfo_Crash.DiscardUnknown(m)
}

var xxx_messageInfo_Crash proto.InternalMessageInfo

func (m *Crash) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Crash) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Crash) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Crash) GetCrashPath() string {
	if m != nil {
		return m.CrashPath
	}
	return ""
}

func (m *Crash) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Crash) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Crash) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *Crash) GetResponsible() string {
	if m != nil {
		return m.Responsible
	}
	return ""
}

func (m *Crash) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Crash) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *Crash) GetCrashedThread() int64 {
	if m != nil {
		return m.CrashedThread
	}
	return 0
}

func (m *Crash) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

func (m *Crash) GetExceptionType() string {
	if m != nil {
		return m.ExceptionType
	}
	return ""
}

func (m *Crash) GetExceptionCodes() string {
	if m != nil {
		return m.ExceptionCodes
	}
	return ""
}

func (m *Crash) GetExceptionNotes() string {
	if m != nil {
		return m.ExceptionNotes
	}
	return ""
}

func (m *Crash) GetRegisters() string {
	if m != nil {
		return m.Registers
	}
	return ""
}

func init() {
	proto.RegisterType((*Crash)(nil), "schemas.Crash")
}

func init() { proto.RegisterFile("pb/crash.proto", fileDescriptor_0bd62cd1ba883daf) }

var fileDescriptor_0bd62cd1ba883daf = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x49, 0x6f, 0xdb, 0x30,
	0x10, 0x85, 0xa1, 0xca, 0x9b, 0xc6, 0xf5, 0x02, 0x1e, 0x0a, 0xa2, 0xe8, 0x22, 0x14, 0x28, 0xea,
	0x4b, 0xad, 0xa2, 0x01, 0x92, 0x43, 0x6e, 0xf1, 0x3d, 0x08, 0x0c, 0x9f, 0x72, 0x11, 0x28, 0x6a,
	0x6c, 0x11, 0x8e, 0x44, 0x86, 0xa4, 0x82, 0xf8, 0xe7, 0xe6, 0x9f, 0x04, 0x1c, 0x3b, 0xb6, 0x93,
	0xdb, 0x9b, 0xef, 0xbd, 0xa1, 0x46, 0xe4, 0xc0, 0xd8, 0x14, 0x99, 0xb4, 0xc2, 0x55, 0x73, 0x63,
	0xb5, 0xd7, 0xac, 0xef, 0x64, 0x85, 0xb5, 0x70, 0xbf, 0x5e, 0x62, 0xe8, 0x2e, 0x82, 0xc1, 0x18,
	0x74, 0xfc, 0xce, 0x20, 0x8f, 0xd2, 0x68, 0x96, 0x2c, 0x49, 0xb3, 0x29, 0xc4, 0x46, 0x95, 0xfc,
	0x53, 0x1a, 0xcd, 0xe2, 0x65, 0x90, 0x21, 0x65, 0x84, 0xaf, 0x78, 0xbc, 0x4f, 0x05, 0xcd, 0xbe,
	0x03, 0xd0, 0xd9, 0x39, 0x39, 0x1d, 0x72, 0x12, 0x22, 0x77, 0xc1, 0xfe, 0x01, 0xa0, 0x4a, 0x6c,
	0xbc, 0x5a, 0x2b, 0xb4, 0xbc, 0x4b, 0xf6, 0x19, 0x61, 0x1c, 0xfa, 0x4f, 0x68, 0x9d, 0xd2, 0x0d,
	0xef, 0x91, 0xf9, 0x56, 0xb2, 0x2f, 0xd0, 0x33, 0xc2, 0x62, 0xe3, 0x79, 0x9f, 0x26, 0x38, 0x54,
	0x2c, 0x85, 0xa1, 0x45, 0x67, 0x74, 0xe3, 0x54, 0xf1, 0x80, 0x7c, 0x40, 0x5d, 0xe7, 0x28, 0x0c,
	0xde, 0xaa, 0x92, 0x27, 0x69, 0x34, 0xeb, 0x2e, 0x83, 0x64, 0x5f, 0x61, 0x50, 0x0a, 0x8f, 0x5e,
	0xd5, 0xc8, 0x81, 0x1a, 0x8e, 0x35, 0xfb, 0x0d, 0x63, 0x1a, 0x17, 0xcb, 0xdc, 0x57, 0x16, 0x45,
	0xc9, 0x87, 0xf4, 0xbd, 0xd1, 0x81, 0xae, 0x08, 0xb2, 0x9f, 0x30, 0x74, 0x5e, 0xc8, 0x6d, 0xee,
	0xad, 0x90, 0xc8, 0x3f, 0xef, 0xff, 0x84, 0xd0, 0x2a, 0x90, 0x70, 0x0e, 0x3e, 0x4b, 0x34, 0x5e,
	0xe9, 0x26, 0xa7, 0xcb, 0x1c, 0x51, 0x66, 0x74, 0xa4, 0xab, 0x70, 0xab, 0x7f, 0x60, 0x72, 0x8a,
	0x49, 0x5d, 0xa2, 0xe3, 0x63, 0xca, 0x9d, 0xba, 0x17, 0x81, 0xbe, 0x0f, 0x36, 0xda, 0xa3, 0xe3,
	0x93, 0x0f, 0xc1, 0xdb, 0x40, 0xd9, 0x37, 0x48, 0x2c, 0x6e, 0x94, 0xf3, 0x68, 0x1d, 0x9f, 0xee,
	0x1f, 0xe0, 0x08, 0x6e, 0xfe, 0xdf, 0xff, 0xdb, 0x28, 0x5f, 0xb5, 0xc5, 0x5c, 0xea, 0x3a, 0xab,
	0x95, 0xdc, 0xa2, 0xb9, 0xba, 0xcc, 0xb4, 0x7b, 0x6c, 0xd1, 0xee, 0xfe, 0xd2, 0x46, 0x14, 0xed,
	0x3a, 0x33, 0xdb, 0xcd, 0xf5, 0x61, 0x2f, 0x8a, 0x1e, 0xd1, 0x8b, 0xd7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x28, 0xb5, 0x61, 0x39, 0x02, 0x00, 0x00,
}
