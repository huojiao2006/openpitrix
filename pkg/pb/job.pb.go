// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateJobRequest struct {
	// required, cluster id
	ClusterId *wrappers.StringValue `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// required, id of app run in cluster
	AppId *wrappers.StringValue `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// id of specific app version
	VersionId *wrappers.StringValue `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// required, describe job's action.eg:[CreateCluster|StartClusters|...]
	JobAction *wrappers.StringValue `protobuf:"bytes,4,opt,name=job_action,json=jobAction,proto3" json:"job_action,omitempty"`
	// required, runtime provide.eg:[qingcloud|aliyun|aws|kubernetes]
	Provider *wrappers.StringValue `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	// required, directive, a json string, describe the info of running the job action
	Directive *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive,proto3" json:"directive,omitempty"`
	// required, runtime id
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,7,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *CreateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRequest.Unmarshal(m, b)
}
func (m *CreateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRequest.Merge(m, src)
}
func (m *CreateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJobRequest.Size(m)
}
func (m *CreateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRequest proto.InternalMessageInfo

func (m *CreateJobRequest) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobRequest) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateJobRequest) GetJobAction() *wrappers.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *CreateJobRequest) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *CreateJobRequest) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *CreateJobRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

type CreateJobResponse struct {
	// id of job created
	JobId *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// id of cluster run job
	ClusterId *wrappers.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// id of app deploy in cluster
	AppId *wrappers.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// id of specific app version
	VersionId *wrappers.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// id of runtime of cluster
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,5,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateJobResponse) Reset()         { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()    {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *CreateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobResponse.Unmarshal(m, b)
}
func (m *CreateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobResponse.Marshal(b, m, deterministic)
}
func (m *CreateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobResponse.Merge(m, src)
}
func (m *CreateJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateJobResponse.Size(m)
}
func (m *CreateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobResponse proto.InternalMessageInfo

func (m *CreateJobResponse) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateJobResponse) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobResponse) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobResponse) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateJobResponse) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

type Job struct {
	// job id
	JobId *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// id of cluster run job
	ClusterId *wrappers.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// id of app deployed in cluster
	AppId *wrappers.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// id of specific app version
	VersionId *wrappers.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// describe job's action eg:[CreateCluster|StartClusters|...]
	JobAction *wrappers.StringValue `protobuf:"bytes,5,opt,name=job_action,json=jobAction,proto3" json:"job_action,omitempty"`
	// status eg.[successful|failed|running|pending]
	Status *wrappers.StringValue `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// error code, if job run failed will return a error code
	ErrorCode *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// directive, a json string, describe the info of running the job action
	Directive *wrappers.StringValue `protobuf:"bytes,8,opt,name=directive,proto3" json:"directive,omitempty"`
	// host name of server
	Executor *wrappers.StringValue `protobuf:"bytes,9,opt,name=executor,proto3" json:"executor,omitempty"`
	// total count of task in job, a job contain one more task
	TaskCount *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=task_count,json=taskCount,proto3" json:"task_count,omitempty"`
	// own path, concat string group_path:user_id
	OwnerPath *wrappers.StringValue `protobuf:"bytes,11,opt,name=owner_path,json=ownerPath,proto3" json:"owner_path,omitempty"`
	// runtime provider eg:[qingcloud|aliyun|aws|kubernetes]
	Provider *wrappers.StringValue `protobuf:"bytes,12,opt,name=provider,proto3" json:"provider,omitempty"`
	// id of runtime of cluster
	RuntimeId *wrappers.StringValue `protobuf:"bytes,13,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	// the time job create
	CreateTime *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// record the status changed time
	StatusTime *timestamp.Timestamp `protobuf:"bytes,15,opt,name=status_time,json=statusTime,proto3" json:"status_time,omitempty"`
	// owner
	Owner                *wrappers.StringValue `protobuf:"bytes,16,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Job) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *Job) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *Job) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *Job) GetJobAction() *wrappers.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *Job) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Job) GetErrorCode() *wrappers.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Job) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Job) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Job) GetTaskCount() *wrappers.UInt32Value {
	if m != nil {
		return m.TaskCount
	}
	return nil
}

func (m *Job) GetOwnerPath() *wrappers.StringValue {
	if m != nil {
		return m.OwnerPath
	}
	return nil
}

func (m *Job) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *Job) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *Job) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Job) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *Job) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

