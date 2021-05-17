package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	logger *log.Logger
	*mongo.Client
}

// NewMongo returns new Mongo client
func NewMongo(logger *log.Logger, ctx context.Context, uri string) *MongoDB {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal(err)
	}

	return &MongoDB{
		logger,
		client,
	}
}
