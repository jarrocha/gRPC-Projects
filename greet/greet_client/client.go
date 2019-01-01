package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/jarrocha/go_grpc/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world from client")

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("gRPC Dial error. ", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	// requesting unary RPC service
	doUnary(c)

	// requesting streaming RPC service
	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("\nStarting Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Michael",
			LastName:  "Jordan",
		},
	}

	resStream, err1 := c.GreetManyTimes(context.Background(), req)
	if err1 != nil {
		log.Println("doServerStreaming GreetManyTimes RPC error", err1)

	}
	for {
		msg, err2 := resStream.Recv()
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
			log.Println("doServerStreaming Stream receive error", err2)
		}

		if msg != nil {
			log.Println(msg)
		}
	}

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("\nStarting Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Michael",
			LastName:  "Jordan",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Println("doUnary Greet RPC error", err)

	}
	log.Println("Reponse from Server: ", resp.Response)
}
