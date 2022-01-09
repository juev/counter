package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	pb "github.com/juev/counter/proto/counter"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	rdb       *redis.Client
	portGrpc  = kingpin.Flag("grpc", "The GRPC port").Default("50051").Int()
	portFiber = kingpin.Flag("http", "The HTTP port").Default("50052").Int()
)

type Implementation struct {
	pb.UnimplementedCounterServer
}

func (s *Implementation) AddDomain(ctx context.Context, domain *pb.Domain) (*pb.Response, error) {
	if !keyExist(ctx, domain.GetDomain()) {
		var index int
		for {
			index = rangeIn()
			if !keyExist(ctx, strconv.Itoa(index)) {
				break
			}
		}
		log.Printf("index: %d", index)
		err := rdb.Set(ctx, domain.Domain, index, 0).Err()
		if err != nil {
			return nil, err
		}
		err = rdb.Set(ctx, strconv.Itoa(index), 0, 0).Err()
		if err != nil {
			return nil, err
		}
		return &pb.Response{
			Status: strconv.Itoa(index),
		}, nil
	}
	index, err := getValue(ctx, domain.GetDomain())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Status: index}, nil
}

func (s *Implementation) RemoveDomain(ctx context.Context, domain *pb.Domain) (*pb.Response, error) {
	index, err := getValue(ctx, domain.GetDomain())
	if err != nil {
		return nil, err
	}
	rdb.Del(ctx, index)
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

	parseInt, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return nil, err
	}

	if domain.Domain == "any" {
		return &pb.Stats{
			Cnt: parseInt,
		}, nil
	}

	val, err := getValue(ctx, index)
	if err != nil {
		return nil, err
	}

	valInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil, err
	}

	return &pb.Stats{
		Cnt: valInt,
	}, nil
}

func getValue(ctx context.Context, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("cannot get `%s` key: %v", key, err)
	}
	return val, nil
}

func rangeIn() int {
	low := 10000000
	hi := 99999999
	return low + rand.Intn(hi-low) //nolint:gosec
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

func runFiber(ctx context.Context, port int) {
	appFiber := fiber.New()

	appFiber.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		if !keyExist(ctx, id) {
			id = "any"
		}

		val := rdb.Incr(ctx, id)
		return c.JSON(fiber.Map{id: val.Val()})
	})

	err := appFiber.Listen(fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("error on handle: %v", err)
	}
}

func runGrpc(port int) {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCounterServer(grpcServer, newCounterServer())
	log.Printf("grpc is running on address: %s", address)
	errServe := grpcServer.Serve(lis)
	if errServe != nil {
		log.Fatalf("failed to start grpc: %v", err)
	}
}

func main() {
	kingpin.Parse()

	ctx := context.Background()

	initRedis(ctx)
	go runFiber(ctx, *portFiber)
	runGrpc(*portGrpc)
}
