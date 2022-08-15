package server

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yushk/healthy_backend/manager/server/helper"
	"github.com/yushk/healthy_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser 创建用户
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (user *pb.User, err error) {
	user, err = s.db.CreateUser(ctx, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// Login 用户登录
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (reply *pb.LoginReply, err error) {
	user, err := s.db.GetUserByName(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, req.Username+" login failed")
	}

	ret := s.db.Authenticate(ctx, req.Username, req.Password)
	if !ret {
		return nil, status.Errorf(codes.Unauthenticated, req.Username+" login failed")
	}

	now := time.Now()
	expiresAt := now.Add(helper.AccessExpiresIn)
	if user.Role == "attendant" {
		expiresAt = now.Add(helper.AccessExpiresIn * 48 * 365)
	}

	jwtClaims := helper.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "cloudstone.com",
			NotBefore: now.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
		UserClaims: helper.UserClaims{
			ID:   user.Id,
			Name: user.Name,
			Role: user.Role,
		},
	}
	token, err := helper.CreateToken(jwtClaims)
	if err != nil {
		return nil, status.Errorf(codes.Internal, req.Username+" login failed")
	}

	reply = &pb.LoginReply{
		Token:  token,
		UserId: user.Id,
	}
	return
}

// Authenticate 认证
func (s *Server) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (reply *pb.AuthenticateReply, err error) {
	claim, err := helper.ParseToken(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}

	reply = &pb.AuthenticateReply{
		Id:     claim.ID,
		Name:   claim.Name,
		Role:   claim.Role,
		Issuer: claim.Issuer,
	}
	return
}

// GetUser 获取用户信息
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (user *pb.User, err error) {
	user, err = s.db.GetUser(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateUser 更新用户信息
func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (user *pb.User, err error) {
	user, err = s.db.UpdateUser(ctx, req.Id, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteUser 删除用户
func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (user *pb.User, err error) {
	user, err = s.db.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.db.RemoveAuthorization(ctx, user.Name)
	return
}

// 用户修改密码
func (s *Server) ModifyUserPassword(ctx context.Context, req *pb.ModifyUserPasswordRequest) (reply *pb.ModifyUserPasswordReply, err error) {
	ret := s.db.Authenticate(ctx, req.Username, req.OldPwd)
	if !ret {
		return nil, status.Errorf(codes.Unauthenticated, req.Username+" oldPassword wrong change failed")
	}

	err = s.db.ChangeAuthorization(ctx, req.Username, req.NewPwd)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, req.Username+" change current password failed")
	}
	reply = &pb.ModifyUserPasswordReply{
		Username: req.Username,
	}
	return
}

// GetUsers
func (s *Server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (reply *pb.GetUsersReply, err error) {
	totalCount, users, err := s.db.GetUsers(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetUsersReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
