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

// MeasureDetailManagerClient is the client API for MeasureDetailManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeasureDetailManagerClient interface {
	CreateMeasureDetail(ctx context.Context, in *CreateMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error)
	GetMeasureDetail(ctx context.Context, in *GetMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error)
	UpdateMeasureDetail(ctx context.Context, in *UpdateMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error)
	DeleteMeasureDetail(ctx context.Context, in *DeleteMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error)
	GetMeasureDetails(ctx context.Context, in *GetMeasureDetailsRequest, opts ...grpc.CallOption) (*GetMeasureDetailsReply, error)
}

type measureDetailManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasureDetailManagerClient(cc grpc.ClientConnInterface) MeasureDetailManagerClient {
	return &measureDetailManagerClient{cc}
}

func (c *measureDetailManagerClient) CreateMeasureDetail(ctx context.Context, in *CreateMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error) {
	out := new(MeasureDetail)
	err := c.cc.Invoke(ctx, "/pb.MeasureDetailManager/CreateMeasureDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureDetailManagerClient) GetMeasureDetail(ctx context.Context, in *GetMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error) {
	out := new(MeasureDetail)
	err := c.cc.Invoke(ctx, "/pb.MeasureDetailManager/GetMeasureDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureDetailManagerClient) UpdateMeasureDetail(ctx context.Context, in *UpdateMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error) {
	out := new(MeasureDetail)
	err := c.cc.Invoke(ctx, "/pb.MeasureDetailManager/UpdateMeasureDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureDetailManagerClient) DeleteMeasureDetail(ctx context.Context, in *DeleteMeasureDetailRequest, opts ...grpc.CallOption) (*MeasureDetail, error) {
	out := new(MeasureDetail)
	err := c.cc.Invoke(ctx, "/pb.MeasureDetailManager/DeleteMeasureDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureDetailManagerClient) GetMeasureDetails(ctx context.Context, in *GetMeasureDetailsRequest, opts ...grpc.CallOption) (*GetMeasureDetailsReply, error) {
	out := new(GetMeasureDetailsReply)
	err := c.cc.Invoke(ctx, "/pb.MeasureDetailManager/GetMeasureDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasureDetailManagerServer is the server API for MeasureDetailManager service.
// All implementations must embed UnimplementedMeasureDetailManagerServer
// for forward compatibility
type MeasureDetailManagerServer interface {
	CreateMeasureDetail(context.Context, *CreateMeasureDetailRequest) (*MeasureDetail, error)
	GetMeasureDetail(context.Context, *GetMeasureDetailRequest) (*MeasureDetail, error)
	UpdateMeasureDetail(context.Context, *UpdateMeasureDetailRequest) (*MeasureDetail, error)
	DeleteMeasureDetail(context.Context, *DeleteMeasureDetailRequest) (*MeasureDetail, error)
	GetMeasureDetails(context.Context, *GetMeasureDetailsRequest) (*GetMeasureDetailsReply, error)
	mustEmbedUnimplementedMeasureDetailManagerServer()
}

// UnimplementedMeasureDetailManagerServer must be embedded to have forward compatible implementations.
type UnimplementedMeasureDetailManagerServer struct {
}

func (UnimplementedMeasureDetailManagerServer) CreateMeasureDetail(context.Context, *CreateMeasureDetailRequest) (*MeasureDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeasureDetail not implemented")
}
func (UnimplementedMeasureDetailManagerServer) GetMeasureDetail(context.Context, *GetMeasureDetailRequest) (*MeasureDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasureDetail not implemented")
}
func (UnimplementedMeasureDetailManagerServer) UpdateMeasureDetail(context.Context, *UpdateMeasureDetailRequest) (*MeasureDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeasureDetail not implemented")
}
func (UnimplementedMeasureDetailManagerServer) DeleteMeasureDetail(context.Context, *DeleteMeasureDetailRequest) (*MeasureDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeasureDetail not implemented")
}
func (UnimplementedMeasureDetailManagerServer) GetMeasureDetails(context.Context, *GetMeasureDetailsRequest) (*GetMeasureDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasureDetails not implemented")
}
func (UnimplementedMeasureDetailManagerServer) mustEmbedUnimplementedMeasureDetailManagerServer() {}

// UnsafeMeasureDetailManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeasureDetailManagerServer will
// result in compilation errors.
type UnsafeMeasureDetailManagerServer interface {
	mustEmbedUnimplementedMeasureDetailManagerServer()
}

func RegisterMeasureDetailManagerServer(s grpc.ServiceRegistrar, srv MeasureDetailManagerServer) {
	s.RegisterService(&MeasureDetailManager_ServiceDesc, srv)
}

func _MeasureDetailManager_CreateMeasureDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeasureDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureDetailManagerServer).CreateMeasureDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MeasureDetailManager/CreateMeasureDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureDetailManagerServer).CreateMeasureDetail(ctx, req.(*CreateMeasureDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureDetailManager_GetMeasureDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasureDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureDetailManagerServer).GetMeasureDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MeasureDetailManager/GetMeasureDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureDetailManagerServer).GetMeasureDetail(ctx, req.(*GetMeasureDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureDetailManager_UpdateMeasureDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMeasureDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureDetailManagerServer).UpdateMeasureDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MeasureDetailManager/UpdateMeasureDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureDetailManagerServer).UpdateMeasureDetail(ctx, req.(*UpdateMeasureDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureDetailManager_DeleteMeasureDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMeasureDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureDetailManagerServer).DeleteMeasureDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MeasureDetailManager/DeleteMeasureDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureDetailManagerServer).DeleteMeasureDetail(ctx, req.(*DeleteMeasureDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureDetailManager_GetMeasureDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasureDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureDetailManagerServer).GetMeasureDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MeasureDetailManager/GetMeasureDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureDetailManagerServer).GetMeasureDetails(ctx, req.(*GetMeasureDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeasureDetailManager_ServiceDesc is the grpc.ServiceDesc for MeasureDetailManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeasureDetailManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MeasureDetailManager",
	HandlerType: (*MeasureDetailManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeasureDetail",
			Handler:    _MeasureDetailManager_CreateMeasureDetail_Handler,
		},
		{
			MethodName: "GetMeasureDetail",
			Handler:    _MeasureDetailManager_GetMeasureDetail_Handler,
		},
		{
			MethodName: "UpdateMeasureDetail",
			Handler:    _MeasureDetailManager_UpdateMeasureDetail_Handler,
		},
		{
			MethodName: "DeleteMeasureDetail",
			Handler:    _MeasureDetailManager_DeleteMeasureDetail_Handler,
		},
		{
			MethodName: "GetMeasureDetails",
			Handler:    _MeasureDetailManager_GetMeasureDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "measure_detail.proto",
}
