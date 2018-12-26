package main

import (
	"context"
	"fmt"
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
	//fmt.Println("Created client: ", c)
	//fmt.Printf("client content: %f ", c)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Michael",
			LastName:  "Jordan",
		},
	}

	resp, _ := c.Greet(context.Background(), req)
	fmt.Println("Reponse from Server: ", resp.Response)
}
