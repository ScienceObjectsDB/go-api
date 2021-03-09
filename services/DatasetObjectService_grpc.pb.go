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
const _ = grpc.SupportPackageIsVersion7

// DatasetObjectsServiceClient is the client API for DatasetObjectsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetObjectsServiceClient interface {
	//CreateObjectHeritage Creates a new object heritage
	CreateObjectHeritage(ctx context.Context, in *CreateObjectHeritageRequest, opts ...grpc.CallOption) (*models.ObjectHeritage, error)
	//CreateObjectGroup Creates a new object group
	CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*models.DatasetObjectEntry, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
}

type datasetObjectsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetObjectsServiceClient(cc grpc.ClientConnInterface) DatasetObjectsServiceClient {
	return &datasetObjectsServiceClient{cc}
}

func (c *datasetObjectsServiceClient) CreateObjectHeritage(ctx context.Context, in *CreateObjectHeritageRequest, opts ...grpc.CallOption) (*models.ObjectHeritage, error) {
	out := new(models.ObjectHeritage)
	err := c.cc.Invoke(ctx, "/DatasetObjectsService/CreateObjectHeritage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*models.DatasetObjectEntry, error) {
	out := new(models.DatasetObjectEntry)
	err := c.cc.Invoke(ctx, "/DatasetObjectsService/CreateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) FinishObjectUpload(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/DatasetObjectsService/FinishObjectUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetObjectsServiceServer is the server API for DatasetObjectsService service.
// All implementations must embed UnimplementedDatasetObjectsServiceServer
// for forward compatibility
type DatasetObjectsServiceServer interface {
	//CreateObjectHeritage Creates a new object heritage
	CreateObjectHeritage(context.Context, *CreateObjectHeritageRequest) (*models.ObjectHeritage, error)
	//CreateObjectGroup Creates a new object group
	CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*models.DatasetObjectEntry, error)
	//FinishObjectUpload Finishes the upload process for an object
	FinishObjectUpload(context.Context, *models.ID) (*models.Empty, error)
	mustEmbedUnimplementedDatasetObjectsServiceServer()
}

// UnimplementedDatasetObjectsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetObjectsServiceServer struct {
}

func (UnimplementedDatasetObjectsServiceServer) CreateObjectHeritage(context.Context, *CreateObjectHeritageRequest) (*models.ObjectHeritage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectHeritage not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*models.DatasetObjectEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) FinishObjectUpload(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishObjectUpload not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) mustEmbedUnimplementedDatasetObjectsServiceServer() {}

// UnsafeDatasetObjectsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetObjectsServiceServer will
// result in compilation errors.
type UnsafeDatasetObjectsServiceServer interface {
	mustEmbedUnimplementedDatasetObjectsServiceServer()
}

func RegisterDatasetObjectsServiceServer(s grpc.ServiceRegistrar, srv DatasetObjectsServiceServer) {
	s.RegisterService(&_DatasetObjectsService_serviceDesc, srv)
}

func _DatasetObjectsService_CreateObjectHeritage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectHeritageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).CreateObjectHeritage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetObjectsService/CreateObjectHeritage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectHeritage(ctx, req.(*CreateObjectHeritageRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/DatasetObjectsService/CreateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, req.(*CreateObjectGroupRequest))
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
		FullMethod: "/DatasetObjectsService/FinishObjectUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasetObjectsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DatasetObjectsService",
	HandlerType: (*DatasetObjectsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectHeritage",
			Handler:    _DatasetObjectsService_CreateObjectHeritage_Handler,
		},
		{
			MethodName: "CreateObjectGroup",
			Handler:    _DatasetObjectsService_CreateObjectGroup_Handler,
		},
		{
			MethodName: "FinishObjectUpload",
			Handler:    _DatasetObjectsService_FinishObjectUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/DatasetObjectService.proto",
}
