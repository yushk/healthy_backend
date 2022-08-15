package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateWorkSubmit(params monitor.CreateWorkSubmitParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.WorkSubmit().CreateWorkSubmit(ctx, &pb.CreateWorkSubmitRequest{
		Data: convert.WorkSubmitV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateWorkSubmitOK().WithPayload(convert.WorkSubmitPb2V1(reply))
}

func DeleteWorkSubmit(params monitor.DeleteWorkSubmitParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.WorkSubmit().DeleteWorkSubmit(ctx, &pb.DeleteWorkSubmitRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteWorkSubmitOK().WithPayload(convert.WorkSubmitPb2V1(reply))
}

func GetWorkSubmit(params monitor.GetWorkSubmitParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.WorkSubmit().GetWorkSubmit(ctx, &pb.GetWorkSubmitRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetWorkSubmitOK().WithPayload(convert.WorkSubmitPb2V1(reply))
}

func UpdateWorkSubmit(params monitor.UpdateWorkSubmitParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.WorkSubmit().UpdateWorkSubmit(ctx, &pb.UpdateWorkSubmitRequest{
		Id:   params.ID,
		Data: convert.WorkSubmitV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateWorkSubmitOK().WithPayload(convert.WorkSubmitPb2V1(reply))
}

// ListWorkSubmit
func GetWorkSubmits(params monitor.GetWorkSubmitsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.WorkSubmit().GetWorkSubmits(ctx, &pb.GetWorkSubmitsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.WorkSubmit{}
	for _, v := range reply.Items {
		items = append(items, convert.WorkSubmitPb2V1(v))
	}

	payload := &v1.WorkSubmits{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetWorkSubmitsOK().WithPayload(payload)
}
