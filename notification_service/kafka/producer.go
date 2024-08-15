package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	ProduceMessages(topic string, message []byte, correlationID,responseTopic string) error
	Close() error
}

type Producer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string) (KafkaProducer, error) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		AllowAutoTopicCreation: true,
	}
	return &Producer{writer: writer}, nil
}

func (p *Producer) ProduceMessages(topic string, message []byte, correlationID,responseTopic string) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: message,
		Headers: []kafka.Header{
			{Key: "correlation_id", Value: []byte(correlationID)},
			{Key: "response_topic", Value: []byte(responseTopic)},
		},
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
