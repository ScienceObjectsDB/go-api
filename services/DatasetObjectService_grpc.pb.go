// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	models "github.com/ScienceObjectsDB/go-api/models"
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
	CreateObjectGroup(ctx context.Context, in *CreateObjectGroupWithRevisionRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error)
	//CreateObjectGroupVersion Creates a new object group version
	AddRevisionToObjectGroup(ctx context.Context, in *AddRevisionToObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error)
	//GetObjectGroup Returns the object group with the given ID
	GetObjectGroup(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error)
	//GetObjectGroupCurrentVersion Returns the head version in the history of a given object group
	GetCurrentObjectGroupRevision(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error)
	GetObjectGroupRevision(ctx context.Context, in *GetObjectGroupRevisionRequest, opts ...grpc.CallOption) (*models.ObjectGroupRevision, error)
	GetObjectGroupRevisions(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*ObjectGroupRevisions, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
	DeleteObjectGroup(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
	DeleteObjectGroupRevision(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
}

type datasetObjectsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetObjectsServiceClient(cc grpc.ClientConnInterface) DatasetObjectsServiceClient {
	return &datasetObjectsServiceClient{cc}
}

func (c *datasetObjectsServiceClient) CreateObjectGroup(ctx context.Context, in *CreateObjectGroupWithRevisionRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error) {
	out := new(GetObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/CreateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) AddRevisionToObjectGroup(ctx context.Context, in *AddRevisionToObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error) {
	out := new(GetObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/AddRevisionToObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroup(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error) {
	out := new(GetObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/GetObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetCurrentObjectGroupRevision(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*GetObjectGroupRevisionResponse, error) {
	out := new(GetObjectGroupRevisionResponse)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/GetCurrentObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroupRevision(ctx context.Context, in *GetObjectGroupRevisionRequest, opts ...grpc.CallOption) (*models.ObjectGroupRevision, error) {
	out := new(models.ObjectGroupRevision)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/GetObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroupRevisions(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*ObjectGroupRevisions, error) {
	out := new(ObjectGroupRevisions)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/GetObjectGroupRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) FinishObjectUpload(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/FinishObjectUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) DeleteObjectGroup(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/DeleteObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) DeleteObjectGroupRevision(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/services.DatasetObjectsService/DeleteObjectGroupRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetObjectsServiceServer is the server API for DatasetObjectsService service.
// All implementations must embed UnimplementedDatasetObjectsServiceServer
// for forward compatibility
type DatasetObjectsServiceServer interface {
	//CreateObjectGroup Creates a new object group
	CreateObjectGroup(context.Context, *CreateObjectGroupWithRevisionRequest) (*GetObjectGroupRevisionResponse, error)
	//CreateObjectGroupVersion Creates a new object group version
	AddRevisionToObjectGroup(context.Context, *AddRevisionToObjectGroupRequest) (*GetObjectGroupRevisionResponse, error)
	//GetObjectGroup Returns the object group with the given ID
	GetObjectGroup(context.Context, *models.ID) (*GetObjectGroupRevisionResponse, error)
	//GetObjectGroupCurrentVersion Returns the head version in the history of a given object group
	GetCurrentObjectGroupRevision(context.Context, *models.ID) (*GetObjectGroupRevisionResponse, error)
	GetObjectGroupRevision(context.Context, *GetObjectGroupRevisionRequest) (*models.ObjectGroupRevision, error)
	GetObjectGroupRevisions(context.Context, *models.ID) (*ObjectGroupRevisions, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(context.Context, *models.ID) (*models.Empty, error)
	DeleteObjectGroup(context.Context, *models.ID) (*models.Empty, error)
	DeleteObjectGroupRevision(context.Context, *models.ID) (*models.Empty, error)
	mustEmbedUnimplementedDatasetObjectsServiceServer()
}

// UnimplementedDatasetObjectsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetObjectsServiceServer struct {
}

func (UnimplementedDatasetObjectsServiceServer) CreateObjectGroup(context.Context, *CreateObjectGroupWithRevisionRequest) (*GetObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) AddRevisionToObjectGroup(context.Context, *AddRevisionToObjectGroupRequest) (*GetObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRevisionToObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroup(context.Context, *models.ID) (*GetObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetCurrentObjectGroupRevision(context.Context, *models.ID) (*GetObjectGroupRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentObjectGroupRevision not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroupRevision(context.Context, *GetObjectGroupRevisionRequest) (*models.ObjectGroupRevision, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupRevision not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroupRevisions(context.Context, *models.ID) (*ObjectGroupRevisions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupRevisions not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) FinishObjectUpload(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishObjectUpload not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) DeleteObjectGroup(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) DeleteObjectGroupRevision(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroupRevision not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) mustEmbedUnimplementedDatasetObjectsServiceServer() {}

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
	in := new(CreateObjectGroupWithRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/CreateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, req.(*CreateObjectGroupWithRevisionRequest))
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
		FullMethod: "/services.DatasetObjectsService/AddRevisionToObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).AddRevisionToObjectGroup(ctx, req.(*AddRevisionToObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/GetObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetCurrentObjectGroupRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetCurrentObjectGroupRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/GetCurrentObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetCurrentObjectGroupRevision(ctx, req.(*models.ID))
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
		FullMethod: "/services.DatasetObjectsService/GetObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevision(ctx, req.(*GetObjectGroupRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroupRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/GetObjectGroupRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroupRevisions(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_FinishObjectUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/FinishObjectUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_DeleteObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/DeleteObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_DeleteObjectGroupRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroupRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatasetObjectsService/DeleteObjectGroupRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroupRevision(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetObjectsService_ServiceDesc is the grpc.ServiceDesc for DatasetObjectsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetObjectsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.DatasetObjectsService",
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
	Metadata: "api/services/DatasetObjectService.proto",
}
