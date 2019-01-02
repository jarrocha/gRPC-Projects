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

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/jarrocha/go_grpc/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	log.Println("Create blog request received.")

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

func (*blogServer) ReadBlog(ctx context.Context,
	req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {

	log.Println("Read blog request received.")

	blog_id := req.GetBlogId()
	log.Println(blog_id)

	oid, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse blog id"))
	}

	filter := bson.M{"_id": oid}
	//filter := bsonx.Doc{{Key: "_id", Value: bsonx.ObjectID(oid)}}
	log.Println(filter)

	res := collection.FindOne(context.Background(), filter)
	data := &blogItem{}
	derr := res.Decode(data)
	if derr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln("Cannot find from blog id: ", derr))
	}

	return &blogpb.ReadBlogResponse{
		Blog: dataToBlobpb(data),
	}, nil

}

func (*blogServer) UpdateBlog(ctx context.Context,
	req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {

	log.Println("Update blog request received.")

	blog := req.GetBlog()
	blog_id := blog.GetId()

	oid, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse blog id"))
	}

	// finc by object id
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	data := &blogItem{}

	if derr := res.Decode(data); derr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln("Cannot find from blog id: ", derr))
	}

	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	_, uerr := collection.ReplaceOne(context.Background(), filter, data)
	if uerr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Cannot update object: ", uerr))
	}

	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlobpb(data),
	}, nil
}

func (*blogServer) DeleteBlog(ctx context.Context,
	req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {

	log.Println("Delete blog request received.")

	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln("Cannot parse blog id"))
	}

	filter := bson.M{"_id": oid}

	_, derr := collection.DeleteOne(context.Background(), filter)
	if derr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Cannot delete from blog id: ", derr))
	}

	return &blogpb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
}

func (*blogServer) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {

	log.Println("List  blog request received.")

	cursor, find_err := collection.Find(context.Background(), nil)
	if find_err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintln("Cannot find blogs"))
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &blogItem{}

		decode_err := cursor.Decode(data)
		if decode_err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintln("Cannot decode blog"))
		}

		resp := &blogpb.ListBlogResponse{
			Blog: dataToBlobpb(data),
		}

		stream.Send(resp)
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintln("Cursor error", err))
	}

	return nil
}

func dataToBlobpb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
		Id:       data.ID.Hex(),
	}
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

	// Register reflection service on gRPC server.
	reflection.Register(srv)

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
