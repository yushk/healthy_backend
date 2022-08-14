package server

import (
	"context"

	"gitee.com/healthy/backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePersonal 创建用户
func (s *Server) CreatePersonal(ctx context.Context, req *pb.CreatePersonalRequest) (data *pb.Personal, err error) {
	data, err = s.db.CreatePersonal(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetPersonal 获取用户信息
func (s *Server) GetPersonal(ctx context.Context, req *pb.GetPersonalRequest) (data *pb.Personal, err error) {
	data, err = s.db.GetPersonal(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdatePersonal 更新用户信息
func (s *Server) UpdatePersonal(ctx context.Context, req *pb.UpdatePersonalRequest) (data *pb.Personal, err error) {
	data, err = s.db.UpdatePersonal(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeletePersonal 删除用户
func (s *Server) DeletePersonal(ctx context.Context, req *pb.DeletePersonalRequest) (data *pb.Personal, err error) {
	data, err = s.db.DeletePersonal(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetPersonals
func (s *Server) GetPersonals(ctx context.Context, req *pb.GetPersonalsRequest) (reply *pb.GetPersonalsReply, err error) {
	totalCount, users, err := s.db.GetPersonals(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetPersonalsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
