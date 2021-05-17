package goa_errors_test

import (
	"errors"
	"net/http"

	"github.com/ebalkanski/goa/internal/service/goa_errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors2", func() {
	var (
		err error
	)

	BeforeEach(func() {
		err = errors.New("ERROR")
	})

	Describe("Goa errors", func() {
		Context("When create BadRequest error", func() {
			It("it should be correct", func() {
				badErr := goa_errors.NewBadRequestError(err)
				Expect(badErr.Message).To(Equal(err.Error()))
				Expect(badErr.StatusCode()).To(Equal(http.StatusBadRequest))
				Expect(badErr.Error()).To(Equal("application error"))
			})
		})

		Context("When create InternalServer error", func() {
			It("it should be correct", func() {
				badErr := goa_errors.NewInternalServerError(err)
				Expect(badErr.Message).To(Equal(err.Error()))
				Expect(badErr.StatusCode()).To(Equal(http.StatusInternalServerError))
				Expect(badErr.Error()).To(Equal("application error"))
			})
		})
	})
})
