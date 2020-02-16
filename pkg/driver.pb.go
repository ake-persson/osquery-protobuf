// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/driver.proto

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

// Details for in-use Windows device drivers. This does not display installed but unused drivers.
type Driver struct {
	// Device ID
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id"`
	// Device name
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name"`
	// Path to driver image file
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image"`
	// Driver description
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// Driver service name
	Service string `protobuf:"bytes,5,opt,name=service,proto3" json:"service"`
	// Driver service registry key
	ServiceKey string `protobuf:"bytes,6,opt,name=service_key,json=serviceKey,proto3" json:"service_key"`
	// Driver version
	Version string `protobuf:"bytes,7,opt,name=version,proto3" json:"version"`
	// Associated inf file
	Inf string `protobuf:"bytes,8,opt,name=inf,proto3" json:"inf"`
	// Device/driver class name
	Class string `protobuf:"bytes,9,opt,name=class,proto3" json:"class"`
	// Driver provider
	Provider string `protobuf:"bytes,10,opt,name=provider,proto3" json:"provider"`
	// Device manufacturer
	Manufacturer string `protobuf:"bytes,11,opt,name=manufacturer,proto3" json:"manufacturer"`
	// Driver key
	DriverKey string `protobuf:"bytes,12,opt,name=driver_key,json=driverKey,proto3" json:"driver_key"`
	// Driver date
	Date int64 `protobuf:"varint,13,opt,name=date,proto3" json:"date"`
	// Whether the driver is signed or not
	Signed               int32    `protobuf:"varint,14,opt,name=signed,proto3" json:"signed"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Driver) Reset()         { *m = Driver{} }
func (m *Driver) String() string { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()    {}
func (*Driver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc778a441993d7f, []int{0}
}

func (m *Driver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Driver.Unmarshal(m, b)
}
func (m *Driver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Driver.Marshal(b, m, deterministic)
}
func (m *Driver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Driver.Merge(m, src)
}
func (m *Driver) XXX_Size() int {
	return xxx_messageInfo_Driver.Size(m)
}
func (m *Driver) XXX_DiscardUnknown() {
	xxx_messageInfo_Driver.DiscardUnknown(m)
}

var xxx_messageInfo_Driver proto.InternalMessageInfo

func (m *Driver) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Driver) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *Driver) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Driver) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Driver) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Driver) GetServiceKey() string {
	if m != nil {
		return m.ServiceKey
	}
	return ""
}

func (m *Driver) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Driver) GetInf() string {
	if m != nil {
		return m.Inf
	}
	return ""
}

func (m *Driver) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *Driver) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Driver) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *Driver) GetDriverKey() string {
	if m != nil {
		return m.DriverKey
	}
	return ""
}

func (m *Driver) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Driver) GetSigned() int32 {
	if m != nil {
		return m.Signed
	}
	return 0
}

func init() {
	proto.RegisterType((*Driver)(nil), "schemas.Driver")
}

func init() { proto.RegisterFile("pb/driver.proto", fileDescriptor_7dc778a441993d7f) }

var fileDescriptor_7dc778a441993d7f = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0xc9, 0x4b, 0x9b, 0x36, 0xd3, 0xbe, 0xf7, 0x64, 0x11, 0x59, 0x14, 0x31, 0xf4, 0xd4,
	0x8b, 0x8d, 0x28, 0xe8, 0xc1, 0x9b, 0x78, 0x11, 0xc1, 0x43, 0x8f, 0x5e, 0xca, 0x66, 0x77, 0x9a,
	0x2e, 0x75, 0x93, 0xb8, 0x9b, 0x04, 0xfa, 0x0f, 0xf8, 0x77, 0xcb, 0xce, 0xa6, 0xa0, 0xb7, 0xf9,
	0x7e, 0x5f, 0x66, 0xe6, 0xcb, 0x0e, 0xfc, 0x6f, 0x8a, 0x5c, 0x59, 0xdd, 0xa3, 0x5d, 0x35, 0xb6,
	0x6e, 0x6b, 0x36, 0x71, 0x72, 0x87, 0x46, 0xb8, 0xc5, 0x57, 0x0c, 0xc9, 0x33, 0x39, 0xec, 0x02,
	0x52, 0x85, 0xbd, 0x96, 0xb8, 0xd1, 0x8a, 0x47, 0x59, 0xb4, 0x4c, 0xd7, 0xd3, 0x00, 0x5e, 0x14,
	0xbb, 0x82, 0xd9, 0x60, 0x56, 0xc2, 0x20, 0xff, 0x43, 0x36, 0x04, 0xf4, 0x26, 0x0c, 0xb2, 0x53,
	0x18, 0x6b, 0x23, 0x4a, 0xe4, 0x31, 0x59, 0x41, 0xb0, 0xcc, 0xb7, 0x39, 0x69, 0x75, 0xd3, 0xea,
	0xba, 0xe2, 0x23, 0xf2, 0x7e, 0x22, 0xc6, 0x61, 0xe2, 0xd0, 0xfa, 0x31, 0x7c, 0x4c, 0xee, 0x51,
	0xfa, 0x95, 0x43, 0xb9, 0xd9, 0xe3, 0x81, 0x27, 0x61, 0xe5, 0x80, 0x5e, 0xf1, 0xe0, 0x5b, 0x7b,
	0xb4, 0xce, 0x0f, 0x9e, 0x84, 0xd6, 0x41, 0xb2, 0x13, 0x88, 0x75, 0xb5, 0xe5, 0x53, 0xa2, 0xbe,
	0xf4, 0xf1, 0xe4, 0x87, 0x70, 0x8e, 0xa7, 0x21, 0x1e, 0x09, 0x76, 0x0e, 0xd3, 0xc6, 0xd6, 0xbd,
	0x56, 0x68, 0x39, 0x84, 0x3f, 0x3e, 0x6a, 0xb6, 0x80, 0xb9, 0x11, 0x55, 0xb7, 0x15, 0xb2, 0xed,
	0x2c, 0x5a, 0x3e, 0x23, 0xff, 0x17, 0x63, 0x97, 0x00, 0xe1, 0x59, 0x29, 0xe1, 0x9c, 0xbe, 0x48,
	0x03, 0xf1, 0x01, 0x19, 0x8c, 0x94, 0x68, 0x91, 0xff, 0xcd, 0xa2, 0x65, 0xbc, 0xa6, 0x9a, 0x9d,
	0x41, 0xe2, 0x74, 0x59, 0xa1, 0xe2, 0xff, 0xb2, 0x68, 0x39, 0x5e, 0x0f, 0xea, 0xe9, 0xf6, 0xfd,
	0xa6, 0xd4, 0xed, 0xae, 0x2b, 0x56, 0xb2, 0x36, 0xb9, 0xd1, 0x72, 0x8f, 0xcd, 0xc3, 0x7d, 0x5e,
	0xbb, 0xcf, 0x0e, 0xed, 0xe1, 0x9a, 0xce, 0x56, 0x74, 0xdb, 0xbc, 0xd9, 0x97, 0x8f, 0xc3, 0xf1,
	0x8a, 0x84, 0xe8, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xc0, 0x29, 0xc5, 0xdf, 0x01,
	0x00, 0x00,
}
