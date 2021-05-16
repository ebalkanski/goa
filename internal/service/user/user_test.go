package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
	"github.com/ebalkanski/goa/internal/service/user"
	"github.com/ebalkanski/goa/internal/service/user/userfakes"
)

func TestNewUser(t *testing.T) {
	repo := &userfakes.FakeUserRepo{}
	svc := user.NewUser(repo)
	assert.IsType(t, &user.User{}, svc)
}

func TestFetchFails(t *testing.T) {
	tests := []struct {
		name     string
		payload  *goauser.FetchPayload
		userStub func(context.Context, string) (*goauser.User, error)

		err error
	}{
		{
			name:    "user is not found",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, user.UserNotFound
			},

			err: user.UserNotFound,
		},
		{
			name:    "user exists",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, user.UserExists
			},

			err: user.UserExists,
		},
		{
			name:    "error in repo",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, errors.New("ERROR")
			},

			err: user.UserNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &userfakes.FakeUserRepo{}
			repo.UserStub = test.userStub

			svc := user.NewUser(repo)
			res, err := svc.Fetch(context.Background(), test.payload)

			assert.Nil(t, res)
			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.err.Error(), gerr.Message)
		})
	}
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name     string
		payload  *goauser.FetchPayload
		userStub func(context.Context, string) (*goauser.User, error)

		res *goauser.User
	}{
		{
			name:    "user is found",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return &goauser.User{
					Name: "Bob",
					Age:  22,
				}, nil
			},

			res: &goauser.User{
				Name: "Bob",
				Age:  22,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &userfakes.FakeUserRepo{}
			repo.UserStub = test.userStub

			svc := user.NewUser(repo)
			res, err := svc.Fetch(context.Background(), test.payload)

			assert.Equal(t, test.res, res)
			assert.Nil(t, err)
		})
	}
}
