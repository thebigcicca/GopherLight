package main

import (
	"fmt"

	"github.com/BrunoCiccarino/GopherLight/req"
	"github.com/BrunoCiccarino/GopherLight/router"
)

// main sets up the application, defines a route, and starts the server.
// It listens on port 3333 and responds to the "/hello" path with a plain text message.
func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/hello".
	app.Get("/hello", func(r *req.Request, w *req.Response) {
		w.Send("Hello, World!")
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
