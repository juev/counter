package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/juev/counter/proto/counter"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Implementation struct {
	pb.UnimplementedCounterServer
}

func (s *Implementation) AddDomain(context.Context, *pb.Domain) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDomain not implemented")
}

func (s *Implementation) RemoveDomain(context.Context, *pb.Domain) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDomain not implemented")
}

func (s *Implementation) GetStat(context.Context, *pb.Domain) (*pb.Stats, error) {
	log.Println("Get request /stats")
	return &(pb.Stats{Cnt: 1}), nil
}

func newCounterServer() *Implementation {
	return &Implementation{}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCounterServer(grpcServer, newCounterServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
