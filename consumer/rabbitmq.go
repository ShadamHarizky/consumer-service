package consumer

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQConsumer(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue) *RabbitMQConsumer {
	return &RabbitMQConsumer{conn: conn, channel: channel, queue: queue}
}

func (c *RabbitMQConsumer) ConsumeMessages(callback func([]byte) error) error {
	messages, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for message := range messages {
			err := callback(message.Body)
			log.Printf("Received message: %s", message.Body)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()

	return nil
}
