package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateMoveData(params monitor.CreateMoveDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MoveData().CreateMoveData(ctx, &pb.CreateMoveDataRequest{
		Data: convert.MoveDataV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateMoveDataOK().WithPayload(convert.MoveDataPb2V1(reply))
}

func DeleteMoveData(params monitor.DeleteMoveDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MoveData().DeleteMoveData(ctx, &pb.DeleteMoveDataRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteMoveDataOK().WithPayload(convert.MoveDataPb2V1(reply))
}

func GetMoveData(params monitor.GetMoveDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MoveData().GetMoveData(ctx, &pb.GetMoveDataRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetMoveDataOK().WithPayload(convert.MoveDataPb2V1(reply))
}

func UpdateMoveData(params monitor.UpdateMoveDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.MoveData().UpdateMoveData(ctx, &pb.UpdateMoveDataRequest{
		Id:   params.ID,
		Data: convert.MoveDataV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateMoveDataOK().WithPayload(convert.MoveDataPb2V1(reply))
}

// ListMoveData
func GetMoveDatas(params monitor.GetMoveDatasParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.MoveData().GetMoveDatas(ctx, &pb.GetMoveDatasRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.MoveData{}
	for _, v := range reply.Items {
		items = append(items, convert.MoveDataPb2V1(v))
	}

	payload := &v1.MoveDatas{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetMoveDatasOK().WithPayload(payload)
}
