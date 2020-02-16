// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/sharing_preference.proto

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

// OS X Sharing preferences.
type SharingPreference struct {
	// 1 If screen sharing is enabled else 0
	ScreenSharing int32 `protobuf:"varint,1,opt,name=screen_sharing,json=screenSharing,proto3" json:"screen_sharing"`
	// 1 If file sharing is enabled else 0
	FileSharing int32 `protobuf:"varint,2,opt,name=file_sharing,json=fileSharing,proto3" json:"file_sharing"`
	// 1 If printer sharing is enabled else 0
	PrinterSharing int32 `protobuf:"varint,3,opt,name=printer_sharing,json=printerSharing,proto3" json:"printer_sharing"`
	// 1 If remote login is enabled else 0
	RemoteLogin int32 `protobuf:"varint,4,opt,name=remote_login,json=remoteLogin,proto3" json:"remote_login"`
	// 1 If remote management is enabled else 0
	RemoteManagement int32 `protobuf:"varint,5,opt,name=remote_management,json=remoteManagement,proto3" json:"remote_management"`
	// 1 If remote apple events are enabled else 0
	RemoteAppleEvents int32 `protobuf:"varint,6,opt,name=remote_apple_events,json=remoteAppleEvents,proto3" json:"remote_apple_events"`
	// 1 If internet sharing is enabled else 0
	InternetSharing int32 `protobuf:"varint,7,opt,name=internet_sharing,json=internetSharing,proto3" json:"internet_sharing"`
	// 1 If bluetooth sharing is enabled for any user else 0
	BluetoothSharing int32 `protobuf:"varint,8,opt,name=bluetooth_sharing,json=bluetoothSharing,proto3" json:"bluetooth_sharing"`
	// 1 If CD or DVD sharing is enabled else 0
	DiscSharing int32 `protobuf:"varint,9,opt,name=disc_sharing,json=discSharing,proto3" json:"disc_sharing"`
	// 1 If content caching is enabled else 0
	ContentCaching       int32    `protobuf:"varint,10,opt,name=content_caching,json=contentCaching,proto3" json:"content_caching"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharingPreference) Reset()         { *m = SharingPreference{} }
func (m *SharingPreference) String() string { return proto.CompactTextString(m) }
func (*SharingPreference) ProtoMessage()    {}
func (*SharingPreference) Descriptor() ([]byte, []int) {
	return fileDescriptor_791ceeaf473177bf, []int{0}
}

func (m *SharingPreference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharingPreference.Unmarshal(m, b)
}
func (m *SharingPreference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharingPreference.Marshal(b, m, deterministic)
}
func (m *SharingPreference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharingPreference.Merge(m, src)
}
func (m *SharingPreference) XXX_Size() int {
	return xxx_messageInfo_SharingPreference.Size(m)
}
func (m *SharingPreference) XXX_DiscardUnknown() {
	xxx_messageInfo_SharingPreference.DiscardUnknown(m)
}

var xxx_messageInfo_SharingPreference proto.InternalMessageInfo

func (m *SharingPreference) GetScreenSharing() int32 {
	if m != nil {
		return m.ScreenSharing
	}
	return 0
}

func (m *SharingPreference) GetFileSharing() int32 {
	if m != nil {
		return m.FileSharing
	}
	return 0
}

func (m *SharingPreference) GetPrinterSharing() int32 {
	if m != nil {
		return m.PrinterSharing
	}
	return 0
}

func (m *SharingPreference) GetRemoteLogin() int32 {
	if m != nil {
		return m.RemoteLogin
	}
	return 0
}

func (m *SharingPreference) GetRemoteManagement() int32 {
	if m != nil {
		return m.RemoteManagement
	}
	return 0
}

func (m *SharingPreference) GetRemoteAppleEvents() int32 {
	if m != nil {
		return m.RemoteAppleEvents
	}
	return 0
}

func (m *SharingPreference) GetInternetSharing() int32 {
	if m != nil {
		return m.InternetSharing
	}
	return 0
}

func (m *SharingPreference) GetBluetoothSharing() int32 {
	if m != nil {
		return m.BluetoothSharing
	}
	return 0
}

func (m *SharingPreference) GetDiscSharing() int32 {
	if m != nil {
		return m.DiscSharing
	}
	return 0
}

func (m *SharingPreference) GetContentCaching() int32 {
	if m != nil {
		return m.ContentCaching
	}
	return 0
}

func init() {
	proto.RegisterType((*SharingPreference)(nil), "schemas.SharingPreference")
}

func init() { proto.RegisterFile("pb/sharing_preference.proto", fileDescriptor_791ceeaf473177bf) }

var fileDescriptor_791ceeaf473177bf = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x4a, 0xf4, 0x30,
	0x18, 0x85, 0x99, 0x6f, 0xbe, 0x99, 0xd1, 0xa8, 0xf3, 0x13, 0x37, 0x05, 0x37, 0x2a, 0x88, 0xca,
	0x60, 0x2b, 0x0a, 0xba, 0x70, 0xa5, 0xe2, 0x4e, 0x41, 0x74, 0xe7, 0xa6, 0xa4, 0xf1, 0x6d, 0x1b,
	0xa6, 0xf9, 0x31, 0x49, 0x05, 0x6f, 0xc9, 0xab, 0x94, 0x26, 0x69, 0xe8, 0xf6, 0x39, 0x4f, 0x0e,
	0x27, 0x09, 0x3a, 0x50, 0x45, 0x66, 0x6a, 0xa2, 0x99, 0xa8, 0x72, 0xa5, 0xa1, 0x04, 0x0d, 0x82,
	0x42, 0xaa, 0xb4, 0xb4, 0x12, 0xcf, 0x0c, 0xad, 0x81, 0x13, 0x73, 0xfc, 0x3b, 0x46, 0xab, 0x77,
	0x6f, 0xbd, 0x46, 0x09, 0x9f, 0xa0, 0xb9, 0xa1, 0x1a, 0x40, 0xe4, 0xa1, 0x21, 0x19, 0x1d, 0x8e,
	0xce, 0x26, 0x6f, 0x7b, 0x9e, 0x86, 0x03, 0xf8, 0x08, 0xed, 0x96, 0xac, 0x81, 0x28, 0xfd, 0x73,
	0xd2, 0x4e, 0xc7, 0x7a, 0xe5, 0x14, 0x2d, 0x94, 0x66, 0xc2, 0x82, 0x8e, 0xd6, 0xd8, 0x59, 0xf3,
	0x80, 0x07, 0x5d, 0x1a, 0xb8, 0xb4, 0x90, 0x37, 0xb2, 0x62, 0x22, 0xf9, 0xef, 0xbb, 0x3c, 0x7b,
	0xee, 0x10, 0x5e, 0xa3, 0x55, 0x50, 0x38, 0x11, 0xa4, 0x02, 0x0e, 0xc2, 0x26, 0x13, 0xe7, 0x2d,
	0x7d, 0xf0, 0x12, 0x39, 0x4e, 0xd1, 0x7e, 0x90, 0x89, 0x52, 0x0d, 0xe4, 0xf0, 0x0d, 0xc2, 0x9a,
	0x64, 0xea, 0xf4, 0xd0, 0x73, 0xdf, 0x25, 0x4f, 0x2e, 0xc0, 0xe7, 0x68, 0xe9, 0xf6, 0x08, 0xb0,
	0x71, 0xe9, 0xcc, 0xc9, 0x8b, 0x9e, 0xf7, 0x53, 0xd7, 0x68, 0x55, 0x34, 0x2d, 0x58, 0x29, 0x6d,
	0x1d, 0xdd, 0x2d, 0xbf, 0x23, 0x06, 0x83, 0x7b, 0x7d, 0x32, 0x43, 0xa3, 0xb7, 0xed, 0xef, 0xd5,
	0xb1, 0xc1, 0x1b, 0x51, 0x29, 0x2c, 0x08, 0x9b, 0x53, 0x42, 0xeb, 0xce, 0x42, 0xfe, 0x8d, 0x02,
	0x7e, 0xf4, 0xf4, 0xe1, 0xea, 0xe3, 0xb2, 0x62, 0xb6, 0x6e, 0x8b, 0x94, 0x4a, 0x9e, 0x71, 0x46,
	0x37, 0xa0, 0x6e, 0x6f, 0x32, 0x69, 0xbe, 0x5a, 0xd0, 0x3f, 0x17, 0xee, 0x6b, 0x8b, 0xb6, 0xcc,
	0xd4, 0xa6, 0xba, 0x0b, 0x1f, 0x5c, 0x4c, 0x1d, 0xbd, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0x0a, 0xc4, 0xb0, 0x0f, 0x02, 0x00, 0x00,
}
