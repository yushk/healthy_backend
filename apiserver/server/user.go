package server

import (
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/user"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateUser(params user.CreateUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().CreateUser(ctx, &pb.CreateUserRequest{
		User: convert.UserV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return user.NewCreateUserOK().WithPayload(convert.UserPb2V1(reply))
}

func DeleteUser(params user.DeleteUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewDeleteUserOK().WithPayload(convert.UserPb2V1(reply))
}

func GetUser(params user.GetUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().GetUser(ctx, &pb.GetUserRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewGetUserOK().WithPayload(convert.UserPb2V1(reply))
}

func Login(params user.LoginParams) middleware.Responder {
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
	return user.NewLoginOK().WithPayload(payload)
}

func Logout(params user.LogoutParams) middleware.Responder {
	return user.NewLogoutOK()
}

func UpdateUser(params user.UpdateUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:   params.ID,
		User: convert.UserV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return user.NewUpdateUserOK().WithPayload(convert.UserPb2V1(reply))
}

func ModifyUserPassword(params user.ModifyUserPasswordParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	_, err := client.User().ModifyUserPassword(ctx, &pb.ModifyUserPasswordRequest{
		Username: *params.Username,
		NewPwd:   params.NewPassword,
		OldPwd:   params.OldPassword,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewModifyUserPasswordOK()
}

// ListUser
func GetUsers(params user.GetUsersParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	limit := int64(0)
	if params.Limit != nil {
		limit = *params.Limit
	}
	skip := int64(0)
	if params.Skip != nil {
		skip = *params.Skip
	}
	query := `{}`
	if params.Query != nil {
		query = *params.Query
	}
	reply, err := client.User().GetUsers(ctx, &pb.GetUsersRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.User{}
	for _, v := range reply.Items {
		items = append(items, convert.UserPb2V1(v))
	}

	payload := &v1.Users{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return user.NewGetUsersOK().WithPayload(payload)
}

func GetUserInfo(params user.GetUserInfoParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	headerContent := params.HTTPRequest.Header.Get("authorization")
	token := strings.Split(headerContent, " ")[1]

	authenticate, err := client.User().Authenticate(ctx, &pb.AuthenticateRequest{
		Token: token,
	})
	if err != nil {
		return Error(err)
	}

	reply, err := client.User().GetUser(ctx, &pb.GetUserRequest{
		Id: authenticate.Id,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewGetUserOK().WithPayload(convert.UserPb2V1(reply))
}
