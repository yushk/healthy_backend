package server

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/healthy/backend/apiserver/restapi/operations/oauth"
	v1 "gitee.com/healthy/backend/apiserver/v1"
	"gitee.com/healthy/backend/pkg/pb"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// RoleAuth 权限认证
func RoleAuth(token string, scopes []string) (*v1.Principal, error) {

	reply, err := client.User().Authenticate(context.Background(), &pb.AuthenticateRequest{
		Token: token,
	})
	if err != nil {
		logrus.Errorln(err)
		return nil, errors.New(http.StatusUnauthorized, err.Error())
	}

	ok := false
	currentRole := reply.Role
	for _, role := range scopes {
		if currentRole == role {
			ok = true
		}
	}
	if !ok {
		return nil, fmt.Errorf("current user no permission to operate")
	}

	// AutoRefreshToken(token)
	return &v1.Principal{
		ID:   reply.Id,
		Name: reply.Name,
		Role: reply.Role,
	}, nil
}

// Token POST /oauth/token
func Token(params oauth.TokenParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()

	reply, err := client.User().Login(ctx, &pb.LoginRequest{
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		return Error(err)
	}

	accessToken := reply.Token["access_token"]
	tokenType := reply.Token["token_type"]
	expiresIn := reply.Token["expires_in"]
	expiresAt := reply.Token["expires_at"]
	payload := &v1.Token{
		AccessToken: accessToken,
		TokenType:   &tokenType,
		ExpiresIn:   expiresIn,
		ExpiresAt:   expiresAt,
	}
	return oauth.NewTokenOK().WithPayload(payload)
}
