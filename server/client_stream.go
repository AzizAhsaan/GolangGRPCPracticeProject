package main

import(
	pb "github.com/AzizAhsaan/GolangGRPC/proto"
	"io"

	"log"

)

func (s *helloServer) SayhelloClientStreaming(stream pb.GreetService_SayhelloClientStreamingServer) error {
	var messages []string
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with names: %v", req.Name)
		messages=append(messages, "Hello", req.Name)
	}
}