package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/healthy_backend/apiserver/server/convert"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func CreateCaseData(params monitor.CreateCaseDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.CaseData().CreateCaseData(ctx, &pb.CreateCaseDataRequest{
		Data: convert.CaseDataV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateCaseDataOK().WithPayload(convert.CaseDataPb2V1(reply))
}

func DeleteCaseData(params monitor.DeleteCaseDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.CaseData().DeleteCaseData(ctx, &pb.DeleteCaseDataRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteCaseDataOK().WithPayload(convert.CaseDataPb2V1(reply))
}

func GetCaseData(params monitor.GetCaseDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.CaseData().GetCaseData(ctx, &pb.GetCaseDataRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetCaseDataOK().WithPayload(convert.CaseDataPb2V1(reply))
}

func UpdateCaseData(params monitor.UpdateCaseDataParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.CaseData().UpdateCaseData(ctx, &pb.UpdateCaseDataRequest{
		Id:   params.ID,
		Data: convert.CaseDataV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateCaseDataOK().WithPayload(convert.CaseDataPb2V1(reply))
}

// ListCaseData
func GetCaseDatas(params monitor.GetCaseDatasParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.CaseData().GetCaseDatas(ctx, &pb.GetCaseDatasRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.CaseData{}
	for _, v := range reply.Items {
		items = append(items, convert.CaseDataPb2V1(v))
	}

	payload := &v1.CaseDatas{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetCaseDatasOK().WithPayload(payload)
}
