# grpcchat
This repo contains a chat server and an API server to support broadcast and receive messages from multiple clients.

## To build proto stub
### Change Directory into internal/chatserver
protoc -I=./proto --go_out=./proto ./proto/service.proto


