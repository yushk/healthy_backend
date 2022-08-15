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

func (m *mongoDB) CreateClass(ctx context.Context, data *pb.Class) (*pb.Class, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CClass().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetClass(ctx context.Context, id string) (data *pb.Class, err error) {
	data = &pb.Class{}
	filter := bson.M{"id": id}
	err = m.CClass().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateClass(ctx context.Context, id string, data *pb.Class) (*pb.Class, error) {
	filter := bson.M{"id": id}

	tmp := &pb.Class{}
	err := m.CClass().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CClass().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteClass(ctx context.Context, id string) (data *pb.Class, err error) {
	data = &pb.Class{}
	filter := bson.M{"id": id}
	err = m.CClass().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetClasses(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.Class, err error) {
	datas = []*pb.Class{}

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
	cursor, err := m.CClass().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Class{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Class Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CClass().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Class Count Documents Error")
		count = int64(0)
	}
	return
}
