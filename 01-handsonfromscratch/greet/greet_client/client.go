package main

import (
	"fmt"
	"log"

	"../../cognologix.com/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, just a client here.")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
