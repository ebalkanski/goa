package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/user"
)

type User struct {
	logger *log.Logger
	*mongo.Collection
}

// NewUser returns new storage client
func NewUser(logger *log.Logger, collection *mongo.Collection) *User {
	return &User{
		logger,
		collection,
	}
}

// User retrieves a user
func (storage *User) User(ctx context.Context, name string) (*goauser.User, error) {
	filter := bson.D{{"name", name}}

	var u goauser.User
	if err := storage.FindOne(ctx, filter).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, user.UserNotFound
		}

		storage.logger.Printf("cannot get user from db: %s\n", err.Error())
		return nil, err
	}

	return &u, nil
}

// Users retrieves all users
func (storage *User) Users(ctx context.Context) ([]*goauser.User, error) {
	var usersResult []*goauser.User

	cur, err := storage.Find(ctx, bson.D{{}})
	defer cur.Close(ctx)
	if err != nil {

		storage.logger.Printf("cannot get users from db: %s\n", err.Error())
		return nil, err
	}

	for cur.Next(ctx) {
		var u goauser.User
		err := cur.Decode(&u)
		if err != nil {
			storage.logger.Printf("cannot decode user: %s\n", err.Error())
			continue
		}
		usersResult = append(usersResult, &u)
	}

	if err := cur.Err(); err != nil {
		storage.logger.Printf("error while processing users: %s\n", err.Error())
		return nil, err
	}

	return usersResult, nil
}

// Create creates a user
func (storage *User) Create(ctx context.Context, u *goauser.User) error {
	ok, err := storage.userExists(ctx, u)

	if err != nil {
		return err
	}
	if ok {
		return user.UserExists
	}

	_, err = storage.InsertOne(ctx, u)
	if err != nil {
		storage.logger.Printf("cannot insert user into db: %s\n", err.Error())
		return err
	}

	return nil
}

// Edit edits a user
func (storage *User) Edit(ctx context.Context, u *goauser.User) error {
	filter := bson.D{{"name", u.Name}}
	update := bson.D{
		{"$set", bson.D{{"age", u.Age}}},
	}

	updateResult, err := storage.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

// Delete deletes a user
func (storage *User) Delete(ctx context.Context, name string) error {
	filter := bson.D{{"name", name}}

	deleteResult, err := storage.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

func (storage *User) userExists(ctx context.Context, u *goauser.User) (bool, error) {
	filter := bson.D{{"name", u.Name}}

	var user goauser.User
	if err := storage.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		storage.logger.Printf("cannot get user from db: %s\n", err.Error())
		return false, err
	}

	return true, nil
}
