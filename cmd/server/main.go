package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "demo-grpc/pb/user"
)

type Server struct {
}

func (s *Server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Status: "success",
		Users: []*pb.Model{
			{
				Id:   "001",
				Name: "Quang Dang",
			},
			{
				Id:   "002",
				Name: "Tony",
			},
		},
	}, nil
}

func (s *Server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &Server{})
	srv.Serve(lis)
}
