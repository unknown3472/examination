package server

import (
	"context"
	"user/config"
	pb "user/genuser"
	"user/storage"
)

type Server struct{
	pb.UnimplementedUserServiceServer
	storage *storage.UserRepo
}

func NewServer(ctx context.Context, cfg config.Config)(*Server, error){
	storage, err := storage.NewUserRepo(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &Server{
		storage: storage,
	}, nil
}


func (s *Server)Register(ctx context.Context, req *pb.UserRequest)(*pb.UserResponse, error){
	res, err := s.storage.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)Login(ctx context.Context, req *pb.LoginRequest)(*pb.LoginResponse, error){
	res, err := s.storage.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}