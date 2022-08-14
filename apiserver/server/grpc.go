package server

import (
	"gitee.com/healthy/backend/pkg/pb"
	"github.com/sirupsen/logrus"
)

// 客户端列表
var client pb.GRPCClient

func RegisterClients() {
	var err error
	client, err = pb.NewGRPCClient()
	if err != nil {
		logrus.WithError(err).Fatalln("connect grpc server error")
	} else {
		logrus.Infoln("connect grpc server ok")
	}
}
