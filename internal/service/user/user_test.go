package user_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
	"github.com/ebalkanski/goa/internal/service/user"
	"github.com/ebalkanski/goa/internal/service/user/userfakes"
)

func userBob() *goauser.User {
	return &goauser.User{
		Name: "Bob",
		Age:  22,
	}
}

func TestNewUser(t *testing.T) {
	storage := &userfakes.FakeUserStorage{}
	svc := user.NewUser(storage)
	assert.IsType(t, &user.User{}, svc)
}

func TestFetchFails(t *testing.T) {
	tests := []struct {
		name     string
		payload  *goauser.FetchPayload
		userStub func(context.Context, string) (*goauser.User, error)

		errText string
	}{
		{
			name:    "user is not found",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, user.UserNotFound
			},

			errText: user.UserNotFound.Error(),
		},
		{
			name:    "user exists",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, user.UserExists
			},

			errText: user.UserExists.Error(),
		},
		{
			name:    "error in storage",
			payload: &goauser.FetchPayload{Name: "Bob"},
			userStub: func(ctx context.Context, s string) (*goauser.User, error) {
				return nil, errors.New("ERROR")
			},

			errText: user.UserNotFound.Error(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.UserStub = test.userStub

			svc := user.NewUser(storage)
			res, err := svc.Fetch(context.Background(), test.payload)

			assert.Nil(t, res)
			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.errText, gerr.Message)
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
				return userBob(), nil
			},

			res: userBob(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.UserStub = test.userStub

			svc := user.NewUser(storage)
			res, err := svc.Fetch(context.Background(), test.payload)

			assert.Equal(t, test.res, res)
			assert.Nil(t, err)
		})
	}
}

func TestFetchAllFails(t *testing.T) {
	tests := []struct {
		name      string
		usersStub func(context.Context) ([]*goauser.User, error)

		errText string
	}{
		{
			name: "cannot get users from storage",
			usersStub: func(ctx context.Context) ([]*goauser.User, error) {
				return nil, errors.New("ERROR")
			},

			errText: "cannot get users",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.UsersStub = test.usersStub

			svc := user.NewUser(storage)
			res, err := svc.FetchAll(context.Background())

			assert.Nil(t, res)
			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.errText, gerr.Message)
			assert.Equal(t, http.StatusInternalServerError, gerr.StatusCode())
		})
	}
}

func TestFetchAll(t *testing.T) {
	tests := []struct {
		name      string
		usersStub func(context.Context) ([]*goauser.User, error)

		res *goauser.Users
	}{
		{
			name: "get users from storage",
			usersStub: func(ctx context.Context) ([]*goauser.User, error) {
				return []*goauser.User{
					userBob(),
				}, nil
			},

			res: &goauser.Users{
				Users: []*goauser.User{
					userBob(),
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.UsersStub = test.usersStub

			svc := user.NewUser(storage)
			res, err := svc.FetchAll(context.Background())

			assert.Equal(t, test.res, res)
			assert.Nil(t, err)
		})
	}
}

func TestCreateFails(t *testing.T) {
	tests := []struct {
		name       string
		createStub func(context.Context, *goauser.User) error
		user       *goauser.User

		errText string
	}{
		{
			name: "such user exists on create",
			createStub: func(context.Context, *goauser.User) error {
				return user.UserExists
			},
			user: userBob(),

			errText: user.UserExists.Error(),
		},
		{
			name: "storage error",
			createStub: func(context.Context, *goauser.User) error {
				return errors.New("ERROR")
			},
			user: userBob(),

			errText: "user cannot be created",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.CreateStub = test.createStub

			svc := user.NewUser(storage)
			err := svc.Create(context.Background(), test.user)

			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.errText, gerr.Message)
			assert.Equal(t, http.StatusBadRequest, gerr.StatusCode())
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name       string
		createStub func(context.Context, *goauser.User) error
		user       *goauser.User
	}{
		{
			name: "user is created",
			createStub: func(context.Context, *goauser.User) error {
				return nil
			},
			user: userBob(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.CreateStub = test.createStub

			svc := user.NewUser(storage)
			err := svc.Create(context.Background(), test.user)

			assert.Nil(t, err)
		})
	}
}

func TestEditFails(t *testing.T) {
	tests := []struct {
		name     string
		editStub func(context.Context, *goauser.User) error
		user     *goauser.User

		errText string
	}{
		{
			name: "such user exists on edit",
			editStub: func(context.Context, *goauser.User) error {
				return user.UserNotFound
			},
			user: userBob(),

			errText: user.UserNotFound.Error(),
		},
		{
			name: "storage error",
			editStub: func(context.Context, *goauser.User) error {
				return errors.New("ERROR")
			},
			user: userBob(),

			errText: "user cannot be edited",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.EditStub = test.editStub

			svc := user.NewUser(storage)
			err := svc.Edit(context.Background(), test.user)

			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.errText, gerr.Message)
			assert.Equal(t, http.StatusBadRequest, gerr.StatusCode())
		})
	}
}

func TestEdit(t *testing.T) {
	tests := []struct {
		name     string
		editStub func(context.Context, *goauser.User) error
		user     *goauser.User
	}{
		{
			name: "user is edited",
			editStub: func(context.Context, *goauser.User) error {
				return nil
			},
			user: userBob(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.EditStub = test.editStub

			svc := user.NewUser(storage)
			err := svc.Edit(context.Background(), test.user)

			assert.Nil(t, err)
		})
	}
}

func TestDeleteFails(t *testing.T) {
	tests := []struct {
		name       string
		payload    *goauser.DeletePayload
		deleteStub func(context.Context, string) error
		user       *goauser.User

		errText string
		status  int
	}{
		{
			name:    "user not found on delete",
			payload: &goauser.DeletePayload{Name: "Bob"},
			deleteStub: func(context.Context, string) error {
				return user.UserNotFound
			},
			user: userBob(),

			errText: user.UserNotFound.Error(),
			status:  http.StatusBadRequest,
		},
		{
			name:    "storage error on delete",
			payload: &goauser.DeletePayload{Name: "Bob"},
			deleteStub: func(context.Context, string) error {
				return errors.New("ERROR")
			},
			user: userBob(),

			errText: "user cannot be deleted",
			status:  http.StatusInternalServerError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.DeleteStub = test.deleteStub

			svc := user.NewUser(storage)
			err := svc.Delete(context.Background(), test.payload)

			assert.IsType(t, &goa_errors.Error{}, err)
			gerr, ok := err.(*goa_errors.Error)
			assert.True(t, ok)
			assert.Equal(t, test.errText, gerr.Message)
			assert.Equal(t, test.status, gerr.StatusCode())
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name       string
		payload    *goauser.DeletePayload
		deleteStub func(context.Context, string) error
		user       *goauser.User
	}{
		{
			name:    "user is edited",
			payload: &goauser.DeletePayload{Name: "Bob"},
			deleteStub: func(context.Context, string) error {
				return nil
			},
			user: userBob(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := &userfakes.FakeUserStorage{}
			storage.DeleteStub = test.deleteStub

			svc := user.NewUser(storage)
			err := svc.Delete(context.Background(), test.payload)

			assert.Nil(t, err)
		})
	}
}
