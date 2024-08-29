package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	Broker string
	Topic  string
}

func NewProducer(b string, t string) *Producer {
	return &Producer{
		Broker: b,
		Topic:  t,
	}
}

func (p *Producer) ProduceMessage(message string) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": p.Broker})
	if err != nil {
		return err
	}
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("event partition error: %v\n", ev.TopicPartition)
				} else {
					log.Printf("delivered message : %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return err
	}
	producer.Flush(15 * 1000)
	return nil
}
