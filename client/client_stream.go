package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AzizAhsaan/GolangGRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming started")
	stream,err := client.SayhelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names :%v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name :name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("send the request with same :%s", name)
		time.Sleep(2 &time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("Error whiel receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}