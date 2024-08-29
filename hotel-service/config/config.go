package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Mongo_Host string
	Mongo_Port string
	Mongo_Database string
	Server_Host string
	Server_Port string
	Room_Collection string
	Description_Collection string
	Hotel_Collectin string
}

func Load(filepath string)(*Config, error){
	err := godotenv.Load(filepath)
	if err != nil {
		return nil, err
	}
	conf := Config{}
	conf.Mongo_Host=os.Getenv("MONGO_HOST")
	conf.Mongo_Port=os.Getenv("MONGO_PORT")
	conf.Mongo_Database=os.Getenv("MONGO_DATABASE")
	conf.Server_Host=os.Getenv("SERVER_HOST")
	conf.Server_Port=os.Getenv("SERVER_PORT")
	conf.Description_Collection=os.Getenv("DESCRIPTION_COLLECTION")
	conf.Hotel_Collectin=os.Getenv("HOTEL_COLLECTION")
	conf.Room_Collection=os.Getenv("ROOM_COLLECTION")
	return &conf, nil
}