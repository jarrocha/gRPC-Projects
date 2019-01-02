## Motivation
This is a collection of small project to understand the theory and implementation of gRPC. The main program is the blog program which is a implements gRPC unary and streaming services through a CRUD API to interact with a local MongoDB database. It is still in an initial phase and under improvement.

## Table of Contents
- [Blog Program](#blog-program)
- [Output](#output)
- [Improvements](#improvements)

## What is gRPC?
It is a free and open-source RPC protocol framework developed by Google and other companies. 
It is currently part of the Cloud Native Computation Foundation. It allows for transparent, configurable, 
and customizable communication framework between endpoints. It uses protocol buffers to configure data structures 
and function calls. I can also handle much of the heavy lifting for developers in regards to authentication, 
load balancing, logging, and monitoring. I believe that one of its biggest strength is that it can be used to 
generate code in many languages, so its adoptability has been very fast plus its open-source and it is managed by a 
neutral foundation, just like the Linux Kernel.

## Why gRPC?
- supported by many programming languages
- low latency due to HTTP/2 transport mechanism
- built-in SSL security
- streaming support
- API oriented instead of Resource Oriented like REST

## Blog Program
This is an example of an implementation of a CRUD API interface with MongoDB. 
The driver in use in the official from [mongodb](https://github.com/mongodb/mongo-go-driver/).

#### Protocol Buffer file RPC services:
```
service BlogService {
    rpc CreateBlog (CreateBlogRequest) returns (CreateBlogResponse);
    rpc ReadBlog (ReadBlogRequest) returns (ReadBlogResponse);          // can return NOT_FOUND
    rpc UpdateBlog (UpdateBlogRequest) returns (UpdateBlogResponse);    // can return NOT_FOUND
    rpc DeleteBlog (DeleteBlogRequest) returns (DeleteBlogResponse);    // can return NOT_FOUND
    rpc ListBlog (ListBlogRequest) returns (stream ListBlogResponse);
}
```

## Output
The following output is done using the [evans](https://github.com/ktr0731/evans) cli.
```
default@127.0.0.1:8000> show service
+-------------+------------+-------------------+--------------------+
|   SERVICE   |    RPC     |    REQUESTTYPE    |    RESPONSETYPE    |
+-------------+------------+-------------------+--------------------+
| BlogService | CreateBlog | CreateBlogRequest | CreateBlogResponse |
|             | ReadBlog   | ReadBlogRequest   | ReadBlogResponse   |
|             | UpdateBlog | UpdateBlogRequest | UpdateBlogResponse |
|             | DeleteBlog | DeleteBlogRequest | DeleteBlogResponse |
|             | ListBlog   | ListBlogRequest   | ListBlogResponse   |
+-------------+------------+-------------------+--------------------+
```

Checking current database content:
```
default.BlogService@127.0.0.1:8000> call ListBlog
{
  "blog": {
    "id": "5c2c186c23ba3c54a83c4e9a",
    "authorId": "Jaime Arrocha",
    "title": "New Blog Post",
    "content": "Test content for blog."
  }
}
{
  "blog": {
    "id": "5c2c5552fe22cfc08f5c0f67",
    "authorId": "Deckard Cain",
    "title": "Deckard Cain Story",
    "content": "Last of the Horadrim"
  }
}
```

## Improvements
- Better client interface.
- Better blog presentation.
- Encryption
- Docker and Kubernetes deployment example

