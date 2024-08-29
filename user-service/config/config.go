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
	Mongo_Database string
	Mongo_Collection string
	Jwt_Secret_Key string
	Expiration_Time string
}

func Load()(*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	cfg.Server_Host=os.Getenv("SERVER_HOST")
	cfg.Server_Port=os.Getenv("SERVER_PORT")
	cfg.Mongo_Host=os.Getenv("MONGO_HOST")
	cfg.Mongo_Port=os.Getenv("MONGO_PORT")
	cfg.Mongo_Database=os.Getenv("MONGO_DATABASE")
	cfg.Mongo_Collection=os.Getenv("MONGO_COLLECTION")
	cfg.Jwt_Secret_Key=os.Getenv("JWT_SECRET_KEY")
	cfg.Expiration_Time=os.Getenv("EXPIRATION_TIME")
	return &cfg, nil
}