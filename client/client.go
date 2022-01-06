package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/juev/counter/proto/counter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func getStats(ctx context.Context, client pb.CounterClient) {
	log.Printf("Getting stats\n")
	stats, err := client.GetStat(ctx, &pb.Domain{Domain: "juev.org"})

	if err != nil {
		log.Fatalf("%v.GetStat(_) = _, %v: ", client, err)
	}
	log.Println(stats)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Cannot close grpc connection")
		}
	}(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := pb.NewCounterClient(conn)
	getStats(ctx, client)
}
