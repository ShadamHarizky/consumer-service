package repository

import (
	"context"

	"github.com/ShadamHarizky/consumer-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	DB *mongo.Database
}

func NewMessageRepository(db *mongo.Database) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) Save(message *model.Message) error {
	collection := r.DB.Collection("message_consumer")
	_, err := collection.InsertOne(context.TODO(), bson.D{{Key: "content", Value: message.Content}, {Key: "source", Value: message.Source}})
	return err
}
