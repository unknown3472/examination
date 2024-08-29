package server

import (
	"booking/config"
	pb "booking/genbooking-service"
	"booking/storage"
	"context"
)

type Server struct{
	pb.UnimplementedBookingServiceServer
	storage *storage.BookingRepo
}

func NewServer(cfg config.Config, ctx context.Context)(*Server, error){
	col, col2, err := storage.Connect(&cfg, ctx)
	if err != nil {
		return nil, err
	}
	b, err := storage.NewBookingRepo(*col, *col2, cfg)
	if err != nil {
		return nil, err
	}
	return &Server{storage: b}, nil
}

func (s *Server)AddBooking(ctx context.Context, req *pb.PostBookingRequest)(*pb.BookingResponse, error){
	res, err := s.storage.AddBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)GetBooking(ctx context.Context, req *pb.BookingRequest)(*pb.BookingResponse, error){
	res, err := s.storage.GetBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *Server)GetUserEmail(ctx context.Context, req *pb.UserBookingRequest)(*pb.UserBookingEmail, error){
	res, err := s.storage.GetuserEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)PutBooking(ctx context.Context, req *pb.PutBookingRequest)(*pb.BookingResponse, error){
	res, err := s.storage.UpdateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)DelBooking(ctx context.Context, req *pb.BookingRequest)(*pb.DeleteBookingResponse, error){
	res, err := s.storage.DeleteBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server)GetUserBooking(ctx context.Context, req *pb.UserBookingRequest)(*pb.BookingResponse, error){
	res, err := s.storage.GetUserBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}