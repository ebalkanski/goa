package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("play", func() {
	Title("Play with goa")
	Description("HTTP service for playing with goa")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

var _ = Service("user", func() {
	Description("The user service process users")

	Error("BadRequest", GoaError)
	Error("InternalServerError", GoaError)

	Method("get", func() {
		Payload(func() {
			Attribute("name", String, func() {
				Example("Bob")
			})
			Required("name")
		})
		Result(User, func() {
			Example(func() {
				Value(Val{"name": "Bob", "age": 25})
			})
		})
		HTTP(func() {
			GET("/user/{name}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalServerError", StatusInternalServerError)
		})
	})

	Method("create", func() {
		Description("Create new user.")
		Payload(User)
		HTTP(func() {
			POST("/user")
			Response(StatusCreated, func() {
				Description("User is created successfully.")
			})
			Response("BadRequest", StatusBadRequest)
			Response("InternalServerError", StatusInternalServerError)
		})
	})
})

var _ = Service("openapi", func() {
	Description("The swagger service serves the API swagger definition.")
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/swagger-ui")
	})
	Files("/openapi.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containing the API swagger definition")
	})
	Files("/{*filepath}", "./internal/swagger/")
})
