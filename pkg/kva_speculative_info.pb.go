// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/kva_speculative_info.proto

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

// Display kernel virtual address and speculative execution information for the system.
type KvaSpeculativeInfo struct {
	// Kernel Virtual Address shadowing is enabled.
	KvaShadowEnabled int32 `protobuf:"varint,1,opt,name=kva_shadow_enabled,json=kvaShadowEnabled,proto3" json:"kva_shadow_enabled"`
	// User pages are marked as global.
	KvaShadowUserGlobal int32 `protobuf:"varint,2,opt,name=kva_shadow_user_global,json=kvaShadowUserGlobal,proto3" json:"kva_shadow_user_global"`
	// Kernel VA PCID flushing optimization is enabled.
	KvaShadowPcid int32 `protobuf:"varint,3,opt,name=kva_shadow_pcid,json=kvaShadowPcid,proto3" json:"kva_shadow_pcid"`
	// Kernel VA INVPCID is enabled.
	KvaShadowInvPcid int32 `protobuf:"varint,4,opt,name=kva_shadow_inv_pcid,json=kvaShadowInvPcid,proto3" json:"kva_shadow_inv_pcid"`
	// Branch Prediction mitigations are enabled.
	BpMitigations int32 `protobuf:"varint,5,opt,name=bp_mitigations,json=bpMitigations,proto3" json:"bp_mitigations"`
	// Branch Predictions are disabled via system policy.
	BpSystemPolDisabled int32 `protobuf:"varint,6,opt,name=bp_system_pol_disabled,json=bpSystemPolDisabled,proto3" json:"bp_system_pol_disabled"`
	// Branch Predictions are disabled due to lack of microcode update.
	BpMicrocodeDisabled int32 `protobuf:"varint,7,opt,name=bp_microcode_disabled,json=bpMicrocodeDisabled,proto3" json:"bp_microcode_disabled"`
	// SPEC_CTRL MSR supported by CPU Microcode.
	CpuSpecCtrlSupported int32 `protobuf:"varint,8,opt,name=cpu_spec_ctrl_supported,json=cpuSpecCtrlSupported,proto3" json:"cpu_spec_ctrl_supported"`
	// Windows uses IBRS.
	IbrsSupportEnabled int32 `protobuf:"varint,9,opt,name=ibrs_support_enabled,json=ibrsSupportEnabled,proto3" json:"ibrs_support_enabled"`
	// Windows uses STIBP.
	StibpSupportEnabled int32 `protobuf:"varint,10,opt,name=stibp_support_enabled,json=stibpSupportEnabled,proto3" json:"stibp_support_enabled"`
	// PRED_CMD MSR supported by CPU Microcode.
	CpuPredCmdSupported  int32    `protobuf:"varint,11,opt,name=cpu_pred_cmd_supported,json=cpuPredCmdSupported,proto3" json:"cpu_pred_cmd_supported"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvaSpeculativeInfo) Reset()         { *m = KvaSpeculativeInfo{} }
func (m *KvaSpeculativeInfo) String() string { return proto.CompactTextString(m) }
func (*KvaSpeculativeInfo) ProtoMessage()    {}
func (*KvaSpeculativeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_094cb70831044401, []int{0}
}

func (m *KvaSpeculativeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvaSpeculativeInfo.Unmarshal(m, b)
}
func (m *KvaSpeculativeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvaSpeculativeInfo.Marshal(b, m, deterministic)
}
func (m *KvaSpeculativeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvaSpeculativeInfo.Merge(m, src)
}
func (m *KvaSpeculativeInfo) XXX_Size() int {
	return xxx_messageInfo_KvaSpeculativeInfo.Size(m)
}
func (m *KvaSpeculativeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KvaSpeculativeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KvaSpeculativeInfo proto.InternalMessageInfo

func (m *KvaSpeculativeInfo) GetKvaShadowEnabled() int32 {
	if m != nil {
		return m.KvaShadowEnabled
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetKvaShadowUserGlobal() int32 {
	if m != nil {
		return m.KvaShadowUserGlobal
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetKvaShadowPcid() int32 {
	if m != nil {
		return m.KvaShadowPcid
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetKvaShadowInvPcid() int32 {
	if m != nil {
		return m.KvaShadowInvPcid
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetBpMitigations() int32 {
	if m != nil {
		return m.BpMitigations
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetBpSystemPolDisabled() int32 {
	if m != nil {
		return m.BpSystemPolDisabled
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetBpMicrocodeDisabled() int32 {
	if m != nil {
		return m.BpMicrocodeDisabled
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetCpuSpecCtrlSupported() int32 {
	if m != nil {
		return m.CpuSpecCtrlSupported
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetIbrsSupportEnabled() int32 {
	if m != nil {
		return m.IbrsSupportEnabled
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetStibpSupportEnabled() int32 {
	if m != nil {
		return m.StibpSupportEnabled
	}
	return 0
}

func (m *KvaSpeculativeInfo) GetCpuPredCmdSupported() int32 {
	if m != nil {
		return m.CpuPredCmdSupported
	}
	return 0
}

func init() {
	proto.RegisterType((*KvaSpeculativeInfo)(nil), "schemas.KvaSpeculativeInfo")
}

func init() { proto.RegisterFile("pb/kva_speculative_info.proto", fileDescriptor_094cb70831044401) }

var fileDescriptor_094cb70831044401 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4b, 0xab, 0xd3, 0x40,
	0x18, 0x86, 0x39, 0x7a, 0x2e, 0x3a, 0xe2, 0x85, 0xe9, 0xc5, 0x6c, 0x04, 0x11, 0x14, 0x17, 0xb6,
	0x29, 0x2d, 0xea, 0xc2, 0x9d, 0x55, 0xa4, 0x48, 0xa1, 0x18, 0xdc, 0xb8, 0x19, 0xe6, 0xd6, 0x74,
	0xc8, 0x24, 0xf3, 0x39, 0x97, 0x48, 0xff, 0x95, 0x3f, 0xf1, 0x90, 0x49, 0x93, 0x86, 0x6e, 0xbf,
	0xf7, 0x79, 0x86, 0x77, 0xe6, 0x1b, 0xf4, 0x0a, 0x58, 0x5a, 0xd4, 0x94, 0x38, 0x90, 0x3c, 0x68,
	0xea, 0x55, 0x2d, 0x89, 0xaa, 0xf6, 0x66, 0x0e, 0xd6, 0x78, 0x83, 0xef, 0x1c, 0x3f, 0xc8, 0x92,
	0xba, 0x37, 0xff, 0xaf, 0x11, 0xfe, 0x59, 0xd3, 0xec, 0x8c, 0x6d, 0xaa, 0xbd, 0xc1, 0x1f, 0x10,
	0x8e, 0xf6, 0x81, 0x0a, 0xf3, 0x8f, 0xc8, 0x8a, 0x32, 0x2d, 0x45, 0x72, 0xf5, 0xfa, 0xea, 0xfd,
	0xcd, 0xaf, 0x17, 0x45, 0x4d, 0xb3, 0x18, 0x7c, 0x6f, 0xe7, 0x78, 0x85, 0xa6, 0x03, 0x3a, 0x38,
	0x69, 0x49, 0xae, 0x0d, 0xa3, 0x3a, 0x79, 0x10, 0x8d, 0x51, 0x6f, 0xfc, 0x76, 0xd2, 0xfe, 0x88,
	0x11, 0x7e, 0x87, 0x9e, 0x0f, 0x24, 0xe0, 0x4a, 0x24, 0x0f, 0x23, 0xfd, 0xb4, 0xa7, 0x77, 0x5c,
	0x09, 0x3c, 0x43, 0xa3, 0x01, 0xa7, 0xaa, 0xba, 0x65, 0xaf, 0x2f, 0xba, 0x6c, 0xaa, 0x3a, 0xe2,
	0x6f, 0xd1, 0x33, 0x06, 0xa4, 0x54, 0x5e, 0xe5, 0xd4, 0x2b, 0x53, 0xb9, 0xe4, 0xa6, 0x3d, 0x95,
	0xc1, 0xf6, 0x3c, 0x6c, 0x2a, 0x33, 0x20, 0xee, 0xe8, 0xbc, 0x2c, 0x09, 0x18, 0x4d, 0x84, 0x72,
	0xed, 0x25, 0x6f, 0xdb, 0xca, 0x0c, 0xb2, 0x18, 0xee, 0x8c, 0xfe, 0x76, 0x8a, 0xf0, 0x12, 0x4d,
	0xe2, 0xd9, 0xdc, 0x1a, 0x6e, 0x84, 0x3c, 0x3b, 0x77, 0x9d, 0xb3, 0xed, 0xb2, 0xde, 0xf9, 0x88,
	0x5e, 0x72, 0x08, 0x71, 0x0f, 0x84, 0x7b, 0xab, 0x89, 0x0b, 0x00, 0xc6, 0x7a, 0x29, 0x92, 0x47,
	0xd1, 0x1a, 0x73, 0x08, 0xcd, 0xf3, 0xaf, 0xbd, 0xd5, 0x59, 0x97, 0xe1, 0x05, 0x1a, 0x2b, 0x66,
	0x5d, 0x47, 0xf7, 0x2b, 0x78, 0x1c, 0x1d, 0xdc, 0x64, 0x27, 0xb8, 0x5b, 0xc2, 0x12, 0x4d, 0x9c,
	0x57, 0xcd, 0xa5, 0x2e, 0x14, 0xd4, 0x96, 0x8b, 0xe1, 0x85, 0xb3, 0x42, 0xd3, 0xa6, 0x1c, 0x58,
	0x29, 0x08, 0x2f, 0xc5, 0xa0, 0xdb, 0x93, 0x56, 0xe2, 0x10, 0x76, 0x56, 0x8a, 0x75, 0x29, 0xfa,
	0x6a, 0x5f, 0x97, 0x7f, 0x16, 0xb9, 0xf2, 0x87, 0xc0, 0xe6, 0xdc, 0x94, 0x69, 0xa9, 0x78, 0x21,
	0xe1, 0xf3, 0xa7, 0xd4, 0xb8, 0xbf, 0x41, 0xda, 0xe3, 0x2c, 0x7e, 0x30, 0x16, 0xf6, 0x29, 0x14,
	0xf9, 0x97, 0xd3, 0x37, 0x63, 0xb7, 0x71, 0xba, 0xba, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x67, 0x15,
	0x01, 0x70, 0x97, 0x02, 0x00, 0x00,
}
