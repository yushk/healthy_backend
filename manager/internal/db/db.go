package db

import (
	"context"
	"time"

	"github.com/yushk/healthy_backend/manager/internal/db/mongodb"
	"github.com/yushk/healthy_backend/pkg/pb"
)

// Type 数据库类型
type Type string

// 数据库类型枚举
const (
	MongoDB Type = "mongodb"
)

// 数据库连接默认的超时时间
const (
	DefaultTimeout = 10 * time.Second
)

// Database 数据库接口
type Database interface {
	// 用户管理接口
	CreateUser(ctx context.Context, user *pb.User) (*pb.User, error)
	GetUser(ctx context.Context, id string) (user *pb.User, err error)
	DeleteUser(ctx context.Context, id string) (user *pb.User, err error)
	Authenticate(ctx context.Context, username, password string) (ret bool)
	UpdateUser(ctx context.Context, id string, user *pb.User) (*pb.User, error)
	GetUsers(ctx context.Context, limit, skip int64, query string) (int64, []*pb.User, error)
	GetUserByName(ctx context.Context, name string) (user *pb.User, err error)
	ChangeAuthorization(ctx context.Context, name, password string) (err error)
	RemoveAuthorization(ctx context.Context, username string) (err error)
	// 个人基础信息管理接口
	CreatePersonal(ctx context.Context, data *pb.Personal) (*pb.Personal, error)
	GetPersonal(ctx context.Context, id string) (data *pb.Personal, err error)
	DeletePersonal(ctx context.Context, id string) (data *pb.Personal, err error)
	UpdatePersonal(ctx context.Context, id string, data *pb.Personal) (*pb.Personal, error)
	GetPersonals(ctx context.Context, limit, skip int64, query string) (int64, []*pb.Personal, error)
	// 病例数据管理接口
	CreateCaseData(ctx context.Context, data *pb.CaseData) (*pb.CaseData, error)
	GetCaseData(ctx context.Context, id string) (data *pb.CaseData, err error)
	DeleteCaseData(ctx context.Context, id string) (data *pb.CaseData, err error)
	UpdateCaseData(ctx context.Context, id string, data *pb.CaseData) (*pb.CaseData, error)
	GetCaseDatas(ctx context.Context, limit, skip int64, query string) (int64, []*pb.CaseData, error)
	// 班级管理接口
	CreateClass(ctx context.Context, data *pb.Class) (*pb.Class, error)
	GetClass(ctx context.Context, id string) (data *pb.Class, err error)
	DeleteClass(ctx context.Context, id string) (data *pb.Class, err error)
	UpdateClass(ctx context.Context, id string, data *pb.Class) (*pb.Class, error)
	GetClasses(ctx context.Context, limit, skip int64, query string) (int64, []*pb.Class, error)
	// 测评数据明细管理接口
	CreateMeasureDetail(ctx context.Context, data *pb.MeasureDetail) (*pb.MeasureDetail, error)
	GetMeasureDetail(ctx context.Context, id string) (data *pb.MeasureDetail, err error)
	DeleteMeasureDetail(ctx context.Context, id string) (data *pb.MeasureDetail, err error)
	UpdateMeasureDetail(ctx context.Context, id string, data *pb.MeasureDetail) (*pb.MeasureDetail, error)
	GetMeasureDetails(ctx context.Context, limit, skip int64, query string) (int64, []*pb.MeasureDetail, error)
	// 测评数据结果分析管理接口
	CreateMeasureResult(ctx context.Context, data *pb.MeasureResult) (*pb.MeasureResult, error)
	GetMeasureResult(ctx context.Context, id string) (data *pb.MeasureResult, err error)
	DeleteMeasureResult(ctx context.Context, id string) (data *pb.MeasureResult, err error)
	UpdateMeasureResult(ctx context.Context, id string, data *pb.MeasureResult) (*pb.MeasureResult, error)
	GetMeasureResults(ctx context.Context, limit, skip int64, query string) (int64, []*pb.MeasureResult, error)
	// 运动数据管理接口
	CreateMoveData(ctx context.Context, data *pb.MoveData) (*pb.MoveData, error)
	GetMoveData(ctx context.Context, id string) (data *pb.MoveData, err error)
	DeleteMoveData(ctx context.Context, id string) (data *pb.MoveData, err error)
	UpdateMoveData(ctx context.Context, id string, data *pb.MoveData) (*pb.MoveData, error)
	GetMoveDatas(ctx context.Context, limit, skip int64, query string) (int64, []*pb.MoveData, error)
	// 运动处方管理接口
	CreateMovePrescription(ctx context.Context, data *pb.MovePrescription) (*pb.MovePrescription, error)
	GetMovePrescription(ctx context.Context, id string) (data *pb.MovePrescription, err error)
	DeleteMovePrescription(ctx context.Context, id string) (data *pb.MovePrescription, err error)
	UpdateMovePrescription(ctx context.Context, id string, data *pb.MovePrescription) (*pb.MovePrescription, error)
	GetMovePrescriptions(ctx context.Context, limit, skip int64, query string) (int64, []*pb.MovePrescription, error)
	// 作业管理接口
	CreateWork(ctx context.Context, data *pb.Work) (*pb.Work, error)
	GetWork(ctx context.Context, id string) (data *pb.Work, err error)
	DeleteWork(ctx context.Context, id string) (data *pb.Work, err error)
	UpdateWork(ctx context.Context, id string, data *pb.Work) (*pb.Work, error)
	GetWorks(ctx context.Context, limit, skip int64, query string) (int64, []*pb.Work, error)
	// 作业提交管理接口
	CreateWorkSubmit(ctx context.Context, data *pb.WorkSubmit) (*pb.WorkSubmit, error)
	GetWorkSubmit(ctx context.Context, id string) (data *pb.WorkSubmit, err error)
	DeleteWorkSubmit(ctx context.Context, id string) (data *pb.WorkSubmit, err error)
	UpdateWorkSubmit(ctx context.Context, id string, data *pb.WorkSubmit) (*pb.WorkSubmit, error)
	GetWorkSubmits(ctx context.Context, limit, skip int64, query string) (int64, []*pb.WorkSubmit, error)

	Close() error
}

// New 创建数据库句柄
func New(name Type) (Database, error) {
	switch name {
	case MongoDB:
		return mongodb.NewClient()
	default:
		return mongodb.NewClient()
	}
}
