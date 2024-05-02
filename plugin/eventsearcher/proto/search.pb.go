// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: proto/search.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FsEventsFilter_Order int32

const (
	FsEventsFilter_DESC FsEventsFilter_Order = 0
	FsEventsFilter_ASC  FsEventsFilter_Order = 1
)

// Enum value maps for FsEventsFilter_Order.
var (
	FsEventsFilter_Order_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	FsEventsFilter_Order_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x FsEventsFilter_Order) Enum() *FsEventsFilter_Order {
	p := new(FsEventsFilter_Order)
	*p = x
	return p
}

func (x FsEventsFilter_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FsEventsFilter_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_search_proto_enumTypes[0].Descriptor()
}

func (FsEventsFilter_Order) Type() protoreflect.EnumType {
	return &file_proto_search_proto_enumTypes[0]
}

func (x FsEventsFilter_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FsEventsFilter_Order.Descriptor instead.
func (FsEventsFilter_Order) EnumDescriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{0, 0}
}

type ProviderEventsFilter_Order int32

const (
	ProviderEventsFilter_DESC ProviderEventsFilter_Order = 0
	ProviderEventsFilter_ASC  ProviderEventsFilter_Order = 1
)

// Enum value maps for ProviderEventsFilter_Order.
var (
	ProviderEventsFilter_Order_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	ProviderEventsFilter_Order_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x ProviderEventsFilter_Order) Enum() *ProviderEventsFilter_Order {
	p := new(ProviderEventsFilter_Order)
	*p = x
	return p
}

func (x ProviderEventsFilter_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderEventsFilter_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_search_proto_enumTypes[1].Descriptor()
}

func (ProviderEventsFilter_Order) Type() protoreflect.EnumType {
	return &file_proto_search_proto_enumTypes[1]
}

func (x ProviderEventsFilter_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderEventsFilter_Order.Descriptor instead.
func (ProviderEventsFilter_Order) EnumDescriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{1, 0}
}

type LogEventsFilter_Order int32

const (
	LogEventsFilter_DESC LogEventsFilter_Order = 0
	LogEventsFilter_ASC  LogEventsFilter_Order = 1
)

// Enum value maps for LogEventsFilter_Order.
var (
	LogEventsFilter_Order_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	LogEventsFilter_Order_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x LogEventsFilter_Order) Enum() *LogEventsFilter_Order {
	p := new(LogEventsFilter_Order)
	*p = x
	return p
}

func (x LogEventsFilter_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogEventsFilter_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_search_proto_enumTypes[2].Descriptor()
}

func (LogEventsFilter_Order) Type() protoreflect.EnumType {
	return &file_proto_search_proto_enumTypes[2]
}

func (x LogEventsFilter_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogEventsFilter_Order.Descriptor instead.
func (LogEventsFilter_Order) EnumDescriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{2, 0}
}

type FsEventsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimestamp int64                `protobuf:"varint,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   int64                `protobuf:"varint,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	Actions        []string             `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Username       string               `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Ip             string               `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	SshCmd         string               `protobuf:"bytes,6,opt,name=ssh_cmd,json=sshCmd,proto3" json:"ssh_cmd,omitempty"`
	Protocols      []string             `protobuf:"bytes,7,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Statuses       []int32              `protobuf:"varint,8,rep,packed,name=statuses,proto3" json:"statuses,omitempty"`
	InstanceIds    []string             `protobuf:"bytes,9,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	Limit          int32                `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	FromId         string               `protobuf:"bytes,11,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	FsProvider     int32                `protobuf:"varint,12,opt,name=fs_provider,json=fsProvider,proto3" json:"fs_provider,omitempty"`
	Bucket         string               `protobuf:"bytes,13,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Endpoint       string               `protobuf:"bytes,14,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Order          FsEventsFilter_Order `protobuf:"varint,15,opt,name=order,proto3,enum=proto.FsEventsFilter_Order" json:"order,omitempty"`
	Role           string               `protobuf:"bytes,16,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *FsEventsFilter) Reset() {
	*x = FsEventsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FsEventsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FsEventsFilter) ProtoMessage() {}

func (x *FsEventsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FsEventsFilter.ProtoReflect.Descriptor instead.
func (*FsEventsFilter) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{0}
}

func (x *FsEventsFilter) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *FsEventsFilter) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *FsEventsFilter) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *FsEventsFilter) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FsEventsFilter) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *FsEventsFilter) GetSshCmd() string {
	if x != nil {
		return x.SshCmd
	}
	return ""
}

func (x *FsEventsFilter) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *FsEventsFilter) GetStatuses() []int32 {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *FsEventsFilter) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

func (x *FsEventsFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FsEventsFilter) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *FsEventsFilter) GetFsProvider() int32 {
	if x != nil {
		return x.FsProvider
	}
	return 0
}

func (x *FsEventsFilter) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FsEventsFilter) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *FsEventsFilter) GetOrder() FsEventsFilter_Order {
	if x != nil {
		return x.Order
	}
	return FsEventsFilter_DESC
}

func (x *FsEventsFilter) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type ProviderEventsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimestamp int64                      `protobuf:"varint,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   int64                      `protobuf:"varint,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	Actions        []string                   `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Username       string                     `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Ip             string                     `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	ObjectTypes    []string                   `protobuf:"bytes,6,rep,name=object_types,json=objectTypes,proto3" json:"object_types,omitempty"`
	ObjectName     string                     `protobuf:"bytes,7,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	InstanceIds    []string                   `protobuf:"bytes,8,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	Limit          int32                      `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	FromId         string                     `protobuf:"bytes,10,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	Order          ProviderEventsFilter_Order `protobuf:"varint,11,opt,name=order,proto3,enum=proto.ProviderEventsFilter_Order" json:"order,omitempty"`
	Role           string                     `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
	OmitObjectData bool                       `protobuf:"varint,13,opt,name=omit_object_data,json=omitObjectData,proto3" json:"omit_object_data,omitempty"`
}

func (x *ProviderEventsFilter) Reset() {
	*x = ProviderEventsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderEventsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderEventsFilter) ProtoMessage() {}

func (x *ProviderEventsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderEventsFilter.ProtoReflect.Descriptor instead.
func (*ProviderEventsFilter) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{1}
}

func (x *ProviderEventsFilter) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *ProviderEventsFilter) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *ProviderEventsFilter) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *ProviderEventsFilter) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ProviderEventsFilter) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ProviderEventsFilter) GetObjectTypes() []string {
	if x != nil {
		return x.ObjectTypes
	}
	return nil
}

func (x *ProviderEventsFilter) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *ProviderEventsFilter) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

func (x *ProviderEventsFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ProviderEventsFilter) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *ProviderEventsFilter) GetOrder() ProviderEventsFilter_Order {
	if x != nil {
		return x.Order
	}
	return ProviderEventsFilter_DESC
}

func (x *ProviderEventsFilter) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ProviderEventsFilter) GetOmitObjectData() bool {
	if x != nil {
		return x.OmitObjectData
	}
	return false
}

type LogEventsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimestamp int64                 `protobuf:"varint,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   int64                 `protobuf:"varint,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	Events         []int32               `protobuf:"varint,3,rep,packed,name=events,proto3" json:"events,omitempty"`
	Username       string                `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Ip             string                `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Protocols      []string              `protobuf:"bytes,6,rep,name=protocols,proto3" json:"protocols,omitempty"`
	InstanceIds    []string              `protobuf:"bytes,7,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	Limit          int32                 `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	FromId         string                `protobuf:"bytes,9,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	Order          LogEventsFilter_Order `protobuf:"varint,10,opt,name=order,proto3,enum=proto.LogEventsFilter_Order" json:"order,omitempty"`
	Role           string                `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *LogEventsFilter) Reset() {
	*x = LogEventsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEventsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventsFilter) ProtoMessage() {}

