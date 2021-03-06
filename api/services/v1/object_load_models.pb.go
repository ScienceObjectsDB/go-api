// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/services/v1/object_load_models.proto

package v1

import (
	v1 "github.com/ScienceObjectsDB/go-api/api/models/v1"
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

type CreateUploadLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUploadLink) Reset() {
	*x = CreateUploadLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadLink) ProtoMessage() {}

func (x *CreateUploadLink) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadLink.ProtoReflect.Descriptor instead.
func (*CreateUploadLink) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUploadLink) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateUploadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUploadLinkRequest) Reset() {
	*x = CreateUploadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadLinkRequest) ProtoMessage() {}

func (x *CreateUploadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateUploadLinkRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUploadLinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateUploadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadLink string `protobuf:"bytes,1,opt,name=upload_link,json=uploadLink,proto3" json:"upload_link,omitempty"`
}

func (x *CreateUploadLinkResponse) Reset() {
	*x = CreateUploadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadLinkResponse) ProtoMessage() {}

func (x *CreateUploadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateUploadLinkResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUploadLinkResponse) GetUploadLink() string {
	if x != nil {
		return x.UploadLink
	}
	return ""
}

type CreateDownloadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDownloadLinkRequest) Reset() {
	*x = CreateDownloadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLinkRequest) ProtoMessage() {}

func (x *CreateDownloadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateDownloadLinkRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDownloadLinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateDownloadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadLink string     `protobuf:"bytes,1,opt,name=upload_link,json=uploadLink,proto3" json:"upload_link,omitempty"`
	Object     *v1.Object `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *CreateDownloadLinkResponse) Reset() {
	*x = CreateDownloadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLinkResponse) ProtoMessage() {}

func (x *CreateDownloadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateDownloadLinkResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDownloadLinkResponse) GetUploadLink() string {
	if x != nil {
		return x.UploadLink
	}
	return ""
}

func (x *CreateDownloadLinkResponse) GetObject() *v1.Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type StartMultipartUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartMultipartUploadRequest) Reset() {
	*x = StartMultipartUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMultipartUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMultipartUploadRequest) ProtoMessage() {}

func (x *StartMultipartUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMultipartUploadRequest.ProtoReflect.Descriptor instead.
func (*StartMultipartUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{5}
}

func (x *StartMultipartUploadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StartMultipartUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *v1.Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *StartMultipartUploadResponse) Reset() {
	*x = StartMultipartUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMultipartUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMultipartUploadResponse) ProtoMessage() {}

func (x *StartMultipartUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMultipartUploadResponse.ProtoReflect.Descriptor instead.
func (*StartMultipartUploadResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{6}
}

func (x *StartMultipartUploadResponse) GetObject() *v1.Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type GetMultipartUploadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadLink string     `protobuf:"bytes,1,opt,name=upload_link,json=uploadLink,proto3" json:"upload_link,omitempty"`
	Object     *v1.Object `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetMultipartUploadLinkResponse) Reset() {
	*x = GetMultipartUploadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipartUploadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipartUploadLinkResponse) ProtoMessage() {}

func (x *GetMultipartUploadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipartUploadLinkResponse.ProtoReflect.Descriptor instead.
func (*GetMultipartUploadLinkResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{7}
}

func (x *GetMultipartUploadLinkResponse) GetUploadLink() string {
	if x != nil {
		return x.UploadLink
	}
	return ""
}

func (x *GetMultipartUploadLinkResponse) GetObject() *v1.Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type GetMultipartUploadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId   string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	UploadPart int64  `protobuf:"varint,2,opt,name=upload_part,json=uploadPart,proto3" json:"upload_part,omitempty"`
}

func (x *GetMultipartUploadLinkRequest) Reset() {
	*x = GetMultipartUploadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipartUploadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipartUploadLinkRequest) ProtoMessage() {}

