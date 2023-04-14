package main

import (
	"context"
	"log"

	pb "github.com/ushieru/grpc-go-course/greet/proto"
)

func (s *Server) Greet(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function wa invoked %#+v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
