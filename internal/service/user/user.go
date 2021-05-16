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

type UserStorage interface {
	User(ctx context.Context, name string) (*goauser.User, error)
	Users(ctx context.Context) ([]*goauser.User, error)
	Create(ctx context.Context, u *goauser.User) error
	Edit(ctx context.Context, u *goauser.User) error
	Delete(ctx context.Context, name string) error
}

type user struct {
	logger  *log.Logger
	storage UserStorage
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger, storage UserStorage) goauser.Service {
	return &user{
		logger,
		storage,
	}
}

// Fetch returns User info
func (s *user) Fetch(ctx context.Context, p *goauser.FetchPayload) (res *goauser.User, err error) {
	u, err := s.storage.User(ctx, p.Name)
	if err != nil {
		if err == UserNotFound || err == UserExists {
			return nil, goa_errors.NewBadRequestError(err)
		}

		return nil, goa_errors.NewBadRequestError(UserNotFound)
	}

	return u, nil
}

// FetchAll returns all Users info
func (s *user) FetchAll(ctx context.Context) (res *goauser.FetchAllResult, err error) {
	users, err := s.storage.Users(ctx)
	if err != nil {
		return nil, goa_errors.NewInternalServerError(errors.New("cannot get users"))
	}

	return &goauser.FetchAllResult{Users: users}, nil
}

// Create creates a new user
func (s *user) Create(ctx context.Context, u *goauser.User) (err error) {
	if err := s.storage.Create(ctx, u); err != nil {
		if err == UserExists {
			return goa_errors.NewBadRequestError(err)
		}
		return goa_errors.NewBadRequestError(errors.New("user cannot be created"))
	}

	return nil
}

// Edit edits a user
func (s *user) Edit(ctx context.Context, u *goauser.User) (err error) {
	return s.storage.Edit(ctx, u)
}

// Delete deletes a user
func (s *user) Delete(ctx context.Context, p *goauser.DeletePayload) (err error) {
	if err := s.storage.Delete(ctx, p.Name); err != nil {
		if err == UserNotFound {
			return goa_errors.NewBadRequestError(err)
		}

		return goa_errors.NewInternalServerError(err)
	}

	return nil
}
