// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_container_port.proto

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

// Docker container ports.
type DockerContainerPort struct {
	// Container ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// Protocol (tcp
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// Port inside the container
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port"`
	// Host IP address on which public port is listening
	HostIp string `protobuf:"bytes,4,opt,name=host_ip,json=hostIp,proto3" json:"host_ip"`
	// Host port
	HostPort             int32    `protobuf:"varint,5,opt,name=host_port,json=hostPort,proto3" json:"host_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerContainerPort) Reset()         { *m = DockerContainerPort{} }
func (m *DockerContainerPort) String() string { return proto.CompactTextString(m) }
func (*DockerContainerPort) ProtoMessage()    {}
func (*DockerContainerPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_b315164894e1b249, []int{0}
}

func (m *DockerContainerPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerContainerPort.Unmarshal(m, b)
}
func (m *DockerContainerPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerContainerPort.Marshal(b, m, deterministic)
}
func (m *DockerContainerPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerContainerPort.Merge(m, src)
}
func (m *DockerContainerPort) XXX_Size() int {
	return xxx_messageInfo_DockerContainerPort.Size(m)
}
func (m *DockerContainerPort) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerContainerPort.DiscardUnknown(m)
}

var xxx_messageInfo_DockerContainerPort proto.InternalMessageInfo

func (m *DockerContainerPort) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DockerContainerPort) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DockerContainerPort) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DockerContainerPort) GetHostIp() string {
	if m != nil {
		return m.HostIp
	}
	return ""
}

func (m *DockerContainerPort) GetHostPort() int32 {
	if m != nil {
		return m.HostPort
	}
	return 0
}

func init() {
	proto.RegisterType((*DockerContainerPort)(nil), "schemas.DockerContainerPort")
}

func init() { proto.RegisterFile("pb/docker_container_port.proto", fileDescriptor_b315164894e1b249) }

var fileDescriptor_b315164894e1b249 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0x4f, 0x86, 0x30,
	0x10, 0xc5, 0x53, 0xfc, 0x3e, 0x90, 0x0e, 0x0e, 0x75, 0xb0, 0x89, 0x89, 0x21, 0x4e, 0x2c, 0x52,
	0xa3, 0x89, 0x0e, 0x6e, 0xea, 0xe2, 0x66, 0x18, 0x5d, 0x08, 0x2d, 0x15, 0x1a, 0x02, 0x77, 0xb6,
	0x65, 0x60, 0xf6, 0x1f, 0x37, 0xbd, 0xf8, 0x6d, 0xef, 0x7e, 0xf7, 0xde, 0xdd, 0xe3, 0x37, 0xa8,
	0xd5, 0x00, 0x66, 0xb6, 0xbe, 0x33, 0xb0, 0xc6, 0xde, 0xad, 0xd6, 0x77, 0x08, 0x3e, 0x36, 0xe8,
	0x21, 0x82, 0x28, 0x82, 0x99, 0xec, 0xd2, 0x87, 0xdb, 0x5f, 0xc6, 0x2f, 0xdf, 0xc9, 0xf8, 0x76,
	0xf2, 0x7d, 0x82, 0x8f, 0xe2, 0x82, 0x67, 0x6e, 0x90, 0xac, 0x62, 0x75, 0xd9, 0x66, 0x6e, 0x10,
	0x82, 0x1f, 0xe2, 0x8e, 0x56, 0x66, 0x44, 0x48, 0x27, 0x96, 0x4e, 0xca, 0xb3, 0x8a, 0xd5, 0xc7,
	0x96, 0xb4, 0xb8, 0xe2, 0xc5, 0x04, 0x21, 0x76, 0x0e, 0xe5, 0x81, 0xac, 0x79, 0x1a, 0x3f, 0x50,
	0x5c, 0xf3, 0x92, 0x16, 0x94, 0x38, 0x52, 0xe2, 0x3c, 0x81, 0xf4, 0xed, 0xf5, 0xe1, 0xeb, 0x7e,
	0x74, 0x71, 0xda, 0x74, 0x63, 0x60, 0x51, 0x8b, 0x33, 0xb3, 0xc5, 0xe7, 0x27, 0x05, 0xe1, 0x67,
	0xb3, 0x7e, 0xbf, 0xa3, 0xce, 0x7a, 0xfb, 0x56, 0x38, 0x8f, 0x2f, 0xff, 0xcd, 0x75, 0x4e, 0xf4,
	0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x57, 0xec, 0x55, 0xeb, 0x00, 0x00, 0x00,
}