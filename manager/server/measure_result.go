package server

import (
	"context"

	"github.com/yushk/healthy_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMeasureResult ...
func (s *Server) CreateMeasureResult(ctx context.Context, req *pb.CreateMeasureResultRequest) (data *pb.MeasureResult, err error) {
	data, err = s.db.CreateMeasureResult(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMeasureResult ...
func (s *Server) GetMeasureResult(ctx context.Context, req *pb.GetMeasureResultRequest) (data *pb.MeasureResult, err error) {
	data, err = s.db.GetMeasureResult(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateMeasureResult ...
func (s *Server) UpdateMeasureResult(ctx context.Context, req *pb.UpdateMeasureResultRequest) (data *pb.MeasureResult, err error) {
	data, err = s.db.UpdateMeasureResult(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteMeasureResult ...
func (s *Server) DeleteMeasureResult(ctx context.Context, req *pb.DeleteMeasureResultRequest) (data *pb.MeasureResult, err error) {
	data, err = s.db.DeleteMeasureResult(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMeasureResults ...
func (s *Server) GetMeasureResults(ctx context.Context, req *pb.GetMeasureResultsRequest) (reply *pb.GetMeasureResultsReply, err error) {
	totalCount, users, err := s.db.GetMeasureResults(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetMeasureResultsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
