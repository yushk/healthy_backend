package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCaseData ...
func (s *Server) CreateCaseData(ctx context.Context, req *pb.CreateCaseDataRequest) (data *pb.CaseData, err error) {
	data, err = s.db.CreateCaseData(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetCaseData ...
func (s *Server) GetCaseData(ctx context.Context, req *pb.GetCaseDataRequest) (data *pb.CaseData, err error) {
	data, err = s.db.GetCaseData(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateCaseData ...
func (s *Server) UpdateCaseData(ctx context.Context, req *pb.UpdateCaseDataRequest) (data *pb.CaseData, err error) {
	data, err = s.db.UpdateCaseData(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteCaseData ...
func (s *Server) DeleteCaseData(ctx context.Context, req *pb.DeleteCaseDataRequest) (data *pb.CaseData, err error) {
	data, err = s.db.DeleteCaseData(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetCaseDatas ...
func (s *Server) GetCaseDatas(ctx context.Context, req *pb.GetCaseDatasRequest) (reply *pb.GetCaseDatasReply, err error) {
	totalCount, users, err := s.db.GetCaseDatas(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetCaseDatasReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
