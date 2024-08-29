package server

import (
	"context"
	"hotel/config"
	pb "hotel/genhotel"
	"hotel/storage"
)

type Server struct{
	pb.UnimplementedHotelServer
	repo storage.StorageRepo
}

func NewServer(cfg config.Config)(*Server, error){
	col1, col2, col3, err := storage.Connect(cfg)
	if err != nil {
		return nil, err
	}
	return &Server{repo: *storage.NewStorageRepo(col1, col2, col3)}, nil
}

func (s *Server)GetHotel(ctx context.Context, req *pb.VoidHotel)(*pb.GetHotelResponse, error){
	res, err := s.repo.GetHotel(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)DescriptionHotel(ctx context.Context, req *pb.DescriptionRequest)(*pb.DescriptionResponse, error){
	res, err := s.repo.GetDescription(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)GetRooms(ctx context.Context, req *pb.DescriptionRequest)(*pb.RoomAvailability, error){
	res, err := s.repo.GetRooms(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}