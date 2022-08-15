package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateWork(params monitor.CreateWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Work().CreateWork(ctx, &pb.CreateWorkRequest{
		Data: convert.WorkV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateWorkOK().WithPayload(convert.WorkPb2V1(reply))
}

func DeleteWork(params monitor.DeleteWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Work().DeleteWork(ctx, &pb.DeleteWorkRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteWorkOK().WithPayload(convert.WorkPb2V1(reply))
}

func GetWork(params monitor.GetWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Work().GetWork(ctx, &pb.GetWorkRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetWorkOK().WithPayload(convert.WorkPb2V1(reply))
}

func UpdateWork(params monitor.UpdateWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Work().UpdateWork(ctx, &pb.UpdateWorkRequest{
		Id:   params.ID,
		Data: convert.WorkV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateWorkOK().WithPayload(convert.WorkPb2V1(reply))
}

// ListWork
func GetWorks(params monitor.GetWorksParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.Work().GetWorks(ctx, &pb.GetWorksRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Work{}
	for _, v := range reply.Items {
		items = append(items, convert.WorkPb2V1(v))
	}

	payload := &v1.Works{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetWorksOK().WithPayload(payload)
}
