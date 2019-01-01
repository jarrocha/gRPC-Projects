package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/jarrocha/go_grpc/blog/blogpb"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type blogServer struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (*blogServer) CreateBlog(ctx context.Context,
	req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Insert error: ", err))
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("ObjectID conversion error"))
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       objID.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func shutdownServer(li net.Listener, s *grpc.Server) {
	log.Println("Stopping Server")
	s.GracefulStop()
	li.Close()
	log.Println("Server Stopped")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Blog Server initializing")

	// start unsecured connection to MongoDB server
	log.Println("Starting database")
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// opens MongDB database and its collection (creates them if they don't exitst)
	collection = client.Database("mydb").Collection("blog")

	// log file creation
	log.Println("Log file setup")

	li, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("Error on port listen. ", err)
	}

	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(srv, &blogServer{})

	go func() {
		log.Println("Blog Server started")
		if err := srv.Serve(li); err != nil {
			log.Fatalln("grpc Server error. ", err)
		}
	}()

	// wait for interrupt signal
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	<-ch
	shutdownServer(li, srv)

}
