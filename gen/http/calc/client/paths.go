// Code generated by goa v3.3.1, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen github.com/ebalkanski/goa/design

package client

import (
	"fmt"
)

// AddCalcPath returns the URL path to the calc service add HTTP endpoint.
func AddCalcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}