func (x *GetMultipartUploadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipartUploadLinkRequest.ProtoReflect.Descriptor instead.
func (*GetMultipartUploadLinkRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{8}
}

func (x *GetMultipartUploadLinkRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *GetMultipartUploadLinkRequest) GetUploadPart() int64 {
	if x != nil {
		return x.UploadPart
	}
	return 0
}

type CompleteMultipartUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId string            `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Parts    []*CompletedParts `protobuf:"bytes,2,rep,name=parts,proto3" json:"parts,omitempty"`
}

func (x *CompleteMultipartUploadRequest) Reset() {
	*x = CompleteMultipartUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteMultipartUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteMultipartUploadRequest) ProtoMessage() {}

func (x *CompleteMultipartUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteMultipartUploadRequest.ProtoReflect.Descriptor instead.
func (*CompleteMultipartUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{9}
}

func (x *CompleteMultipartUploadRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *CompleteMultipartUploadRequest) GetParts() []*CompletedParts {
	if x != nil {
		return x.Parts
	}
	return nil
}

type CompleteMultipartUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteMultipartUploadResponse) Reset() {
	*x = CompleteMultipartUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteMultipartUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteMultipartUploadResponse) ProtoMessage() {}

func (x *CompleteMultipartUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteMultipartUploadResponse.ProtoReflect.Descriptor instead.
func (*CompleteMultipartUploadResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{10}
}

type CompletedParts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etag string `protobuf:"bytes,1,opt,name=etag,proto3" json:"etag,omitempty"`
	Part int64  `protobuf:"varint,2,opt,name=part,proto3" json:"part,omitempty"`
}

func (x *CompletedParts) Reset() {
	*x = CompletedParts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_object_load_models_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedParts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedParts) ProtoMessage() {}

func (x *CompletedParts) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_object_load_models_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedParts.ProtoReflect.Descriptor instead.
func (*CompletedParts) Descriptor() ([]byte, []int) {
	return file_api_services_v1_object_load_models_proto_rawDescGZIP(), []int{11}
}

func (x *CompletedParts) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *CompletedParts) GetPart() int64 {
	if x != nil {
		return x.Part
	}
	return 0
}

var File_api_services_v1_object_load_models_proto protoreflect.FileDescriptor

var file_api_services_v1_object_load_models_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x2b, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2d, 0x0a, 0x1b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x1c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x70, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x22, 0x74, 0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x73, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x22, 0x21, 0x0a, 0x1f, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x42, 0x85, 0x01, 0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x17, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_v1_object_load_models_proto_rawDescOnce sync.Once
	file_api_services_v1_object_load_models_proto_rawDescData = file_api_services_v1_object_load_models_proto_rawDesc
)

func file_api_services_v1_object_load_models_proto_rawDescGZIP() []byte {
	file_api_services_v1_object_load_models_proto_rawDescOnce.Do(func() {
		file_api_services_v1_object_load_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_v1_object_load_models_proto_rawDescData)
	})
	return file_api_services_v1_object_load_models_proto_rawDescData
}

var file_api_services_v1_object_load_models_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_services_v1_object_load_models_proto_goTypes = []interface{}{
	(*CreateUploadLink)(nil),                // 0: api.services.v1.CreateUploadLink
	(*CreateUploadLinkRequest)(nil),         // 1: api.services.v1.CreateUploadLinkRequest
	(*CreateUploadLinkResponse)(nil),        // 2: api.services.v1.CreateUploadLinkResponse
	(*CreateDownloadLinkRequest)(nil),       // 3: api.services.v1.CreateDownloadLinkRequest
	(*CreateDownloadLinkResponse)(nil),      // 4: api.services.v1.CreateDownloadLinkResponse
	(*StartMultipartUploadRequest)(nil),     // 5: api.services.v1.StartMultipartUploadRequest
	(*StartMultipartUploadResponse)(nil),    // 6: api.services.v1.StartMultipartUploadResponse
	(*GetMultipartUploadLinkResponse)(nil),  // 7: api.services.v1.GetMultipartUploadLinkResponse
	(*GetMultipartUploadLinkRequest)(nil),   // 8: api.services.v1.GetMultipartUploadLinkRequest
	(*CompleteMultipartUploadRequest)(nil),  // 9: api.services.v1.CompleteMultipartUploadRequest
	(*CompleteMultipartUploadResponse)(nil), // 10: api.services.v1.CompleteMultipartUploadResponse
	(*CompletedParts)(nil),                  // 11: api.services.v1.CompletedParts
	(*v1.Object)(nil),                       // 12: api.models.v1.Object
}
var file_api_services_v1_object_load_models_proto_depIdxs = []int32{
	12, // 0: api.services.v1.CreateDownloadLinkResponse.object:type_name -> api.models.v1.Object
	12, // 1: api.services.v1.StartMultipartUploadResponse.object:type_name -> api.models.v1.Object
	12, // 2: api.services.v1.GetMultipartUploadLinkResponse.object:type_name -> api.models.v1.Object
	11, // 3: api.services.v1.CompleteMultipartUploadRequest.parts:type_name -> api.services.v1.CompletedParts
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_services_v1_object_load_models_proto_init() }
func file_api_services_v1_object_load_models_proto_init() {
	if File_api_services_v1_object_load_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_v1_object_load_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadLink); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadLinkRequest); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadLinkResponse); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLinkRequest); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLinkResponse); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMultipartUploadRequest); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMultipartUploadResponse); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipartUploadLinkResponse); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipartUploadLinkRequest); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteMultipartUploadRequest); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteMultipartUploadResponse); i {
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
		file_api_services_v1_object_load_models_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedParts); i {
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
			RawDescriptor: file_api_services_v1_object_load_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_services_v1_object_load_models_proto_goTypes,
		DependencyIndexes: file_api_services_v1_object_load_models_proto_depIdxs,
		MessageInfos:      file_api_services_v1_object_load_models_proto_msgTypes,
	}.Build()
	File_api_services_v1_object_load_models_proto = out.File
	file_api_services_v1_object_load_models_proto_rawDesc = nil
	file_api_services_v1_object_load_models_proto_goTypes = nil
	file_api_services_v1_object_load_models_proto_depIdxs = nil
}
