package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateWorkSubmit ...
func (s *Server) CreateWorkSubmit(ctx context.Context, req *pb.CreateWorkSubmitRequest) (data *pb.WorkSubmit, err error) {
	data, err = s.db.CreateWorkSubmit(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetWorkSubmit ...
func (s *Server) GetWorkSubmit(ctx context.Context, req *pb.GetWorkSubmitRequest) (data *pb.WorkSubmit, err error) {
	data, err = s.db.GetWorkSubmit(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateWorkSubmit ...
func (s *Server) UpdateWorkSubmit(ctx context.Context, req *pb.UpdateWorkSubmitRequest) (data *pb.WorkSubmit, err error) {
	data, err = s.db.UpdateWorkSubmit(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteWorkSubmit ...
func (s *Server) DeleteWorkSubmit(ctx context.Context, req *pb.DeleteWorkSubmitRequest) (data *pb.WorkSubmit, err error) {
	data, err = s.db.DeleteWorkSubmit(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetWorkSubmits ...
func (s *Server) GetWorkSubmits(ctx context.Context, req *pb.GetWorkSubmitsRequest) (reply *pb.GetWorkSubmitsReply, err error) {
	totalCount, users, err := s.db.GetWorkSubmits(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetWorkSubmitsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
