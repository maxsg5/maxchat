package main

import (
	"context"
	"log"
	"net"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-simplechat/server/api"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) GetMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Message: "Hello from gRPC server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Load server's certificate and private key from the scripts directory
	//TODO: Right now this is for local development. This will need to be changed for production.
	cert, err := tls.LoadX509KeyPair("../scripts/localhost.crt", "../scripts/localhost.key")
	if err != nil {
		log.Fatalf("failed to load server's certificate and key: %v", err)
	}

	// Create the TLS credentials for transport
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
	})

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterMyServiceServer(s, &server{})
    log.Println("Server started and listening on localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
