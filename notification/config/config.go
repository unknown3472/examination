package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Server_Host string
	Server_Port string
	Broker string
	Topic string
	Group_Id string
	Email_Sender string
	Email_Password string
}


func Load()(*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	cfg.Server_Host=os.Getenv("SERVER_HOST")
	cfg.Server_Port=os.Getenv("SERVER_PORT")
	cfg.Broker=os.Getenv("BROKER")
	cfg.Topic=os.Getenv("TOPIC")
	cfg.Group_Id=os.Getenv("GROUP_ID")
	cfg.Email_Sender=os.Getenv("EMAIL_SENDER")
	cfg.Email_Password=os.Getenv("EMAIL_PASSWORD")
	return &cfg, nil
}