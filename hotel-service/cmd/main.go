package main

import (
	"hotel/config"
	"hotel/server"
	"log"
	"net"
	"path/filepath"

	pb "hotel/genhotel"

	"google.golang.org/grpc"
)

func main() {
	root := filepath.Join(".env")
	cfg, err := config.Load(root)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("hotel service listening on %s:%s", cfg.Server_Host, cfg.Server_Port)
	lis, err := net.Listen("tcp", cfg.Server_Host+":"+cfg.Server_Port)
	if err != nil {
		log.Fatal(err)
	}
	grpc_server := grpc.NewServer()
	server, err := server.NewServer(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterHotelServer(grpc_server, server)
	grpc_server.Serve(lis)
}
