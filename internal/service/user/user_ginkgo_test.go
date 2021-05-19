package user_test

import (
	"context"
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
	"github.com/ebalkanski/goa/internal/service/user"
	"github.com/ebalkanski/goa/internal/service/user/userfakes"
)

func testUser() *goauser.User {
	return &goauser.User{
		Name: "Bob",
		Age:  22,
	}
}

var _ = Describe("User service", func() {
	var (
		storage *userfakes.FakeUserStorage
		svc     goauser.Service
	)

	BeforeEach(func() {
		storage = &userfakes.FakeUserStorage{}
		svc = user.NewUser(storage)
	})

	Describe("Constructor", func() {
		Context("when use the constructor", func() {
			It("it should create a user service", func() {
				Expect(svc, user.User{})
			})
		})
	})

	Describe("Get user endpoint", func() {
		var (
			payload *goauser.FetchPayload
			res     *goauser.User
			err     error
		)

		BeforeEach(func() {
			payload = &goauser.FetchPayload{Name: "Bob"}
		})

		JustBeforeEach(func() {
			res, err = svc.Fetch(context.Background(), payload)
		})

		Context("when fetch a user and it is not found in storage", func() {
			BeforeEach(func() {
				storage.UserStub = func(ctx context.Context, s string) (*goauser.User, error) {
					return nil, user.UserNotFound
				}
			})

			It("it should return an error", func() {
				Expect(res, nil)
				Expect(err, &goa_errors.Error{})

				gerr, ok := err.(*goa_errors.Error)
				Expect(ok).To(BeTrue())
				Expect(gerr.Message).To(Equal(user.UserNotFound.Error()))
				Expect(gerr.StatusCode()).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when fetch a user and it exists", func() {
			BeforeEach(func() {
				storage.UserStub = func(ctx context.Context, s string) (*goauser.User, error) {
					return nil, user.UserExists
				}
			})

			It("it should return an error", func() {
				Expect(res, nil)
				Expect(err, &goa_errors.Error{})

				gerr, ok := err.(*goa_errors.Error)
				Expect(ok).To(BeTrue())
				Expect(gerr.Message).To(Equal(user.UserExists.Error()))
				Expect(gerr.StatusCode()).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when fetch a user and there is an error in storage", func() {
			BeforeEach(func() {
				storage.UserStub = func(ctx context.Context, s string) (*goauser.User, error) {
					return nil, errors.New("ERR")
				}
			})

			It("it should return an error", func() {
				Expect(res, nil)
				Expect(err, &goa_errors.Error{})

				gerr, ok := err.(*goa_errors.Error)
				Expect(ok).To(BeTrue())
				Expect(gerr.Message).To(Equal(user.UserNotFound.Error()))
				Expect(gerr.StatusCode()).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when fetch a user and user is found", func() {
			var bob *goauser.User

			BeforeEach(func() {
				bob := testUser()
				storage.UserStub = func(ctx context.Context, s string) (*goauser.User, error) {
					return bob, nil
				}
			})

			It("it should be returned", func() {
				Expect(res, bob)
				Expect(err, nil)
			})
		})
	})
})
