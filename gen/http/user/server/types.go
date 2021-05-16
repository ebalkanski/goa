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

// EditRequestBody is the type of the "user" service "edit" endpoint HTTP
// request body.
type EditRequestBody struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// FetchResponseBody is the type of the "user" service "fetch" endpoint HTTP
// response body.
type FetchResponseBody struct {
	Name string `form:"name" json:"name" xml:"name"`
	Age  int    `form:"age" json:"age" xml:"age"`
}

// FetchAllResponseBody is the type of the "user" service "fetchAll" endpoint
// HTTP response body.
type FetchAllResponseBody struct {
	Users []*UserResponseBody `form:"users,omitempty" json:"users,omitempty" xml:"users,omitempty"`
}

// FetchBadRequestResponseBody is the type of the "user" service "fetch"
// endpoint HTTP response body for the "BadRequest" error.
type FetchBadRequestResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// FetchInternalServerErrorResponseBody is the type of the "user" service
// "fetch" endpoint HTTP response body for the "InternalServerError" error.
type FetchInternalServerErrorResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// FetchAllInternalServerErrorResponseBody is the type of the "user" service
// "fetchAll" endpoint HTTP response body for the "InternalServerError" error.
type FetchAllInternalServerErrorResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateBadRequestResponseBody is the type of the "user" service "create"
// endpoint HTTP response body for the "BadRequest" error.
type CreateBadRequestResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateInternalServerErrorResponseBody is the type of the "user" service
// "create" endpoint HTTP response body for the "InternalServerError" error.
type CreateInternalServerErrorResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// EditBadRequestResponseBody is the type of the "user" service "edit" endpoint
// HTTP response body for the "BadRequest" error.
type EditBadRequestResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// EditInternalServerErrorResponseBody is the type of the "user" service "edit"
// endpoint HTTP response body for the "InternalServerError" error.
type EditInternalServerErrorResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteBadRequestResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "BadRequest" error.
type DeleteBadRequestResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteInternalServerErrorResponseBody is the type of the "user" service
// "delete" endpoint HTTP response body for the "InternalServerError" error.
type DeleteInternalServerErrorResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	Name string `form:"name" json:"name" xml:"name"`
	Age  int    `form:"age" json:"age" xml:"age"`
}

// NewFetchResponseBody builds the HTTP response body from the result of the
// "fetch" endpoint of the "user" service.
func NewFetchResponseBody(res *user.User) *FetchResponseBody {
	body := &FetchResponseBody{
		Name: res.Name,
		Age:  res.Age,
	}
	return body
}

// NewFetchAllResponseBody builds the HTTP response body from the result of the
// "fetchAll" endpoint of the "user" service.
func NewFetchAllResponseBody(res *user.FetchAllResult) *FetchAllResponseBody {
	body := &FetchAllResponseBody{}
	if res.Users != nil {
		body.Users = make([]*UserResponseBody, len(res.Users))
		for i, val := range res.Users {
			body.Users[i] = marshalUserUserToUserResponseBody(val)
		}
	}
	return body
}

// NewFetchBadRequestResponseBody builds the HTTP response body from the result
// of the "fetch" endpoint of the "user" service.
func NewFetchBadRequestResponseBody(res *user.GoaError) *FetchBadRequestResponseBody {
	body := &FetchBadRequestResponseBody{
		Message: res.Message,
	}
	return body
}

// NewFetchInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "fetch" endpoint of the "user" service.
func NewFetchInternalServerErrorResponseBody(res *user.GoaError) *FetchInternalServerErrorResponseBody {
	body := &FetchInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewFetchAllInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "fetchAll" endpoint of the "user" service.
func NewFetchAllInternalServerErrorResponseBody(res *user.GoaError) *FetchAllInternalServerErrorResponseBody {
	body := &FetchAllInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "user" service.
func NewCreateBadRequestResponseBody(res *user.GoaError) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Message: res.Message,
	}
	return body
}

// NewCreateInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "create" endpoint of the "user" service.
func NewCreateInternalServerErrorResponseBody(res *user.GoaError) *CreateInternalServerErrorResponseBody {
	body := &CreateInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewEditBadRequestResponseBody builds the HTTP response body from the result
// of the "edit" endpoint of the "user" service.
func NewEditBadRequestResponseBody(res *user.GoaError) *EditBadRequestResponseBody {
	body := &EditBadRequestResponseBody{
		Message: res.Message,
	}
	return body
}

// NewEditInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "edit" endpoint of the "user" service.
func NewEditInternalServerErrorResponseBody(res *user.GoaError) *EditInternalServerErrorResponseBody {
	body := &EditInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteBadRequestResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "user" service.
func NewDeleteBadRequestResponseBody(res *user.GoaError) *DeleteBadRequestResponseBody {
	body := &DeleteBadRequestResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "delete" endpoint of the "user" service.
func NewDeleteInternalServerErrorResponseBody(res *user.GoaError) *DeleteInternalServerErrorResponseBody {
	body := &DeleteInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewFetchPayload builds a user service fetch endpoint payload.
func NewFetchPayload(name string) *user.FetchPayload {
	v := &user.FetchPayload{}
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

// NewEditUser builds a user service edit endpoint payload.
func NewEditUser(body *EditRequestBody) *user.User {
	v := &user.User{
		Name: *body.Name,
		Age:  *body.Age,
	}

	return v
}

// NewDeletePayload builds a user service delete endpoint payload.
func NewDeletePayload(name string) *user.DeletePayload {
	v := &user.DeletePayload{}
	v.Name = name

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

// ValidateEditRequestBody runs the validations defined on EditRequestBody
func ValidateEditRequestBody(body *EditRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Age == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("age", "body"))
	}
	return
}
