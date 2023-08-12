// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: metadata/proto/metadata.proto

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

type SetModificationTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId        string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	ObjectPath       string `protobuf:"bytes,2,opt,name=object_path,json=objectPath,proto3" json:"object_path,omitempty"`
	ModificationTime int64  `protobuf:"varint,3,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
}

func (x *SetModificationTimeRequest) Reset() {
	*x = SetModificationTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetModificationTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetModificationTimeRequest) ProtoMessage() {}

func (x *SetModificationTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetModificationTimeRequest.ProtoReflect.Descriptor instead.
func (*SetModificationTimeRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *SetModificationTimeRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *SetModificationTimeRequest) GetObjectPath() string {
	if x != nil {
		return x.ObjectPath
	}
	return ""
}

func (x *SetModificationTimeRequest) GetModificationTime() int64 {
	if x != nil {
		return x.ModificationTime
	}
	return 0
}

type GetModificationTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId  string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	ObjectPath string `protobuf:"bytes,2,opt,name=object_path,json=objectPath,proto3" json:"object_path,omitempty"`
}

func (x *GetModificationTimeRequest) Reset() {
	*x = GetModificationTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModificationTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModificationTimeRequest) ProtoMessage() {}

func (x *GetModificationTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModificationTimeRequest.ProtoReflect.Descriptor instead.
func (*GetModificationTimeRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *GetModificationTimeRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *GetModificationTimeRequest) GetObjectPath() string {
	if x != nil {
		return x.ObjectPath
	}
	return ""
}

type GetModificationTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModificationTime int64 `protobuf:"varint,1,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
}

func (x *GetModificationTimeResponse) Reset() {
	*x = GetModificationTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModificationTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModificationTimeResponse) ProtoMessage() {}

func (x *GetModificationTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModificationTimeResponse.ProtoReflect.Descriptor instead.
func (*GetModificationTimeResponse) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *GetModificationTimeResponse) GetModificationTime() int64 {
	if x != nil {
		return x.ModificationTime
	}
	return 0
}

type GetModificationTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId  string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	FolderPath string `protobuf:"bytes,2,opt,name=folder_path,json=folderPath,proto3" json:"folder_path,omitempty"`
}

func (x *GetModificationTimesRequest) Reset() {
	*x = GetModificationTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModificationTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModificationTimesRequest) ProtoMessage() {}

func (x *GetModificationTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModificationTimesRequest.ProtoReflect.Descriptor instead.
func (*GetModificationTimesRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *GetModificationTimesRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *GetModificationTimesRequest) GetFolderPath() string {
	if x != nil {
		return x.FolderPath
	}
	return ""
}

type GetModificationTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the file name (not the full path) is the map key and the modification time is the map value
	Pairs map[string]int64 `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetModificationTimesResponse) Reset() {
	*x = GetModificationTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModificationTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModificationTimesResponse) ProtoMessage() {}

func (x *GetModificationTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModificationTimesResponse.ProtoReflect.Descriptor instead.
func (*GetModificationTimesResponse) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *GetModificationTimesResponse) GetPairs() map[string]int64 {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type RemoveMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId  string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	ObjectPath string `protobuf:"bytes,2,opt,name=object_path,json=objectPath,proto3" json:"object_path,omitempty"`
}

func (x *RemoveMetadataRequest) Reset() {
	*x = RemoveMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMetadataRequest) ProtoMessage() {}

func (x *RemoveMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMetadataRequest.ProtoReflect.Descriptor instead.
func (*RemoveMetadataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveMetadataRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *RemoveMetadataRequest) GetObjectPath() string {
	if x != nil {
		return x.ObjectPath
	}
	return ""
}

type GetFoldersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *GetFoldersRequest) Reset() {
	*x = GetFoldersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoldersRequest) ProtoMessage() {}

func (x *GetFoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoldersRequest.ProtoReflect.Descriptor instead.
func (*GetFoldersRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{6}
}

func (x *GetFoldersRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *GetFoldersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetFoldersRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type GetFoldersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []string `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *GetFoldersResponse) Reset() {
	*x = GetFoldersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_metadata_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoldersResponse) ProtoMessage() {}

