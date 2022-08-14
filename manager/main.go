package main

import (
	"flag"

	"gitee.com/healthy/backend/manager/server"
	"github.com/sirupsen/logrus"
)

// inject by go build
var (
	Version   = "0.0.0"
	BuildTime = "2020-01-13-0802 UTC"
)

var (
	debug   = flag.Bool("debug", false, "set log level to debug")
	version = flag.Bool("v", false, "print version")
)

func main() {
	flag.Parse()
	if *version == true {
		logrus.Infoln("Version:", Version)
		logrus.Infoln("BuildTime:", BuildTime)
		return
	}
	if *debug == true {
		logrus.SetLevel(logrus.DebugLevel)
	}
	s := server.NewServer()
	defer s.Close()

	logrus.Infoln("device is starting...")
	if err := s.Serve(); err != nil {
		logrus.Fatalln(err)
	}
}
