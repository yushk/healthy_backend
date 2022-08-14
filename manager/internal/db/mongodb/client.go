package mongodb

import (
	"context"
	"fmt"

	config2 "gitee.com/healthy/backend/pkg/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Mongo数据库名和表名定义
const (
	dbName                     = "healthy"
	collectionUser             = "user"              // 用户表
	collectionAuth             = "auth"              // 用户认证(含认证信息)
	collectionPersonal         = "personal"          // 个人基础信息表
	collectionCaseData         = "case_data"         // 病例数据表
	collectionClass            = "class"             // 班级表
	collectionMeasureDetail    = "measure_detail"    // 测评数据明细表
	collectionMeasureResult    = "measure_result"    // 测评数据结果分析表
	collectionMoveData         = "move_data"         // 运动数据表
	collectionMovePrescription = "move_prescription" // 运动处方表
	collectionWork             = "work"              // 作业表
	collectionWorkSubmit       = "work_submit"       // 作业提交表
)

type mongoDB struct {
	client *mongo.Client
}

func NewClient() (*mongoDB, error) {
	mongoURI := config2.GetString(config2.DBAddress)
	username := config2.GetString(config2.DBUsername)
	password := config2.GetString(config2.DBPassword)
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	if username != "" {
		auth := options.Credential{
			Username: username,
			Password: password,
		}
		clientOptions.SetAuth(auth)
	}
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logrus.Fatalln(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	// Create index
	err = createEnsureIndex(client)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln("Connected to MongoDB!")

	return &mongoDB{
		client: client,
	}, nil
}

func (m *mongoDB) Close() error {
	if m.client != nil {
		return m.client.Disconnect(context.TODO())
	}
	return nil
}

// DB 数据库句柄
func (m *mongoDB) DB(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return m.client.Database(name)
}

// C 集合句柄
func (m *mongoDB) C(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.client.Database(dbName).Collection(name, opts...)
}

// CAuth 认证集合句柄
func (m *mongoDB) CAuth() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionAuth)
}

// CUser 用户集合句柄
func (m *mongoDB) CUser() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionUser)
}

// CPersonal 个人基础信息集合句柄
func (m *mongoDB) CPersonal() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionPersonal)
}

// CCaseData 病例数据表集合句柄
func (m *mongoDB) CCaseData() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionCaseData)
}

// CClass 班级表集合句柄
func (m *mongoDB) CClass() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionClass)
}

// CMeasureDetail 测评数据明细表集合句柄
func (m *mongoDB) CMeasureDetail() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionMeasureDetail)
}

// CMeasureResult 测评数据结果分析表集合句柄
func (m *mongoDB) CMeasureResult() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionMeasureResult)
}

// CMoveData 运动数据表集合句柄
func (m *mongoDB) CMoveData() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionMoveData)
}

// CMovePrescription 作业表集合句柄
func (m *mongoDB) CMovePrescription() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionMovePrescription)
}

// CWork 运动处方表集合句柄
func (m *mongoDB) CWork() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionWork)
}

// CWorkSubmit 作业提交表集合句柄
func (m *mongoDB) CWorkSubmit() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionWorkSubmit)
}

// createEnsureIndex  创建唯一索引
func createEnsureIndex(client *mongo.Client) (err error) {
	collections := map[string][]string{
		collectionUser:     {"name"},
		collectionPersonal: {"name"},
		// collectionMT:         {"parkId", "deviceId", "createdAt"},
	}

	for collect, rowArr := range collections {
		switch len(rowArr) {
		case 1:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		case 2:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
						{Key: rowArr[1], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		case 3:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
						{Key: rowArr[1], Value: bsonx.Int32(1)},
						{Key: rowArr[2], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		default:
			err = fmt.Errorf("should implement more")
		}
		if err != nil {
			logrus.Fatalln(err)
		}
	}
	return
}
