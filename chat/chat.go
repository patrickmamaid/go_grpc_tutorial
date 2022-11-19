package chat

import (
	"context"
	"log"
)

type Server struct {
}

// SayHello hooks into *pb.go
// both in client.go and server.go
// this will not work without you setting up the binding by doing:
// this in:
// client.go -->
//
// c := chat.NewChatServiceClient(conn)
//
// or
// server.go -->
//
// s := chat.Server{}
// grpcServer := grpc.NewServer()
// chat.RegisterChatServiceServer(grpcServer, &s)
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recieved msg from client: %s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}

// mustEmbedUnimplementedChatServiceServer unused but needed because of the pbgo interface requirement
func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	panic("implement me")
}
