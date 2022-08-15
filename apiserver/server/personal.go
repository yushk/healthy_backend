package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/user"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreatePersonal(params user.CreatePersonalParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Personal().CreatePersonal(ctx, &pb.CreatePersonalRequest{
		Data: convert.PersonalV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return user.NewCreatePersonalOK().WithPayload(convert.PersonalPb2V1(reply))
}

func DeletePersonal(params user.DeletePersonalParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Personal().DeletePersonal(ctx, &pb.DeletePersonalRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewDeletePersonalOK().WithPayload(convert.PersonalPb2V1(reply))
}

func GetPersonal(params user.GetPersonalParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Personal().GetPersonal(ctx, &pb.GetPersonalRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewGetPersonalOK().WithPayload(convert.PersonalPb2V1(reply))
}

func UpdatePersonal(params user.UpdatePersonalParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Personal().UpdatePersonal(ctx, &pb.UpdatePersonalRequest{
		Id:   params.ID,
		Data: convert.PersonalV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return user.NewUpdatePersonalOK().WithPayload(convert.PersonalPb2V1(reply))
}

// ListPersonal
func GetPersonals(params user.GetPersonalsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.Personal().GetPersonals(ctx, &pb.GetPersonalsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Personal{}
	for _, v := range reply.Items {
		items = append(items, convert.PersonalPb2V1(v))
	}

	payload := &v1.Personals{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return user.NewGetPersonalsOK().WithPayload(payload)
}