func (x *LogEventsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventsFilter.ProtoReflect.Descriptor instead.
func (*LogEventsFilter) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{2}
}

func (x *LogEventsFilter) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *LogEventsFilter) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *LogEventsFilter) GetEvents() []int32 {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *LogEventsFilter) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LogEventsFilter) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LogEventsFilter) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *LogEventsFilter) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

func (x *LogEventsFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LogEventsFilter) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *LogEventsFilter) GetOrder() LogEventsFilter_Order {
	if x != nil {
		return x.Order
	}
	return LogEventsFilter_DESC
}

func (x *LogEventsFilter) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // JSON serialized response to return
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_search_proto protoreflect.FileDescriptor

var file_proto_search_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x04, 0x0a, 0x0e,
	0x46, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x73, 0x68, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x73, 0x68, 0x43, 0x6d, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x22,
	0xd3, 0x03, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6f, 0x6d, 0x69, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x53, 0x43, 0x10, 0x01, 0x22, 0xf7, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x22,
	0x24, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd8, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1c, 0x5a, 0x1a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_proto_rawDescOnce sync.Once
	file_proto_search_proto_rawDescData = file_proto_search_proto_rawDesc
)

func file_proto_search_proto_rawDescGZIP() []byte {
	file_proto_search_proto_rawDescOnce.Do(func() {
		file_proto_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_proto_rawDescData)
	})
	return file_proto_search_proto_rawDescData
}

var file_proto_search_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_search_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_search_proto_goTypes = []interface{}{
	(FsEventsFilter_Order)(0),       // 0: proto.FsEventsFilter.Order
	(ProviderEventsFilter_Order)(0), // 1: proto.ProviderEventsFilter.Order
	(LogEventsFilter_Order)(0),      // 2: proto.LogEventsFilter.Order
	(*FsEventsFilter)(nil),          // 3: proto.FsEventsFilter
	(*ProviderEventsFilter)(nil),    // 4: proto.ProviderEventsFilter
	(*LogEventsFilter)(nil),         // 5: proto.LogEventsFilter
	(*SearchResponse)(nil),          // 6: proto.SearchResponse
}
var file_proto_search_proto_depIdxs = []int32{
	0, // 0: proto.FsEventsFilter.order:type_name -> proto.FsEventsFilter.Order
	1, // 1: proto.ProviderEventsFilter.order:type_name -> proto.ProviderEventsFilter.Order
	2, // 2: proto.LogEventsFilter.order:type_name -> proto.LogEventsFilter.Order
	3, // 3: proto.Searcher.SearchFsEvents:input_type -> proto.FsEventsFilter
	4, // 4: proto.Searcher.SearchProviderEvents:input_type -> proto.ProviderEventsFilter
	5, // 5: proto.Searcher.SearchLogEvents:input_type -> proto.LogEventsFilter
	6, // 6: proto.Searcher.SearchFsEvents:output_type -> proto.SearchResponse
	6, // 7: proto.Searcher.SearchProviderEvents:output_type -> proto.SearchResponse
	6, // 8: proto.Searcher.SearchLogEvents:output_type -> proto.SearchResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_search_proto_init() }
func file_proto_search_proto_init() {
	if File_proto_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FsEventsFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderEventsFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEventsFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_search_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_proto_goTypes,
		DependencyIndexes: file_proto_search_proto_depIdxs,
		EnumInfos:         file_proto_search_proto_enumTypes,
		MessageInfos:      file_proto_search_proto_msgTypes,
	}.Build()
	File_proto_search_proto = out.File
	file_proto_search_proto_rawDesc = nil
	file_proto_search_proto_goTypes = nil
	file_proto_search_proto_depIdxs = nil
}
