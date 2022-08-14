package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMovePrescription ...
func (s *Server) CreateMovePrescription(ctx context.Context, req *pb.CreateMovePrescriptionRequest) (data *pb.MovePrescription, err error) {
	data, err = s.db.CreateMovePrescription(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMovePrescription ...
func (s *Server) GetMovePrescription(ctx context.Context, req *pb.GetMovePrescriptionRequest) (data *pb.MovePrescription, err error) {
	data, err = s.db.GetMovePrescription(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateMovePrescription ...
func (s *Server) UpdateMovePrescription(ctx context.Context, req *pb.UpdateMovePrescriptionRequest) (data *pb.MovePrescription, err error) {
	data, err = s.db.UpdateMovePrescription(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteMovePrescription ...
func (s *Server) DeleteMovePrescription(ctx context.Context, req *pb.DeleteMovePrescriptionRequest) (data *pb.MovePrescription, err error) {
	data, err = s.db.DeleteMovePrescription(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetMovePrescriptions ...
func (s *Server) GetMovePrescriptions(ctx context.Context, req *pb.GetMovePrescriptionsRequest) (reply *pb.GetMovePrescriptionsReply, err error) {
	totalCount, users, err := s.db.GetMovePrescriptions(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetMovePrescriptionsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
