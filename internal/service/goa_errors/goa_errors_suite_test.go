package goa_errors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoaErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goa Custom Errors Suite")
}
