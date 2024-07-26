package main

import (
	"context"
	"log"
	"net"
	pb "poll-lite/id-genereator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeIdGeneratorServiceServer
}

func (s *server) GenerateIds(ctx context.Context, in *pb.GenerateIdsRequest) (*pb.GenerateIdsResponse, error) {
	return &pb.GenerateIdsResponse{Message: "Hello, World! "}, nil
}

func (s *server) TrackIdUsage(ctx context.Context, in *pb.TrackIdUsageRequest) (*pb.TrackIdUsageResponse, error) {
	return &pb.TrackIdUsageResponse{Message: "Hello, World! "}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterIdGeneratorServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
