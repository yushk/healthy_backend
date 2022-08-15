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

func (m *mongoDB) CreateMeasureResult(ctx context.Context, data *pb.MeasureResult) (*pb.MeasureResult, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CMeasureResult().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetMeasureResult(ctx context.Context, id string) (data *pb.MeasureResult, err error) {
	data = &pb.MeasureResult{}
	filter := bson.M{"id": id}
	err = m.CMeasureResult().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateMeasureResult(ctx context.Context, id string, data *pb.MeasureResult) (*pb.MeasureResult, error) {
	filter := bson.M{"id": id}

	tmp := &pb.MeasureResult{}
	err := m.CMeasureResult().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CMeasureResult().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteMeasureResult(ctx context.Context, id string) (data *pb.MeasureResult, err error) {
	data = &pb.MeasureResult{}
	filter := bson.M{"id": id}
	err = m.CMeasureResult().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetMeasureResults(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.MeasureResult, err error) {
	datas = []*pb.MeasureResult{}

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
	cursor, err := m.CMeasureResult().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.MeasureResult{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("MeasureResult Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CMeasureResult().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("MeasureResult Count Documents Error")
		count = int64(0)
	}
	return
}
