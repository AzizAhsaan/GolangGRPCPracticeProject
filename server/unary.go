package main 
import (
	pb "github.com/AzizAhsaan/GolangGRPC/proto"
	"context"
)

func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse {
		Message : "Hello",
	}, nil
}