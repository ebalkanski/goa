// Code generated by goa v3.3.1, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen github.com/ebalkanski/goa/design

package user

import (
	"context"
)

// The user service process users
type Service interface {
	// Fetch user.
	Get(context.Context, *GetPayload) (res *User, err error)
	// Create new user.
	Create(context.Context, *User) (err error)
	// Edit user.
	Edit(context.Context, *User) (err error)
	// Delete user.
	Delete(context.Context, *DeletePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"get", "create", "edit", "delete"}

// GetPayload is the payload type of the user service get method.
type GetPayload struct {
	Name string
}

// User is the result type of the user service get method.
type User struct {
	Name string
	Age  int
}

// DeletePayload is the payload type of the user service delete method.
type DeletePayload struct {
	Name string
}

// GoaError is the error returned from services.
type GoaError struct {
	Message string
}

// Error returns an error description.
func (e *GoaError) Error() string {
	return "GoaError is the error returned from services."
}

// ErrorName returns "GoaError".
func (e *GoaError) ErrorName() string {
	return "BadRequest"
}
