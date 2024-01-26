package main

import (
	"github.com/stewie/internal/application"
	"github.com/stewie/internal/pb"
	"github.com/stewie/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	err := application.Configure()
	if err != nil {
		log.Fatal(err)
		return
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterJiraServer(s, server.NewServer())

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
