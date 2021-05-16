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

	Method("fetch", func() {
		Description("Fetch user.")
		Meta("swagger:summary", "")
		Payload(func() {
			Attribute("name", String, func() {
				Example("Bob")
			})
			Required("name")
		})
		Result(User)
		HTTP(func() {
			GET("/user/{name}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalServerError", StatusInternalServerError)
		})
	})

	Method("fetchAll", func() {
		Description("Fetch all users.")
		Meta("swagger:summary", "")
		//Result(func() {
		//	Attribute("users", ArrayOf(User), func() {
		//		Example([]user.User{
		//			{Name: "Bob", Age: 25},
		//			{Name: "John", Age: 33},
		//		})
		//	})
		//})
		Result(Users)
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
			Response("InternalServerError", StatusInternalServerError)
		})
	})

	Method("create", func() {
		Description("Create new user.")
		Meta("swagger:summary", "")
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

	Method("edit", func() {
		Description("Edit user.")
		Meta("swagger:summary", "")
		Payload(User)
		HTTP(func() {
			PUT("/user")
			Response(StatusNoContent, func() {
				Description("User is edited successfully.")
			})
			Response("BadRequest", StatusBadRequest)
			Response("InternalServerError", StatusInternalServerError)
		})
	})

	Method("delete", func() {
		Description("Delete user.")
		Meta("swagger:summary", "")
		Payload(func() {
			Attribute("name", String, func() {
				Example("Bob")
			})
			Required("name")
		})
		HTTP(func() {
			DELETE("/user/{name}")
			Response(StatusNoContent)
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
