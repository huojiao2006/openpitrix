// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/metadata.proto

package pbtypes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Metadata struct {
	// Instance host                     = _; // softlink, ignore
	Cluster              *Cluster             `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster"`
	Hosts                map[string]*Instance `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AddingHosts          map[string]*Instance `protobuf:"bytes,3,rep,name=adding_hosts,json=addingHosts,proto3" json:"adding_hosts" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DeletingHosts        map[string]*Instance `protobuf:"bytes,4,rep,name=deleting_hosts,json=deletingHosts,proto3" json:"deleting_hosts" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env                  map[string]string    `protobuf:"bytes,5,rep,name=env,proto3" json:"env" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Links                map[string]string    `protobuf:"bytes,6,rep,name=links,proto3" json:"links" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VerticalScalingRoles string               `protobuf:"bytes,7,opt,name=vertical_scaling_roles,json=verticalScalingRoles,proto3" json:"vertical_scaling_roles"`
	Cmd                  *CommandInfo         `protobuf:"bytes,8,opt,name=cmd,proto3" json:"cmd"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *Metadata) GetHosts() map[string]*Instance {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Metadata) GetAddingHosts() map[string]*Instance {
	if m != nil {
		return m.AddingHosts
	}
	return nil
}

func (m *Metadata) GetDeletingHosts() map[string]*Instance {
	if m != nil {
		return m.DeletingHosts
	}
	return nil
}

func (m *Metadata) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Metadata) GetLinks() map[string]string {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Metadata) GetVerticalScalingRoles() string {
	if m != nil {
		return m.VerticalScalingRoles
	}
	return ""
}

func (m *Metadata) GetCmd() *CommandInfo {
	if m != nil {
		return m.Cmd
	}
	return nil
}

// key: /self/cluster
type Cluster struct {
	AppId                string                `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ClusterId            string                `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id"`
	UserId               string                `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id"`
	GlobalUuid           string                `protobuf:"bytes,4,opt,name=global_uuid,json=globalUuid,proto3" json:"global_uuid"`
	Subnet               string                `protobuf:"bytes,5,opt,name=subnet,proto3" json:"subnet"`
	Zone                 string                `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone"`
	Endpoints            map[string]*Endpoint  `protobuf:"bytes,7,rep,name=endpoints,proto3" json:"endpoints" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReservedIps          map[string]*IPAddress `protobuf:"bytes,8,rep,name=reserved_ips,json=reservedIps,proto3" json:"reserved_ips" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ApiServer            *Endpoint             `protobuf:"bytes,9,opt,name=api_server,json=apiServer,proto3" json:"api_server"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Cluster) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Cluster) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Cluster) GetGlobalUuid() string {
	if m != nil {
		return m.GlobalUuid
	}
	return ""
}

func (m *Cluster) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Cluster) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Cluster) GetEndpoints() map[string]*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Cluster) GetReservedIps() map[string]*IPAddress {
	if m != nil {
		return m.ReservedIps
	}
	return nil
}

func (m *Cluster) GetApiServer() *Endpoint {
	if m != nil {
		return m.ApiServer
	}
	return nil
}

// key: /self/host
// key: /self/hosts/[role name]/[instance_id]
// key: /self/adding-hosts/[role name]/[instance_id]
// key: /self/deleting-hosts/[role name]/[instance_id]
// key: /[instance_id]
type Instance struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Ip                   string                `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip"`
	Port                 int32                 `protobuf:"varint,3,opt,name=port,proto3" json:"port"`
	Eip                  string                `protobuf:"bytes,4,opt,name=eip,proto3" json:"eip"`
	Mac                  string                `protobuf:"bytes,5,opt,name=mac,proto3" json:"mac"`
	Sid                  string                `protobuf:"bytes,6,opt,name=sid,proto3" json:"sid"`
	Gid                  string                `protobuf:"bytes,7,opt,name=gid,proto3" json:"gid"`
	Gsid                 string                `protobuf:"bytes,8,opt,name=gsid,proto3" json:"gsid"`
	NodeId               string                `protobuf:"bytes,9,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	InstanceId           string                `protobuf:"bytes,10,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
	Cpu                  int32                 `protobuf:"varint,11,opt,name=cpu,proto3" json:"cpu"`
	Gpu                  int32                 `protobuf:"varint,12,opt,name=gpu,proto3" json:"gpu"`
	Memory               int32                 `protobuf:"varint,13,opt,name=memory,proto3" json:"memory"`
	VolumeSize           int32                 `protobuf:"varint,14,opt,name=volume_size,json=volumeSize,proto3" json:"volume_size"`
	InstanceClass        string                `protobuf:"bytes,15,opt,name=instance_class,json=instanceClass,proto3" json:"instance_class"`
	GpuClass             string                `protobuf:"bytes,16,opt,name=gpu_class,json=gpuClass,proto3" json:"gpu_class"`
	VolumeClass          string                `protobuf:"bytes,17,opt,name=volume_class,json=volumeClass,proto3" json:"volume_class"`
	PhysicalMachine      string                `protobuf:"bytes,18,opt,name=physical_machine,json=physicalMachine,proto3" json:"physical_machine"`
	Role                 string                `protobuf:"bytes,19,opt,name=role,proto3" json:"role"`
	PubKey               string                `protobuf:"bytes,20,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	Token                string                `protobuf:"bytes,21,opt,name=token,proto3" json:"token"`
	ReservedIps          map[string]*IPAddress `protobuf:"bytes,22,rep,name=reserved_ips,json=reservedIps,proto3" json:"reserved_ips" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{2}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Instance) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Instance) GetEip() string {
	if m != nil {
		return m.Eip
	}
	return ""
}

