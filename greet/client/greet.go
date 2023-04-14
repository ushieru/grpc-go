package main

import (
	"context"
	"log"

	pb "github.com/ushieru/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet wa invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ushieru",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
