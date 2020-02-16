// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_image_layer.proto

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

// Docker image layers information.
type DockerImageLayer struct {
	// Image ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// Layer ID
	LayerId string `protobuf:"bytes,2,opt,name=layer_id,json=layerId,proto3" json:"layer_id"`
	// Layer Order (1 = base layer)
	LayerOrder           int32    `protobuf:"varint,3,opt,name=layer_order,json=layerOrder,proto3" json:"layer_order"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerImageLayer) Reset()         { *m = DockerImageLayer{} }
func (m *DockerImageLayer) String() string { return proto.CompactTextString(m) }
func (*DockerImageLayer) ProtoMessage()    {}
func (*DockerImageLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f38d80234bfb725, []int{0}
}

func (m *DockerImageLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerImageLayer.Unmarshal(m, b)
}
func (m *DockerImageLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerImageLayer.Marshal(b, m, deterministic)
}
func (m *DockerImageLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerImageLayer.Merge(m, src)
}
func (m *DockerImageLayer) XXX_Size() int {
	return xxx_messageInfo_DockerImageLayer.Size(m)
}
func (m *DockerImageLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerImageLayer.DiscardUnknown(m)
}

var xxx_messageInfo_DockerImageLayer proto.InternalMessageInfo

func (m *DockerImageLayer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DockerImageLayer) GetLayerId() string {
	if m != nil {
		return m.LayerId
	}
	return ""
}

func (m *DockerImageLayer) GetLayerOrder() int32 {
	if m != nil {
		return m.LayerOrder
	}
	return 0
}

func init() {
	proto.RegisterType((*DockerImageLayer)(nil), "schemas.DockerImageLayer")
}

func init() { proto.RegisterFile("pb/docker_image_layer.proto", fileDescriptor_1f38d80234bfb725) }

var fileDescriptor_1f38d80234bfb725 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x48, 0xd2, 0x4f,
	0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0x8a, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x8d, 0xcf, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c,
	0x56, 0x8a, 0xe3, 0x12, 0x70, 0x01, 0x2b, 0xf2, 0x04, 0xa9, 0xf1, 0x01, 0x29, 0x11, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe4,
	0xe2, 0x00, 0xeb, 0x8d, 0xcf, 0x4c, 0x91, 0x60, 0x02, 0x8b, 0xb2, 0x83, 0xf9, 0x9e, 0x29, 0x42,
	0xf2, 0x5c, 0xdc, 0x10, 0xa9, 0xfc, 0xa2, 0x94, 0xd4, 0x22, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x2e, 0xb0, 0x90, 0x3f, 0x48, 0xc4, 0xc9, 0x28, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x37, 0x33, 0x39, 0x3b, 0xb5, 0xc0, 0xdc, 0x4c, 0x3f, 0xbf,
	0xb8, 0xb0, 0x34, 0xb5, 0xa8, 0x52, 0x17, 0xec, 0x9a, 0xa4, 0xd2, 0x34, 0xfd, 0x82, 0xec, 0x74,
	0x6b, 0xa8, 0x9b, 0x92, 0xd8, 0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xe3,
	0x8d, 0xe3, 0xc2, 0x00, 0x00, 0x00,
}
