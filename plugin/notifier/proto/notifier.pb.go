// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: proto/notifier.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp         int64             `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Action            string            `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Username          string            `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	FsPath            string            `protobuf:"bytes,4,opt,name=fs_path,json=fsPath,proto3" json:"fs_path,omitempty"`
	FsTargetPath      string            `protobuf:"bytes,5,opt,name=fs_target_path,json=fsTargetPath,proto3" json:"fs_target_path,omitempty"`
	SshCmd            string            `protobuf:"bytes,6,opt,name=ssh_cmd,json=sshCmd,proto3" json:"ssh_cmd,omitempty"`
	FileSize          int64             `protobuf:"varint,7,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Protocol          string            `protobuf:"bytes,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Status            int32             `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Ip                string            `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty"`
	VirtualPath       string            `protobuf:"bytes,11,opt,name=virtual_path,json=virtualPath,proto3" json:"virtual_path,omitempty"`
	VirtualTargetPath string            `protobuf:"bytes,12,opt,name=virtual_target_path,json=virtualTargetPath,proto3" json:"virtual_target_path,omitempty"`
	SessionId         string            `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	FsProvider        int32             `protobuf:"varint,14,opt,name=fs_provider,json=fsProvider,proto3" json:"fs_provider,omitempty"`
	Bucket            string            `protobuf:"bytes,15,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Endpoint          string            `protobuf:"bytes,16,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	OpenFlags         int32             `protobuf:"varint,17,opt,name=open_flags,json=openFlags,proto3" json:"open_flags,omitempty"`
	Role              string            `protobuf:"bytes,18,opt,name=role,proto3" json:"role,omitempty"`
	Elapsed           int64             `protobuf:"varint,19,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	Metadata          map[string]string `protobuf:"bytes,20,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FsEvent) Reset() {
	*x = FsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FsEvent) ProtoMessage() {}

func (x *FsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FsEvent.ProtoReflect.Descriptor instead.
func (*FsEvent) Descriptor() ([]byte, []int) {
	return file_proto_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *FsEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *FsEvent) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FsEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FsEvent) GetFsPath() string {
	if x != nil {
		return x.FsPath
	}
	return ""
}

func (x *FsEvent) GetFsTargetPath() string {
	if x != nil {
		return x.FsTargetPath
	}
	return ""
}

func (x *FsEvent) GetSshCmd() string {
	if x != nil {
		return x.SshCmd
	}
	return ""
}

func (x *FsEvent) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FsEvent) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *FsEvent) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FsEvent) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *FsEvent) GetVirtualPath() string {
	if x != nil {
		return x.VirtualPath
	}
	return ""
}

func (x *FsEvent) GetVirtualTargetPath() string {
	if x != nil {
		return x.VirtualTargetPath
	}
	return ""
}

func (x *FsEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *FsEvent) GetFsProvider() int32 {
	if x != nil {
		return x.FsProvider
	}
	return 0
}

func (x *FsEvent) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FsEvent) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *FsEvent) GetOpenFlags() int32 {
	if x != nil {
		return x.OpenFlags
	}
	return 0
}

func (x *FsEvent) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *FsEvent) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *FsEvent) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ProviderEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp  int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Action     string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	ObjectType string `protobuf:"bytes,3,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	Username   string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Ip         string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	ObjectName string `protobuf:"bytes,6,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ObjectData []byte `protobuf:"bytes,7,opt,name=object_data,json=objectData,proto3" json:"object_data,omitempty"` // object JSON serialized
	Role       string `protobuf:"bytes,8,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *ProviderEvent) Reset() {
	*x = ProviderEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderEvent) ProtoMessage() {}

func (x *ProviderEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderEvent.ProtoReflect.Descriptor instead.
func (*ProviderEvent) Descriptor() ([]byte, []int) {
	return file_proto_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *ProviderEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ProviderEvent) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ProviderEvent) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *ProviderEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ProviderEvent) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ProviderEvent) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *ProviderEvent) GetObjectData() []byte {
	if x != nil {
		return x.ObjectData
	}
	return nil
}

func (x *ProviderEvent) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Event     int32  `protobuf:"varint,2,opt,name=event,proto3" json:"event,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Protocol  string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip        string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Message   string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Role      string `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_proto_notifier_proto_rawDescGZIP(), []int{2}
}

func (x *LogEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogEvent) GetEvent() int32 {
	if x != nil {
		return x.Event
	}
	return 0
}

func (x *LogEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LogEvent) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *LogEvent) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LogEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEvent) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_proto_notifier_proto protoreflect.FileDescriptor

var file_proto_notifier_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x05, 0x0a, 0x07, 0x46,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x73, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x73, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x73, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x73, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x73, 0x68, 0x5f, 0x63,
	0x6d, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x73, 0x68, 0x43, 0x6d, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x73, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe8, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0xbd,
	0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x46, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x41, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x17,
	0x5a, 0x15, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notifier_proto_rawDescOnce sync.Once
	file_proto_notifier_proto_rawDescData = file_proto_notifier_proto_rawDesc
)

func file_proto_notifier_proto_rawDescGZIP() []byte {
	file_proto_notifier_proto_rawDescOnce.Do(func() {
		file_proto_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notifier_proto_rawDescData)
	})
	return file_proto_notifier_proto_rawDescData
}

var file_proto_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_notifier_proto_goTypes = []interface{}{
	(*FsEvent)(nil),       // 0: proto.FsEvent
	(*ProviderEvent)(nil), // 1: proto.ProviderEvent
	(*LogEvent)(nil),      // 2: proto.LogEvent
	nil,                   // 3: proto.FsEvent.MetadataEntry
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_proto_notifier_proto_depIdxs = []int32{
	3, // 0: proto.FsEvent.metadata:type_name -> proto.FsEvent.MetadataEntry
	0, // 1: proto.Notifier.SendFsEvent:input_type -> proto.FsEvent
	1, // 2: proto.Notifier.SendProviderEvent:input_type -> proto.ProviderEvent
	2, // 3: proto.Notifier.SendLogEvent:input_type -> proto.LogEvent
	4, // 4: proto.Notifier.SendFsEvent:output_type -> google.protobuf.Empty
	4, // 5: proto.Notifier.SendProviderEvent:output_type -> google.protobuf.Empty
	4, // 6: proto.Notifier.SendLogEvent:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_notifier_proto_init() }
func file_proto_notifier_proto_init() {
	if File_proto_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FsEvent); i {
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
		file_proto_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderEvent); i {
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
		file_proto_notifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEvent); i {
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
			RawDescriptor: file_proto_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notifier_proto_goTypes,
		DependencyIndexes: file_proto_notifier_proto_depIdxs,
		MessageInfos:      file_proto_notifier_proto_msgTypes,
	}.Build()
	File_proto_notifier_proto = out.File
	file_proto_notifier_proto_rawDesc = nil
	file_proto_notifier_proto_goTypes = nil
	file_proto_notifier_proto_depIdxs = nil
}
