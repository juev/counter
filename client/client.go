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
	serverAddr          = app.Flag("addr", "The server address in the format of host:port").Default("127.0.0.1:50051").String()
	addDomainCommand    = app.Command("add", "Add domain to counter")
	addDomainArg        = addDomainCommand.Arg("domain", "Domain name").Required().String()
	removeDomainCommand = app.Command("remove", "Remove domain from counter")
	removeDomainArg     = removeDomainCommand.Arg("domain", "Domain for remove").Required().String()
	statCommand         = app.Command("stat", "Get stat for domain")
	statArg             = statCommand.Arg("domain", "Domain for stats").Required().String()
)

func getStats(ctx context.Context, domain string) (*pb.Stats, error) {
	client := initClient()
	response, err := client.GetStat(ctx, &pb.Domain{Domain: domain})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func addDomain(ctx context.Context, domain string) (*pb.Response, error) {
	client := initClient()
	response, err := client.AddDomain(ctx, &pb.Domain{Domain: domain})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func removeDomain(ctx context.Context, domain string) (*pb.Response, error) {
	client := initClient()
	response, err := client.RemoveDomain(ctx, &pb.Domain{Domain: domain})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func initClient() pb.CounterClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return pb.NewCounterClient(conn)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case addDomainCommand.FullCommand():
		log.Println("addDomainCommand")
		response, err := addDomain(ctx, *addDomainArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println(response.Status)
	case removeDomainCommand.FullCommand():
		println("removeDomainCommand")
		response, err := removeDomain(ctx, *removeDomainArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println(response.Status)

	case statCommand.FullCommand():
		println("statCommand")
		response, err := getStats(ctx, *statArg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("%s: %d", *statArg, response.Cnt)
	}
}
