package design

import (
	. "goa.design/goa/v3/dsl"
)

var User = Type("User", func() {
	Description("User representation")
	Attribute("name", String, func() {
		Example("Bob")
	})
	Attribute("age", Int, func() {
		Example(25)
	})
	Required("name", "age")
})

var GoaError = Type("GoaError", func() {
	Description("GoaError is the error returned from services.")
	Field(1, "message", String, func() {
		Example("error")
	})
	Required("message")
})

var BadRequest = Type("BadRequest", func() {
	Description("BadRequest is the error returned when request data is invalid.")
	Reference(GoaError)
})

var InternalServerError = Type("InternalServerError", func() {
	Description("InternalServerError is the error returned when internal error occurs.")
	Reference(GoaError)
})
