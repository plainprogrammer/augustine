package main

import (
	"github.com/plainprogrammer/augustine/augustine"
	"github.com/plainprogrammer/augustine/middlewares"
	"log"
)

func main() {
	app := augustine.New()

	log.Println("Setting up middleware stack...")
	app.MiddlewareStack = []augustine.Middleware{middlewares.HelloWorld}
	log.Println("Finished setting up middleware stack...")

	log.Println("Starting augustine!")
	app.Run()
}
