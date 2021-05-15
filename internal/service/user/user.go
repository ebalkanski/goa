package user

import (
	"context"
	"errors"
	"log"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
)

var UserNotFound = errors.New("user is not found")
var UserExists = errors.New("such user already exists")

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
		if err == UserNotFound || err == UserExists {
			return nil, goa_errors.NewBadRequestError(err)
		}

		return nil, goa_errors.NewBadRequestError(UserNotFound)
	}

	return u, nil
}

// Create creates a new user
func (s *user) Create(ctx context.Context, u *goauser.User) (err error) {
	if err := s.storage.CreateUser(ctx, u); err != nil {
		if err == UserExists {
			return goa_errors.NewBadRequestError(err)
		}
		return goa_errors.NewBadRequestError(errors.New("user cannot be created"))
	}

	return nil
}
