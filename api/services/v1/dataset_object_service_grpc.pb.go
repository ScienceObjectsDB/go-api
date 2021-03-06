// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DatasetObjectsServiceClient is the client API for DatasetObjectsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetObjectsServiceClient interface {
	//CreateObjectGroup Creates a new object group
	CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error)
	//CreateObjectGroupVersion Creates a new object group version
	AddRevisionToObjectGroup(ctx context.Context, in *AddRevisionToObjectGroupRequest, opts ...grpc.CallOption) (*AddRevisionToObjectGroupResponse, error)
	//GetObjectGroup Returns the object group with the given ID
	GetObjectGroup(ctx context.Context, in *GetObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupResponse, error)
	//GetObjectGroupCurrentVersion Returns the head version in the history of a given object group
	GetCurrentObjectGroupRevision(ctx context.Context, in *GetCurrentObjectGroupRevisionRequest, opts ...grpc.CallOption) (*GetCurrentObjectGroupRevisionResponse, error)
	GetObjectGroupRevision(ctx context.Context, in *GetObjectGroupRevisionRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error)
	GetObjectGroupRevisions(ctx context.Context, in *GetObjectGroupRevisionsRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionsResponse, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(ctx context.Context, in *FinishObjectUploadRequest, opts ...grpc.CallOption) (*FinishObjectUploadResponse, error)
	DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error)
	DeleteObjectGroupRevision(ctx context.Context, in *DeleteObjectGroupRevisionRequest, opts ...grpc.CallOption) (*DeleteObjectGroupRevisionResponse, error)
}

type datasetObjectsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetObjectsServiceClient(cc grpc.ClientConnInterface) DatasetObjectsServiceClient {
	return &datasetObjectsServiceClient{cc}
}

