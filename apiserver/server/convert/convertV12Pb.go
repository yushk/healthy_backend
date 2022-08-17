package convert

import (
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
	"github.com/yushk/healthy_backend/pkg/pb"
)

func PersonalV12Pb(v1Personal *v1.Personal) (pbPersonal *pb.Personal) {
	class := []*pb.Class{}
	if len(pbPersonal.Class) > 0 {
		for _, v := range v1Personal.Class {
			class = append(class, ClassV12Pb(v))
		}
	}
	pbPersonal = &pb.Personal{
		Id:     v1Personal.ID,
		Class:  class,
		Type:   v1Personal.Type,
		Name:   v1Personal.Name,
		Userid: v1Personal.Userid,
		Gender: v1Personal.Gender,
		Birth:  v1Personal.Birth,
	}
	return
}

func UserV12Pb(v1User *v1.User) (pbUser *pb.User) {
	pbUser = &pb.User{
		Id:    v1User.ID,
		Name:  v1User.Name,
		Ps:    v1User.Ps,
		Role:  v1User.Role,
		Phone: v1User.Phone,
		Email: v1User.Email,
	}
	return
}

func CaseDataV12Pb(v1CaseData *v1.CaseData) (pbCaseData *pb.CaseData) {
	pbCaseData = &pb.CaseData{
		Id:     v1CaseData.ID,
		Userid: v1CaseData.Userid,
	}
	return
}

func ClassV12Pb(v1Class *v1.Class) (pbClass *pb.Class) {
	pbClass = &pb.Class{
		Id:      v1Class.ID,
		Grade:   v1Class.Grade,
		Faculty: v1Class.Faculty,
		Number:  v1Class.Number,
		Subject: v1Class.Subject,
	}
	return
}

func MeasureDetailV12Pb(v1MeasureDetail *v1.MeasureDetail) (pbMeasureDetail *pb.MeasureDetail) {
	pbMeasureDetail = &pb.MeasureDetail{
		Id:     v1MeasureDetail.ID,
		Userid: v1MeasureDetail.Userid,
	}
	return
}

func MeasureResultV12Pb(v1MeasureResult *v1.MeasureResult) (pbMeasureResult *pb.MeasureResult) {
	pbMeasureResult = &pb.MeasureResult{
		Id:     v1MeasureResult.ID,
		Userid: v1MeasureResult.Userid,
	}
	return
}

func MoveDataV12Pb(v1MoveData *v1.MoveData) (pbMoveData *pb.MoveData) {
	pbMoveData = &pb.MoveData{
		Id:     v1MoveData.ID,
		Userid: v1MoveData.Userid,
	}
	return
}

func MovePrescriptionV12Pb(v1MovePrescription *v1.MovePrescription) (pbMovePrescription *pb.MovePrescription) {
	pbMovePrescription = &pb.MovePrescription{
		Id:     v1MovePrescription.ID,
		Userid: v1MovePrescription.Userid,
	}
	return
}

func WorkV12Pb(v1Work *v1.Work) (pbWork *pb.Work) {
	pbWork = &pb.Work{
		Id:      v1Work.ID,
		Creater: v1Work.Creater,
		Receive: v1Work.Receive,
		Content: v1Work.Content,
		Answer:  v1Work.Answer,
	}
	return
}

func WorkSubmitV12Pb(v1WorkSubmit *v1.WorkSubmit) (pbWorkSubmit *pb.WorkSubmit) {
	pbWorkSubmit = &pb.WorkSubmit{
		Id:        v1WorkSubmit.ID,
		Studentid: v1WorkSubmit.Studentid,
		Taskid:    v1WorkSubmit.Taskid,
		Content:   v1WorkSubmit.Content,
	}
	return
}
