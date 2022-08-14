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

func (m *mongoDB) CreateMoveData(ctx context.Context, data *pb.MoveData) (*pb.MoveData, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CMoveData().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetMoveData(ctx context.Context, id string) (data *pb.MoveData, err error) {
	data = &pb.MoveData{}
	filter := bson.M{"id": id}
	err = m.CMoveData().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateMoveData(ctx context.Context, id string, data *pb.MoveData) (*pb.MoveData, error) {
	filter := bson.M{"id": id}

	tmp := &pb.MoveData{}
	err := m.CMoveData().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	data.Id = tmp.Id
	update := bson.M{"$set": data}
	err = m.CMoveData().FindOneAndUpdate(ctx, filter, update).Err()
	return data, err
}

func (m *mongoDB) DeleteMoveData(ctx context.Context, id string) (data *pb.MoveData, err error) {
	data = &pb.MoveData{}
	filter := bson.M{"id": id}
	err = m.CMoveData().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetMoveDatas(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.MoveData, err error) {
	datas = []*pb.MoveData{}

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
	cursor, err := m.CMoveData().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.MoveData{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("MoveData Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CMoveData().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("MoveData Count Documents Error")
		count = int64(0)
	}
	return
}
