package server

import (
	"gitee.com/healthy/backend/apiserver/restapi/operations/monitor"
	"gitee.com/healthy/backend/apiserver/server/convert"
	v1 "gitee.com/healthy/backend/apiserver/v1"
	"gitee.com/healthy/backend/pkg/pb"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func CreateMovePrescription(params monitor.CreateMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MovePrescription().CreateMovePrescription(ctx, &pb.CreateMovePrescriptionRequest{
		Data: convert.MovePrescriptionV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateMovePrescriptionOK().WithPayload(convert.MovePrescriptionPb2V1(reply))
}

func DeleteMovePrescription(params monitor.DeleteMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MovePrescription().DeleteMovePrescription(ctx, &pb.DeleteMovePrescriptionRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteMovePrescriptionOK().WithPayload(convert.MovePrescriptionPb2V1(reply))
}

func GetMovePrescription(params monitor.GetMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MovePrescription().GetMovePrescription(ctx, &pb.GetMovePrescriptionRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetMovePrescriptionOK().WithPayload(convert.MovePrescriptionPb2V1(reply))
}

func UpdateMovePrescription(params monitor.UpdateMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MovePrescription().UpdateMovePrescription(ctx, &pb.UpdateMovePrescriptionRequest{
		Id:   params.ID,
		Data: convert.MovePrescriptionV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateMovePrescriptionOK().WithPayload(convert.MovePrescriptionPb2V1(reply))
}

// ListMovePrescription
func GetMovePrescriptions(params monitor.GetMovePrescriptionsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.MovePrescription().GetMovePrescriptions(ctx, &pb.GetMovePrescriptionsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.MovePrescription{}
	for _, v := range reply.Items {
		items = append(items, convert.MovePrescriptionPb2V1(v))
	}

	payload := &v1.MovePrescriptions{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetMovePrescriptionsOK().WithPayload(payload)
}
