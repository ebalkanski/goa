package user_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ebalkanski/goa/internal/service/user"
	"github.com/ebalkanski/goa/internal/service/user/userfakes"
)

func Test(t *testing.T) {
	logger := log.Default()
	repo := &userfakes.FakeUserRepo{}

	svc := user.NewUser(logger, repo)
	assert.IsType(t, &user.User{}, svc)
}
