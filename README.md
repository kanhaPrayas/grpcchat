# grpcchat
This repo contains a chat server written using protobuf and GRPC to support broadcast and receive messages from multiple clients. There are features like Chat over a Private room, Ignore an User's messages.  
This also has a api service which can fetch all messages and post messages over HTTP.

## To run Chat Server
Update the config.yaml. Without proper values(setting up chatlog value is mandatory) server will not come up. 
chatlog: Location of the file where the messages will be written
serverip: Server IP of the GRPC chat server
serverport: Server Port of the GRPC chat server
apiserverport: API Server Port of the GRPC chat server

cd into cmd/chatserver/server
go run main.go 

## To run Chat Client
cd into cmd/chatserver/client

### With only Name of client. This will connect to default chat room 
go run main.go -N Prayas 

### With Name of client and Room name. 
go run main.go -N Prayas -R Private

### With Name of client and Room name with option to ignore messages from particular user
go run main.go -N Prayas -R Private -B Prayas

## Chat Application over HTTP API

## Running API server
cd into cmd/api/server
go run main.go

### Get messages

#### URL : http://127.0.0.1:8000/messages

Note: Please update the above url as per your config

### POST messages

#### URL : http://127.0.0.1:8000/messages 
Note: Please update the above url as per your config
#### Content-Type : application/json
#### Body: 
{
  "name": "Prayas",
  "room": "Private",
  "message": "Hello I am over HTTP"
}


## To build proto stub in case of any change in protobuf. For demo/test please avoid and running following
### Change Directory into internal/chatserver
####
protoc -I=./internal/chatserver/proto --go_out=./internal/chatserver/proto ./internal/chatserver/proto/service.proto
####
protoc -I=./internal/chatserver/proto --go-grpc_out=./internal/chatserver/proto ./internal/chatserver/proto/service.proto

Comment out mustEmbedUnimplementedBroadcastServer() method inside BroadcastServer interface. As we dont need it.

