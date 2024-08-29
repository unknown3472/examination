package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Server_Host string
	Server_Port string
	Mongo_Host string
	Mongo_Port string
	Mongo_Collection string
	Mongo_Database string
	Hotel_Server_Host string
	Hotel_Server_Port string
	Mongo_User_collection string
	Mongo_User_Database string
}

func Load()(*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	cfg.Server_Host=os.Getenv("SERVER_HOST")
	cfg.Server_Port=os.Getenv("SERVER_PORT")
	cfg.Mongo_Database=os.Getenv("MONGO_DATABASE")
	cfg.Mongo_Collection=os.Getenv("MONGO_COLLECTION")
	cfg.Mongo_Host=os.Getenv("MONGO_HOST")
	cfg.Mongo_Port=os.Getenv("MONGO_PORT")
	cfg.Hotel_Server_Host=os.Getenv("HOTEL_SERVER_HOST")
	cfg.Hotel_Server_Port=os.Getenv("HOTEL_SERVER_PORT")
	cfg.Mongo_User_collection=os.Getenv("MONGO_USER_COLLECTION")
	cfg.Mongo_User_Database=os.Getenv("MONGO_USER_DATABASE")
	return &cfg, nil
}