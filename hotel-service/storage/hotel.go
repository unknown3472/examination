package storage

import (
	"context"
	pb "hotel/genhotel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageRepo struct {
	col1 *mongo.Collection
	col2 *mongo.Collection
	col3 *mongo.Collection
}

func NewStorageRepo(col *mongo.Collection, col2 *mongo.Collection, col3 *mongo.Collection) *StorageRepo {
	return &StorageRepo{
		col1: col,
		col2: col2,
		col3: col3,
	}
}

func (s *StorageRepo) GetHotel(ctx context.Context) (*pb.GetHotelResponse, error) {
	cursor, err := s.col1.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var hotels []*pb.HotelInfo
	for cursor.Next(ctx) {
		var hotel *pb.HotelInfo
		if err := cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetHotelResponse{Hotels: hotels}, nil
}

func (s *StorageRepo) GetDescription(ctx context.Context, id *pb.DescriptionRequest) (*pb.DescriptionResponse, error) {
	var res pb.DescriptionResponse
	filter := bson.D{{
		Key: "hotelId", Value: id.GetHotelId(),
	}}
	err := s.col2.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *StorageRepo) GetRooms(ctx context.Context, id *pb.DescriptionRequest) (*pb.RoomAvailability, error) {
	filter := bson.D{{
		Key: "hotelId", Value: id.GetHotelId(),
	}}
	cursor, err := s.col3.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*pb.RoomInfo
	for cursor.Next(ctx){
		var room *pb.RoomInfo
		if err := cursor.Decode(&room); err!=nil{
			return nil, err
		}
		rooms=append(rooms, room)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &pb.RoomAvailability{Rooms:rooms}, nil
}
