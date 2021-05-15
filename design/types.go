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
