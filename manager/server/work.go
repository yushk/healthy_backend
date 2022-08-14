package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateWork ...
func (s *Server) CreateWork(ctx context.Context, req *pb.CreateWorkRequest) (data *pb.Work, err error) {
	data, err = s.db.CreateWork(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetWork ...
func (s *Server) GetWork(ctx context.Context, req *pb.GetWorkRequest) (data *pb.Work, err error) {
	data, err = s.db.GetWork(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateWork ...
func (s *Server) UpdateWork(ctx context.Context, req *pb.UpdateWorkRequest) (data *pb.Work, err error) {
	data, err = s.db.UpdateWork(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteWork ...
func (s *Server) DeleteWork(ctx context.Context, req *pb.DeleteWorkRequest) (data *pb.Work, err error) {
	data, err = s.db.DeleteWork(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetWorks ...
func (s *Server) GetWorks(ctx context.Context, req *pb.GetWorksRequest) (reply *pb.GetWorksReply, err error) {
	totalCount, users, err := s.db.GetWorks(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetWorksReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
