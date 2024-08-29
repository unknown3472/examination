package main

import (
	"log"
	"n9/config"
	"n9/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(*cfg)
	r := handler.NewApi(h)
	r.Run(cfg.Server_Host + ":" + cfg.Server_Port)
}
