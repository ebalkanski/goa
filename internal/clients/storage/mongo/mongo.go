package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDB struct {
	logger *log.Logger
	db     string
	*mongo.Client
}

// NewMongo returns new Mongo client
func NewMongo(logger *log.Logger, ctx context.Context, uri string, db string) *mongoDB {
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal(err)
	}

	return &mongoDB{
		logger,
		db,
		client,
	}
}

// User retrieves a user
func (m mongoDB) User() {
	m.logger.Println("implement me")
}

// CreateUser creates a user
func (m mongoDB) CreateUser() {
	m.logger.Println("implement me")
}
