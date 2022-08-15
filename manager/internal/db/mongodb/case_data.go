package mongodb

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/pkg/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreateCaseData(ctx context.Context, data *pb.CaseData) (*pb.CaseData, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CCaseData().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetCaseData(ctx context.Context, id string) (data *pb.CaseData, err error) {
	data = &pb.CaseData{}
	filter := bson.M{"id": id}
	err = m.CCaseData().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateCaseData(ctx context.Context, id string, data *pb.CaseData) (*pb.CaseData, error) {
	filter := bson.M{"id": id}

	tmp := &pb.CaseData{}
	err := m.CCaseData().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CCaseData().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteCaseData(ctx context.Context, id string) (data *pb.CaseData, err error) {
	data = &pb.CaseData{}
	filter := bson.M{"id": id}
	err = m.CCaseData().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetCaseDatas(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.CaseData, err error) {
	datas = []*pb.CaseData{}

	filter := bson.M{}
	if query != "" {
		err = json.Unmarshal([]byte(query), &filter)
		if err != nil {
			return
		}
	}

	findOption := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	cursor, err := m.CCaseData().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.CaseData{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("CaseData Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CCaseData().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("CaseData Count Documents Error")
		count = int64(0)
	}
	return
}
