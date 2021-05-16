package user

import (
	"context"
	"errors"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
)

var UserNotFound = errors.New("user is not found")
var UserExists = errors.New("such user already exists")

//go:generate counterfeiter . UserRepo
type UserRepo interface {
	User(ctx context.Context, name string) (*goauser.User, error)
	Users(ctx context.Context) ([]*goauser.User, error)
	Create(ctx context.Context, u *goauser.User) error
	Edit(ctx context.Context, u *goauser.User) error
	Delete(ctx context.Context, name string) error
}

type User struct {
	repo UserRepo
}

// NewUser returns the user service implementation.
func NewUser(repo UserRepo) goauser.Service {
	return &User{
		repo,
	}
}

// Fetch returns User info
func (svc *User) Fetch(ctx context.Context, p *goauser.FetchPayload) (res *goauser.User, err error) {
	user, err := svc.repo.User(ctx, p.Name)
	if err != nil {
		if err == UserNotFound || err == UserExists {
			return nil, goa_errors.NewBadRequestError(err)
		}

		return nil, goa_errors.NewBadRequestError(UserNotFound)
	}

	return user, nil
}

// FetchAll returns all Users info
func (svc *User) FetchAll(ctx context.Context) (res *goauser.Users, err error) {
	users, err := svc.repo.Users(ctx)
	if err != nil {
		return nil, goa_errors.NewInternalServerError(errors.New("cannot get users"))
	}

	return &goauser.Users{Users: users}, nil
}

// Create creates a new user
func (svc *User) Create(ctx context.Context, u *goauser.User) (err error) {
	if err := svc.repo.Create(ctx, u); err != nil {
		if err == UserExists {
			return goa_errors.NewBadRequestError(err)
		}
		return goa_errors.NewBadRequestError(errors.New("user cannot be created"))
	}

	return nil
}

// Edit edits a user
func (svc *User) Edit(ctx context.Context, u *goauser.User) (err error) {
	return svc.repo.Edit(ctx, u)
}

// Delete deletes a user
func (svc *User) Delete(ctx context.Context, p *goauser.DeletePayload) (err error) {
	if err := svc.repo.Delete(ctx, p.Name); err != nil {
		if err == UserNotFound {
			return goa_errors.NewBadRequestError(err)
		}

		return goa_errors.NewInternalServerError(err)
	}

	return nil
}
