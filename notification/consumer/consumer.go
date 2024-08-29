package consumer

import (
	// "log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


type Consumer struct {
	Broker   string
	Topic    string
	Group_id string
}

func NewConsumer(broker, topic, group_id string) *Consumer {
	return &Consumer{
		Broker:   broker,
		Topic:    topic,
		Group_id: group_id,
	}
}

func (c *Consumer) ConsumeMessage() (to string, err error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": c.Broker,
		"group.id":          c.Group_id,
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return "", err
	}
	err = consumer.Subscribe(c.Topic, nil)
	if err != nil {
		return "", err
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			return string(msg.Value), nil
		} else {
			return "", err
		}
	}
}
