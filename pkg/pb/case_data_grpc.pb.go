// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CaseDataManagerClient is the client API for CaseDataManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseDataManagerClient interface {
	CreateCaseData(ctx context.Context, in *CreateCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error)
	GetCaseData(ctx context.Context, in *GetCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error)
	UpdateCaseData(ctx context.Context, in *UpdateCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error)
	DeleteCaseData(ctx context.Context, in *DeleteCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error)
	GetCaseDatas(ctx context.Context, in *GetCaseDatasRequest, opts ...grpc.CallOption) (*GetCaseDatasReply, error)
}

type caseDataManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseDataManagerClient(cc grpc.ClientConnInterface) CaseDataManagerClient {
	return &caseDataManagerClient{cc}
}

func (c *caseDataManagerClient) CreateCaseData(ctx context.Context, in *CreateCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error) {
	out := new(CaseData)
	err := c.cc.Invoke(ctx, "/pb.CaseDataManager/CreateCaseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseDataManagerClient) GetCaseData(ctx context.Context, in *GetCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error) {
	out := new(CaseData)
	err := c.cc.Invoke(ctx, "/pb.CaseDataManager/GetCaseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseDataManagerClient) UpdateCaseData(ctx context.Context, in *UpdateCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error) {
	out := new(CaseData)
	err := c.cc.Invoke(ctx, "/pb.CaseDataManager/UpdateCaseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseDataManagerClient) DeleteCaseData(ctx context.Context, in *DeleteCaseDataRequest, opts ...grpc.CallOption) (*CaseData, error) {
	out := new(CaseData)
	err := c.cc.Invoke(ctx, "/pb.CaseDataManager/DeleteCaseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseDataManagerClient) GetCaseDatas(ctx context.Context, in *GetCaseDatasRequest, opts ...grpc.CallOption) (*GetCaseDatasReply, error) {
	out := new(GetCaseDatasReply)
	err := c.cc.Invoke(ctx, "/pb.CaseDataManager/GetCaseDatas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseDataManagerServer is the server API for CaseDataManager service.
// All implementations must embed UnimplementedCaseDataManagerServer
// for forward compatibility
type CaseDataManagerServer interface {
	CreateCaseData(context.Context, *CreateCaseDataRequest) (*CaseData, error)
	GetCaseData(context.Context, *GetCaseDataRequest) (*CaseData, error)
	UpdateCaseData(context.Context, *UpdateCaseDataRequest) (*CaseData, error)
	DeleteCaseData(context.Context, *DeleteCaseDataRequest) (*CaseData, error)
	GetCaseDatas(context.Context, *GetCaseDatasRequest) (*GetCaseDatasReply, error)
	mustEmbedUnimplementedCaseDataManagerServer()
}

// UnimplementedCaseDataManagerServer must be embedded to have forward compatible implementations.
type UnimplementedCaseDataManagerServer struct {
}

func (UnimplementedCaseDataManagerServer) CreateCaseData(context.Context, *CreateCaseDataRequest) (*CaseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCaseData not implemented")
}
func (UnimplementedCaseDataManagerServer) GetCaseData(context.Context, *GetCaseDataRequest) (*CaseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaseData not implemented")
}
func (UnimplementedCaseDataManagerServer) UpdateCaseData(context.Context, *UpdateCaseDataRequest) (*CaseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCaseData not implemented")
}
func (UnimplementedCaseDataManagerServer) DeleteCaseData(context.Context, *DeleteCaseDataRequest) (*CaseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCaseData not implemented")
}
func (UnimplementedCaseDataManagerServer) GetCaseDatas(context.Context, *GetCaseDatasRequest) (*GetCaseDatasReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaseDatas not implemented")
}
func (UnimplementedCaseDataManagerServer) mustEmbedUnimplementedCaseDataManagerServer() {}

// UnsafeCaseDataManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseDataManagerServer will
// result in compilation errors.
type UnsafeCaseDataManagerServer interface {
	mustEmbedUnimplementedCaseDataManagerServer()
}

func RegisterCaseDataManagerServer(s grpc.ServiceRegistrar, srv CaseDataManagerServer) {
	s.RegisterService(&CaseDataManager_ServiceDesc, srv)
}

func _CaseDataManager_CreateCaseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaseDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseDataManagerServer).CreateCaseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CaseDataManager/CreateCaseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseDataManagerServer).CreateCaseData(ctx, req.(*CreateCaseDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseDataManager_GetCaseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaseDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseDataManagerServer).GetCaseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CaseDataManager/GetCaseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseDataManagerServer).GetCaseData(ctx, req.(*GetCaseDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseDataManager_UpdateCaseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaseDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseDataManagerServer).UpdateCaseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CaseDataManager/UpdateCaseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseDataManagerServer).UpdateCaseData(ctx, req.(*UpdateCaseDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseDataManager_DeleteCaseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCaseDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseDataManagerServer).DeleteCaseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CaseDataManager/DeleteCaseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseDataManagerServer).DeleteCaseData(ctx, req.(*DeleteCaseDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseDataManager_GetCaseDatas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaseDatasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseDataManagerServer).GetCaseDatas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CaseDataManager/GetCaseDatas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseDataManagerServer).GetCaseDatas(ctx, req.(*GetCaseDatasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaseDataManager_ServiceDesc is the grpc.ServiceDesc for CaseDataManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaseDataManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CaseDataManager",
	HandlerType: (*CaseDataManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCaseData",
			Handler:    _CaseDataManager_CreateCaseData_Handler,
		},
		{
			MethodName: "GetCaseData",
			Handler:    _CaseDataManager_GetCaseData_Handler,
		},
		{
			MethodName: "UpdateCaseData",
			Handler:    _CaseDataManager_UpdateCaseData_Handler,
		},
		{
			MethodName: "DeleteCaseData",
			Handler:    _CaseDataManager_DeleteCaseData_Handler,
		},
		{
			MethodName: "GetCaseDatas",
			Handler:    _CaseDataManager_GetCaseDatas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "case_data.proto",
}
