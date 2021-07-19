package main

import (
	"fmt"
	"net"
	"os"

	"github.com/kanhaPrayas/grpcchat/conf"
	proto "github.com/kanhaPrayas/grpcchat/internal/chatserver/proto"
	server "github.com/kanhaPrayas/grpcchat/internal/chatserver/server"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

const (
	config = "../../../config.yaml"
)

//Init variables to be used through out the main
var grpcLog glog.LoggerV2
var fileDesc *os.File
var err error
var cnf *conf.Conf

func init() {
	cnf = &conf.Conf{}
	cnf = cnf.GetConf(config)
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	// If the file doesn't exist, create it, or append to the file
	fileDesc, err = os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		grpcLog.Fatal(err)
	}
	if err := os.Truncate(cnf.ChatLog, 0); err != nil {
		grpcLog.Errorf("Failed to truncate: %v", err)
	}
}
func main() {
	var connection []*server.Connection

	//Create a GRPC server
	server := &server.Server{
		Connection: connection,
		FileDesc:   fileDesc,
		GrpcLog:    grpcLog,
	}

	grpcServer := grpc.NewServer()

	tcpPort := fmt.Sprintf(":%d", cnf.Serverport)
	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		grpcLog.Fatalf("error creating the server %v", err)
	}

	grpcLog.Infof("Starting server at port :%d", cnf.Serverport)

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)

	if err := fileDesc.Close(); err != nil {
		grpcLog.Fatal(err)
	}

	err = os.Remove(cnf.ChatLog)
	if err != nil {
		grpcLog.Fatal(err)
	}
}
