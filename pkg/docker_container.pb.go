// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_container.proto

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

// Docker containers information.
type DockerContainer struct {
	// Container ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// Container name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// Docker image (name) used to launch this container
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image"`
	// Docker image ID
	ImageId string `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id"`
	// Command with arguments
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command"`
	// Time of creation as UNIX time
	Created int64 `protobuf:"varint,6,opt,name=created,proto3" json:"created"`
	// Container state (created
	State string `protobuf:"bytes,7,opt,name=state,proto3" json:"state"`
	// Container status information
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	// Identifier of the initial process
	Pid int64 `protobuf:"varint,9,opt,name=pid,proto3" json:"pid"`
	// Container path
	Path string `protobuf:"bytes,10,opt,name=path,proto3" json:"path"`
	// Container entrypoint(s)
	ConfigEntrypoint string `protobuf:"bytes,11,opt,name=config_entrypoint,json=configEntrypoint,proto3" json:"config_entrypoint"`
	// Container start time as string
	StartedAt string `protobuf:"bytes,12,opt,name=started_at,json=startedAt,proto3" json:"started_at"`
	// Container finish time as string
	FinishedAt string `protobuf:"bytes,13,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at"`
	// Is the container privileged
	Privileged int32 `protobuf:"varint,14,opt,name=privileged,proto3" json:"privileged"`
	// List of container security options
	SecurityOptions string `protobuf:"bytes,15,opt,name=security_options,json=securityOptions,proto3" json:"security_options"`
	// Container environmental variables
	EnvVariables string `protobuf:"bytes,16,opt,name=env_variables,json=envVariables,proto3" json:"env_variables"`
	// Is the root filesystem mounted as read only
	ReadonlyRootfs int32 `protobuf:"varint,17,opt,name=readonly_rootfs,json=readonlyRootfs,proto3" json:"readonly_rootfs"`
	// cgroup namespace
	CgroupNamespace string `protobuf:"bytes,18,opt,name=cgroup_namespace,json=cgroupNamespace,proto3" json:"cgroup_namespace"`
	// IPC namespace
	IpcNamespace string `protobuf:"bytes,19,opt,name=ipc_namespace,json=ipcNamespace,proto3" json:"ipc_namespace"`
	// Mount namespace
	MntNamespace string `protobuf:"bytes,20,opt,name=mnt_namespace,json=mntNamespace,proto3" json:"mnt_namespace"`
	// Network namespace
	NetNamespace string `protobuf:"bytes,21,opt,name=net_namespace,json=netNamespace,proto3" json:"net_namespace"`
	// PID namespace
	PidNamespace string `protobuf:"bytes,22,opt,name=pid_namespace,json=pidNamespace,proto3" json:"pid_namespace"`
	// User namespace
	UserNamespace string `protobuf:"bytes,23,opt,name=user_namespace,json=userNamespace,proto3" json:"user_namespace"`
	// UTS namespace
	UtsNamespace         string   `protobuf:"bytes,24,opt,name=uts_namespace,json=utsNamespace,proto3" json:"uts_namespace"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerContainer) Reset()         { *m = DockerContainer{} }
func (m *DockerContainer) String() string { return proto.CompactTextString(m) }
func (*DockerContainer) ProtoMessage()    {}
func (*DockerContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb2637c84090c86, []int{0}
}

func (m *DockerContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerContainer.Unmarshal(m, b)
}
func (m *DockerContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerContainer.Marshal(b, m, deterministic)
}
func (m *DockerContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerContainer.Merge(m, src)
}
func (m *DockerContainer) XXX_Size() int {
	return xxx_messageInfo_DockerContainer.Size(m)
}
func (m *DockerContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerContainer.DiscardUnknown(m)
}

var xxx_messageInfo_DockerContainer proto.InternalMessageInfo

func (m *DockerContainer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DockerContainer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerContainer) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DockerContainer) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *DockerContainer) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *DockerContainer) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *DockerContainer) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *DockerContainer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DockerContainer) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *DockerContainer) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DockerContainer) GetConfigEntrypoint() string {
	if m != nil {
		return m.ConfigEntrypoint
	}
	return ""
}

func (m *DockerContainer) GetStartedAt() string {
	if m != nil {
		return m.StartedAt
	}
	return ""
}

func (m *DockerContainer) GetFinishedAt() string {
	if m != nil {
		return m.FinishedAt
	}
	return ""
}

func (m *DockerContainer) GetPrivileged() int32 {
	if m != nil {
		return m.Privileged
	}
	return 0
}

func (m *DockerContainer) GetSecurityOptions() string {
	if m != nil {
		return m.SecurityOptions
	}
	return ""
}

func (m *DockerContainer) GetEnvVariables() string {
	if m != nil {
		return m.EnvVariables
	}
	return ""
}

func (m *DockerContainer) GetReadonlyRootfs() int32 {
	if m != nil {
		return m.ReadonlyRootfs
	}
	return 0
}

func (m *DockerContainer) GetCgroupNamespace() string {
	if m != nil {
		return m.CgroupNamespace
	}
	return ""
}

func (m *DockerContainer) GetIpcNamespace() string {
	if m != nil {
		return m.IpcNamespace
	}
	return ""
}

func (m *DockerContainer) GetMntNamespace() string {
	if m != nil {
		return m.MntNamespace
	}
	return ""
}

func (m *DockerContainer) GetNetNamespace() string {
	if m != nil {
		return m.NetNamespace
	}
	return ""
}

func (m *DockerContainer) GetPidNamespace() string {
	if m != nil {
		return m.PidNamespace
	}
	return ""
}

func (m *DockerContainer) GetUserNamespace() string {
	if m != nil {
		return m.UserNamespace
	}
	return ""
}

func (m *DockerContainer) GetUtsNamespace() string {
	if m != nil {
		return m.UtsNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerContainer)(nil), "schemas.DockerContainer")
}

func init() { proto.RegisterFile("pb/docker_container.proto", fileDescriptor_2bb2637c84090c86) }

var fileDescriptor_2bb2637c84090c86 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x86, 0x95, 0xa6, 0x49, 0x1a, 0xb7, 0xf9, 0xa8, 0x29, 0xc5, 0x3d, 0x00, 0x11, 0x08, 0x11,
	0x84, 0x68, 0x10, 0x48, 0x70, 0xe0, 0x54, 0x3e, 0x0e, 0x5c, 0x40, 0xca, 0x81, 0x03, 0x97, 0x95,
	0xd7, 0x9e, 0x6c, 0xac, 0x64, 0x6d, 0x63, 0xcf, 0x46, 0xca, 0xaf, 0x07, 0xed, 0x38, 0x5b, 0x36,
	0xb7, 0x99, 0xe7, 0x7d, 0x34, 0xef, 0x4a, 0x2b, 0xb3, 0x1b, 0x9f, 0x2f, 0xb4, 0x53, 0x1b, 0x08,
	0x99, 0x72, 0x16, 0xa5, 0xb1, 0x10, 0x6e, 0x7d, 0x70, 0xe8, 0xf8, 0x20, 0xaa, 0x35, 0x94, 0x32,
	0x3e, 0xfb, 0xdb, 0x63, 0x93, 0xaf, 0xe4, 0x7c, 0x69, 0x14, 0x3e, 0x66, 0x27, 0x46, 0x8b, 0xce,
	0xac, 0x33, 0x1f, 0x2e, 0x4f, 0x8c, 0xe6, 0x9c, 0x9d, 0x5a, 0x59, 0x82, 0x38, 0x21, 0x42, 0x33,
	0xbf, 0x62, 0x3d, 0x53, 0xca, 0x02, 0x44, 0x97, 0x60, 0x5a, 0xf8, 0x0d, 0x3b, 0xa3, 0x21, 0x33,
	0x5a, 0x9c, 0x52, 0x30, 0xa0, 0xfd, 0xbb, 0xe6, 0x82, 0x0d, 0x94, 0x2b, 0x4b, 0x69, 0xb5, 0xe8,
	0xa5, 0xe4, 0xb0, 0x52, 0x12, 0x40, 0x22, 0x68, 0xd1, 0x9f, 0x75, 0xe6, 0xdd, 0x65, 0xb3, 0xd6,
	0x25, 0x11, 0x25, 0x82, 0x18, 0xa4, 0x12, 0x5a, 0xf8, 0x35, 0xeb, 0xd7, 0x43, 0x15, 0xc5, 0x19,
	0xe1, 0xc3, 0xc6, 0xa7, 0xac, 0xeb, 0x8d, 0x16, 0x43, 0xba, 0x51, 0x8f, 0xf5, 0x87, 0x7b, 0x89,
	0x6b, 0xc1, 0xd2, 0x87, 0xd7, 0x33, 0x7f, 0xcd, 0x2e, 0x95, 0xb3, 0x2b, 0x53, 0x64, 0x60, 0x31,
	0xec, 0xbd, 0x33, 0x16, 0xc5, 0x39, 0x09, 0xd3, 0x14, 0x7c, 0xbb, 0xe7, 0xfc, 0x31, 0x63, 0x11,
	0x65, 0x40, 0xd0, 0x99, 0x44, 0x71, 0x41, 0xd6, 0xf0, 0x40, 0xee, 0x90, 0x3f, 0x65, 0xe7, 0x2b,
	0x63, 0x4d, 0x5c, 0xa7, 0x7c, 0x44, 0x39, 0x6b, 0xd0, 0x1d, 0xf2, 0x27, 0x8c, 0xf9, 0x60, 0x76,
	0x66, 0x0b, 0x05, 0x68, 0x31, 0x9e, 0x75, 0xe6, 0xbd, 0x65, 0x8b, 0xf0, 0x57, 0x6c, 0x1a, 0x41,
	0x55, 0xc1, 0xe0, 0x3e, 0x73, 0x1e, 0x8d, 0xb3, 0x51, 0x4c, 0xe8, 0xca, 0xa4, 0xe1, 0x3f, 0x13,
	0xe6, 0xcf, 0xd9, 0x08, 0xec, 0x2e, 0xdb, 0xc9, 0x60, 0x64, 0xbe, 0x85, 0x28, 0xa6, 0xe4, 0x5d,
	0x80, 0xdd, 0xfd, 0x6a, 0x18, 0x7f, 0xc9, 0x26, 0x01, 0xa4, 0x76, 0x76, 0xbb, 0xcf, 0x82, 0x73,
	0xb8, 0x8a, 0xe2, 0x92, 0x4a, 0xc7, 0x0d, 0x5e, 0x12, 0xad, 0x8b, 0x55, 0x11, 0x5c, 0xe5, 0xb3,
	0xfa, 0x6f, 0x46, 0x2f, 0x15, 0x08, 0x9e, 0x8a, 0x13, 0xff, 0xd1, 0xe0, 0xba, 0xd8, 0x78, 0xd5,
	0xf2, 0x1e, 0xa4, 0x62, 0xe3, 0xd5, 0x91, 0x54, 0x5a, 0x6c, 0x49, 0x57, 0x49, 0x2a, 0x2d, 0x1e,
	0x49, 0x16, 0xda, 0xd2, 0xc3, 0x24, 0x59, 0x38, 0x96, 0xbc, 0xd1, 0x2d, 0xe9, 0x3a, 0x49, 0xde,
	0xe8, 0xff, 0xd2, 0x0b, 0x36, 0xae, 0x22, 0x84, 0x96, 0xf5, 0x88, 0xac, 0x51, 0x4d, 0x8f, 0x6e,
	0x55, 0x18, 0x5b, 0x96, 0x48, 0xb7, 0x2a, 0x8c, 0xf7, 0xd2, 0xe7, 0x77, 0xbf, 0xdf, 0x16, 0x06,
	0xd7, 0x55, 0x7e, 0xab, 0x5c, 0xb9, 0x28, 0x8d, 0xda, 0x80, 0xff, 0xf8, 0x61, 0xe1, 0xe2, 0x9f,
	0x0a, 0xc2, 0xfe, 0x0d, 0xbd, 0x97, 0xbc, 0x5a, 0x2d, 0xfc, 0xa6, 0xf8, 0x74, 0x78, 0x35, 0x79,
	0x9f, 0xe8, 0xfb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xdc, 0xe4, 0x21, 0x62, 0x03, 0x00,
	0x00,
}
