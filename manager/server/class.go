package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateClass ...
func (s *Server) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (data *pb.Class, err error) {
	data, err = s.db.CreateClass(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetClass ...
func (s *Server) GetClass(ctx context.Context, req *pb.GetClassRequest) (data *pb.Class, err error) {
	data, err = s.db.GetClass(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateClass ...
func (s *Server) UpdateClass(ctx context.Context, req *pb.UpdateClassRequest) (data *pb.Class, err error) {
	data, err = s.db.UpdateClass(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteClass ...
func (s *Server) DeleteClass(ctx context.Context, req *pb.DeleteClassRequest) (data *pb.Class, err error) {
	data, err = s.db.DeleteClass(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetClasses ...
func (s *Server) GetClasses(ctx context.Context, req *pb.GetClassesRequest) (reply *pb.GetClassesReply, err error) {
	totalCount, users, err := s.db.GetClasses(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetClassesReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
