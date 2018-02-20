package main

import (
	"github.com/plainprogrammer/augustine/application"
	"github.com/plainprogrammer/augustine/application/middlewares"
	"log"
)

func main() {
	app := application.New()

	log.Println("Setting up middleware stack...")
	app.MiddlewareStack = []application.Middleware{middlewares.HelloWorld}
	log.Println("Finished setting up middleware stack...")

	log.Println("Starting application!")
	app.Run()
}
