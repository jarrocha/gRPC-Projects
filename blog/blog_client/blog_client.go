package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jarrocha/go_grpc/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world from client")

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("gRPC Dial error. ", err)
	}
	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Jaime Arrocha",
			Title:    "New Blog Post",
			Content:  "Test content for blog.",
		},
	}

	resp, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Println("Error creating post", err)
	}
	fmt.Println("Post created successully!!\n", resp.GetBlog())

}
