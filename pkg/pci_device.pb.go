// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/pci_device.proto

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

// PCI devices active on the host system.
type PciDevice struct {
	// PCI Device used slot
	PciSlot string `protobuf:"bytes,1,opt,name=pci_slot,json=pciSlot,proto3" json:"pci_slot"`
	// PCI Device class
	PciClass string `protobuf:"bytes,2,opt,name=pci_class,json=pciClass,proto3" json:"pci_class"`
	// PCI Device used driver
	Driver string `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver"`
	// PCI Device vendor
	Vendor string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor"`
	// Hex encoded PCI Device vendor identifier
	VendorId string `protobuf:"bytes,5,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id"`
	// PCI Device model
	Model string `protobuf:"bytes,6,opt,name=model,proto3" json:"model"`
	// Hex encoded PCI Device model identifier
	ModelId string `protobuf:"bytes,7,opt,name=model_id,json=modelId,proto3" json:"model_id"`
	// PCI Device subsystem
	Subsystem string `protobuf:"bytes,8,opt,name=subsystem,proto3" json:"subsystem"`
	// 1 If PCI device is express else 0
	Express int32 `protobuf:"varint,9,opt,name=express,proto3" json:"express"`
	// 1 If PCI device is thunderbolt else 0
	Thunderbolt int32 `protobuf:"varint,10,opt,name=thunderbolt,proto3" json:"thunderbolt"`
	// 1 If PCI device is removable else 0
	Removable int32 `protobuf:"varint,11,opt,name=removable,proto3" json:"removable"`
	// PCI Device class ID in hex format
	PciClassId string `protobuf:"bytes,12,opt,name=pci_class_id,json=pciClassId,proto3" json:"pci_class_id"`
	// PCI Device  subclass in hex format
	PciSubclassId string `protobuf:"bytes,13,opt,name=pci_subclass_id,json=pciSubclassId,proto3" json:"pci_subclass_id"`
	// PCI Device subclass
	PciSubclass string `protobuf:"bytes,14,opt,name=pci_subclass,json=pciSubclass,proto3" json:"pci_subclass"`
	// Vendor ID of PCI device subsystem
	SubsystemVendorId string `protobuf:"bytes,15,opt,name=subsystem_vendor_id,json=subsystemVendorId,proto3" json:"subsystem_vendor_id"`
	// Vendor of PCI device subsystem
	SubsystemVendor string `protobuf:"bytes,16,opt,name=subsystem_vendor,json=subsystemVendor,proto3" json:"subsystem_vendor"`
	// Model ID of PCI device subsystem
	SubsystemModelId string `protobuf:"bytes,17,opt,name=subsystem_model_id,json=subsystemModelId,proto3" json:"subsystem_model_id"`
	// Device description of PCI device subsystem
	SubsystemModel       string   `protobuf:"bytes,18,opt,name=subsystem_model,json=subsystemModel,proto3" json:"subsystem_model"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PciDevice) Reset()         { *m = PciDevice{} }
func (m *PciDevice) String() string { return proto.CompactTextString(m) }
func (*PciDevice) ProtoMessage()    {}
func (*PciDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_11799b0446d7b0a5, []int{0}
}

func (m *PciDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PciDevice.Unmarshal(m, b)
}
func (m *PciDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PciDevice.Marshal(b, m, deterministic)
}
func (m *PciDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PciDevice.Merge(m, src)
}
func (m *PciDevice) XXX_Size() int {
	return xxx_messageInfo_PciDevice.Size(m)
}
func (m *PciDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_PciDevice.DiscardUnknown(m)
}

var xxx_messageInfo_PciDevice proto.InternalMessageInfo

func (m *PciDevice) GetPciSlot() string {
	if m != nil {
		return m.PciSlot
	}
	return ""
}

func (m *PciDevice) GetPciClass() string {
	if m != nil {
		return m.PciClass
	}
	return ""
}

func (m *PciDevice) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *PciDevice) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *PciDevice) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *PciDevice) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *PciDevice) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *PciDevice) GetSubsystem() string {
	if m != nil {
		return m.Subsystem
	}
	return ""
}

func (m *PciDevice) GetExpress() int32 {
	if m != nil {
		return m.Express
	}
	return 0
}

func (m *PciDevice) GetThunderbolt() int32 {
	if m != nil {
		return m.Thunderbolt
	}
	return 0
}

func (m *PciDevice) GetRemovable() int32 {
	if m != nil {
		return m.Removable
	}
	return 0
}

func (m *PciDevice) GetPciClassId() string {
	if m != nil {
		return m.PciClassId
	}
	return ""
}

func (m *PciDevice) GetPciSubclassId() string {
	if m != nil {
		return m.PciSubclassId
	}
	return ""
}

func (m *PciDevice) GetPciSubclass() string {
	if m != nil {
		return m.PciSubclass
	}
	return ""
}

func (m *PciDevice) GetSubsystemVendorId() string {
	if m != nil {
		return m.SubsystemVendorId
	}
	return ""
}

func (m *PciDevice) GetSubsystemVendor() string {
	if m != nil {
		return m.SubsystemVendor
	}
	return ""
}

func (m *PciDevice) GetSubsystemModelId() string {
	if m != nil {
		return m.SubsystemModelId
	}
	return ""
}

func (m *PciDevice) GetSubsystemModel() string {
	if m != nil {
		return m.SubsystemModel
	}
	return ""
}

func init() {
	proto.RegisterType((*PciDevice)(nil), "schemas.PciDevice")
}

func init() { proto.RegisterFile("pb/pci_device.proto", fileDescriptor_11799b0446d7b0a5) }

var fileDescriptor_11799b0446d7b0a5 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x3b, 0xef, 0xd3, 0x30,
	0x14, 0xc5, 0x55, 0xf8, 0xb7, 0x69, 0x6e, 0x1f, 0x69, 0x5d, 0x84, 0x8c, 0x60, 0x08, 0x0c, 0x50,
	0x24, 0x68, 0x10, 0x48, 0x30, 0xb0, 0x01, 0x4b, 0x06, 0x24, 0x54, 0x24, 0x06, 0x96, 0xaa, 0x7e,
	0xd0, 0x5a, 0x4d, 0xea, 0x60, 0x27, 0x11, 0xfd, 0xa6, 0x7c, 0x1c, 0xe4, 0xeb, 0x3c, 0x4a, 0x37,
	0x9f, 0x73, 0x7e, 0xb9, 0xb9, 0x3e, 0x32, 0xac, 0x0a, 0x96, 0x14, 0x5c, 0xed, 0x84, 0xac, 0x15,
	0x97, 0x9b, 0xc2, 0xe8, 0x52, 0x93, 0xc0, 0xf2, 0xa3, 0xcc, 0xf7, 0xf6, 0xd9, 0xdf, 0x3b, 0x08,
	0xbf, 0x71, 0xf5, 0x05, 0x43, 0xf2, 0x08, 0xc6, 0x0e, 0xb5, 0x99, 0x2e, 0xe9, 0x20, 0x1e, 0xac,
	0xc3, 0x6d, 0x50, 0x70, 0xf5, 0x3d, 0xd3, 0x25, 0x79, 0x0c, 0xa1, 0x8b, 0x78, 0xb6, 0xb7, 0x96,
	0xde, 0xc3, 0xcc, 0xb1, 0x9f, 0x9d, 0x26, 0x0f, 0x61, 0x24, 0x8c, 0xaa, 0xa5, 0xa1, 0xf7, 0x31,
	0x69, 0x94, 0xf3, 0x6b, 0x79, 0x16, 0xda, 0xd0, 0x3b, 0xef, 0x7b, 0xe5, 0x86, 0xf9, 0xd3, 0x4e,
	0x09, 0x3a, 0xf4, 0xc3, 0xbc, 0x91, 0x0a, 0xf2, 0x00, 0x86, 0xb9, 0x16, 0x32, 0xa3, 0x23, 0x0c,
	0xbc, 0x70, 0xab, 0xe1, 0xc1, 0x7d, 0x11, 0xf8, 0xd5, 0x50, 0xa7, 0x82, 0x3c, 0x81, 0xd0, 0x56,
	0xcc, 0x5e, 0x6c, 0x29, 0x73, 0x3a, 0xc6, 0xac, 0x37, 0x08, 0x85, 0x40, 0xfe, 0x29, 0x8c, 0xb4,
	0x96, 0x86, 0xf1, 0x60, 0x3d, 0xdc, 0xb6, 0x92, 0xc4, 0x30, 0x29, 0x8f, 0xd5, 0x59, 0x48, 0xc3,
	0x74, 0x56, 0x52, 0xc0, 0xf4, 0xda, 0x72, 0x93, 0x8d, 0xcc, 0x75, 0xbd, 0x67, 0x99, 0xa4, 0x13,
	0xcc, 0x7b, 0x83, 0xc4, 0x30, 0xed, 0x2a, 0x71, 0x6b, 0x4d, 0xf1, 0xd7, 0xd0, 0xb6, 0x92, 0x0a,
	0xf2, 0x1c, 0x22, 0xec, 0xb3, 0x62, 0x1d, 0x34, 0x43, 0x68, 0xe6, 0x6a, 0x6d, 0xdc, 0x54, 0x90,
	0xa7, 0x7e, 0x52, 0xcb, 0xd1, 0x39, 0x42, 0x93, 0x2b, 0x88, 0x6c, 0x60, 0xd5, 0xdd, 0x69, 0xd7,
	0x97, 0x17, 0x21, 0xb9, 0xec, 0xa2, 0x1f, 0x6d, 0x8b, 0x2f, 0x61, 0x71, 0xcb, 0xd3, 0x05, 0xc2,
	0xd1, 0x0d, 0x4c, 0x5e, 0x01, 0xe9, 0xd1, 0xae, 0xe4, 0x25, 0xc2, 0xfd, 0x90, 0xaf, 0x4d, 0xdb,
	0x2f, 0x20, 0xba, 0xa1, 0x29, 0x41, 0x74, 0xfe, 0x3f, 0xfa, 0xe9, 0xed, 0xcf, 0x37, 0x07, 0x55,
	0x1e, 0x2b, 0xb6, 0xe1, 0x3a, 0x4f, 0x72, 0xc5, 0x4f, 0xb2, 0xf8, 0xf0, 0x3e, 0xd1, 0xf6, 0x77,
	0x25, 0xcd, 0xe5, 0x35, 0x3e, 0x44, 0x56, 0xfd, 0x4a, 0x8a, 0xd3, 0xe1, 0x63, 0xf3, 0x1c, 0xd9,
	0x08, 0xdd, 0x77, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x3e, 0x8a, 0x1b, 0xb5, 0x02, 0x00,
	0x00,
}
