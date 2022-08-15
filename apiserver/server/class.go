package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateClass(params monitor.CreateClassParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Class().CreateClass(ctx, &pb.CreateClassRequest{
		Data: convert.ClassV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateClassOK().WithPayload(convert.ClassPb2V1(reply))
}

func DeleteClass(params monitor.DeleteClassParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Class().DeleteClass(ctx, &pb.DeleteClassRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteClassOK().WithPayload(convert.ClassPb2V1(reply))
}

func GetClass(params monitor.GetClassParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Class().GetClass(ctx, &pb.GetClassRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetClassOK().WithPayload(convert.ClassPb2V1(reply))
}

func UpdateClass(params monitor.UpdateClassParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Class().UpdateClass(ctx, &pb.UpdateClassRequest{
		Id:   params.ID,
		Data: convert.ClassV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateClassOK().WithPayload(convert.ClassPb2V1(reply))
}

// ListClass
func GetClasses(params monitor.GetClassesParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.Class().GetClasses(ctx, &pb.GetClassesRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Class{}
	for _, v := range reply.Items {
		items = append(items, convert.ClassPb2V1(v))
	}

	payload := &v1.Classes{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetClassesOK().WithPayload(payload)
}
