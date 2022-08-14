package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gitee.com/healthy/backend/pkg/pb"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type authorization struct {
	Username   string
	Credential string
}

func (m *mongoDB) Authenticate(ctx context.Context, username, password string) bool {
	logrus.Debugln("Authenticate", username, password)

	filter := bson.M{"username": username}
	auth := authorization{}
	err := m.CAuth().FindOne(ctx, filter).Decode(&auth)
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Credential), []byte(password))
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	return true
}

// AddAuthorization ...
func (m *mongoDB) AddAuthorization(ctx context.Context, username, password string) (err error) {
	// unique value check
	filter := bson.M{"username": username}
	count, err := m.CAuth().CountDocuments(ctx, filter)
	if err != nil {
		return
	}

	if count != 0 {
		return errors.New(username + " is already existed")
	}
	logrus.Debugln("Add Authorization", username, password)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	auth := authorization{
		Username:   username,
		Credential: string(bytes),
	}
	_, err = m.CAuth().InsertOne(ctx, &auth)
	return
}

// RemoveAuthorization ...
func (m *mongoDB) RemoveAuthorization(ctx context.Context, username string) (err error) {
	filter := bson.M{"username": username}
	_, err = m.CAuth().DeleteMany(ctx, filter)
	return
}

func (m *mongoDB) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	if user.Id == "" {
		user.Id = xid.New().String()
	}

	err := m.AddAuthorization(ctx, user.Name, user.Ps)
	if err != nil {
		logrus.Errorln(err)
		return user, err
	}
	user.Ps = ""
	_, err = m.CUser().InsertOne(ctx, user)
	if err != nil {
		logrus.Errorln(err)
		return user, err
	}
	return user, nil
}

func (m *mongoDB) GetUser(ctx context.Context, id string) (user *pb.User, err error) {
	user = &pb.User{}
	filter := bson.M{"id": id}
	err = m.CUser().FindOne(ctx, filter).Decode(user)
	return
}

func (m *mongoDB) UpdateUser(ctx context.Context, id string, user *pb.User) (*pb.User, error) {
	filter := bson.M{"id": id}

	tmp := &pb.User{}
	err := m.CUser().FindOne(ctx, filter).Decode(tmp)
	if err != nil {
		return nil, err
	}
	if user.Name != tmp.Name {
		return nil, fmt.Errorf("can not modify username")
	}
	user.Id = tmp.Id
	update := bson.M{"$set": user}
	err = m.CUser().FindOneAndUpdate(ctx, filter, update).Err()
	return user, err
}

func (m *mongoDB) DeleteUser(ctx context.Context, id string) (user *pb.User, err error) {
	user = &pb.User{}
	filter := bson.M{"id": id}
	err = m.CUser().FindOneAndDelete(ctx, filter).Decode(user)
	// 删除鉴权信息
	m.RemoveAuthorization(ctx, user.Name)
	return
}

func (m *mongoDB) GetUsers(ctx context.Context, limit, skip int64, query string) (count int64, users []*pb.User, err error) {
	users = []*pb.User{}

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
	cursor, err := m.CUser().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		user := &pb.User{}
		err := cursor.Decode(user)
		if err != nil {
			logrus.WithError(err).Errorln("User Decode Error")
			return 0, users, err
		}
		users = append(users, user)
	}
	count, err = m.CUser().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("User Count Documents Error")
		count = int64(0)
	}
	return
}

// ChangeAuthorization 修改权限
func (m *mongoDB) ChangeAuthorization(ctx context.Context, username, password string) (err error) {
	// unique value check
	filter := bson.M{"username": username}
	logrus.Debugln("change Authorization", username, password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	auth := authorization{
		Username:   username,
		Credential: string(bytes),
	}
	update := bson.M{"$set": auth}
	err = m.CAuth().FindOneAndUpdate(ctx, filter, update).Err()
	return
}

func (m *mongoDB) GetUserByName(ctx context.Context, name string) (user *pb.User, err error) {
	user = &pb.User{}
	filter := bson.M{"name": name}
	err = m.CUser().FindOne(ctx, filter).Decode(user)
	return
}
