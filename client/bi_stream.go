package main

import (
	"context"
	"log"
	"time"
	"io"
	pb "github.com/AzizAhsaan/GolangGRPC/proto"
)

func callHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidrectional Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil{
		log.Fatalf("could not send anmes: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for{
			message, err:=stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil{
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name:= range names.Names {
		req := &pb.HelloRequest {
			Name:name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("Error whiule sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidrectional streaming finished")
}