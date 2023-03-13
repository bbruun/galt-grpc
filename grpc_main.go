package main

import (
	"log"
	"net"
	"sync"

	minion "github.com/bbruun/galt/proto"
	"github.com/bbruun/galt/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func registerMinionServices(grpcServer *grpc.Server) {
	reflection.Register(grpcServer) // enable reflection from client side
	service := &myMinionServer{}
	minion.RegisterMinionServiceServer(grpcServer, service)

}

func gRPCServer(wg sync.WaitGroup) {
	defer wg.Done()
	grpcListener, err := net.Listen("tcp", server.Configuration.GetListenAddress())
	if err != nil {
		log.Fatalf("error setting up net.Listen(): %s", err)
	}

	grpcServer := grpc.NewServer()
	registerMinionServices(grpcServer)

	err = grpcServer.Serve(grpcListener)
	if err != nil {
		log.Fatalf("failed to start gRPC server on port 4505: %s", err)
	}

}
