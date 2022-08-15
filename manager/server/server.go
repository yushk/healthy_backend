package server

import (
	"context"
	"net"
	"sync"

	config2 "github.com/yushk/healthy_backend/pkg/config"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/yushk/healthy_backend/manager/internal/db"
	"github.com/yushk/healthy_backend/manager/server/helper"
	"github.com/yushk/healthy_backend/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// RoleRoot 超级管理员角色
	RoleRoot = "root"
)

type Server struct {
	pb.UnimplementedUserManagerServer
	pb.UnimplementedPersonalManagerServer
	pb.UnimplementedCaseDataManagerServer
	pb.UnimplementedClassManagerServer
	pb.UnimplementedMeasureDetailManagerServer
	pb.UnimplementedMeasureResultManagerServer
	pb.UnimplementedMoveDataManagerServer
	pb.UnimplementedMovePrescriptionManagerServer
	pb.UnimplementedWorkManagerServer
	pb.UnimplementedWorkSubmitManagerServer

	// GRPCHost   string `long:"grpc-host" description:"the IP to listen on for grpc" default:"0.0.0.0" env:"GRPC_HOST"`
	// GRPCPort   int    `long:"grpc-port" description:"the port to listen on for grpc's insecure connections" default:"30031" env:"GRPC_PORT"`
	grpcServer *grpc.Server

	// db 数据库client
	db   db.Database
	cron *cron.Cron

	LogLevel string `long:"log-level" description:"the log level of logrus" default:"debug" env:"LOG_LEVEL"`
}

func NewServer() *Server {
	return &Server{
		cron: cron.New(),
	}
}

func (s *Server) Close() (err error) {
	s.grpcServer.GracefulStop()
	_ = s.db.Close()
	s.cron.Stop()
	return
}

func (s *Server) Serve() (err error) {
	// 初始化数据库
	err = s.initDB()
	if err != nil {
		logrus.Fatalln("init db err", err)
		return
	}

	// 定时任务
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.grpcServe()
	}()
	wg.Wait()
	logrus.Infoln("Server Exited ...")
	return
}

func (s *Server) grpcServe() error {
	network := config2.GetString(config2.ManagerProtocol)
	if network == "" {
		network = "tcp"
	}
	host := config2.GetString(config2.ManagerHost)
	if host == "" {
		host = "127.0.0.1"
	}
	port := config2.GetString(config2.ManagerPort)
	if port == "" {
		port = "30031"
	}
	listener, err := net.Listen(network, net.JoinHostPort(host, port))
	if err != nil {
		logrus.Fatalf("failed to system listen: %v", err)
		return err
	}
	opts := []grpc.ServerOption{
		// StatsHandler returns a ServerOption that sets the stats handler for the server.
		grpc.StatsHandler(&helper.StatsHandler{}),
	}
	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterUserManagerServer(s.grpcServer, s)
	pb.RegisterPersonalManagerServer(s.grpcServer, s)
	pb.RegisterCaseDataManagerServer(s.grpcServer, s)
	pb.RegisterClassManagerServer(s.grpcServer, s)
	pb.RegisterMeasureDetailManagerServer(s.grpcServer, s)
	pb.RegisterMeasureResultManagerServer(s.grpcServer, s)
	pb.RegisterMoveDataManagerServer(s.grpcServer, s)
	pb.RegisterMovePrescriptionManagerServer(s.grpcServer, s)
	pb.RegisterWorkManagerServer(s.grpcServer, s)
	pb.RegisterWorkSubmitManagerServer(s.grpcServer, s)
	// grpc cli
	reflection.Register(s.grpcServer)

	logrus.Infoln("system server started, addr: ", listener.Addr().String())
	err = s.grpcServer.Serve(listener)
	if err != nil {
		logrus.Fatalf("failed to system serve: %v", err)
		return err
	}
	logrus.Infoln("system server existed ...")
	return nil
}

func (s *Server) initDB() error {
	client, err := db.New(db.MongoDB)
	if err != nil {
		return err
	}
	s.db = client
	// 初始化超级管理员(超级管理员从环境变量或配置文件识别)
	rootid := config2.GetString(config2.RootUserID)
	rootname := config2.GetString(config2.RootUsername)
	password := config2.GetString(config2.RootPassword)
	ctx := context.TODO()
	_, err = client.GetUserByName(ctx, rootname)
	if err != nil {
		logrus.Debugln("get root user error", err)
		_, err = client.CreateUser(ctx, &pb.User{
			Id:   rootid,
			Name: rootname,
			Role: RoleRoot,
			Ps:   password,
		})
		if err != nil {
			logrus.Errorln("create root user error", err)
			return nil
		}
		return err
	}
	return err
}
