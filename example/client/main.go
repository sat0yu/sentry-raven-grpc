package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/sat0yu/sentry-raven-grpc/example/mirror"
	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("missing the request body")
	}
	requestBody := os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMirrorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{RequestBody: requestBody})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.ResponseBody)
}
