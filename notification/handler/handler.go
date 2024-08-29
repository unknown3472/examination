package handler

import (
	"log"
	"net/http"

	// "net/mail"

	"n9/config"
	"n9/consumer"
	"n9/mail"

	"github.com/gorilla/websocket"
)

type Handler struct {
	Consumer *consumer.Consumer
	Mailer   *mail.Mailer
}

func NewHandler(cfg config.Config) *Handler {
	return &Handler{
		Consumer: consumer.NewConsumer(cfg.Broker, cfg.Topic, cfg.Group_Id),
		Mailer:   mail.NewMailer(cfg),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		to, err := h.Consumer.ConsumeMessage()
		if err != nil {
			log.Println("Error consuming message:", err.Error())
			return
		}
		h.Mailer.Mail("booking process ended successfully", to)
		err = conn.WriteMessage(websocket.TextMessage, []byte(to))
		if err != nil {
			log.Println("Error sending message to WebSocket:", err.Error())
			return
		}
	}
}
