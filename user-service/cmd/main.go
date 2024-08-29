package main

import (
	"context"
	"log"
	"net"
	"user/config"
	pb "user/genuser"
	"user/server"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user service listening on %s:%s", cfg.Server_Host, cfg.Server_Port)
	lis, err := net.Listen("tcp", cfg.Server_Host+":"+cfg.Server_Port)
	if err != nil {
		log.Fatal(err)
	}
	serve := grpc.NewServer()
	s, err := server.NewServer(context.Background(), *cfg)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterUserServiceServer(serve, s)
	serve.Serve(lis)
}