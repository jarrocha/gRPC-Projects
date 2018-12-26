package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jarrocha/go_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with ", in)

	firstname := in.GetGreeting().GetFirstName()
	greetresponse := "Hello " + firstname

	resp := &greetpb.GreetResponse{Response: greetresponse}

	return resp, nil
}

func main() {
	li, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("Error on listen port. ", err)
	}
	defer li.Close()

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(li); err != nil {
		log.Fatalln("gRPC failed Serve error. ", err)
	}

}
