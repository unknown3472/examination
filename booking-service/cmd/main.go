package main

import (
	"booking/config"
	pb "booking/genbooking-service"
	"booking/server"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("booking service listening on %s:%s", cfg.Server_Host, cfg.Server_Port)
	lis, err := net.Listen("tcp", cfg.Server_Host+":"+cfg.Server_Port)
	if err != nil {
		log.Fatal(err)
	}
	s, err := server.NewServer(*cfg, context.Background())
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterBookingServiceServer(server, s)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}