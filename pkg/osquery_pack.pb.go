// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/osquery_pack.proto

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

// Information about the current query packs that are loaded in osquery.
type OsqueryPack struct {
	// The given name for this query pack
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Platforms this query is supported on
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform"`
	// Minimum osquery version that this query will run on
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version"`
	// Shard restriction limit
	Shard int32 `protobuf:"varint,4,opt,name=shard,proto3" json:"shard"`
	// The number of times that the discovery query used cached values since the last time the config was reloaded
	DiscoveryCacheHits int32 `protobuf:"varint,5,opt,name=discovery_cache_hits,json=discoveryCacheHits,proto3" json:"discovery_cache_hits"`
	// The number of times that the discovery queries have been executed since the last time the config was reloaded
	DiscoveryExecutions int32 `protobuf:"varint,6,opt,name=discovery_executions,json=discoveryExecutions,proto3" json:"discovery_executions"`
	// Whether this pack is active (the version
	Active               int32    `protobuf:"varint,7,opt,name=active,proto3" json:"active"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OsqueryPack) Reset()         { *m = OsqueryPack{} }
func (m *OsqueryPack) String() string { return proto.CompactTextString(m) }
func (*OsqueryPack) ProtoMessage()    {}
func (*OsqueryPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3c7dab43a6aee9, []int{0}
}

func (m *OsqueryPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OsqueryPack.Unmarshal(m, b)
}
func (m *OsqueryPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OsqueryPack.Marshal(b, m, deterministic)
}
func (m *OsqueryPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsqueryPack.Merge(m, src)
}
func (m *OsqueryPack) XXX_Size() int {
	return xxx_messageInfo_OsqueryPack.Size(m)
}
func (m *OsqueryPack) XXX_DiscardUnknown() {
	xxx_messageInfo_OsqueryPack.DiscardUnknown(m)
}

var xxx_messageInfo_OsqueryPack proto.InternalMessageInfo

func (m *OsqueryPack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OsqueryPack) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *OsqueryPack) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OsqueryPack) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *OsqueryPack) GetDiscoveryCacheHits() int32 {
	if m != nil {
		return m.DiscoveryCacheHits
	}
	return 0
}

func (m *OsqueryPack) GetDiscoveryExecutions() int32 {
	if m != nil {
		return m.DiscoveryExecutions
	}
	return 0
}

func (m *OsqueryPack) GetActive() int32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func init() {
	proto.RegisterType((*OsqueryPack)(nil), "schemas.OsqueryPack")
}

func init() { proto.RegisterFile("pb/osquery_pack.proto", fileDescriptor_2f3c7dab43a6aee9) }

var fileDescriptor_2f3c7dab43a6aee9 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0x95, 0xef, 0x6b, 0x12, 0x30, 0xdb, 0x51, 0x90, 0xc5, 0x54, 0x31, 0x75, 0xa1, 0x29,
	0x20, 0xc1, 0xc0, 0x06, 0x42, 0x62, 0x03, 0x75, 0x64, 0xa9, 0x9c, 0xeb, 0xb5, 0xb6, 0x82, 0x63,
	0xe3, 0x73, 0x22, 0xf8, 0xcb, 0xfc, 0x0a, 0x54, 0xb7, 0x8d, 0xd8, 0xfc, 0xdc, 0xf3, 0xbe, 0x92,
	0xf5, 0x8a, 0x33, 0x5f, 0x57, 0x8e, 0x3f, 0x3b, 0x0a, 0xdf, 0x4b, 0xaf, 0xb0, 0x99, 0xf9, 0xe0,
	0xa2, 0x83, 0x92, 0x51, 0x93, 0x55, 0x7c, 0xf9, 0x93, 0x89, 0x93, 0xd7, 0x9d, 0x7f, 0x53, 0xd8,
	0x00, 0x88, 0x51, 0xab, 0x2c, 0xc9, 0x6c, 0x92, 0x4d, 0x8f, 0x17, 0xe9, 0x0d, 0x17, 0xe2, 0xc8,
	0x7f, 0xa8, 0xb8, 0x76, 0xc1, 0xca, 0x7f, 0xe9, 0x3e, 0x30, 0x48, 0x51, 0xf6, 0x14, 0xd8, 0xb8,
	0x56, 0xfe, 0x4f, 0xea, 0x80, 0x30, 0x16, 0x39, 0x6b, 0x15, 0x56, 0x72, 0x34, 0xc9, 0xa6, 0xf9,
	0x62, 0x07, 0x30, 0x17, 0xe3, 0x95, 0x61, 0x74, 0xfd, 0xf6, 0x43, 0xa8, 0x50, 0xd3, 0x52, 0x9b,
	0xc8, 0x32, 0x4f, 0x21, 0x18, 0xdc, 0xd3, 0x56, 0xbd, 0x98, 0xc8, 0x70, 0xfd, 0xb7, 0x41, 0x5f,
	0x84, 0x5d, 0x34, 0xae, 0x65, 0x59, 0xa4, 0xc6, 0xe9, 0xe0, 0x9e, 0x07, 0x05, 0xe7, 0xa2, 0x50,
	0x18, 0x4d, 0x4f, 0xb2, 0x4c, 0xa1, 0x3d, 0x3d, 0xde, 0xbc, 0xcf, 0x37, 0x26, 0xea, 0xae, 0x9e,
	0xa1, 0xb3, 0x95, 0x35, 0xd8, 0x90, 0xbf, 0xbf, 0x3b, 0xec, 0x73, 0x95, 0xa6, 0xa9, 0xbb, 0x75,
	0xe5, 0x9b, 0xcd, 0xc3, 0x7e, 0xa0, 0xba, 0x48, 0xd7, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0xd9, 0xc7, 0x6a, 0x49, 0x01, 0x00, 0x00,
}
