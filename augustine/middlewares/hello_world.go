package middlewares

import (
	"github.com/plainprogrammer/augustine/augustine"
	"net/http"
)

var HelloWorld = augustine.NewMiddleware(
	"HelloWorld",
	func(env augustine.Environment) augustine.Environment {
		result := augustine.Environment{}

		result.Status = http.StatusOK
		result.Content = append(env.Content, "Hello, World!")

		return result
	},
	func(Environment augustine.Environment) augustine.Environment { return Environment })
