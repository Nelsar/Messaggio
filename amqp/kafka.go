package amqp

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"messaggio.com/configuration"
	"messaggio.com/models"
)

func ProducerHanler(context context.Context, event models.Event) error {

	kafkaWriter := getKafkaWriter()

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("%d", event.UserId)),
		Value: []byte(fmt.Sprintf("%v", event)),
	}

	err := kafkaWriter.WriteMessages(context, msg)

	return err

}

func getKafkaWriter() *kafka.Writer {
	c := configuration.GetConfiguration()

	return &kafka.Writer{
		Addr:     kafka.TCP(c.KafkaUrl),
		Topic:    c.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

}
