package main

import pb "github.com/conradwt/grpc-go-course/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
