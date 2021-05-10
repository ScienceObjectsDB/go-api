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

// ProjectAPIClient is the client API for ProjectAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectAPIClient interface {
	//CreateProject creates a new projects
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*models.Project, error)
	//AddUserToProject Adds a new user to a given project
	AddUserToProject(ctx context.Context, in *AddUserToProjectRequest, opts ...grpc.CallOption) (*models.Project, error)
	CreateAPIToken(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.APIToken, error)
	//GetProjectDatasets Returns all datasets that belong to a certain project
	GetProjectDatasets(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*DatasetList, error)
	//GetUserProjects Returns all projects that a specified user has access to
	GetUserProjects(ctx context.Context, in *models.Empty, opts ...grpc.CallOption) (*ProjectList, error)
	GetProject(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Project, error)
	GetAPIToken(ctx context.Context, in *models.Empty, opts ...grpc.CallOption) (*APITokenList, error)
	//DeleteProject Deletes a specific project
	//Will also delete all associated resources (Datasets/Objects/etc...) both from objects storage and the database
	DeleteProject(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
	DeleteAPIToken(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error)
}

type projectAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAPIClient(cc grpc.ClientConnInterface) ProjectAPIClient {
	return &projectAPIClient{cc}
}

func (c *projectAPIClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*models.Project, error) {
	out := new(models.Project)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) AddUserToProject(ctx context.Context, in *AddUserToProjectRequest, opts ...grpc.CallOption) (*models.Project, error) {
	out := new(models.Project)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/AddUserToProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) CreateAPIToken(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.APIToken, error) {
	out := new(models.APIToken)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/CreateAPIToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetProjectDatasets(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*DatasetList, error) {
	out := new(DatasetList)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/GetProjectDatasets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetUserProjects(ctx context.Context, in *models.Empty, opts ...grpc.CallOption) (*ProjectList, error) {
	out := new(ProjectList)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/GetUserProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetProject(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Project, error) {
	out := new(models.Project)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetAPIToken(ctx context.Context, in *models.Empty, opts ...grpc.CallOption) (*APITokenList, error) {
	out := new(APITokenList)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/GetAPIToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) DeleteProject(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) DeleteAPIToken(ctx context.Context, in *models.ID, opts ...grpc.CallOption) (*models.Empty, error) {
	out := new(models.Empty)
	err := c.cc.Invoke(ctx, "/services.ProjectAPI/DeleteAPIToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAPIServer is the server API for ProjectAPI service.
// All implementations must embed UnimplementedProjectAPIServer
// for forward compatibility
type ProjectAPIServer interface {
	//CreateProject creates a new projects
	CreateProject(context.Context, *CreateProjectRequest) (*models.Project, error)
	//AddUserToProject Adds a new user to a given project
	AddUserToProject(context.Context, *AddUserToProjectRequest) (*models.Project, error)
	CreateAPIToken(context.Context, *models.ID) (*models.APIToken, error)
	//GetProjectDatasets Returns all datasets that belong to a certain project
	GetProjectDatasets(context.Context, *models.ID) (*DatasetList, error)
	//GetUserProjects Returns all projects that a specified user has access to
	GetUserProjects(context.Context, *models.Empty) (*ProjectList, error)
	GetProject(context.Context, *models.ID) (*models.Project, error)
	GetAPIToken(context.Context, *models.Empty) (*APITokenList, error)
	//DeleteProject Deletes a specific project
	//Will also delete all associated resources (Datasets/Objects/etc...) both from objects storage and the database
	DeleteProject(context.Context, *models.ID) (*models.Empty, error)
	DeleteAPIToken(context.Context, *models.ID) (*models.Empty, error)
	mustEmbedUnimplementedProjectAPIServer()
}

// UnimplementedProjectAPIServer must be embedded to have forward compatible implementations.
type UnimplementedProjectAPIServer struct {
}

func (UnimplementedProjectAPIServer) CreateProject(context.Context, *CreateProjectRequest) (*models.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectAPIServer) AddUserToProject(context.Context, *AddUserToProjectRequest) (*models.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToProject not implemented")
}
func (UnimplementedProjectAPIServer) CreateAPIToken(context.Context, *models.ID) (*models.APIToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIToken not implemented")
}
func (UnimplementedProjectAPIServer) GetProjectDatasets(context.Context, *models.ID) (*DatasetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectDatasets not implemented")
}
func (UnimplementedProjectAPIServer) GetUserProjects(context.Context, *models.Empty) (*ProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProjects not implemented")
}
func (UnimplementedProjectAPIServer) GetProject(context.Context, *models.ID) (*models.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectAPIServer) GetAPIToken(context.Context, *models.Empty) (*APITokenList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIToken not implemented")
}
func (UnimplementedProjectAPIServer) DeleteProject(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectAPIServer) DeleteAPIToken(context.Context, *models.ID) (*models.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPIToken not implemented")
}
func (UnimplementedProjectAPIServer) mustEmbedUnimplementedProjectAPIServer() {}

// UnsafeProjectAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectAPIServer will
// result in compilation errors.
type UnsafeProjectAPIServer interface {
	mustEmbedUnimplementedProjectAPIServer()
}

func RegisterProjectAPIServer(s grpc.ServiceRegistrar, srv ProjectAPIServer) {
	s.RegisterService(&ProjectAPI_ServiceDesc, srv)
}

func _ProjectAPI_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_AddUserToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).AddUserToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/AddUserToProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).AddUserToProject(ctx, req.(*AddUserToProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_CreateAPIToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).CreateAPIToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/CreateAPIToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).CreateAPIToken(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetProjectDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetProjectDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/GetProjectDatasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetProjectDatasets(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetUserProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetUserProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/GetUserProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetUserProjects(ctx, req.(*models.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetProject(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetAPIToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetAPIToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/GetAPIToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetAPIToken(ctx, req.(*models.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).DeleteProject(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_DeleteAPIToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).DeleteAPIToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProjectAPI/DeleteAPIToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).DeleteAPIToken(ctx, req.(*models.ID))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectAPI_ServiceDesc is the grpc.ServiceDesc for ProjectAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProjectAPI",
	HandlerType: (*ProjectAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectAPI_CreateProject_Handler,
		},
		{
			MethodName: "AddUserToProject",
			Handler:    _ProjectAPI_AddUserToProject_Handler,
		},
		{
			MethodName: "CreateAPIToken",
			Handler:    _ProjectAPI_CreateAPIToken_Handler,
		},
		{
			MethodName: "GetProjectDatasets",
			Handler:    _ProjectAPI_GetProjectDatasets_Handler,
		},
		{
			MethodName: "GetUserProjects",
			Handler:    _ProjectAPI_GetUserProjects_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectAPI_GetProject_Handler,
		},
		{
			MethodName: "GetAPIToken",
			Handler:    _ProjectAPI_GetAPIToken_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectAPI_DeleteProject_Handler,
		},
		{
			MethodName: "DeleteAPIToken",
			Handler:    _ProjectAPI_DeleteAPIToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/ProjectService.proto",
}
