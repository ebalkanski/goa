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

var _ = Service("play", func() {
	Description("The play service is a sandbox for goa testing")

	Method("add", func() {
		Payload(func() {
			Attribute("a", Int, "Left operand", func() {
				Example(1)
			})
			Attribute("b", Int, "Right operand", func() {
				Example(2)
			})
			Required("a", "b")
		})

		Result(Int, func() {
			Example(3)
		})

		HTTP(func() {
			GET("/add/{a}/{b}")
			Response(StatusOK)
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
		Path("/swagger-ui")
	})
	Files("/openapi.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containing the API swagger definition")
	})
	Files("/{*filepath}", "./internal/swagger/")
})
