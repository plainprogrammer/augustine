package middlewares

import (
	"github.com/plainprogrammer/augustine/application"
	"net/http"
)

var HelloWorld = application.NewMiddleware(
	"HelloWorld",
	func(env application.Environment) application.Environment {
		result := application.Environment{}

		result.Status = http.StatusOK
		result.Content = append(env.Content, "Hello, World!")

		return result
	},
	func(Environment application.Environment) application.Environment { return Environment })
