# grpcchat
This repo contains a chat server and an API server to support broadcast and receive messages from multiple clients.

## To build proto stub in case of any change in protobuf.
### Change Directory into internal/chatserver
####
protoc -I=./proto --go_out=./proto ./proto/service.proto
####
protoc -I=./proto --go-grpc_out=./proto ./proto/service.proto

## To run Chat Server
cd into cmd/chatserver/server
go run main.go 

## To run Chat Client
cd into cmd/chatserver/client

### With only Name of client. This will connect to default chat room 
go run main.go -N Prayas 

### With Name of client and Room name. 
go run main.go -N Prayas -R Private