func (c *datasetObjectsServiceClient) CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error) {
	out := new(CreateObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/CreateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) AddRevisionToObjectGroup(ctx context.Context, in *AddRevisionToObjectGroupRequest, opts ...grpc.CallOption) (*AddRevisionToObjectGroupResponse, error) {
	out := new(AddRevisionToObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/AddRevisionToObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroup(ctx context.Context, in *GetObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupResponse, error) {
	out := new(GetObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/GetObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetCurrentObjectGroupRevision(ctx context.Context, in *GetCurrentObjectGroupRevisionRequest, opts ...grpc.CallOption) (*GetCurrentObjectGroupRevisionResponse, error) {
	out := new(GetCurrentObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/GetCurrentObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroupRevision(ctx context.Context, in *GetObjectGroupRevisionRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error) {
	out := new(GetObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/GetObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroupRevisions(ctx context.Context, in *GetObjectGroupRevisionsRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionsResponse, error) {
	out := new(GetObjectGroupRevisionsResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/GetObjectGroupRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) FinishObjectUpload(ctx context.Context, in *FinishObjectUploadRequest, opts ...grpc.CallOption) (*FinishObjectUploadResponse, error) {
	out := new(FinishObjectUploadResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/FinishObjectUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error) {
	out := new(DeleteObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/DeleteObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) DeleteObjectGroupRevision(ctx context.Context, in *DeleteObjectGroupRevisionRequest, opts ...grpc.CallOption) (*DeleteObjectGroupRevisionResponse, error) {
	out := new(DeleteObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/DeleteObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetObjectsServiceServer is the server API for DatasetObjectsService service.
// All implementations should embed UnimplementedDatasetObjectsServiceServer
// for forward compatibility
type DatasetObjectsServiceServer interface {
	//CreateObjectGroup Creates a new object group
	CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error)
	//CreateObjectGroupVersion Creates a new object group version
	AddRevisionToObjectGroup(context.Context, *AddRevisionToObjectGroupRequest) (*AddRevisionToObjectGroupResponse, error)
	//GetObjectGroup Returns the object group with the given ID
	GetObjectGroup(context.Context, *GetObjectGroupRequest) (*GetObjectGroupResponse, error)
	//GetObjectGroupCurrentVersion Returns the head version in the history of a given object group
	GetCurrentObjectGroupRevision(context.Context, *GetCurrentObjectGroupRevisionRequest) (*GetCurrentObjectGroupRevisionResponse, error)
	GetObjectGroupRevision(context.Context, *GetObjectGroupRevisionRequest) (*GetObjectGroupRevisionResponse, error)
	GetObjectGroupRevisions(context.Context, *GetObjectGroupRevisionsRequest) (*GetObjectGroupRevisionsResponse, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(context.Context, *FinishObjectUploadRequest) (*FinishObjectUploadResponse, error)
	DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error)
	DeleteObjectGroupRevision(context.Context, *DeleteObjectGroupRevisionRequest) (*DeleteObjectGroupRevisionResponse, error)
}

// UnimplementedDatasetObjectsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDatasetObjectsServiceServer struct {
}

func (UnimplementedDatasetObjectsServiceServer) CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) AddRevisionToObjectGroup(context.Context, *AddRevisionToObjectGroupRequest) (*AddRevisionToObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRevisionToObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroup(context.Context, *GetObjectGroupRequest) (*GetObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetCurrentObjectGroupRevision(context.Context, *GetCurrentObjectGroupRevisionRequest) (*GetCurrentObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentObjectGroupRevision not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroupRevision(context.Context, *GetObjectGroupRevisionRequest) (*GetObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupRevision not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroupRevisions(context.Context, *GetObjectGroupRevisionsRequest) (*GetObjectGroupRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupRevisions not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) FinishObjectUpload(context.Context, *FinishObjectUploadRequest) (*FinishObjectUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishObjectUpload not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) DeleteObjectGroupRevision(context.Context, *DeleteObjectGroupRevisionRequest) (*DeleteObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroupRevision not implemented")
}

// UnsafeDatasetObjectsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetObjectsServiceServer will
// result in compilation errors.
type UnsafeDatasetObjectsServiceServer interface {
	mustEmbedUnimplementedDatasetObjectsServiceServer()
}

func RegisterDatasetObjectsServiceServer(s grpc.ServiceRegistrar, srv DatasetObjectsServiceServer) {
	s.RegisterService(&DatasetObjectsService_ServiceDesc, srv)
}

func _DatasetObjectsService_CreateObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/CreateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, req.(*CreateObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_AddRevisionToObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRevisionToObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).AddRevisionToObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/AddRevisionToObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).AddRevisionToObjectGroup(ctx, req.(*AddRevisionToObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/GetObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, req.(*GetObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetCurrentObjectGroupRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentObjectGroupRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetCurrentObjectGroupRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/GetCurrentObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetCurrentObjectGroupRevision(ctx, req.(*GetCurrentObjectGroupRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroupRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/GetObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevision(ctx, req.(*GetObjectGroupRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroupRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/GetObjectGroupRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevisions(ctx, req.(*GetObjectGroupRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_FinishObjectUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishObjectUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/FinishObjectUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, req.(*FinishObjectUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_DeleteObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/DeleteObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, req.(*DeleteObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_DeleteObjectGroupRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectGroupRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroupRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/DeleteObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroupRevision(ctx, req.(*DeleteObjectGroupRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetObjectsService_ServiceDesc is the grpc.ServiceDesc for DatasetObjectsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetObjectsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.v1.DatasetObjectsService",
	HandlerType: (*DatasetObjectsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectGroup",
			Handler:    _DatasetObjectsService_CreateObjectGroup_Handler,
		},
		{
			MethodName: "AddRevisionToObjectGroup",
			Handler:    _DatasetObjectsService_AddRevisionToObjectGroup_Handler,
		},
		{
			MethodName: "GetObjectGroup",
			Handler:    _DatasetObjectsService_GetObjectGroup_Handler,
		},
		{
			MethodName: "GetCurrentObjectGroupRevision",
			Handler:    _DatasetObjectsService_GetCurrentObjectGroupRevision_Handler,
		},
		{
			MethodName: "GetObjectGroupRevision",
			Handler:    _DatasetObjectsService_GetObjectGroupRevision_Handler,
		},
		{
			MethodName: "GetObjectGroupRevisions",
			Handler:    _DatasetObjectsService_GetObjectGroupRevisions_Handler,
		},
		{
			MethodName: "FinishObjectUpload",
			Handler:    _DatasetObjectsService_FinishObjectUpload_Handler,
		},
		{
			MethodName: "DeleteObjectGroup",
			Handler:    _DatasetObjectsService_DeleteObjectGroup_Handler,
		},
		{
			MethodName: "DeleteObjectGroupRevision",
			Handler:    _DatasetObjectsService_DeleteObjectGroupRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/v1/dataset_object_service.proto",
}
