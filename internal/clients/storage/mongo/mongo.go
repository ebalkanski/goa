package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/user"
)

type mongoDB struct {
	logger *log.Logger
	db     string
	*mongo.Client
}

// NewMongo returns new Mongo client
func NewMongo(logger *log.Logger, ctx context.Context, uri string, db string) *mongoDB {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
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
func (m mongoDB) User(ctx context.Context, name string) (*goauser.User, error) {
	users := m.Database(m.db).Collection("users")
	filter := bson.D{{"name", name}}

	var u goauser.User
	if err := users.FindOne(ctx, filter).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, user.UserNotFound
		}

		m.logger.Printf("cannot get user from db: %s\n", err.Error())
		return nil, err
	}

	return &u, nil
}

// CreateUser creates a user
func (m mongoDB) Create(ctx context.Context, u *goauser.User) error {
	ok, err := m.userExists(ctx, u)

	if err != nil {
		return err
	}
	if ok {
		return user.UserExists
	}

	users := m.Database(m.db).Collection("users")
	_, err = users.InsertOne(ctx, u)
	if err != nil {
		m.logger.Printf("cannot insert user into db: %s\n", err.Error())
		return err
	}

	return nil
}

// Edit edits a user
func (m mongoDB) Edit(ctx context.Context, u *goauser.User) error {
	users := m.Database(m.db).Collection("users")
	filter := bson.D{{"name", u.Name}}
	update := bson.D{
		{"$set", bson.D{{"age", u.Age}}},
	}

	updateResult, err := users.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

// Delete deletes a user
func (m mongoDB) Delete(ctx context.Context, name string) error {
	users := m.Database(m.db).Collection("users")
	filter := bson.D{{"name", name}}

	deleteResult, err := users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

func (m mongoDB) userExists(ctx context.Context, u *goauser.User) (bool, error) {
	users := m.Database(m.db).Collection("users")
	filter := bson.D{{"name", u.Name}}

	var user goauser.User
	if err := users.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		m.logger.Printf("cannot get user from db: %s\n", err.Error())
		return false, err
	}

	return true, nil
}
