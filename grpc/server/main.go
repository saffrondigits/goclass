package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	calc := &CalCServer{}
	proto.RegisterCalculatorServiceServer(server, calc)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}

}

type CalCServer struct{}

func (c *CalCServer) Add(context.Context, *proto.AddRequest) (*proto.AddResponse, error) {

}
