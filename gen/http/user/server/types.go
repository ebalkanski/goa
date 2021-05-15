// Code generated by goa v3.3.1, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/ebalkanski/goa/design

package server

import (
	user "github.com/ebalkanski/goa/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	Name string `form:"name" json:"name" xml:"name"`
	Age  int    `form:"age" json:"age" xml:"age"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "user" service.
func NewGetResponseBody(res *user.User) *GetResponseBody {
	body := &GetResponseBody{
		Name: res.Name,
		Age:  res.Age,
	}
	return body
}

// NewGetPayload builds a user service get endpoint payload.
func NewGetPayload(name string) *user.GetPayload {
	v := &user.GetPayload{}
	v.Name = name

	return v
}

// NewCreateUser builds a user service create endpoint payload.
func NewCreateUser(body *CreateRequestBody) *user.User {
	v := &user.User{
		Name: *body.Name,
		Age:  *body.Age,
	}

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Age == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("age", "body"))
	}
	return
}