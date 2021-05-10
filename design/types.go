package design

import (
	. "goa.design/goa/v3/dsl"
)

var User = Type("User", func() {
	Description("User representation")
	Attribute("name", String)
	Attribute("age", Int)
	Required("name", "age")
})
