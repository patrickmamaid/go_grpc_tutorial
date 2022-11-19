package main

import (
	"context"
	"github.com/patrickmamaid/go_grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// main()
// the intent here is to communicate with a remote server on port 9000
// to USE the chat.* ability we created with protobuf protoc grpc within the package chat
func main() {

	// @@@@@@@<<<<<<<<<<<<<<<<< CONNECTION PREPARATION >>>>>>>>>>>>>>>>>@@@@@@@@@@@
	// because later on, grpc.Dial() returns a pointer to grpc.ClientConn, we make it here first
	// but realistically, you do not need to do this because of conn, err := grpc.Dial
	// you can comment this out :), im leaving it here for clarity
	var conn *grpc.ClientConn

	// Dial creates a client connection to the given target, this line is in preparation for
	//          c := chat.NewChatServiceClient(conn) later on.
	// what we will get here is, a configured *conn which has insecure creds and a hosttarget:port
	//
	// grpc.WithTransportCredentials(insecure.NewCredentials()) is used because grpc.WithInsecure() is now deprecated
	// I have found this while looking through package grpc
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	// catch any connection errors here and print it
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	// in the package chat, we have NewChatServiceClient() (constructor) in chat_grpc.pb.go (autogenerated)
	// what this line is doing is essentially: "I want to bind the chat service client, with a connection config"
	// as a result you get a c ChatServiceClient
	c := chat.NewChatServiceClient(conn)

	// @@@@@@@<<<<<<<<<<<<<<<<< END PREPARATION >>>>>>>>>>>>>>>>>@@@@@@@@@@@

	// message := chat.Message{Body: "hello from the client"}
	// IS the actual outbound rpc call with a string message as specified in chat.proto:
	//  syntax = "proto3" ;
	//  option go_package = "./chat";
	//  package chat;
	//  message Message{
	//  	string body = 1;
	//  }
	//  service ChatService {
	//  	rpc SayHello(Message) returns (Message) {}
	//  }
	message := chat.Message{Body: "hello from the client"}

	// for every rpc send, you get a request. here we handle it by just printing it:
	response, err := c.SayHello(context.Background(), &message)
	// always check for errors first
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	log.Printf("server responded with: %s", response)

}
