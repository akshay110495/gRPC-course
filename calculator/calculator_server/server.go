package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/akshay110495/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	firstOperand := req.GetOpearnds().GetFirstOp()
	secondOperand := req.GetOpearnds().GetSecondOp()
	result := firstOperand + secondOperand
	response := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
