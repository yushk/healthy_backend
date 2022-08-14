package mongodb

import (
	"context"
	"encoding/json"
	"fmt"

	"gitee.com/healthy/backend/pkg/pb"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreatePersonal(ctx context.Context, data *pb.Personal) (*pb.Personal, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CPersonal().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetPersonal(ctx context.Context, id string) (data *pb.Personal, err error) {
	data = &pb.Personal{}
	filter := bson.M{"id": id}
	err = m.CPersonal().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdatePersonal(ctx context.Context, id string, data *pb.Personal) (*pb.Personal, error) {
	filter := bson.M{"id": id}

	tmp := &pb.Personal{}
	err := m.CPersonal().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	if data.Name != tmp.Name {
		return nil, fmt.Errorf("can not modify username")
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CPersonal().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeletePersonal(ctx context.Context, id string) (data *pb.Personal, err error) {
	data = &pb.Personal{}
	filter := bson.M{"id": id}
	err = m.CPersonal().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetPersonals(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.Personal, err error) {
	datas = []*pb.Personal{}

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
	cursor, err := m.CPersonal().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Personal{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Personal Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CPersonal().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Personal Count Documents Error")
		count = int64(0)
	}
	return
}
