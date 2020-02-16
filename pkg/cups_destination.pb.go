// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/cups_destination.proto

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

// Returns all configured printers.
type CupsDestination struct {
	// Name of the printer
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Option name
	OptionName string `protobuf:"bytes,2,opt,name=option_name,json=optionName,proto3" json:"option_name"`
	// Option value
	OptionValue          string   `protobuf:"bytes,3,opt,name=option_value,json=optionValue,proto3" json:"option_value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CupsDestination) Reset()         { *m = CupsDestination{} }
func (m *CupsDestination) String() string { return proto.CompactTextString(m) }
func (*CupsDestination) ProtoMessage()    {}
func (*CupsDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4a721d5ed77726, []int{0}
}

func (m *CupsDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CupsDestination.Unmarshal(m, b)
}
func (m *CupsDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CupsDestination.Marshal(b, m, deterministic)
}
func (m *CupsDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CupsDestination.Merge(m, src)
}
func (m *CupsDestination) XXX_Size() int {
	return xxx_messageInfo_CupsDestination.Size(m)
}
func (m *CupsDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_CupsDestination.DiscardUnknown(m)
}

var xxx_messageInfo_CupsDestination proto.InternalMessageInfo

func (m *CupsDestination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CupsDestination) GetOptionName() string {
	if m != nil {
		return m.OptionName
	}
	return ""
}

func (m *CupsDestination) GetOptionValue() string {
	if m != nil {
		return m.OptionValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CupsDestination)(nil), "schemas.CupsDestination")
}

func init() { proto.RegisterFile("pb/cups_destination.proto", fileDescriptor_9d4a721d5ed77726) }

var fileDescriptor_9d4a721d5ed77726 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0xd2, 0x4f,
	0x2e, 0x2d, 0x28, 0x8e, 0x4f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0xca,
	0xe4, 0xe2, 0x77, 0x2e, 0x2d, 0x28, 0x76, 0x41, 0xa8, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xe4, 0xb9, 0xb8, 0xf3, 0x0b,
	0x40, 0xb2, 0xf1, 0x60, 0x29, 0x26, 0xb0, 0x14, 0x17, 0x44, 0xc8, 0x0f, 0xa4, 0x40, 0x91, 0x8b,
	0x07, 0xaa, 0xa0, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x19, 0xac, 0x02, 0xaa, 0x29, 0x0c, 0x24,
	0xe4, 0x64, 0x14, 0x65, 0x90, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x9b, 0x99, 0x9c, 0x9d, 0x5a, 0x60, 0x6e, 0xa6, 0x9f, 0x5f, 0x5c, 0x58, 0x9a, 0x5a, 0x54, 0xa9,
	0x0b, 0x76, 0x58, 0x52, 0x69, 0x9a, 0x7e, 0x41, 0x76, 0xba, 0x35, 0xd4, 0x79, 0x49, 0x6c, 0x60,
	0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x27, 0x51, 0x17, 0xcb, 0x00, 0x00, 0x00,
}
