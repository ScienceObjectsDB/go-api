// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/services/DatasetObjectService.proto

package services

import (
	models "github.com/ScienceObjectsDB/go-api/models"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_services_DatasetObjectService_proto protoreflect.FileDescriptor

var file_api_services_DatasetObjectService_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc5,
	0x08, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x9f, 0x01, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x29, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x6a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x44, 0x1a,
	0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x44, 0x1a, 0x28, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x8b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6f,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x12,
	0x51, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49,
	0x44, 0x1a, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x3a,
	0x01, 0x2a, 0x12, 0x50, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x76, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x16, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50,
	0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_services_DatasetObjectService_proto_goTypes = []interface{}{
	(*CreateObjectGroupWithRevisionRequest)(nil), // 0: services.CreateObjectGroupWithRevisionRequest
	(*AddRevisionToObjectGroupRequest)(nil),      // 1: services.AddRevisionToObjectGroupRequest
	(*models.ID)(nil),                            // 2: models.ID
	(*GetObjectGroupRevisionRequest)(nil),        // 3: services.GetObjectGroupRevisionRequest
	(*GetObjectGroupRevisionResponse)(nil),       // 4: services.GetObjectGroupRevisionResponse
	(*models.ObjectGroupRevision)(nil),           // 5: models.ObjectGroupRevision
	(*ObjectGroupRevisions)(nil),                 // 6: services.ObjectGroupRevisions
	(*models.Empty)(nil),                         // 7: models.Empty
}
var file_api_services_DatasetObjectService_proto_depIdxs = []int32{
	0, // 0: services.DatasetObjectsService.CreateObjectGroup:input_type -> services.CreateObjectGroupWithRevisionRequest
	1, // 1: services.DatasetObjectsService.AddRevisionToObjectGroup:input_type -> services.AddRevisionToObjectGroupRequest
	2, // 2: services.DatasetObjectsService.GetObjectGroup:input_type -> models.ID
	2, // 3: services.DatasetObjectsService.GetCurrentObjectGroupRevision:input_type -> models.ID
	3, // 4: services.DatasetObjectsService.GetObjectGroupRevision:input_type -> services.GetObjectGroupRevisionRequest
	2, // 5: services.DatasetObjectsService.GetObjectGroupRevisions:input_type -> models.ID
	2, // 6: services.DatasetObjectsService.FinishObjectUpload:input_type -> models.ID
	2, // 7: services.DatasetObjectsService.DeleteObjectGroup:input_type -> models.ID
	2, // 8: services.DatasetObjectsService.DeleteObjectGroupRevision:input_type -> models.ID
	4, // 9: services.DatasetObjectsService.CreateObjectGroup:output_type -> services.GetObjectGroupRevisionResponse
	4, // 10: services.DatasetObjectsService.AddRevisionToObjectGroup:output_type -> services.GetObjectGroupRevisionResponse
	4, // 11: services.DatasetObjectsService.GetObjectGroup:output_type -> services.GetObjectGroupRevisionResponse
	4, // 12: services.DatasetObjectsService.GetCurrentObjectGroupRevision:output_type -> services.GetObjectGroupRevisionResponse
	5, // 13: services.DatasetObjectsService.GetObjectGroupRevision:output_type -> models.ObjectGroupRevision
	6, // 14: services.DatasetObjectsService.GetObjectGroupRevisions:output_type -> services.ObjectGroupRevisions
	7, // 15: services.DatasetObjectsService.FinishObjectUpload:output_type -> models.Empty
	7, // 16: services.DatasetObjectsService.DeleteObjectGroup:output_type -> models.Empty
	7, // 17: services.DatasetObjectsService.DeleteObjectGroupRevision:output_type -> models.Empty
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_services_DatasetObjectService_proto_init() }
func file_api_services_DatasetObjectService_proto_init() {
	if File_api_services_DatasetObjectService_proto != nil {
		return
	}
	file_api_services_DatasetObjectServiceModels_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_DatasetObjectService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_DatasetObjectService_proto_goTypes,
		DependencyIndexes: file_api_services_DatasetObjectService_proto_depIdxs,
	}.Build()
	File_api_services_DatasetObjectService_proto = out.File
	file_api_services_DatasetObjectService_proto_rawDesc = nil
	file_api_services_DatasetObjectService_proto_goTypes = nil
	file_api_services_DatasetObjectService_proto_depIdxs = nil
}
