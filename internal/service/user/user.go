package user

import (
	"context"
	"log"

	goauser "github.com/ebalkanski/goa/gen/user"
)

type userStorage interface {
	User()
	CreateUser()
}

type user struct {
	logger  *log.Logger
	storage userStorage
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger, storage userStorage) goauser.Service {
	return &user{
		logger,
		storage,
	}
}

// Get returns User info
func (s *user) Get(ctx context.Context, payload *goauser.GetPayload) (res *goauser.User, err error) {
	return &goauser.User{
		Name: "John",
		Age:  33,
	}, nil
}

// Create creates a new user
func (s *user) Create(ctx context.Context, payload *goauser.User) (err error) {
	return nil
}
