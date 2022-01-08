package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/juev/counter/proto/counter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                 = kingpin.New("client", "Client for counter")
	serverAddr          = app.Flag("addr", "The server address in the format of host:port").Default("localhost:50051").String()
	addDomainCommand    = app.Command("add", "Add domain to counter")
	addDomainArg        = addDomainCommand.Arg("domain", "Domain name").Required().String()
	removeDomainCommand = app.Command("remove", "Remove domain from counter")
	removeDomainArg     = removeDomainCommand.Arg("domain", "Domain for remove").Required().String()
	statCommand         = app.Command("stat", "Get stat for domain")
	statArg             = statCommand.Arg("domain", "Domain for stats").Required().String()
)

func getStats(ctx context.Context, client pb.CounterClient, domain string) (*pb.Stats, error) {
	response, err := client.GetStat(ctx, &pb.Domain{Domain: domain})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func addDomain(ctx context.Context, client pb.CounterClient, domain string) (*pb.Response, error) {
	response, err := client.AddDomain(ctx, &pb.Domain{Domain: domain})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func removeDomain(ctx context.Context, client pb.CounterClient, domain string) (*pb.Response, error) {
	response, err := client.RemoveDomain(ctx, &pb.Domain{Domain: domain})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func initClient(server string) pb.CounterClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Cannot close grpc connection")
		}
	}(conn)

	return pb.NewCounterClient(conn)
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case addDomainCommand.FullCommand():
		log.Println("addDomainCommand")
		client := initClient(*serverAddr)
		log.Printf("client: %s", *serverAddr)
		response, err := addDomain(ctx, client, *addDomainArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println(response.Status)
	case removeDomainCommand.FullCommand():
		println("removeDomainCommand")
		client := initClient(*serverAddr)
		response, err := removeDomain(ctx, client, *removeDomainArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println(response.Status)

	case statCommand.FullCommand():
		println("statCommand")
		client := initClient(*serverAddr)
		response, err := getStats(ctx, client, *statArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("%s: %d", *statArg, response.Cnt)
	}
}
