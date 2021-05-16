package mongo_repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	goauser "github.com/ebalkanski/goa/gen/user"
	storage "github.com/ebalkanski/goa/internal/clients/storage/mongo"
	"github.com/ebalkanski/goa/internal/service/user"
)

type UserRepo struct {
	logger *log.Logger
	*mongo.Collection
}

// NewUserRepo returns new repository client
func NewUserRepo(logger *log.Logger, m *storage.MongoDB, db string) *UserRepo {
	collection := m.Database(db).Collection("users")
	return &UserRepo{
		logger,
		collection,
	}
}

// User retrieves a user
func (repo *UserRepo) User(ctx context.Context, name string) (*goauser.User, error) {
	filter := bson.D{{"name", name}}

	var u goauser.User
	if err := repo.FindOne(ctx, filter).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, user.UserNotFound
		}

		repo.logger.Printf("cannot get user from db: %s\n", err.Error())
		return nil, err
	}

	return &u, nil
}

// Users retrieves all users
func (repo *UserRepo) Users(ctx context.Context) ([]*goauser.User, error) {
	var usersResult []*goauser.User

	cur, err := repo.Find(ctx, bson.D{{}})
	defer cur.Close(ctx)
	if err != nil {

		repo.logger.Printf("cannot get users from db: %s\n", err.Error())
		return nil, err
	}

	for cur.Next(ctx) {
		var u goauser.User
		err := cur.Decode(&u)
		if err != nil {
			repo.logger.Printf("cannot decode user: %s\n", err.Error())
			continue
		}
		usersResult = append(usersResult, &u)
	}

	if err := cur.Err(); err != nil {
		repo.logger.Printf("error while processing users: %s\n", err.Error())
		return nil, err
	}

	return usersResult, nil
}

// Create creates a user
func (repo *UserRepo) Create(ctx context.Context, u *goauser.User) error {
	ok, err := repo.userExists(ctx, u)

	if err != nil {
		return err
	}
	if ok {
		return user.UserExists
	}

	_, err = repo.InsertOne(ctx, u)
	if err != nil {
		repo.logger.Printf("cannot insert user into db: %s\n", err.Error())
		return err
	}

	return nil
}

// Edit edits a user
func (repo *UserRepo) Edit(ctx context.Context, u *goauser.User) error {
	filter := bson.D{{"name", u.Name}}
	update := bson.D{
		{"$set", bson.D{{"age", u.Age}}},
	}

	updateResult, err := repo.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

// Delete deletes a user
func (repo *UserRepo) Delete(ctx context.Context, name string) error {
	filter := bson.D{{"name", name}}

	deleteResult, err := repo.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 1 {
		return nil
	}

	return user.UserNotFound
}

func (repo *UserRepo) userExists(ctx context.Context, u *goauser.User) (bool, error) {
	filter := bson.D{{"name", u.Name}}

	var user goauser.User
	if err := repo.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		repo.logger.Printf("cannot get user from db: %s\n", err.Error())
		return false, err
	}

	return true, nil
}
