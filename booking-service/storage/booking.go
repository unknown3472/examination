package storage

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"booking/config"
	pb "booking/genbooking-service"
	hotel "booking/genhotel"

	user "booking/genuser"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingRepo struct {
	col   mongo.Collection
	col_user mongo.Collection
	hotel hotel.HotelClient
}

func NewBookingRepo(col mongo.Collection, col2 mongo.Collection, cfg config.Config) (*BookingRepo, error) {
	con, err := grpc.NewClient(cfg.Hotel_Server_Host+":"+cfg.Hotel_Server_Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := hotel.NewHotelClient(con)
	return &BookingRepo{col: col, col_user: col2, hotel: client}, nil
}

func (b *BookingRepo) AddBooking(ctx context.Context, req *pb.PostBookingRequest) (*pb.BookingResponse, error) {
	r := hotel.DescriptionRequest{HotelId: req.HotelId}
	res, err := b.hotel.DescriptionHotel(ctx, &r)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key: "userid",Value: req.UserId,
	}}
	user_res := user.UserResponse{}
	err = b.col_user.FindOne(ctx, filter).Decode(&user_res)
	if err != nil {
		return nil, err
	}

	response := pb.BookingResponse{}
	namespace := uuid.New()
	uuidV5 := uuid.NewSHA1(namespace, []byte(req.UserId))
	response.BookingId = uuidV5.String()

	ch := 0
	if res.HotelId == req.HotelId {
		for _, i := range res.Rooms {
			if i.RoomType == req.RoomType {
				diff, err := Difference(req.CheckInDate, req.CheckOutDate)
				if err != nil {
					return nil, err
				}
				response.TotalAmount = int32((*diff) * int(i.PricePerNight))
				response.CheckInDate = req.CheckInDate
				response.CheckOutDate = req.CheckOutDate
				response.RoomType = req.RoomType
				response.HotelId = req.HotelId
				response.UserId = req.UserId
				response.Status = "approved"
				_, err = b.col.InsertOne(ctx, response)
				if err != nil {
					return nil, err
				}
				ch = 1
			}
		}
	}

	if ch == 0 {
		return nil, errors.New("do not matching with data of hotel")
	}

	return &response, nil
}

func (b *BookingRepo) UpdateBooking(ctx context.Context, req *pb.PutBookingRequest) (*pb.BookingResponse, error) {
	filter := bson.D{{
		Key: "bookingid", Value: req.BookingId,
	}}

	var booking pb.BookingResponse
	b.col.FindOne(ctx, filter).Decode(&booking)

	ch := 0
	r := hotel.DescriptionRequest{HotelId: booking.HotelId}
	res, err := b.hotel.DescriptionHotel(ctx, &r)
	if err != nil {
		return nil, err
	}
	var totalAmount int32
	if res.HotelId == booking.HotelId {
		for _, i := range res.Rooms {
			if i.RoomType == booking.RoomType {
				diff, err := Difference(req.CheckInDate, req.CheckOutDate)
				if err != nil {
					return nil, err
				}
				totalAmount = int32((*diff) * int(i.PricePerNight))
				ch = 1
			}
		}
	}
	if ch == 0 {
		return nil, errors.New("do not matching with data of hotel")
	}

	upd := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "checkindate", Value: req.CheckInDate},
			{Key: "checkoutdate", Value: req.CheckOutDate},
			{Key: "totalamount", Value: totalAmount},
			{Key: "status", Value: req.Status},
		},
	}}
	_, err = b.col.UpdateOne(ctx, filter, upd)
	if err != nil {
		return nil, err
	}
	var response pb.BookingResponse
	err = b.col.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (b *BookingRepo) GetBooking(ctx context.Context, req *pb.BookingRequest) (*pb.BookingResponse, error) {
	filter := bson.D{{
		Key: "bookingid", Value: req.BookingId,
	}}
	res := pb.BookingResponse{}
	err := b.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (b *BookingRepo) GetUserBooking(ctx context.Context, req *pb.UserBookingRequest) (*pb.BookingResponse, error) {
	filter := bson.D{{
		Key: "userid", Value: req.UserId,
	}}
	var res pb.BookingResponse
	err := b.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}


func (b *BookingRepo)GetuserEmail(ctx context.Context, req *pb.UserBookingRequest)(*pb.UserBookingEmail, error){
	filter := bson.D{{
		Key: "userid",Value: req.UserId,
	}}
	res := user.UserResponse{}
	err := b.col_user.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &pb.UserBookingEmail{Email: res.Email}, nil
}

func (b *BookingRepo) DeleteBooking(ctx context.Context, req *pb.BookingRequest) (*pb.DeleteBookingResponse, error) {
	filter := bson.D{{
		Key: "bookingid", Value: req.BookingId,
	}}
	_, err := b.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookingResponse{Message: "deleted", BookingId: req.GetBookingId()}, nil
}

func Difference(time1, time2 string) (*int, error) {
	layout := "2006-01-02"
	date1, err := time.Parse(layout, time1)
	if err != nil {
		return nil, err
	}
	date2, err := time.Parse(layout, time2)
	if err != nil {
		return nil, err
	}
	duration := date2.Sub(date1)
	day := int(duration.Hours() / 24)
	return &day, nil
}
