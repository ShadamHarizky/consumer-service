package consumer

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConsumer struct {
	client  *redis.Client
	channel string
}

func NewRedisConsumer(client *redis.Client, ch string) *RedisConsumer {
	return &RedisConsumer{
		client:  client,
		channel: ch,
	}
}

func (c *RedisConsumer) ConsumeMessages(callback func(string) error) error {
	pubsub := c.client.Subscribe(context.Background(), c.channel)
	channel := pubsub.Channel()

	go func() {
		for message := range channel {
			log.Printf("Received message: %s", message.Payload)
			err := callback(message.Payload)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()

	return nil
}
