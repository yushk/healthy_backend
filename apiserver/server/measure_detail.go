package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateMeasureDetail(params monitor.CreateMeasureDetailParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureDetail().CreateMeasureDetail(ctx, &pb.CreateMeasureDetailRequest{
		Data: convert.MeasureDetailV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateMeasureDetailOK().WithPayload(convert.MeasureDetailPb2V1(reply))
}

func DeleteMeasureDetail(params monitor.DeleteMeasureDetailParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureDetail().DeleteMeasureDetail(ctx, &pb.DeleteMeasureDetailRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteMeasureDetailOK().WithPayload(convert.MeasureDetailPb2V1(reply))
}

func GetMeasureDetail(params monitor.GetMeasureDetailParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureDetail().GetMeasureDetail(ctx, &pb.GetMeasureDetailRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetMeasureDetailOK().WithPayload(convert.MeasureDetailPb2V1(reply))
}

func UpdateMeasureDetail(params monitor.UpdateMeasureDetailParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MeasureDetail().UpdateMeasureDetail(ctx, &pb.UpdateMeasureDetailRequest{
		Id:   params.ID,
		Data: convert.MeasureDetailV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateMeasureDetailOK().WithPayload(convert.MeasureDetailPb2V1(reply))
}

// ListMeasureDetail
func GetMeasureDetails(params monitor.GetMeasureDetailsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.MeasureDetail().GetMeasureDetails(ctx, &pb.GetMeasureDetailsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.MeasureDetail{}
	for _, v := range reply.Items {
		items = append(items, convert.MeasureDetailPb2V1(v))
	}

	payload := &v1.MeasureDetails{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetMeasureDetailsOK().WithPayload(payload)
}
