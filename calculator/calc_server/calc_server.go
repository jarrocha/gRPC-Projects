package main

import (
	"context"
	"log"
	"net"

	"github.com/jarrocha/grpc-go/calculator/calcpb"
	"google.golang.org/grpc"
)

func performSum(x, y float32) float32 {
	return x + y
}

func performSub(x, y float32) float32 {
	return x - y
}

func performMul(x, y float32) float32 {
	return x * y
}

func performDiv(x, y float32) float32 {
	if y == 0 {
		return 0
	}

	return x / y
}

type calcService struct{}

func (c *calcService) Calculate(ctx context.Context, in *calcpb.OperRequest) (*calcpb.OperRespond, error) {
	var result float32
	resp := &calcpb.OperRespond{Result: result}
	oper := in.GetOperation().GetOperator()
	num1 := in.GetOperation().GetNumber1()
	num2 := in.GetOperation().GetNumber2()

	switch oper {
	case 1:
		resp.Result = performSum(num1, num2)
	case 2:
		resp.Result = performSub(num1, num2)
	case 3:
		resp.Result = performMul(num1, num2)
	case 4:
		resp.Result = performDiv(num1, num2)
	}

	return resp, nil
}

func main() {

	li, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("Error on port listen. ", err)
	}

	srv := grpc.NewServer()
	calcpb.RegisterOperServiceServer(srv, &calcService{})

	if err := srv.Serve(li); err != nil {
		log.Fatalln("grpc Serve error. ", err)
	}

}
