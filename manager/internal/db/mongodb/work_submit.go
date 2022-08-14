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

func (m *mongoDB) CreateWorkSubmit(ctx context.Context, data *pb.WorkSubmit) (*pb.WorkSubmit, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CWorkSubmit().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetWorkSubmit(ctx context.Context, id string) (data *pb.WorkSubmit, err error) {
	data = &pb.WorkSubmit{}
	filter := bson.M{"id": id}
	err = m.CWorkSubmit().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateWorkSubmit(ctx context.Context, id string, data *pb.WorkSubmit) (*pb.WorkSubmit, error) {
	filter := bson.M{"id": id}

	tmp := &pb.WorkSubmit{}
	err := m.CWorkSubmit().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CWorkSubmit().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteWorkSubmit(ctx context.Context, id string) (data *pb.WorkSubmit, err error) {
	data = &pb.WorkSubmit{}
	filter := bson.M{"id": id}
	err = m.CWorkSubmit().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetWorkSubmits(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.WorkSubmit, err error) {
	datas = []*pb.WorkSubmit{}

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
	cursor, err := m.CWorkSubmit().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.WorkSubmit{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("WorkSubmit Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CWorkSubmit().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("WorkSubmit Count Documents Error")
		count = int64(0)
	}
	return
}
