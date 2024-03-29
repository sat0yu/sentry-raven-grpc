package main

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net"

	pb "github.com/sat0yu/sentry-raven-grpc/example/mirror"
	raven "github.com/sat0yu/sentry-raven-grpc"
	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

type server struct{}

// Echo implements mirror.Mirror
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.RequestBody)

	if len(in.RequestBody) > 5 {
		log.Printf("too long! the error is reported to Sentry")
		return nil, errors.New("too long")
	}

	var result []byte
	for _, ch := range []byte(in.RequestBody) {
		result = append([]byte{ch}, result...)
	}
	return &pb.EchoResponse{ResponseBody: string(result)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: replace your Sentry DSN
	dsn := "https://XXXXXXXXXXX:YYYYYYYYYY.io/123456789"
	ravenClient := raven.NewClient(dsn)
	option := raven.SentryRavenInterceptorOption(ravenClient)
	s := grpc.NewServer(option)

	pb.RegisterMirrorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
