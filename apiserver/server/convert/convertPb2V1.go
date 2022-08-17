package convert

import (
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func PersonalPb2V1(pbPersonal *pb.Personal) (v1Personal *v1.Personal) {
	if pbPersonal == nil {
		return
	}
	class := []*v1.Class{}
	if len(pbPersonal.Class) > 0 {
		for _, v := range pbPersonal.Class {
			class = append(class, ClassPb2V1(v))
		}
	}
	v1Personal = &v1.Personal{
		ID:     pbPersonal.Id,
		Class:  class,
		Type:   pbPersonal.Type,
		Name:   pbPersonal.Name,
		Userid: pbPersonal.Userid,
		Gender: pbPersonal.Gender,
		Birth:  pbPersonal.Birth,
	}
	return
}

func UserPb2V1(pbUser *pb.User) (v1User *v1.User) {
	if pbUser == nil {
		return
	}

	v1User = &v1.User{
		ID:    pbUser.Id,
		Name:  pbUser.Name,
		Ps:    pbUser.Ps,
		Role:  pbUser.Role,
		Phone: pbUser.Phone,
		Email: pbUser.Email,
	}
	return
}

func CaseDataPb2V1(pbCaseData *pb.CaseData) (v1CaseData *v1.CaseData) {
	if pbCaseData == nil {
		return
	}

	v1CaseData = &v1.CaseData{
		ID:     pbCaseData.Id,
		Userid: pbCaseData.Userid,
	}
	return
}

func ClassPb2V1(pbClass *pb.Class) (v1Class *v1.Class) {
	if pbClass == nil {
		return
	}

	v1Class = &v1.Class{
		ID:      pbClass.Id,
		Grade:   pbClass.Grade,
		Faculty: pbClass.Faculty,
		Number:  pbClass.Number,
		Subject: pbClass.Subject,
	}
	return
}

func MeasureDetailPb2V1(pbMeasureDetail *pb.MeasureDetail) (v1MeasureDetail *v1.MeasureDetail) {
	if pbMeasureDetail == nil {
		return
	}

	v1MeasureDetail = &v1.MeasureDetail{
		ID:     pbMeasureDetail.Id,
		Userid: pbMeasureDetail.Userid,
	}
	return
}

func MeasureResultPb2V1(pbMeasureResult *pb.MeasureResult) (v1MeasureResult *v1.MeasureResult) {
	if pbMeasureResult == nil {
		return
	}

	v1MeasureResult = &v1.MeasureResult{
		ID:     pbMeasureResult.Id,
		Userid: pbMeasureResult.Userid,
	}
	return
}

func MoveDataPb2V1(pbMoveData *pb.MoveData) (v1MoveData *v1.MoveData) {
	if pbMoveData == nil {
		return
	}

	v1MoveData = &v1.MoveData{
		ID:     pbMoveData.Id,
		Userid: pbMoveData.Userid,
	}
	return
}

func MovePrescriptionPb2V1(pbMovePrescription *pb.MovePrescription) (v1MovePrescription *v1.MovePrescription) {
	if pbMovePrescription == nil {
		return
	}

	v1MovePrescription = &v1.MovePrescription{
		ID:     pbMovePrescription.Id,
		Userid: pbMovePrescription.Userid,
	}
	return
}

func WorkPb2V1(pbWork *pb.Work) (v1Work *v1.Work) {
	if pbWork == nil {
		return
	}

	v1Work = &v1.Work{
		ID:      pbWork.Id,
		Creater: pbWork.Creater,
		Receive: pbWork.Receive,
		Content: pbWork.Content,
		Answer:  pbWork.Answer,
	}
	return
}

func WorkSubmitPb2V1(pbWorkSubmit *pb.WorkSubmit) (v1WorkSubmit *v1.WorkSubmit) {
	if pbWorkSubmit == nil {
		return
	}

	v1WorkSubmit = &v1.WorkSubmit{
		ID:        pbWorkSubmit.Id,
		Studentid: pbWorkSubmit.Studentid,
		Taskid:    pbWorkSubmit.Taskid,
		Content:   pbWorkSubmit.Content,
	}
	return
}
