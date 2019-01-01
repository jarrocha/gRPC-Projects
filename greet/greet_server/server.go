package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/jarrocha/go_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with ", req)

	firstname := req.GetGreeting().GetFirstName()
	greetresponse := "Hello " + firstname

	resp := &greetpb.GreetResponse{Response: greetresponse}

	return resp, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Println("GreetManyTimes function was invoked with ", req)

	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		response := "Hello " + firstname + "reponse #" + strconv.Itoa(i+1)
		res := &greetpb.GreetManytimesResponse{Response: response}
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil
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