type DescribeJobsRequest struct {
	// query key, support these fields(job_id, cluster_id, app_id, version_id, executor, provider, status, owner)
	SearchWord *wrappers.StringValue `protobuf:"bytes,1,opt,name=search_word,json=searchWord,proto3" json:"search_word,omitempty"`
	// sort key, order by sort_key, default create_time
	SortKey *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
	// value = 0 sort ASC, value = 1 sort DESC
	Reverse *wrappers.BoolValue `protobuf:"bytes,3,opt,name=reverse,proto3" json:"reverse,omitempty"`
	// data limit per page, default value 20, max value 200
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// data offset, default 0
	Offset uint32 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// select column to display
	DisplayColumns []string `protobuf:"bytes,6,rep,name=display_columns,json=displayColumns,proto3" json:"display_columns,omitempty"`
	// job ids
	JobId []string `protobuf:"bytes,11,rep,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// cluster id
	ClusterId *wrappers.StringValue `protobuf:"bytes,12,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// app id
	AppId *wrappers.StringValue `protobuf:"bytes,13,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// specific app version id to filter result
	VersionId *wrappers.StringValue `protobuf:"bytes,14,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// host name of server
	Executor *wrappers.StringValue `protobuf:"bytes,15,opt,name=executor,proto3" json:"executor,omitempty"`
	// runtime provider eg.[qingcloud|aliyun|aws|kubernetes]
	Provider *wrappers.StringValue `protobuf:"bytes,16,opt,name=provider,proto3" json:"provider,omitempty"`
	// runtime id
	RuntimeId *wrappers.StringValue `protobuf:"bytes,17,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	// status eg.[successful|failed|running|pending]
	Status []string `protobuf:"bytes,18,rep,name=status,proto3" json:"status,omitempty"`
	// owner
	Owner                []string `protobuf:"bytes,19,rep,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeJobsRequest) Reset()         { *m = DescribeJobsRequest{} }
func (m *DescribeJobsRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeJobsRequest) ProtoMessage()    {}
func (*DescribeJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *DescribeJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeJobsRequest.Unmarshal(m, b)
}
func (m *DescribeJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeJobsRequest.Marshal(b, m, deterministic)
}
func (m *DescribeJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeJobsRequest.Merge(m, src)
}
func (m *DescribeJobsRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeJobsRequest.Size(m)
}
func (m *DescribeJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeJobsRequest proto.InternalMessageInfo

func (m *DescribeJobsRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

func (m *DescribeJobsRequest) GetSortKey() *wrappers.StringValue {
	if m != nil {
		return m.SortKey
	}
	return nil
}

func (m *DescribeJobsRequest) GetReverse() *wrappers.BoolValue {
	if m != nil {
		return m.Reverse
	}
	return nil
}

func (m *DescribeJobsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeJobsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeJobsRequest) GetDisplayColumns() []string {
	if m != nil {
		return m.DisplayColumns
	}
	return nil
}

func (m *DescribeJobsRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeJobsRequest) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *DescribeJobsRequest) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *DescribeJobsRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *DescribeJobsRequest) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeJobsRequest) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *DescribeJobsRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *DescribeJobsRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeJobsRequest) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

