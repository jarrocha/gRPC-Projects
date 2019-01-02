package main

import (
	"context"
	"fmt"
	"io"
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
	blog_id := resp.GetBlog().GetId()
	req3 := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       blog_id,
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

	// delete blog
	req4 := &blogpb.DeleteBlogRequest{
		BlogId: blog_id,
	}
	resp4, err4 := c.DeleteBlog(context.Background(), req4)
	if err4 != nil {
		log.Println("\nError deleting post", err4)
	} else {
		fmt.Println("\nPost deleted successully!!\n", resp4)
	}

	// list blog
	listStream, err5 := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err5 != nil {
		log.Println("\nError listing blogs", err5)
	} else {
		fmt.Println("\nListing all blogs in database")
		for {
			resp, recv_err := listStream.Recv()
			if recv_err != nil {
				if recv_err == io.EOF {
					break
				}
				log.Println("ListBlog RPC Stream receive error", recv_err)
			}

			if resp != nil {
				fmt.Println(resp.GetBlog())
			}
		}
	}

}
