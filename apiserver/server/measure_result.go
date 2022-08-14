package server

import (
	"gitee.com/healthy/backend/apiserver/restapi/operations/monitor"
	"gitee.com/healthy/backend/apiserver/server/convert"
	v1 "gitee.com/healthy/backend/apiserver/v1"
	"gitee.com/healthy/backend/pkg/pb"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func CreateMeasureResult(params monitor.CreateMeasureResultParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureResult().CreateMeasureResult(ctx, &pb.CreateMeasureResultRequest{
		Data: convert.MeasureResultV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateMeasureResultOK().WithPayload(convert.MeasureResultPb2V1(reply))
}

func DeleteMeasureResult(params monitor.DeleteMeasureResultParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureResult().DeleteMeasureResult(ctx, &pb.DeleteMeasureResultRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteMeasureResultOK().WithPayload(convert.MeasureResultPb2V1(reply))
}

func GetMeasureResult(params monitor.GetMeasureResultParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureResult().GetMeasureResult(ctx, &pb.GetMeasureResultRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetMeasureResultOK().WithPayload(convert.MeasureResultPb2V1(reply))
}

func UpdateMeasureResult(params monitor.UpdateMeasureResultParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureResult().UpdateMeasureResult(ctx, &pb.UpdateMeasureResultRequest{
		Id:   params.ID,
		Data: convert.MeasureResultV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateMeasureResultOK().WithPayload(convert.MeasureResultPb2V1(reply))
}

// ListMeasureResult
func GetMeasureResults(params monitor.GetMeasureResultsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.MeasureResult().GetMeasureResults(ctx, &pb.GetMeasureResultsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.MeasureResult{}
	for _, v := range reply.Items {
		items = append(items, convert.MeasureResultPb2V1(v))
	}

	payload := &v1.MeasureResults{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetMeasureResultsOK().WithPayload(payload)
}
