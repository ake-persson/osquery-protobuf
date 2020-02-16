// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/load_average.proto

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

// Displays information about the system wide load averages.
type LoadAverage struct {
	// Period over which the average is calculated.
	Period string `protobuf:"bytes,1,opt,name=period,proto3" json:"period"`
	// Load average over the specified period.
	Average              string   `protobuf:"bytes,2,opt,name=average,proto3" json:"average"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadAverage) Reset()         { *m = LoadAverage{} }
func (m *LoadAverage) String() string { return proto.CompactTextString(m) }
func (*LoadAverage) ProtoMessage()    {}
func (*LoadAverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa4b7582a8e58f66, []int{0}
}

func (m *LoadAverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadAverage.Unmarshal(m, b)
}
func (m *LoadAverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadAverage.Marshal(b, m, deterministic)
}
func (m *LoadAverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadAverage.Merge(m, src)
}
func (m *LoadAverage) XXX_Size() int {
	return xxx_messageInfo_LoadAverage.Size(m)
}
func (m *LoadAverage) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadAverage.DiscardUnknown(m)
}

var xxx_messageInfo_LoadAverage proto.InternalMessageInfo

func (m *LoadAverage) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *LoadAverage) GetAverage() string {
	if m != nil {
		return m.Average
	}
	return ""
}

func init() {
	proto.RegisterType((*LoadAverage)(nil), "schemas.LoadAverage")
}

func init() { proto.RegisterFile("pb/load_average.proto", fileDescriptor_aa4b7582a8e58f66) }

var fileDescriptor_aa4b7582a8e58f66 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x48, 0xd2, 0xcf,
	0xc9, 0x4f, 0x4c, 0x89, 0x4f, 0x2c, 0x4b, 0x2d, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0xb2, 0xe7, 0xe2, 0xf6, 0xc9,
	0x4f, 0x4c, 0x71, 0x84, 0xc8, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa4, 0x16, 0x65, 0xe6, 0xa7, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0x50, 0x03, 0x24, 0x98,
	0xc0, 0x12, 0x30, 0xae, 0x93, 0x51, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x6e, 0x66, 0x72, 0x76, 0x6a, 0x81, 0xb9, 0x99, 0x7e, 0x7e, 0x71, 0x61, 0x69,
	0x6a, 0x51, 0xa5, 0x2e, 0xd8, 0xba, 0xa4, 0xd2, 0x34, 0xfd, 0x82, 0xec, 0x74, 0x6b, 0xa8, 0xa5,
	0x49, 0x6c, 0x60, 0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xb6, 0x2a, 0x2c, 0x9d,
	0x00, 0x00, 0x00,
}
