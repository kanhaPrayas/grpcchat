package main

import (
	"flag"

	client "github.com/kanhaPrayas/grpcchat/internal/chatserver/client"
)

func main() {

	//Read inputs from the user for name and Chat room name
	name := flag.String("N", "Prayas", "The name of the user")
	room_name := flag.String("R", "default", "The name of the chat room")
	flag.Parse()

	client := &client.Client{
		Name:     *name,
		RoomName: *room_name,
	}
	client.Exec(*name, *room_name)
}