func (x *GetFoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_metadata_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoldersResponse.ProtoReflect.Descriptor instead.
func (*GetFoldersResponse) Descriptor() ([]byte, []int) {
	return file_metadata_proto_metadata_proto_rawDescGZIP(), []int{7}
}

func (x *GetFoldersResponse) GetFolders() []string {
	if x != nil {
		return x.Folders
	}
	return nil
}

var File_metadata_proto_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_metadata_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x5c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4a, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x22, 0x9e, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x50, 0x61, 0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x15, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x32, 0xa6, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x50, 0x0a,
	0x13, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_metadata_proto_rawDescData = file_metadata_proto_metadata_proto_rawDesc
)

func file_metadata_proto_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_metadata_proto_rawDescData)
	})
	return file_metadata_proto_metadata_proto_rawDescData
}

var file_metadata_proto_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_metadata_proto_metadata_proto_goTypes = []interface{}{
	(*SetModificationTimeRequest)(nil),   // 0: proto.SetModificationTimeRequest
	(*GetModificationTimeRequest)(nil),   // 1: proto.GetModificationTimeRequest
	(*GetModificationTimeResponse)(nil),  // 2: proto.GetModificationTimeResponse
	(*GetModificationTimesRequest)(nil),  // 3: proto.GetModificationTimesRequest
	(*GetModificationTimesResponse)(nil), // 4: proto.GetModificationTimesResponse
	(*RemoveMetadataRequest)(nil),        // 5: proto.RemoveMetadataRequest
	(*GetFoldersRequest)(nil),            // 6: proto.GetFoldersRequest
	(*GetFoldersResponse)(nil),           // 7: proto.GetFoldersResponse
	nil,                                  // 8: proto.GetModificationTimesResponse.PairsEntry
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
}
var file_metadata_proto_metadata_proto_depIdxs = []int32{
	8, // 0: proto.GetModificationTimesResponse.pairs:type_name -> proto.GetModificationTimesResponse.PairsEntry
	0, // 1: proto.Metadata.SetModificationTime:input_type -> proto.SetModificationTimeRequest
	1, // 2: proto.Metadata.GetModificationTime:input_type -> proto.GetModificationTimeRequest
	3, // 3: proto.Metadata.GetModificationTimes:input_type -> proto.GetModificationTimesRequest
	5, // 4: proto.Metadata.RemoveMetadata:input_type -> proto.RemoveMetadataRequest
	6, // 5: proto.Metadata.GetFolders:input_type -> proto.GetFoldersRequest
	9, // 6: proto.Metadata.SetModificationTime:output_type -> google.protobuf.Empty
	2, // 7: proto.Metadata.GetModificationTime:output_type -> proto.GetModificationTimeResponse
	4, // 8: proto.Metadata.GetModificationTimes:output_type -> proto.GetModificationTimesResponse
	9, // 9: proto.Metadata.RemoveMetadata:output_type -> google.protobuf.Empty
	7, // 10: proto.Metadata.GetFolders:output_type -> proto.GetFoldersResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metadata_proto_metadata_proto_init() }
func file_metadata_proto_metadata_proto_init() {
	if File_metadata_proto_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetModificationTimeRequest); i {
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
		file_metadata_proto_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModificationTimeRequest); i {
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
		file_metadata_proto_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModificationTimeResponse); i {
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
		file_metadata_proto_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModificationTimesRequest); i {
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
		file_metadata_proto_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModificationTimesResponse); i {
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
		file_metadata_proto_metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMetadataRequest); i {
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
		file_metadata_proto_metadata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoldersRequest); i {
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
		file_metadata_proto_metadata_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoldersResponse); i {
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
			RawDescriptor: file_metadata_proto_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_proto_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_metadata_proto_depIdxs,
		MessageInfos:      file_metadata_proto_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto_metadata_proto = out.File
	file_metadata_proto_metadata_proto_rawDesc = nil
	file_metadata_proto_metadata_proto_goTypes = nil
	file_metadata_proto_metadata_proto_depIdxs = nil
}
