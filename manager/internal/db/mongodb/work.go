package mongodb

import (
	"context"
	"encoding/json"

	"gitee.com/healthy/backend/pkg/pb"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreateWork(ctx context.Context, data *pb.Work) (*pb.Work, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CWork().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetWork(ctx context.Context, id string) (data *pb.Work, err error) {
	data = &pb.Work{}
	filter := bson.M{"id": id}
	err = m.CWork().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateWork(ctx context.Context, id string, data *pb.Work) (*pb.Work, error) {
	filter := bson.M{"id": id}

	tmp := &pb.Work{}
	err := m.CWork().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CWork().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteWork(ctx context.Context, id string) (data *pb.Work, err error) {
	data = &pb.Work{}
	filter := bson.M{"id": id}
	err = m.CWork().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetWorks(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.Work, err error) {
	datas = []*pb.Work{}

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
	cursor, err := m.CWork().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Work{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Work Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CWork().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Work Count Documents Error")
		count = int64(0)
	}
	return
}
