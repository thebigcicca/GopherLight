package main

import (
	"fmt"

	"github.com/BrunoCiccarino/GopherLight/req"
	"github.com/BrunoCiccarino/GopherLight/router"
)

// main sets up the application, defines a route that checks for an Authorization header,
// and responds with an appropriate message based on the header's presence.
// It listens on port 3333 and responds to the "/auth" path.
func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/auth".
	app.Route("GET", "/auth", func(r *req.Request, w *req.Response) {
		authHeader := r.Header("Authorization")
		if authHeader == "" {
			w.Status(401).Send("Unauthorized")
		} else {
			w.Send("Authorized: " + authHeader)
		}
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
