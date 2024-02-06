package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func NewDatabase(uri string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Database{Client: client}, nil
}

func (db *Database) Close() {
	if db.Client != nil {
		db.Client.Disconnect(context.TODO())
	}
}
