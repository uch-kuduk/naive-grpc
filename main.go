package main

import (
	"context"
	"fmt"
	"github.com/uch-kuduk/naive-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type MyServer struct {
	pb.UnimplementedSampleServerServer // по дефолту на MyServer получаем все методы
}

func (MyServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Name: "Hello " + req.Name}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	s := MyServer{}
	pb.RegisterSampleServerServer(grpcServer, s)
	reflection.Register(grpcServer) // ОБЛЕГЧАЕТ работу c evans
	lis, _ := net.Listen("tcp", ":8081")
	fmt.Printf("grpc on port 8081")
	grpcServer.Serve(lis)
}
