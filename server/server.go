package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
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
		var index *big.Int
		for {
			index, _ := rand.Prime(rand.Reader, 64)
			if !keyExist(ctx, index.String()) {
				break
			}
		}
		err := rdb.Set(ctx, domain.Domain, index, 0).Err()
		if err != nil {
			return nil, err
		}
		err = rdb.Set(ctx, index.String(), 0, 0).Err()
		if err != nil {
			return nil, err
		}
		return &pb.Response{
			Status: index.String(),
		}, nil
	}
	index, err := getValue(ctx, domain.GetDomain())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Status: strconv.FormatInt(index, 10)}, nil
}

func (s *Implementation) RemoveDomain(ctx context.Context, domain *pb.Domain) (*pb.Response, error) {
	index, err := getValue(ctx, domain.GetDomain())
	if err != nil {
		return nil, err
	}
	rdb.Del(ctx, strconv.FormatInt(index, 10))
	rdb.Del(ctx, domain.GetDomain())
	return &pb.Response{
		Status: `domain deleted`,
	}, nil
}

func (s *Implementation) GetStat(ctx context.Context, domain *pb.Domain) (*pb.Stats, error) {
	index, err := getValue(ctx, domain.GetDomain())
	if err != nil {
		return nil, err
	}

	val, err := getValue(ctx, strconv.FormatInt(index, 10))
	if err != nil {
		return nil, err
	}

	return &pb.Stats{
		Cnt: val,
	}, nil
}

func getValue(ctx context.Context, key string) (int64, error) {
	val, err := rdb.Get(ctx, key).Int64()
	if err != nil {
		return 0, fmt.Errorf("cannot get `%s` key: %v", key, err)
	}
	return val, nil
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
	if err == nil || err != redis.Nil {
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