func (m *Instance) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Instance) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *Instance) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Instance) GetGsid() string {
	if m != nil {
		return m.Gsid
	}
	return ""
}

func (m *Instance) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Instance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Instance) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *Instance) GetGpu() int32 {
	if m != nil {
		return m.Gpu
	}
	return 0
}

func (m *Instance) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Instance) GetVolumeSize() int32 {
	if m != nil {
		return m.VolumeSize
	}
	return 0
}

func (m *Instance) GetInstanceClass() string {
	if m != nil {
		return m.InstanceClass
	}
	return ""
}

func (m *Instance) GetGpuClass() string {
	if m != nil {
		return m.GpuClass
	}
	return ""
}

func (m *Instance) GetVolumeClass() string {
	if m != nil {
		return m.VolumeClass
	}
	return ""
}

func (m *Instance) GetPhysicalMachine() string {
	if m != nil {
		return m.PhysicalMachine
	}
	return ""
}

func (m *Instance) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Instance) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Instance) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Instance) GetReservedIps() map[string]*IPAddress {
	if m != nil {
		return m.ReservedIps
	}
	return nil
}

type CommandInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Cmd                  string   `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd"`
	Timeout              int32    `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandInfo) Reset()         { *m = CommandInfo{} }
func (m *CommandInfo) String() string { return proto.CompactTextString(m) }
func (*CommandInfo) ProtoMessage()    {}
func (*CommandInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{3}
}

func (m *CommandInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandInfo.Unmarshal(m, b)
}
func (m *CommandInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandInfo.Marshal(b, m, deterministic)
}
func (m *CommandInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandInfo.Merge(m, src)
}
func (m *CommandInfo) XXX_Size() int {
	return xxx_messageInfo_CommandInfo.Size(m)
}
func (m *CommandInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommandInfo proto.InternalMessageInfo

func (m *CommandInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CommandInfo) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *CommandInfo) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type Endpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{4}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type IPAddress struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_08fa6b2d4fb98c3c, []int{5}
}

func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (m *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(m, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

func (m *IPAddress) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "metadata.types.Metadata")
	proto.RegisterMapType((map[string]*Instance)(nil), "metadata.types.Metadata.AddingHostsEntry")
	proto.RegisterMapType((map[string]*Instance)(nil), "metadata.types.Metadata.DeletingHostsEntry")
	proto.RegisterMapType((map[string]string)(nil), "metadata.types.Metadata.EnvEntry")
	proto.RegisterMapType((map[string]*Instance)(nil), "metadata.types.Metadata.HostsEntry")
	proto.RegisterMapType((map[string]string)(nil), "metadata.types.Metadata.LinksEntry")
	proto.RegisterType((*Cluster)(nil), "metadata.types.Cluster")
	proto.RegisterMapType((map[string]*Endpoint)(nil), "metadata.types.Cluster.EndpointsEntry")
	proto.RegisterMapType((map[string]*IPAddress)(nil), "metadata.types.Cluster.ReservedIpsEntry")
	proto.RegisterType((*Instance)(nil), "metadata.types.Instance")
	proto.RegisterMapType((map[string]*IPAddress)(nil), "metadata.types.Instance.ReservedIpsEntry")
	proto.RegisterType((*CommandInfo)(nil), "metadata.types.CommandInfo")
	proto.RegisterType((*Endpoint)(nil), "metadata.types.Endpoint")
	proto.RegisterType((*IPAddress)(nil), "metadata.types.IPAddress")
}

func init() { proto.RegisterFile("metadata/types/metadata.proto", fileDescriptor_08fa6b2d4fb98c3c) }

var fileDescriptor_08fa6b2d4fb98c3c = []byte{
	// 925 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x6f, 0x23, 0x35,
	0x10, 0x57, 0x93, 0xe6, 0xcf, 0x4e, 0xda, 0x5c, 0x30, 0xbd, 0x9e, 0xe9, 0xe9, 0x44, 0x5b, 0x04,
	0xea, 0x09, 0x91, 0x88, 0x3b, 0x74, 0x14, 0x78, 0x3a, 0xee, 0x2a, 0x11, 0xdd, 0x1d, 0x42, 0x5b,
	0x81, 0xe0, 0x5e, 0x22, 0x67, 0x6d, 0x52, 0xab, 0xbb, 0xb6, 0xb5, 0xde, 0x8d, 0x48, 0x3f, 0x0f,
	0xef, 0x7c, 0x30, 0xbe, 0x04, 0x1a, 0xdb, 0x9b, 0x6c, 0xda, 0x04, 0x04, 0x3a, 0x9e, 0x32, 0xf3,
	0x9b, 0x99, 0x9f, 0x1d, 0xcf, 0xcf, 0xe3, 0x85, 0x47, 0x99, 0x28, 0x18, 0x67, 0x05, 0x1b, 0x15,
	0x0b, 0x23, 0xec, 0xa8, 0x72, 0x87, 0x26, 0xd7, 0x85, 0x26, 0xfd, 0xa5, 0xef, 0xc2, 0xa7, 0x7f,
	0xb6, 0xa1, 0xfb, 0x26, 0x40, 0xe4, 0x73, 0xe8, 0x24, 0x69, 0x69, 0x0b, 0x91, 0xd3, 0x9d, 0xe3,
	0x9d, 0xb3, 0xde, 0x93, 0x07, 0xc3, 0xf5, 0xf4, 0xe1, 0x0b, 0x1f, 0x8e, 0xab, 0x3c, 0xf2, 0x15,
	0xb4, 0xae, 0xb4, 0x2d, 0x2c, 0x6d, 0x1c, 0x37, 0xcf, 0x7a, 0x4f, 0x3e, 0xba, 0x5d, 0x50, 0x71,
	0x0f, 0xbf, 0xc3, 0xac, 0x0b, 0x55, 0xe4, 0x8b, 0xd8, 0x57, 0x90, 0xd7, 0xb0, 0xc7, 0x38, 0x97,
	0x6a, 0x36, 0xf1, 0x0c, 0x4d, 0xc7, 0xf0, 0x78, 0x2b, 0xc3, 0x73, 0x97, 0x5c, 0xe3, 0xe9, 0xb1,
	0x15, 0x42, 0x62, 0xe8, 0x73, 0x91, 0x8a, 0x62, 0xc5, 0xb7, 0xeb, 0xf8, 0x3e, 0xdd, 0xca, 0xf7,
	0x32, 0xa4, 0xd7, 0x18, 0xf7, 0x79, 0x1d, 0x23, 0x4f, 0xa1, 0x29, 0xd4, 0x9c, 0xb6, 0x1c, 0xd1,
	0xc9, 0x56, 0xa2, 0x0b, 0x35, 0xf7, 0xe5, 0x98, 0x8d, 0x27, 0x92, 0x4a, 0x75, 0x6d, 0x69, 0xfb,
	0x1f, 0x4e, 0xe4, 0x35, 0x66, 0x85, 0x13, 0x71, 0x15, 0xe4, 0x0b, 0x38, 0x9c, 0x8b, 0xbc, 0x90,
	0x09, 0x4b, 0x27, 0x36, 0x61, 0x29, 0xfe, 0x97, 0x5c, 0xa7, 0xc2, 0xd2, 0xce, 0xf1, 0xce, 0x59,
	0x14, 0x1f, 0x54, 0xd1, 0x4b, 0x1f, 0x8c, 0x31, 0x46, 0x3e, 0x83, 0x66, 0x92, 0x71, 0xda, 0x75,
	0x1d, 0x7b, 0x78, 0xa7, 0x63, 0x3a, 0xcb, 0x98, 0xe2, 0x63, 0xf5, 0xab, 0x8e, 0x31, 0xef, 0x28,
	0x06, 0x58, 0xfd, 0x63, 0x32, 0x80, 0xe6, 0xb5, 0x58, 0xb8, 0x76, 0x47, 0x31, 0x9a, 0x64, 0x08,
	0xad, 0x39, 0x4b, 0x4b, 0x41, 0x1b, 0x8e, 0x90, 0xde, 0x26, 0x1c, 0x2b, 0x5b, 0x30, 0x95, 0x88,
	0xd8, 0xa7, 0x7d, 0xdd, 0x38, 0xdf, 0x39, 0xfa, 0x19, 0x06, 0xb7, 0xbb, 0xf3, 0x8e, 0x98, 0xdf,
	0x02, 0xb9, 0xdb, 0xa7, 0x77, 0xc4, 0xfd, 0x0c, 0xba, 0x55, 0xeb, 0x36, 0x30, 0x1e, 0xd4, 0x19,
	0xa3, 0x7a, 0xdd, 0x39, 0xc0, 0xaa, 0x77, 0xff, 0xa6, 0xf2, 0xf4, 0xf7, 0x5d, 0xe8, 0x84, 0x2b,
	0x44, 0xee, 0x43, 0x9b, 0x19, 0x33, 0x91, 0x3c, 0x94, 0xb6, 0x98, 0x31, 0x63, 0x4e, 0x1e, 0x01,
	0x84, 0xbb, 0x85, 0x21, 0xcf, 0x10, 0x05, 0x64, 0xcc, 0xc9, 0x03, 0xe8, 0x94, 0xd6, 0xc7, 0x9a,
	0x2e, 0xd6, 0x46, 0x77, 0xcc, 0xc9, 0x87, 0xd0, 0x9b, 0xa5, 0x7a, 0xca, 0xd2, 0x49, 0x59, 0x4a,
	0x4e, 0x77, 0x5d, 0x10, 0x3c, 0xf4, 0x63, 0x29, 0x39, 0x39, 0x84, 0xb6, 0x2d, 0xa7, 0x4a, 0x14,
	0xb4, 0xe5, 0x0b, 0xbd, 0x47, 0x08, 0xec, 0xde, 0x68, 0x25, 0x68, 0xdb, 0xa1, 0xce, 0x26, 0x2f,
	0x21, 0x12, 0x8a, 0x1b, 0x2d, 0x55, 0x81, 0xda, 0x43, 0x1d, 0x7f, 0xb2, 0x65, 0x14, 0x0c, 0x2f,
	0xaa, 0x44, 0x2f, 0xe5, 0x55, 0x21, 0x79, 0x05, 0x7b, 0xb9, 0xb0, 0x22, 0x9f, 0x0b, 0x3e, 0x91,
	0xc6, 0xd2, 0xae, 0x23, 0x3a, 0xdb, 0x46, 0x14, 0x87, 0xdc, 0xb1, 0xa9, 0xee, 0x77, 0xbe, 0x42,
	0xc8, 0x97, 0x00, 0xcc, 0xc8, 0x89, 0x03, 0x72, 0x1a, 0x6d, 0xee, 0x72, 0xb5, 0x97, 0x38, 0x62,
	0x46, 0x5e, 0xba, 0xd4, 0xa3, 0x9f, 0xa0, 0xbf, 0xbe, 0xc5, 0xff, 0xa0, 0x9e, 0x25, 0x6f, 0x4d,
	0x05, 0xbf, 0xc0, 0xe0, 0xf6, 0x8e, 0x37, 0x30, 0x8f, 0xd6, 0x99, 0x3f, 0xb8, 0xa3, 0xcb, 0x1f,
	0x9e, 0x73, 0x9e, 0x0b, 0x6b, 0xeb, 0x32, 0xf9, 0xa3, 0x05, 0xdd, 0x4a, 0xb0, 0xa4, 0x0f, 0x8d,
	0xa5, 0x46, 0x1a, 0x92, 0x3b, 0xdf, 0x04, 0x61, 0x34, 0xa4, 0xc1, 0xfe, 0x19, 0x9d, 0x17, 0x4e,
	0x0e, 0xad, 0xd8, 0xd9, 0xb8, 0x0f, 0x21, 0x4d, 0x10, 0x01, 0x9a, 0x88, 0x64, 0x2c, 0x09, 0xad,
	0x47, 0x13, 0x11, 0x2b, 0x79, 0x68, 0x3b, 0x9a, 0x88, 0xcc, 0x24, 0x0f, 0xb3, 0x06, 0x4d, 0xe4,
	0x9e, 0x61, 0x52, 0xd7, 0x6b, 0x03, 0x6d, 0x54, 0xa0, 0xd2, 0x5c, 0xa0, 0x02, 0x23, 0x2f, 0x24,
	0x74, 0xbd, 0x02, 0x65, 0xd8, 0x34, 0x06, 0xc1, 0x2b, 0xb0, 0x82, 0xc6, 0x8e, 0x3f, 0x31, 0x25,
	0xed, 0xb9, 0x8d, 0xa2, 0xe9, 0x56, 0x34, 0x25, 0xdd, 0xf3, 0xc8, 0xcc, 0x94, 0xa8, 0xd2, 0x4c,
	0x64, 0x3a, 0x5f, 0xd0, 0x7d, 0x07, 0x06, 0x0f, 0xc9, 0xe7, 0x3a, 0x2d, 0x33, 0x31, 0xb1, 0xf2,
	0x46, 0xd0, 0xbe, 0x0b, 0x82, 0x87, 0x2e, 0xe5, 0x8d, 0x20, 0x1f, 0x43, 0x7f, 0xb9, 0x7a, 0x92,
	0x32, 0x6b, 0xe9, 0x3d, 0xb7, 0x81, 0xfd, 0x0a, 0x7d, 0x81, 0x20, 0x79, 0x08, 0xd1, 0xcc, 0x94,
	0x21, 0x63, 0xe0, 0x32, 0xba, 0x33, 0x53, 0xfa, 0xe0, 0x09, 0xec, 0x85, 0x45, 0x7c, 0xfc, 0x3d,
	0x17, 0x0f, 0x0b, 0xfb, 0x94, 0xc7, 0x30, 0x30, 0x57, 0x0b, 0xeb, 0x46, 0x74, 0xc6, 0x92, 0x2b,
	0xa9, 0x04, 0x25, 0x2e, 0xed, 0x5e, 0x85, 0xbf, 0xf1, 0x30, 0x1e, 0x1e, 0x0e, 0x6f, 0xfa, 0xbe,
	0x3f, 0x3c, 0xb4, 0xf1, 0xf0, 0x4c, 0x39, 0x9d, 0xa0, 0x48, 0x0e, 0xfc, 0xe1, 0x99, 0x72, 0xfa,
	0xca, 0xcf, 0x8c, 0x42, 0x5f, 0x0b, 0x45, 0xef, 0xfb, 0x61, 0xe0, 0x1c, 0x7c, 0x22, 0xd7, 0x6e,
	0xd0, 0xe1, 0xe6, 0x27, 0xb2, 0xd2, 0xca, 0xdf, 0x5f, 0xa1, 0xff, 0x53, 0xb1, 0x63, 0xe8, 0xd5,
	0x1e, 0x9a, 0x3b, 0x9a, 0x1d, 0xf8, 0x27, 0xca, 0x8b, 0x16, 0x4d, 0x42, 0xa1, 0x53, 0xc8, 0x4c,
	0xe8, 0xb2, 0x12, 0x6e, 0xe5, 0x9e, 0x7e, 0x8f, 0x53, 0xd9, 0x5f, 0x37, 0x3c, 0x42, 0x7c, 0xcb,
	0x03, 0x93, 0xb3, 0xc9, 0x11, 0x74, 0xdd, 0xa7, 0x4c, 0xa2, 0xd3, 0x40, 0xb8, 0xf4, 0x37, 0xdd,
	0x85, 0xd3, 0x13, 0x88, 0x96, 0x5b, 0x5e, 0x8d, 0xe6, 0x9d, 0xda, 0x68, 0xfe, 0xf6, 0xfc, 0xed,
	0x33, 0x6d, 0x84, 0x32, 0xb2, 0xc8, 0xe5, 0x6f, 0x43, 0xa9, 0x47, 0x2b, 0x6f, 0x64, 0xae, 0x67,
	0x23, 0x33, 0x1d, 0xad, 0x7f, 0x55, 0x7d, 0x63, 0xa6, 0xee, 0x77, 0xda, 0x76, 0x4b, 0x3f, 0xfd,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xca, 0xd1, 0x13, 0x98, 0x76, 0x09, 0x00, 0x00,
}
