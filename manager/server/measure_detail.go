package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMeasureDetail ...
func (s *Server) CreateMeasureDetail(ctx context.Context, req *pb.CreateMeasureDetailRequest) (data *pb.MeasureDetail, err error) {
	data, err = s.db.CreateMeasureDetail(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMeasureDetail ...
func (s *Server) GetMeasureDetail(ctx context.Context, req *pb.GetMeasureDetailRequest) (data *pb.MeasureDetail, err error) {
	data, err = s.db.GetMeasureDetail(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateMeasureDetail ...
func (s *Server) UpdateMeasureDetail(ctx context.Context, req *pb.UpdateMeasureDetailRequest) (data *pb.MeasureDetail, err error) {
	data, err = s.db.UpdateMeasureDetail(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteMeasureDetail ...
func (s *Server) DeleteMeasureDetail(ctx context.Context, req *pb.DeleteMeasureDetailRequest) (data *pb.MeasureDetail, err error) {
	data, err = s.db.DeleteMeasureDetail(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMeasureDetails ...
func (s *Server) GetMeasureDetails(ctx context.Context, req *pb.GetMeasureDetailsRequest) (reply *pb.GetMeasureDetailsReply, err error) {
	totalCount, users, err := s.db.GetMeasureDetails(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetMeasureDetailsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
