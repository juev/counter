package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
	pb "github.com/juev/counter/proto/counter"
	"google.golang.org/grpc"
)

var (
	rdb  *redis.Client
	port = flag.Int("port", 50051, "The server port")
)

type Implementation struct {
	pb.UnimplementedCounterServer
}

func (s *Implementation) AddDomain(ctx context.Context, domain *pb.Domain) (*pb.Response, error) {
	if !keyExist(ctx, domain.GetDomain()) {
		err := rdb.Set(ctx, domain.Domain, 0, 0).Err()
		if err != nil {
			return nil, err
		}
		return &pb.Response{
			Status: `ok`,
		}, nil
	}
	return &pb.Response{Status: `domain already exist`}, nil
}

func (s *Implementation) RemoveDomain(ctx context.Context, domain *pb.Domain) (*pb.Response, error) {
	rdb.Del(ctx, domain.GetDomain())
	return &pb.Response{
		Status: `domain deleted`,
	}, nil
}

func (s *Implementation) GetStat(ctx context.Context, domain *pb.Domain) (*pb.Stats, error) {
	if keyExist(ctx, domain.GetDomain()) {
		val, err := rdb.Get(ctx, domain.GetDomain()).Result()
		switch {
		case err == redis.Nil:
			return &pb.Stats{Cnt: 0}, nil
		case err != nil:
			log.Fatalf("cannot get `%s` key: %v", domain.GetDomain(), err)
		default:
			result, _ := strconv.Atoi(val)
			return &pb.Stats{Cnt: int64(result)}, nil
		}
	}
	log.Println("Get request /stats")
	return &(pb.Stats{Cnt: 1}), nil
}

func newCounterServer() *Implementation {
	return &Implementation{}
}

func initRedis(ctx context.Context) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if !keyExist(ctx, "any") {
		err := rdb.Set(ctx, "any", 0, 0).Err()
		if err != nil {
			log.Fatalf("cannot set `any` key on redis: %v", err)
		}
	}
}

func keyExist(ctx context.Context, key string) bool {
	_, err := rdb.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return false
	case err != nil:
		log.Fatalf("cannot get `%s` key: %v", key, err)
	default:
		return true
	}
	return false
}

func main() {
	ctx := context.Background()

	initRedis(ctx)

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
