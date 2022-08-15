package server

import (
	"context"

	"github.com/yushk/healthy_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMoveData ...
func (s *Server) CreateMoveData(ctx context.Context, req *pb.CreateMoveDataRequest) (data *pb.MoveData, err error) {
	data, err = s.db.CreateMoveData(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMoveData ...
func (s *Server) GetMoveData(ctx context.Context, req *pb.GetMoveDataRequest) (data *pb.MoveData, err error) {
	data, err = s.db.GetMoveData(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateMoveData ...
func (s *Server) UpdateMoveData(ctx context.Context, req *pb.UpdateMoveDataRequest) (data *pb.MoveData, err error) {
	data, err = s.db.UpdateMoveData(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteMoveData ...
func (s *Server) DeleteMoveData(ctx context.Context, req *pb.DeleteMoveDataRequest) (data *pb.MoveData, err error) {
	data, err = s.db.DeleteMoveData(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMoveDatas ...
func (s *Server) GetMoveDatas(ctx context.Context, req *pb.GetMoveDatasRequest) (reply *pb.GetMoveDatasReply, err error) {
	totalCount, users, err := s.db.GetMoveDatas(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetMoveDatasReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
