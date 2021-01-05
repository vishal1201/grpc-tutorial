package main

import (
	"context"
	"fmt"
	"log"

	"../calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Service Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)
	req := &calculatorpb.CalculatorRequest{
		Number1: 5,
		Number2: 12,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum RPC: %v", res.Result)
}
