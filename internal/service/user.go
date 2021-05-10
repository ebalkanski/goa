package service

import (
	"context"
	"log"

	"github.com/ebalkanski/goa/gen/user"
)

type usersvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersvc{logger}
}

func (s *usersvc) Get(ctx context.Context, payload *user.GetPayload) (res *user.User, err error) {
	return &user.User{
		Name: "John",
		Age:  33,
	}, nil
}

func (s *usersvc) Create(ctx context.Context, payload *user.User) (err error) {
	return nil
}
