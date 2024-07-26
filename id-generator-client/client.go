package idgeneratorclient

import (
	"context"
	"log"
	pb "poll-lite/id-genereator"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const target = "localhost:50051"

func GenerateIds(count uint32) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()

	c := pb.NewIdGeneratorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//-----------------------
	r, err := c.GenerateIds(ctx, &pb.GenerateIdsRequest{Count: count})
	if err != nil {
		log.Fatalf("error calling function GenerateIds: %v", err)
	}

	log.Printf("Response from gRPC server's GenerateIds function: %v", r.GetIds())
}

func TrackIdUsage(id string, isUsed bool) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()

	c := pb.NewIdGeneratorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//-----------------------
	r, err := c.TrackIdUsage(ctx, &pb.TrackIdUsageRequest{Id: id, IsUsed: isUsed})
	if err != nil {
		log.Fatalf("error calling function TrackIdUsage: %v", err)
	}

	log.Printf("Response from gRPC server's TrackIdUsage function: %v", r.String())
}
