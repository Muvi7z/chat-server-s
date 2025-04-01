package main

import (
	"github.com/Muvi7z/chat-auth-s/gen/api/chat_v1"
	"github.com/Muvi7z/chat-auth-s/internal/api/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50041")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, chat.NewImplementationChat())

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
