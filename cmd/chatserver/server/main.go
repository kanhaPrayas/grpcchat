package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	conf "github.com/kanhaPrayas/grpcchat/conf"
	proto "github.com/kanhaPrayas/grpcchat/internal/chatserver/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

const (
	config = "../../../config.yaml"
)

var grpcLog glog.LoggerV2
var f *os.File
var err error
var cnf *conf.Conf

func init() {
	cnf = &conf.Conf{}
	cnf = cnf.GetConf(config)
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	// If the file doesn't exist, create it, or append to the file
	f, err = os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Truncate(cnf.ChatLog, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Server struct {
	Connection []*Connection
}

func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	if _, err := f.Write([]byte(msg.Content + "\n")); err != nil {
		log.Fatal(err)
	}
	for _, conn := range s.Connection {

		wait.Add(1)

		go func(msg *proto.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream)

				if err != nil {
					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &proto.Close{}, nil
}

func main() {
	var connections []*Connection

	server := &Server{connections}

	grpcServer := grpc.NewServer()

	tcpPort := fmt.Sprintf(":%d", cnf.Serverport)
	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Infof("Starting server at port :%d", cnf.Serverport)

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	err = os.Remove(cnf.ChatLog)
	if err != nil {
		log.Fatal(err)
	}
}
