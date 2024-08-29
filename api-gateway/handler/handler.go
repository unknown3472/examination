package handler

import (
	"api/config"
	booking "api/genbooking-service"
	hotel "api/genhotel"
	user "api/genuser"
	"api/producer"
	_ "api/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct{
	Hotel hotel.HotelClient
	Booking booking.BookingServiceClient
	User user.UserServiceClient
	Cfg config.Config
	Producer producer.Producer
}

func NewHandler(cfg config.Config)(*Handler, error){
	con_hotel, err := grpc.NewClient(cfg.Hotel_Server_Host+":"+cfg.Hotel_Server_Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	con_booking, err := grpc.NewClient(cfg.Booking_Server_Host+":"+cfg.Booking_Server_Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	con_user, err := grpc.NewClient(cfg.User_Server_Host+":"+cfg.User_server_Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	p := producer.NewProducer(cfg.Kafka_Broker, cfg.Kafka_Topic)

	return &Handler{
		Hotel: hotel.NewHotelClient(con_hotel),
		Booking: booking.NewBookingServiceClient(con_booking),
		User: user.NewUserServiceClient(con_user),
		Cfg: cfg,
		Producer: *p,
	}, nil
}