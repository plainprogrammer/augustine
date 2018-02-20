package application

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Application struct {
	Server          *http.Server
	Router          interface{}
	MiddlewareStack []Middleware
}

func (app *Application) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	env := Environment{}
	log.Printf("Received request for %s\n", request.RequestURI)

	for _, middleware := range app.MiddlewareStack {
		env = middleware.OnRequest(env)
		log.Printf("Processed request behavior for %s\n", middleware.Name())
	}

	for i := len(app.MiddlewareStack) - 1; i >= 0; i-- {
		middleware := app.MiddlewareStack[i]
		env = middleware.OnResponse(env)
		log.Printf("Processed response behavior for %s\n", middleware.Name())
	}

	for header, value := range env.Headers {
		responseWriter.Header().Set(header, value)
		log.Println("Set headers for response")
	}

	responseWriter.WriteHeader(env.Status)
	io.WriteString(responseWriter, strings.Join(env.Content, ""))
	log.Println("Sent response")
}

func (app *Application) Run() {
	app.Server.Handler = app
	log.Fatal(app.Server.ListenAndServe())
}

func New() *Application {
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Application{
		Server: server,
	}
}
