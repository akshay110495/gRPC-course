package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akshay110495/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a clinet")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	req := &calculatorpb.CalculatorRequest{
		Opearnds: &calculatorpb.Operands{
			FirstOp:  3,
			SecondOp: 10,
		},
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error wile calling greet RPC: %v\n", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}
