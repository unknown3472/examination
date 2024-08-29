package main

import (
	"api/config"
	"api/handler"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_"api/docs"
)


// @title Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	h, err := handler.NewHandler(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := handler.NewGin(h)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		TLSConfig:    &tls.Config{CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}},
	}
	go func() {
		log.Println("Starting HTTPS server...")
		if err := srv.ListenAndServeTLS("./certificate/localhost.pem", "./certificate/localhost-key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Server is running...")

	signalReceived := <-sigChan
	log.Println("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Println("Graceful shutdown complete.")
}


// swag init -g cmd/main.go -o docs
