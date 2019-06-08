package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sat0yu/sentry-raven-grpc/example/mirror"
	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

type server struct{}

// Echo implements mirror.Mirror
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.RequestBody)
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
	s := grpc.NewServer()
	pb.RegisterMirrorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
