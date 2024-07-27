package main

import (
	"context"
	"log"
	"net"
	"poll-lite/id-generator-server/generator"
	"poll-lite/id-generator-server/tracker"
	pb "poll-lite/id-genereator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeIdGeneratorServiceServer
}

func (s *server) GenerateIds(ctx context.Context, in *pb.GenerateIdsRequest) (*pb.GenerateIdsResponse, error) {

	keys, err := generator.GenerateKeys(in.Count)
	return &pb.GenerateIdsResponse{Ids: keys}, err
}

func (s *server) TrackIdUsage(ctx context.Context, in *pb.TrackIdUsageRequest) (*pb.TrackIdUsageResponse, error) {
	err := tracker.TrackIdUsage(in.Id, in.IsUsed)
	return &pb.TrackIdUsageResponse{}, err
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
