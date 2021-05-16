package goa_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError(t *testing.T) {
	err := errors.New("ERROR")
	badErr := NewBadRequestError(err)
	assert.Equal(t, err.Error(), badErr.Message)
	assert.Equal(t, "application error", badErr.Error())
	assert.Equal(t, http.StatusBadRequest, badErr.StatusCode())
}

func TestNewInternalServerError(t *testing.T) {
	err := errors.New("ERROR")
	intErr := NewInternalServerError(err)
	assert.Equal(t, err.Error(), intErr.Message)
	assert.Equal(t, "application error", intErr.Error())
	assert.Equal(t, http.StatusInternalServerError, intErr.StatusCode())
}
