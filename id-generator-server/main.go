package main

import (
	"context"
	"log"
	"net"
	db "poll-lite/id-generator-server/database"
	"poll-lite/id-generator-server/generator"
	idgen "poll-lite/id-genereator"

	"google.golang.org/grpc"
)

type server struct {
	idgen.UnsafeIdGeneratorServiceServer
}

func (s *server) GenerateIds(ctx context.Context, in *idgen.GenerateIdsRequest) (*idgen.GenerateIdsResponse, error) {
	result := []string{}

	requestedCount := int(in.Count)
	var err error = nil

	for len(result) < requestedCount {
		var keys []string
		keys, err = generator.GenerateKeys(requestedCount - len(result))

		if err != nil {
			break
		}

		for _, key := range keys {
			err = db.Add(key, true, false)
			if err == nil {
				result = append(result, key)
			}
		}
	}

	return &idgen.GenerateIdsResponse{Ids: result}, err
}

func (s *server) TrackIdUsage(ctx context.Context, in *idgen.TrackIdUsageRequest) (*idgen.TrackIdUsageResponse, error) {
	err := db.SetIsUsed(in.Id, in.IsUsed)
	if !in.IsUsed {
		err = db.SetIsBooked(in.Id, false)
	}
	return &idgen.TrackIdUsageResponse{}, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	idgen.RegisterIdGeneratorServiceServer(s, &server{})

	//db.InitDatabase(":memory:")
	db.InitDatabase("unique_keys.db")

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
