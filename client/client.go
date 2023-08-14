package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-simplechat/server/api" // Adjust this import path based on your actual project structure
)

func main() {
	// Load the server's CA certificate
	//TODO: Right now this is for local development. This will need to be changed for production.
	creds, err := credentials.NewClientTLSFromFile("../scripts/server.crt", "")
	if err != nil {
		log.Fatalf("failed to load server's CA certificate: %v", err)
	}

	// Set up a connection to the server with the credentials
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	// Example call to the server
	response, err := c.GetMessage(context.Background(), &pb.MessageRequest{})
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
