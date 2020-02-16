// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/interface_detail.proto

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

// Detailed information and stats of network interfaces.
type InterfaceDetail struct {
	// Interface name
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface"`
	// MAC of interface (optional)
	Mac string `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac"`
	// Interface type (includes virtual)
	Type int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type"`
	// Network MTU
	Mtu int32 `protobuf:"varint,4,opt,name=mtu,proto3" json:"mtu"`
	// Metric based on the speed of the interface
	Metric int32 `protobuf:"varint,5,opt,name=metric,proto3" json:"metric"`
	// Flags (netdevice) for the device
	Flags int32 `protobuf:"varint,6,opt,name=flags,proto3" json:"flags"`
	// Input packets
	Ipackets int64 `protobuf:"varint,7,opt,name=ipackets,proto3" json:"ipackets"`
	// Output packets
	Opackets int64 `protobuf:"varint,8,opt,name=opackets,proto3" json:"opackets"`
	// Input bytes
	Ibytes int64 `protobuf:"varint,9,opt,name=ibytes,proto3" json:"ibytes"`
	// Output bytes
	Obytes int64 `protobuf:"varint,10,opt,name=obytes,proto3" json:"obytes"`
	// Input errors
	Ierrors int64 `protobuf:"varint,11,opt,name=ierrors,proto3" json:"ierrors"`
	// Output errors
	Oerrors int64 `protobuf:"varint,12,opt,name=oerrors,proto3" json:"oerrors"`
	// Input drops
	Idrops int64 `protobuf:"varint,13,opt,name=idrops,proto3" json:"idrops"`
	// Output drops
	Odrops int64 `protobuf:"varint,14,opt,name=odrops,proto3" json:"odrops"`
	// Packet Collisions detected
	Collisions int64 `protobuf:"varint,15,opt,name=collisions,proto3" json:"collisions"`
	// Time of last device modification (optional)
	LastChange int64 `protobuf:"varint,16,opt,name=last_change,json=lastChange,proto3" json:"last_change"`
	// Interface speed in Mb/s
	LinkSpeed int64 `protobuf:"varint,17,opt,name=link_speed,json=linkSpeed,proto3" json:"link_speed"`
	// PCI slot number
	PciSlot string `protobuf:"bytes,18,opt,name=pci_slot,json=pciSlot,proto3" json:"pci_slot"`
	// The friendly display name of the interface.
	FriendlyName string `protobuf:"bytes,19,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name"`
	// Short description of the object a one-line string.
	Description string `protobuf:"bytes,20,opt,name=description,proto3" json:"description"`
	// Name of the network adapter's manufacturer.
	Manufacturer string `protobuf:"bytes,21,opt,name=manufacturer,proto3" json:"manufacturer"`
	// Name of the network connection as it appears in the Network Connections Control Panel program.
	ConnectionId string `protobuf:"bytes,22,opt,name=connection_id,json=connectionId,proto3" json:"connection_id"`
	// State of the network adapter connection to the network.
	ConnectionStatus string `protobuf:"bytes,23,opt,name=connection_status,json=connectionStatus,proto3" json:"connection_status"`
	// Indicates whether the adapter is enabled or not.
	Enabled int32 `protobuf:"varint,24,opt,name=enabled,proto3" json:"enabled"`
	// Indicates whether the adapter is a physical or a logical adapter.
	PhysicalAdapter int32 `protobuf:"varint,25,opt,name=physical_adapter,json=physicalAdapter,proto3" json:"physical_adapter"`
	// Estimate of the current bandwidth in bits per second.
	Speed int32 `protobuf:"varint,26,opt,name=speed,proto3" json:"speed"`
	// The name of the service the network adapter uses.
	Service string `protobuf:"bytes,27,opt,name=service,proto3" json:"service"`
	// If TRUE
	DhcpEnabled int32 `protobuf:"varint,28,opt,name=dhcp_enabled,json=dhcpEnabled,proto3" json:"dhcp_enabled"`
	// Expiration date and time for a leased IP address that was assigned to the computer by the dynamic host configuration protocol (DHCP) server.
	DhcpLeaseExpires string `protobuf:"bytes,29,opt,name=dhcp_lease_expires,json=dhcpLeaseExpires,proto3" json:"dhcp_lease_expires"`
	// Date and time the lease was obtained for the IP address assigned to the computer by the dynamic host configuration protocol (DHCP) server.
	DhcpLeaseObtained string `protobuf:"bytes,30,opt,name=dhcp_lease_obtained,json=dhcpLeaseObtained,proto3" json:"dhcp_lease_obtained"`
	// IP address of the dynamic host configuration protocol (DHCP) server.
	DhcpServer string `protobuf:"bytes,31,opt,name=dhcp_server,json=dhcpServer,proto3" json:"dhcp_server"`
	// Organization name followed by a period and an extension that indicates the type of organization
	DnsDomain string `protobuf:"bytes,32,opt,name=dns_domain,json=dnsDomain,proto3" json:"dns_domain"`
	// Array of DNS domain suffixes to be appended to the end of host names during name resolution.
	DnsDomainSuffixSearchOrder string `protobuf:"bytes,33,opt,name=dns_domain_suffix_search_order,json=dnsDomainSuffixSearchOrder,proto3" json:"dns_domain_suffix_search_order"`
	// Host name used to identify the local computer for authentication by some utilities.
	DnsHostName string `protobuf:"bytes,34,opt,name=dns_host_name,json=dnsHostName,proto3" json:"dns_host_name"`
	// Array of server IP addresses to be used in querying for DNS servers.
	DnsServerSearchOrder string   `protobuf:"bytes,35,opt,name=dns_server_search_order,json=dnsServerSearchOrder,proto3" json:"dns_server_search_order"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceDetail) Reset()         { *m = InterfaceDetail{} }
func (m *InterfaceDetail) String() string { return proto.CompactTextString(m) }
func (*InterfaceDetail) ProtoMessage()    {}
func (*InterfaceDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_3175305c5c677267, []int{0}
}

func (m *InterfaceDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceDetail.Unmarshal(m, b)
}
func (m *InterfaceDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceDetail.Marshal(b, m, deterministic)
}
func (m *InterfaceDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceDetail.Merge(m, src)
}
func (m *InterfaceDetail) XXX_Size() int {
	return xxx_messageInfo_InterfaceDetail.Size(m)
}
func (m *InterfaceDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceDetail.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceDetail proto.InternalMessageInfo

func (m *InterfaceDetail) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *InterfaceDetail) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *InterfaceDetail) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *InterfaceDetail) GetMtu() int32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceDetail) GetMetric() int32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *InterfaceDetail) GetFlags() int32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *InterfaceDetail) GetIpackets() int64 {
	if m != nil {
		return m.Ipackets
	}
	return 0
}

func (m *InterfaceDetail) GetOpackets() int64 {
	if m != nil {
		return m.Opackets
	}
	return 0
}

func (m *InterfaceDetail) GetIbytes() int64 {
	if m != nil {
		return m.Ibytes
	}
	return 0
}

func (m *InterfaceDetail) GetObytes() int64 {
	if m != nil {
		return m.Obytes
	}
	return 0
}

func (m *InterfaceDetail) GetIerrors() int64 {
	if m != nil {
		return m.Ierrors
	}
	return 0
}

func (m *InterfaceDetail) GetOerrors() int64 {
	if m != nil {
		return m.Oerrors
	}
	return 0
}

func (m *InterfaceDetail) GetIdrops() int64 {
	if m != nil {
		return m.Idrops
	}
	return 0
}

func (m *InterfaceDetail) GetOdrops() int64 {
	if m != nil {
		return m.Odrops
	}
	return 0
}

func (m *InterfaceDetail) GetCollisions() int64 {
	if m != nil {
		return m.Collisions
	}
	return 0
}

func (m *InterfaceDetail) GetLastChange() int64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *InterfaceDetail) GetLinkSpeed() int64 {
	if m != nil {
		return m.LinkSpeed
	}
	return 0
}

func (m *InterfaceDetail) GetPciSlot() string {
	if m != nil {
		return m.PciSlot
	}
	return ""
}

func (m *InterfaceDetail) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *InterfaceDetail) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InterfaceDetail) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *InterfaceDetail) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *InterfaceDetail) GetConnectionStatus() string {
	if m != nil {
		return m.ConnectionStatus
	}
	return ""
}

func (m *InterfaceDetail) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *InterfaceDetail) GetPhysicalAdapter() int32 {
	if m != nil {
		return m.PhysicalAdapter
	}
	return 0
}

func (m *InterfaceDetail) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *InterfaceDetail) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *InterfaceDetail) GetDhcpEnabled() int32 {
	if m != nil {
		return m.DhcpEnabled
	}
	return 0
}

func (m *InterfaceDetail) GetDhcpLeaseExpires() string {
	if m != nil {
		return m.DhcpLeaseExpires
	}
	return ""
}

func (m *InterfaceDetail) GetDhcpLeaseObtained() string {
	if m != nil {
		return m.DhcpLeaseObtained
	}
	return ""
}

func (m *InterfaceDetail) GetDhcpServer() string {
	if m != nil {
		return m.DhcpServer
	}
	return ""
}

func (m *InterfaceDetail) GetDnsDomain() string {
	if m != nil {
		return m.DnsDomain
	}
	return ""
}

func (m *InterfaceDetail) GetDnsDomainSuffixSearchOrder() string {
	if m != nil {
		return m.DnsDomainSuffixSearchOrder
	}
	return ""
}

func (m *InterfaceDetail) GetDnsHostName() string {
	if m != nil {
		return m.DnsHostName
	}
	return ""
}

func (m *InterfaceDetail) GetDnsServerSearchOrder() string {
	if m != nil {
		return m.DnsServerSearchOrder
	}
	return ""
}

func init() {
	proto.RegisterType((*InterfaceDetail)(nil), "schemas.InterfaceDetail")
}

func init() { proto.RegisterFile("pb/interface_detail.proto", fileDescriptor_3175305c5c677267) }

var fileDescriptor_3175305c5c677267 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x94, 0x4f, 0x6f, 0x13, 0x3b,
	0x14, 0xc5, 0x95, 0xd7, 0x3f, 0x49, 0x9c, 0xf4, 0x35, 0x75, 0xfb, 0x5a, 0xb7, 0xaf, 0x7f, 0xd2,
	0x74, 0x53, 0x04, 0x34, 0x08, 0x04, 0x2c, 0x58, 0x51, 0x5a, 0x89, 0x4a, 0x88, 0x4a, 0xcd, 0x8e,
	0xcd, 0xc8, 0xb1, 0x6f, 0x12, 0x2b, 0x33, 0xb6, 0xb1, 0x3d, 0xa8, 0xf9, 0xd2, 0x7c, 0x06, 0xe4,
	0xeb, 0x99, 0x24, 0xec, 0x7c, 0x7e, 0xe7, 0xd8, 0xc7, 0x99, 0xb9, 0x19, 0x72, 0x6c, 0xc7, 0x43,
	0xa5, 0x03, 0xb8, 0x09, 0x17, 0x90, 0x49, 0x08, 0x5c, 0xe5, 0x37, 0xd6, 0x99, 0x60, 0x68, 0xd3,
	0x8b, 0x19, 0x14, 0xdc, 0x0f, 0x7e, 0xb7, 0xc8, 0xee, 0x43, 0x9d, 0xb9, 0xc3, 0x08, 0x3d, 0x25,
	0xed, 0xe5, 0x36, 0xd6, 0xe8, 0x37, 0xae, 0xdb, 0x4f, 0x2b, 0x40, 0x7b, 0x64, 0xa3, 0xe0, 0x82,
	0xfd, 0x83, 0x3c, 0x2e, 0x29, 0x25, 0x9b, 0x61, 0x61, 0x81, 0x6d, 0xf4, 0x1b, 0xd7, 0x5b, 0x4f,
	0xb8, 0xc6, 0x54, 0x28, 0xd9, 0x26, 0xa2, 0xb8, 0xa4, 0x87, 0x64, 0xbb, 0x80, 0xe0, 0x94, 0x60,
	0x5b, 0x08, 0x2b, 0x45, 0x0f, 0xc8, 0xd6, 0x24, 0xe7, 0x53, 0xcf, 0xb6, 0x11, 0x27, 0x41, 0x4f,
	0x48, 0x4b, 0x59, 0x2e, 0xe6, 0x10, 0x3c, 0x6b, 0xf6, 0x1b, 0xd7, 0x1b, 0x4f, 0x4b, 0x1d, 0x3d,
	0x53, 0x7b, 0xad, 0xe4, 0xd5, 0x3a, 0xb6, 0xa8, 0xf1, 0x22, 0x80, 0x67, 0x6d, 0x74, 0x2a, 0x15,
	0xb9, 0x49, 0x9c, 0x24, 0x9e, 0x14, 0x65, 0xa4, 0xa9, 0xc0, 0x39, 0xe3, 0x3c, 0xeb, 0xa0, 0x51,
	0xcb, 0xe8, 0x98, 0xca, 0xe9, 0x26, 0xa7, 0x92, 0xd8, 0x21, 0x9d, 0xb1, 0x9e, 0xed, 0x54, 0x1d,
	0xa8, 0xb0, 0x23, 0xf1, 0x7f, 0xab, 0x8e, 0xc4, 0xcf, 0x09, 0x11, 0x26, 0xcf, 0x95, 0x57, 0x46,
	0x7b, 0xb6, 0x8b, 0xde, 0x1a, 0xa1, 0x17, 0xa4, 0x93, 0x73, 0x1f, 0x32, 0x31, 0xe3, 0x7a, 0x0a,
	0xac, 0x97, 0x02, 0x11, 0x7d, 0x41, 0x42, 0xcf, 0x08, 0xc9, 0x95, 0x9e, 0x67, 0xde, 0x02, 0x48,
	0xb6, 0x87, 0x7e, 0x3b, 0x92, 0x51, 0x04, 0xf4, 0x98, 0xb4, 0xac, 0x50, 0x99, 0xcf, 0x4d, 0x60,
	0x14, 0x5f, 0x4b, 0xd3, 0x0a, 0x35, 0xca, 0x4d, 0xa0, 0x57, 0x64, 0x67, 0xe2, 0x14, 0x68, 0x99,
	0x2f, 0x32, 0xcd, 0x0b, 0x60, 0xfb, 0xe8, 0x77, 0x6b, 0xf8, 0x9d, 0x17, 0x40, 0xfb, 0xa4, 0x23,
	0xc1, 0x0b, 0xa7, 0x6c, 0x50, 0x46, 0xb3, 0x03, 0x8c, 0xac, 0x23, 0x3a, 0x20, 0xdd, 0x82, 0xeb,
	0x72, 0xc2, 0x45, 0x28, 0x1d, 0x38, 0xf6, 0x5f, 0x3a, 0x65, 0x9d, 0xc5, 0x2a, 0x61, 0xb4, 0x06,
	0x11, 0x77, 0x64, 0x4a, 0xb2, 0xc3, 0x14, 0x5a, 0xc1, 0x07, 0x49, 0x5f, 0x92, 0xbd, 0xb5, 0x90,
	0x0f, 0x3c, 0x94, 0x9e, 0x1d, 0x61, 0xb0, 0xb7, 0x32, 0x46, 0xc8, 0xe3, 0x1b, 0x00, 0xcd, 0xc7,
	0x39, 0x48, 0xc6, 0x70, 0x36, 0x6a, 0x49, 0x5f, 0x90, 0x9e, 0x9d, 0x2d, 0xbc, 0x12, 0x3c, 0xcf,
	0xb8, 0xe4, 0x36, 0x80, 0x63, 0xc7, 0x18, 0xd9, 0xad, 0xf9, 0xe7, 0x84, 0xe3, 0x78, 0xa5, 0xc7,
	0x76, 0x92, 0xc6, 0x0b, 0x45, 0x3c, 0xda, 0x83, 0xfb, 0xa5, 0x04, 0xb0, 0xff, 0xd3, 0x13, 0xab,
	0x24, 0xbd, 0x24, 0x5d, 0x39, 0x13, 0x36, 0xab, 0x9b, 0x4f, 0x71, 0x5b, 0x27, 0xb2, 0xfb, 0xaa,
	0xfd, 0x15, 0xa1, 0x18, 0xc9, 0x81, 0x7b, 0xc8, 0xe0, 0xd9, 0x2a, 0x07, 0x9e, 0x9d, 0xa5, 0x5f,
	0x11, 0x9d, 0x6f, 0xd1, 0xb8, 0x4f, 0x9c, 0xde, 0x90, 0xfd, 0xb5, 0xb4, 0x19, 0x07, 0xae, 0x34,
	0x48, 0x76, 0x8e, 0xf1, 0xbd, 0x65, 0xfc, 0xb1, 0x32, 0xe2, 0x34, 0x60, 0x3e, 0x5e, 0x08, 0x1c,
	0xbb, 0xc0, 0x1c, 0x89, 0x68, 0x84, 0x24, 0x4e, 0x83, 0xd4, 0x3e, 0x93, 0xa6, 0xe0, 0x4a, 0xb3,
	0x7e, 0xfa, 0x7f, 0x4a, 0xed, 0xef, 0x10, 0xd0, 0x5b, 0x72, 0xbe, 0xb2, 0x33, 0x5f, 0x4e, 0x26,
	0xea, 0x39, 0xf3, 0xc0, 0x9d, 0x98, 0x65, 0xc6, 0x49, 0x70, 0xec, 0x12, 0xb7, 0x9c, 0x2c, 0xb7,
	0x8c, 0x30, 0x33, 0xc2, 0xc8, 0x63, 0x4c, 0xd0, 0x01, 0xd9, 0x89, 0x67, 0xcc, 0x8c, 0x0f, 0x69,
	0x6c, 0x06, 0xd5, 0x4c, 0x68, 0xff, 0xd5, 0xf8, 0x80, 0x53, 0xf3, 0x9e, 0x1c, 0xc5, 0x4c, 0xba,
	0xe6, 0xdf, 0x05, 0x57, 0x98, 0x3e, 0x90, 0xda, 0xa7, 0x2b, 0xaf, 0x1d, 0x7d, 0xfb, 0xf6, 0xc7,
	0x9b, 0xa9, 0x0a, 0xb3, 0x72, 0x7c, 0x23, 0x4c, 0x31, 0x2c, 0x94, 0x98, 0x83, 0xfd, 0xf8, 0x61,
	0x68, 0xfc, 0xcf, 0x12, 0xdc, 0xe2, 0x35, 0x7e, 0x9e, 0xc6, 0xe5, 0x64, 0x68, 0xe7, 0xd3, 0x4f,
	0xd5, 0x47, 0x6a, 0xbc, 0x8d, 0xf4, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x52, 0x33,
	0x05, 0xd1, 0x04, 0x00, 0x00,
}
