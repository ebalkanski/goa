package user

import (
	"context"
	"errors"
	"log"

	goauser "github.com/ebalkanski/goa/gen/user"
)

var UserNotFound = errors.New("user is not found")

type userStorage interface {
	User(ctx context.Context, name string) (*goauser.User, error)
	CreateUser(ctx context.Context, u *goauser.User) error
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
func (s *user) Get(ctx context.Context, p *goauser.GetPayload) (res *goauser.User, err error) {
	u, err := s.storage.User(ctx, p.Name)
	if err != nil {
		if err == UserNotFound {
			return nil, err
		}

		return nil, err
	}

	return u, nil
}

// Create creates a new user
func (s *user) Create(ctx context.Context, u *goauser.User) (err error) {
	if err := s.storage.CreateUser(ctx, u); err != nil {
		return err
	}
	return nil
}
