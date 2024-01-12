package main

import (
	"log"
	pb "github.com/AzizAhsaan/GolangGRPC/proto"
	"time"
	"errors"
)

func(s *helloServer)  SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("Got request with names : %v", req.Names)
	if stream == nil {
        return errors.New("stream is nil")
    }
	for _, name := range req.Names {
		res := &pb.HelloResponse {
			Message: "Hello" +name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}