package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"

	"github.com/ShadamHarizky/consumer-service/consumer"
	"github.com/ShadamHarizky/consumer-service/model"
	"github.com/ShadamHarizky/consumer-service/repository"
	"github.com/ShadamHarizky/consumer-service/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize MongoDB
	mongoDB, err := repository.NewDatabase(os.Getenv("MONGODB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDB.Close()

	db := mongoDB.Client.Database(os.Getenv("DB_NAME"))
	messageRepo := repository.NewMessageRepository(db)
	consumerService := service.NewConsumerService(*messageRepo)

	// Initialize RabbitMQ Consumer
	rabbitMQConn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitMQChannel.Close()

	rabbitMQQueue, err := rabbitMQChannel.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitMQConsumer := consumer.NewRabbitMQConsumer(rabbitMQConn, rabbitMQChannel, rabbitMQQueue)

	// Initialize Redis Consumer
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})
	defer redisClient.Close()

	redisConsumer := consumer.NewRedisConsumer(redisClient, os.Getenv("REDIS_CHANNEL"))

	// Start consuming messages
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		rabbitMQConsumer.ConsumeMessages(func(body []byte) error {
			content := string(body)
			return consumerService.ProcessMessage(model.Message{
				Content: content,
				Source:  "rabbitmq",
			})
		})
	}()

	go func() {
		defer wg.Done()
		redisConsumer.ConsumeMessages(func(payload string) error {
			return consumerService.ProcessMessage(model.Message{
				Content: payload,
				Source:  "redis",
			})
		})
	}()

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Give some time for cleanup before exiting
	time.Sleep(2 * time.Second)

	// Close connections and exit
	log.Println("Shutting down...")
	wg.Wait()
	log.Println("Shutdown complete.")
}
