package main

import (
	"fmt"

	"github.com/BrunoCiccarino/GopherLight/middleware"
	"github.com/BrunoCiccarino/GopherLight/req"
	"github.com/BrunoCiccarino/GopherLight/router"
)

// Define the "/hello" route handler
func HelloHandler(req *req.Request, res *req.Response) {
	res.Send("Hello, World!")
}

func main() {

	app := router.NewApp()

	// Use the cors middleware
	app.Use(middleware.CORSMiddleware(middleware.DefaultCORSOptions))

	// Register routes
	app.Get("/hello", HelloHandler)

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
