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

	// create blog
	resp, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Println("Error creating post", err)
	}
	fmt.Println("\nPost created successully!!\n", resp.GetBlog())

	// read blog
	req2 := &blogpb.ReadBlogRequest{BlogId: resp.GetBlog().GetId()}
	resp2, err2 := c.ReadBlog(context.Background(), req2)
	if err2 != nil {
		log.Println("Error reading post", err2)
	} else {
		fmt.Println("\nPost read successully!!\n", resp2)
	}

	// update blog
	req3 := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       resp.GetBlog().GetId(),
			AuthorId: "Deckard Cain",
			Title:    "Deckard Cain Story",
			Content:  "Last of the Horadrim",
		},
	}
	resp3, err3 := c.UpdateBlog(context.Background(), req3)
	if err3 != nil {
		log.Println("\nError reading post", err3)
	} else {
		fmt.Println("\nPost updated successully!!\n", resp3)
	}

}
