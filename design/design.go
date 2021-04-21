package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

var DivByZero = Type("DivByZero", func() {
	Description("DivByZero is the error returned when using value 0 as divisor.")
	Field(1, "errors", ArrayOf(String), "division by zero leads to infinity.", func() {
		Example([]string{
			"Err1",
			"Err2",
		})
	})
	Required("errors")
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")

	Error("div_by_zero", DivByZero, func() {
		Description("DivByZero is the error returned by the service methods when the right operand is 0.")
	})

	// Method describes a service method (endpoint)
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand", func() {
				Example(1)
			})
			Attribute("b", Int, "Right operand", func() {
				Example(0)
			})
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int, func() {
			Example(3)
		})

		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
			Response("div_by_zero", StatusBadRequest)
		})
	})

	Method("rate", func() {
		Payload(func() {
			Attribute("id", Int, func() {
				Example(1)
			})
			Attribute("rates", MapOf(String, Float64), func() {
				Example(map[string]float64{"a": 1.1, "b": 2.2})
			})
		})
		HTTP(func() {
			POST("/rate/{id}")
			Body("rates")
		})
	})
})

var _ = Service("openapi", func() {
	Description("The swagger service serves the API swagger definition.")
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/openapi")
	})
	Files("/openapi.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containing the API swagger definition")
	})
	Files("/{*filepath}", "swagger/")
})
