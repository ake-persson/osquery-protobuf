// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/account_policy_datum.proto

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

// Additional OS X user account data from the AccountPolicy section of OpenDirectory.
type AccountPolicyDatum struct {
	// User ID
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// When the account was first created
	CreationTime float64 `protobuf:"fixed64,2,opt,name=creation_time,json=creationTime,proto3" json:"creation_time"`
	// The number of times the user failed to login with the correct password. Resets after a correct password is entered
	FailedLoginCount int64 `protobuf:"varint,3,opt,name=failed_login_count,json=failedLoginCount,proto3" json:"failed_login_count"`
	// The time of the last failed login attempt. Resets after a correct password is entered
	FailedLoginTimestamp float64 `protobuf:"fixed64,4,opt,name=failed_login_timestamp,json=failedLoginTimestamp,proto3" json:"failed_login_timestamp"`
	// The time the password was last changed
	PasswordLastSetTime  float64  `protobuf:"fixed64,5,opt,name=password_last_set_time,json=passwordLastSetTime,proto3" json:"password_last_set_time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountPolicyDatum) Reset()         { *m = AccountPolicyDatum{} }
func (m *AccountPolicyDatum) String() string { return proto.CompactTextString(m) }
func (*AccountPolicyDatum) ProtoMessage()    {}
func (*AccountPolicyDatum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4b80c24b71ae3f, []int{0}
}

func (m *AccountPolicyDatum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountPolicyDatum.Unmarshal(m, b)
}
func (m *AccountPolicyDatum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountPolicyDatum.Marshal(b, m, deterministic)
}
func (m *AccountPolicyDatum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPolicyDatum.Merge(m, src)
}
func (m *AccountPolicyDatum) XXX_Size() int {
	return xxx_messageInfo_AccountPolicyDatum.Size(m)
}
func (m *AccountPolicyDatum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPolicyDatum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPolicyDatum proto.InternalMessageInfo

func (m *AccountPolicyDatum) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AccountPolicyDatum) GetCreationTime() float64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *AccountPolicyDatum) GetFailedLoginCount() int64 {
	if m != nil {
		return m.FailedLoginCount
	}
	return 0
}

func (m *AccountPolicyDatum) GetFailedLoginTimestamp() float64 {
	if m != nil {
		return m.FailedLoginTimestamp
	}
	return 0
}

func (m *AccountPolicyDatum) GetPasswordLastSetTime() float64 {
	if m != nil {
		return m.PasswordLastSetTime
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountPolicyDatum)(nil), "schemas.AccountPolicyDatum")
}

func init() { proto.RegisterFile("pb/account_policy_datum.proto", fileDescriptor_6b4b80c24b71ae3f) }

var fileDescriptor_6b4b80c24b71ae3f = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x89, 0xe3, 0x0f, 0x04, 0x85, 0x12, 0xa5, 0xcc, 0x46, 0x28, 0xba, 0xe9, 0x42, 0x3b,
	0x62, 0x45, 0x17, 0xae, 0xfc, 0x59, 0x76, 0x21, 0xb5, 0x2b, 0x37, 0x21, 0x93, 0x49, 0xa7, 0xa1,
	0x93, 0xb9, 0x71, 0xee, 0x0d, 0xd2, 0xf7, 0xf5, 0x41, 0x24, 0x19, 0x07, 0x74, 0x17, 0xce, 0xf9,
	0xce, 0x47, 0xb8, 0xfc, 0xdc, 0x97, 0x85, 0xd2, 0x1a, 0x42, 0x4b, 0xd2, 0x43, 0x63, 0xf5, 0x4e,
	0x56, 0x8a, 0x82, 0x9b, 0xf9, 0x0e, 0x08, 0xc4, 0x11, 0xea, 0x8d, 0x71, 0x0a, 0x2f, 0xbe, 0x19,
	0x17, 0x4f, 0x3d, 0xf7, 0x96, 0xb0, 0xd7, 0x48, 0x89, 0x11, 0xcf, 0x82, 0xad, 0x72, 0x36, 0x61,
	0xd3, 0x6c, 0x19, 0x9f, 0xe2, 0x92, 0x9f, 0xe8, 0xce, 0x28, 0xb2, 0xd0, 0x4a, 0xb2, 0xce, 0xe4,
	0x7b, 0x13, 0x36, 0x65, 0xcb, 0xe3, 0x21, 0x5c, 0x59, 0x67, 0xc4, 0x15, 0x17, 0x6b, 0x65, 0x1b,
	0x53, 0xc9, 0x06, 0x6a, 0xdb, 0xca, 0xe4, 0xcd, 0xb3, 0x64, 0x19, 0xf5, 0xcd, 0x22, 0x16, 0x2f,
	0x31, 0x17, 0x77, 0x7c, 0xfc, 0x8f, 0x8e, 0x5a, 0x24, 0xe5, 0x7c, 0xbe, 0x9f, 0xdc, 0x67, 0x7f,
	0x16, 0xab, 0xa1, 0x13, 0x73, 0x3e, 0xf6, 0x0a, 0xf1, 0x0b, 0xba, 0x4a, 0x36, 0x0a, 0x49, 0xa2,
	0xa1, 0xfe, 0x47, 0x07, 0x69, 0x75, 0x3a, 0xb4, 0x0b, 0x85, 0xf4, 0x6e, 0x28, 0x2e, 0x9f, 0x6f,
	0x3f, 0x6e, 0x6a, 0x4b, 0x9b, 0x50, 0xce, 0x34, 0xb8, 0xc2, 0x59, 0xbd, 0x35, 0xfe, 0xe1, 0xbe,
	0x00, 0xfc, 0x0c, 0xa6, 0xdb, 0x5d, 0xa7, 0xa3, 0x94, 0x61, 0x5d, 0xf8, 0x6d, 0xfd, 0xf8, 0x7b,
	0x9a, 0xf2, 0x30, 0xa5, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x80, 0x39, 0x41, 0x4b,
	0x01, 0x00, 0x00,
}