type DescribeJobsResponse struct {
	// total count of job
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// list of job
	JobSet               []*Job   `protobuf:"bytes,2,rep,name=job_set,json=jobSet,proto3" json:"job_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeJobsResponse) Reset()         { *m = DescribeJobsResponse{} }
func (m *DescribeJobsResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeJobsResponse) ProtoMessage()    {}
func (*DescribeJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{4}
}

func (m *DescribeJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeJobsResponse.Unmarshal(m, b)
}
func (m *DescribeJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeJobsResponse.Marshal(b, m, deterministic)
}
func (m *DescribeJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeJobsResponse.Merge(m, src)
}
func (m *DescribeJobsResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeJobsResponse.Size(m)
}
func (m *DescribeJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeJobsResponse proto.InternalMessageInfo

func (m *DescribeJobsResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeJobsResponse) GetJobSet() []*Job {
	if m != nil {
		return m.JobSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateJobRequest)(nil), "openpitrix.CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "openpitrix.CreateJobResponse")
	proto.RegisterType((*Job)(nil), "openpitrix.Job")
	proto.RegisterType((*DescribeJobsRequest)(nil), "openpitrix.DescribeJobsRequest")
	proto.RegisterType((*DescribeJobsResponse)(nil), "openpitrix.DescribeJobsResponse")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0x4f, 0x6f, 0x23, 0x35,
	0x14, 0x57, 0x92, 0xa6, 0x49, 0x9c, 0xfe, 0x5b, 0x6f, 0x59, 0x8d, 0xa2, 0x42, 0xad, 0x5e, 0x08,
	0xd2, 0x6c, 0x22, 0xd2, 0x95, 0x58, 0xb1, 0xe2, 0xb0, 0x5b, 0x24, 0x68, 0x11, 0x12, 0xca, 0xf2,
	0x47, 0xe2, 0x32, 0xf2, 0xcc, 0xbc, 0x24, 0xce, 0x4e, 0xc7, 0xc6, 0xf6, 0x34, 0xdb, 0x2b, 0x12,
	0x12, 0x67, 0xb8, 0xf2, 0x19, 0x90, 0x90, 0xb8, 0x73, 0xe6, 0xcc, 0x57, 0xe0, 0x03, 0xf0, 0x11,
	0x90, 0xed, 0x99, 0x74, 0xba, 0xdd, 0x8a, 0x99, 0x72, 0xdc, 0x53, 0x34, 0xf6, 0xef, 0x97, 0xf7,
	0xfc, 0x7e, 0xef, 0xf7, 0x6c, 0xd4, 0x5b, 0xf2, 0x70, 0x24, 0x24, 0xd7, 0x1c, 0x23, 0x2e, 0x20,
	0x15, 0x4c, 0x4b, 0xf6, 0x72, 0xf0, 0xce, 0x9c, 0xf3, 0x79, 0x02, 0x63, 0xbb, 0x13, 0x66, 0xb3,
	0xf1, 0x4a, 0x52, 0x21, 0x40, 0x2a, 0x87, 0x1d, 0x1c, 0xbe, 0xba, 0xaf, 0xd9, 0x39, 0x28, 0x4d,
	0xcf, 0x45, 0x0e, 0x38, 0xc8, 0x01, 0x54, 0xb0, 0x31, 0x4d, 0x53, 0xae, 0xa9, 0x66, 0x3c, 0x2d,
	0xe8, 0xbe, 0xfd, 0x89, 0x1e, 0xce, 0x21, 0x7d, 0xa8, 0x56, 0x74, 0x3e, 0x07, 0x39, 0xe6, 0xc2,
	0x22, 0x6e, 0xa2, 0x8f, 0x7e, 0x6b, 0xa1, 0xbd, 0x13, 0x09, 0x54, 0xc3, 0x19, 0x0f, 0xa7, 0xf0,
	0x5d, 0x06, 0x4a, 0xe3, 0x27, 0x08, 0x45, 0x49, 0xa6, 0x34, 0xc8, 0x80, 0xc5, 0x5e, 0x83, 0x34,
	0x86, 0xfd, 0xc9, 0xc1, 0xc8, 0x45, 0x1d, 0x15, 0x69, 0x8d, 0x9e, 0x6b, 0xc9, 0xd2, 0xf9, 0xd7,
	0x34, 0xc9, 0x60, 0xda, 0xcb, 0xf1, 0xa7, 0x31, 0x3e, 0x46, 0x9b, 0x54, 0x08, 0x43, 0x6c, 0x56,
	0x20, 0xb6, 0xa9, 0x10, 0xa7, 0xb1, 0x89, 0x78, 0x01, 0x52, 0x31, 0x9e, 0x1a, 0x62, 0xab, 0x4a,
	0xc4, 0x1c, 0xef, 0xc8, 0x4b, 0x1e, 0x06, 0x34, 0x32, 0x07, 0xf3, 0x36, 0xaa, 0x90, 0x97, 0x3c,
	0x7c, 0x6a, 0xe1, 0xf8, 0x31, 0xea, 0x0a, 0xc9, 0x2f, 0x58, 0x0c, 0xd2, 0x6b, 0x57, 0xa0, 0xae,
	0xd1, 0xf8, 0x43, 0xd4, 0x8b, 0x99, 0x84, 0x48, 0xb3, 0x0b, 0xf0, 0x36, 0xab, 0x44, 0x5d, 0xc3,
	0x4d, 0xca, 0x32, 0x4b, 0x8d, 0xb0, 0xe6, 0xbc, 0x9d, 0x2a, 0xe4, 0x1c, 0x7f, 0x1a, 0x1f, 0xfd,
	0xde, 0x44, 0xf7, 0x4a, 0x9a, 0x29, 0xc1, 0x53, 0x05, 0xa6, 0xee, 0xa6, 0x0a, 0x15, 0x05, 0x6b,
	0x2f, 0x79, 0xe8, 0x4a, 0x57, 0x52, 0xba, 0x79, 0x57, 0xa5, 0x5b, 0x77, 0x55, 0x7a, 0xa3, 0xb6,
	0xd2, 0xa5, 0xb2, 0xb5, 0xeb, 0x95, 0xed, 0xd7, 0x0e, 0x6a, 0x9d, 0xf1, 0xf0, 0x4d, 0x29, 0x54,
	0xc9, 0x12, 0xed, 0x7a, 0x96, 0x78, 0x84, 0x36, 0x95, 0xa6, 0x3a, 0x53, 0x95, 0xba, 0x3a, 0xc7,
	0x9a, 0x90, 0x20, 0x25, 0x97, 0x41, 0xc4, 0x63, 0xb8, 0xb5, 0xa5, 0xbf, 0x3a, 0x4d, 0xf5, 0xf1,
	0x24, 0x0f, 0x69, 0xf1, 0x27, 0x3c, 0x86, 0xeb, 0x5e, 0xea, 0xd6, 0xf3, 0xd2, 0x63, 0xd4, 0x85,
	0x97, 0x10, 0x65, 0x9a, 0x4b, 0xaf, 0x57, 0xc5, 0xc1, 0x05, 0xda, 0xa4, 0xac, 0xa9, 0x7a, 0x11,
	0x44, 0x3c, 0x4b, 0xb5, 0x87, 0xaa, 0xa4, 0x6c, 0xf0, 0x27, 0x06, 0x6e, 0xc8, 0x7c, 0x95, 0x82,
	0x0c, 0x04, 0xd5, 0x0b, 0xaf, 0x5f, 0x25, 0x67, 0x8b, 0xff, 0x82, 0xea, 0xc5, 0xb5, 0xa9, 0xb3,
	0x55, 0x6b, 0xea, 0x5c, 0xb7, 0xc0, 0x76, 0x2d, 0x0b, 0xe0, 0x27, 0xa8, 0x1f, 0xd9, 0xc1, 0x11,
	0x98, 0x05, 0x6f, 0xc7, 0xb2, 0x07, 0x37, 0xd8, 0x5f, 0x16, 0x17, 0xce, 0x14, 0x39, 0xb8, 0x59,
	0x30, 0x64, 0x27, 0xb5, 0x23, 0xef, 0xfe, 0x37, 0xd9, 0xc1, 0x2d, 0x79, 0x82, 0xda, 0xf6, 0xf4,
	0xde, 0x5e, 0x15, 0x07, 0x58, 0xe8, 0xd1, 0x1f, 0x6d, 0x74, 0xff, 0x63, 0x50, 0x91, 0x64, 0xa1,
	0x99, 0x74, 0xaa, 0xb8, 0x9e, 0x3e, 0x42, 0x7d, 0x05, 0x54, 0x46, 0x8b, 0x60, 0xc5, 0x65, 0x35,
	0x17, 0x23, 0x47, 0xf8, 0x86, 0xcb, 0x18, 0x7f, 0x80, 0xba, 0x8a, 0x4b, 0x1d, 0xbc, 0x80, 0xcb,
	0x4a, 0x46, 0xee, 0x18, 0xf4, 0x67, 0x70, 0x89, 0x1f, 0xa1, 0x8e, 0x04, 0xe3, 0x31, 0xc8, 0x7d,
	0x7c, 0xf3, 0xf0, 0xcf, 0x38, 0x4f, 0x72, 0x56, 0x0e, 0xc5, 0xfb, 0xa8, 0x9d, 0xb0, 0x73, 0xa6,
	0xad, 0x85, 0xb7, 0xa7, 0xee, 0x03, 0x3f, 0x40, 0x9b, 0x7c, 0x36, 0x53, 0xa0, 0xad, 0x39, 0xb7,
	0xa7, 0xf9, 0x17, 0x7e, 0x17, 0xed, 0xc6, 0x4c, 0x89, 0x84, 0x5e, 0x06, 0x11, 0x4f, 0xb2, 0xf3,
	0xd4, 0x98, 0xb0, 0x35, 0xec, 0x4d, 0x77, 0xf2, 0xe5, 0x13, 0xb7, 0x8a, 0xdf, 0x5a, 0x4f, 0xb1,
	0xbe, 0xdd, 0x7f, 0xed, 0x9c, 0xda, 0xba, 0xeb, 0x9c, 0xda, 0xbe, 0xeb, 0x9c, 0xda, 0xa9, 0x37,
	0xa7, 0xca, 0xde, 0xdd, 0xad, 0xe5, 0xdd, 0xb2, 0x83, 0xf6, 0xfe, 0x87, 0x83, 0xee, 0xd5, 0x73,
	0xd0, 0x83, 0xf5, 0x6c, 0xc4, 0xb6, 0xec, 0xc5, 0xf4, 0xdb, 0x2f, 0xfa, 0xfb, 0xbe, 0x53, 0xc3,
	0x75, 0x30, 0x45, 0xfb, 0xd7, 0x1b, 0x38, 0xbf, 0xab, 0x0f, 0x51, 0x5f, 0x73, 0x4d, 0x93, 0x7c,
	0xf2, 0x34, 0x6c, 0x0b, 0x20, 0xbb, 0xe4, 0x86, 0xcb, 0x10, 0x75, 0x8c, 0xba, 0xa6, 0x3f, 0x9a,
	0xa4, 0x35, 0xec, 0x4f, 0x76, 0x47, 0x57, 0x2f, 0xc8, 0x91, 0xb9, 0xf6, 0x8d, 0xfa, 0xcf, 0x41,
	0x4f, 0xfe, 0x6c, 0x22, 0x74, 0xc6, 0xc3, 0xcf, 0x69, 0x4a, 0xe7, 0x20, 0xf1, 0xa7, 0xa8, 0xb7,
	0x7e, 0x1a, 0xe0, 0x83, 0x32, 0xe9, 0xd5, 0x57, 0xde, 0xe0, 0xed, 0x5b, 0x76, 0xf3, 0x1c, 0xff,
	0x69, 0xa0, 0xad, 0x72, 0xf2, 0xf8, 0xb0, 0x8c, 0x7f, 0x8d, 0x2f, 0x07, 0xe4, 0x76, 0x80, 0xfb,
	0xcf, 0xa3, 0x5f, 0x1a, 0x3f, 0x3d, 0xfd, 0xb1, 0x81, 0x7f, 0x68, 0x7c, 0x02, 0x9a, 0x2c, 0x79,
	0xe8, 0x93, 0x19, 0x4b, 0x34, 0x48, 0xb2, 0x62, 0x7a, 0x41, 0xf4, 0x02, 0x14, 0x90, 0x19, 0x83,
	0x24, 0x56, 0x43, 0xd7, 0xdc, 0x3e, 0xb9, 0xea, 0x66, 0x9f, 0xb8, 0xe6, 0xf4, 0xc9, 0x55, 0xbf,
	0xf9, 0xa4, 0x68, 0x08, 0x9f, 0x14, 0x02, 0xfb, 0xc4, 0x69, 0xe2, 0x13, 0x2b, 0xc2, 0x7b, 0x3e,
	0x89, 0x61, 0x46, 0xb3, 0x44, 0x13, 0x09, 0x3a, 0x93, 0x29, 0xa1, 0x49, 0x62, 0x82, 0xab, 0xef,
	0xff, 0xfa, 0xfb, 0xe7, 0x26, 0xc2, 0xdd, 0xf1, 0xc5, 0xfb, 0x63, 0xf3, 0xfd, 0x6c, 0xe3, 0xdb,
	0xa6, 0x08, 0xc3, 0x4d, 0xdb, 0x03, 0xc7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x03, 0x9a, 0x48,
	0x87, 0xbf, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobManagerClient is the client API for JobManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobManagerClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	// Get job, filter with these fields(job_id, cluster_id, app_id, version_id, executor, provider, status, owner), default return all jobs
	DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error)
}

type jobManagerClient struct {
	cc *grpc.ClientConn
}

func NewJobManagerClient(cc *grpc.ClientConn) JobManagerClient {
	return &jobManagerClient{cc}
}

func (c *jobManagerClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.JobManager/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobManagerClient) DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error) {
	out := new(DescribeJobsResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.JobManager/DescribeJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobManagerServer is the server API for JobManager service.
type JobManagerServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	// Get job, filter with these fields(job_id, cluster_id, app_id, version_id, executor, provider, status, owner), default return all jobs
	DescribeJobs(context.Context, *DescribeJobsRequest) (*DescribeJobsResponse, error)
}

func RegisterJobManagerServer(s *grpc.Server, srv JobManagerServer) {
	s.RegisterService(&_JobManager_serviceDesc, srv)
}

func _JobManager_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobManager_DescribeJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).DescribeJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/DescribeJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).DescribeJobs(ctx, req.(*DescribeJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.JobManager",
	HandlerType: (*JobManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobManager_CreateJob_Handler,
		},
		{
			MethodName: "DescribeJobs",
			Handler:    _JobManager_DescribeJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
