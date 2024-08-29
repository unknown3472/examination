package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Hotel_Server_Host string
	Hotel_Server_Port string
	Booking_Server_Host string
	Booking_Server_Port string
	User_Server_Host string
	User_server_Port string
	JWT_key string
	Kafka_Broker string
	Kafka_Topic string
}

func Load()(*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	cfg.Hotel_Server_Host=os.Getenv("HOTEL_SERVER_HOST")
	cfg.Hotel_Server_Port=os.Getenv("HOTEL_SERVER_PORT")
	cfg.Booking_Server_Host=os.Getenv("BOOKING_SERVER_HOST")
	cfg.Booking_Server_Port=os.Getenv("BOOKING_SERVER_PORT")
	cfg.User_Server_Host=os.Getenv("USER_SERVER_HOST")
	cfg.User_server_Port=os.Getenv("USER_SERVER_PORT")
	cfg.JWT_key=os.Getenv("JWT_KEY")
	cfg.Kafka_Broker=os.Getenv("KAFKA_BROKER")
	cfg.Kafka_Topic=os.Getenv("KAFKA_TOPIC")
	return &cfg, nil
}