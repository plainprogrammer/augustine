package middlewares_test

import (
	. "github.com/plainprogrammer/augustine/augustine/middlewares"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/plainprogrammer/augustine/augustine"
	"net/http"
)

var _ = Describe("HelloWorld", func() {
	Describe("Name()", func() {
		It("returns 'HelloWorld'", func() {
			subject := HelloWorld.Name()

			Expect(subject).To(Equal("HelloWorld"))
		})
	})

	Describe("OnRequest()", func() {
		It("sets status code of environment to http.StatusOK", func() {
			environment := augustine.Environment{}
			subject := HelloWorld.OnRequest(environment)

			Expect(subject.Status).To(Equal(http.StatusOK))
		})

		It("appends 'Hello, World!' to content of environment", func() {
			environment := augustine.Environment{Content: []string{"Example"}}
			subject := HelloWorld.OnRequest(environment)

			Expect(subject.Content).To(Equal([]string{"Example", "Hello, World!"}))
		})
	})

	Describe("OnResponse()", func() {
		It("returns the environment unmodified", func() {
			environment := augustine.Environment{}
			subject := HelloWorld.OnResponse(environment)

			Expect(subject).To(Equal(environment))
		})
	})
})
