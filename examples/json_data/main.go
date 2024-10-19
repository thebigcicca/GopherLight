package main

import (
	"fmt"

	"github.com/BrunoCiccarino/express-go/req"
	"github.com/BrunoCiccarino/express-go/router"
)

// main sets up the application and defines a route that returns JSON data.
// It listens on port 3333 and responds to the "/json" path with a JSON object.
func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/json".
	app.Route("/json", func(r *req.Request, w *req.Response) {
		data := map[string]string{
			"message": "Hello, JSON",
		}
		w.JSON(data) // Send the JSON response
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
