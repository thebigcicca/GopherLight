## Docs

Hey folks, first I would like to thank you for choosing to use our project. Even though he is small, we did it with great enthusiasm! To start using it you first have to have go installed, let's assume you already have it. Then install the main modules of the framework, which are req and router

```bash
go get github.com/BrunoCiccarino/GopherLight/router
go get github.com/BrunoCiccarino/GopherLight/req
```

Already downloaded? Phew! Now we can make our first hello world.

```go
package main

import (
	"fmt"
	"github.com/BrunoCiccarino/GopherLight/router"
	"github.com/BrunoCiccarino/GopherLight/req"
)


func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/hello".
	app.Get("/hello", func(r *req.Request, w *req.Response) {
		w.Send("Hello, World!")
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
```

Pretty simple, right? And there’s way more we can do with GopherLight. Keep reading for a full breakdown of HTTP methods and our Request and Response tools.

Supported HTTP Methods
Here’s the list of HTTP methods you can use with router.App. Each of these allows you to set up routes to handle different types of requests. Let’s dive in!

### GET
* Usage: `app.Get(path, handler)`

Retrieves data without modifying anything.
Example: Fetching a list of items or reading user details.

### POST
* Usage: `app.Post(path, handler)`

Sends data to create a new resource.
Example: Submitting a form or adding a new item to a list.

### PUT
Usage: `app.Put(path, handler)`

Updates or replaces a resource. It’s an “overwrite” action.
Example: Updating a full user profile.

### DELETE
Usage: `app.Delete(path, handler)`

Deletes a resource.
Example: Removing a user or deleting a post.

### PATCH
Usage: `app.Patch(path, handler)`

Partially updates a resource without replacing everything.
Example: Updating just the email on a user profile.

### OPTIONS
Usage: `app.Options(path, handler)`

Returns allowed HTTP methods for a URL, mainly for CORS preflight requests.

### HEAD
Usage: `app.Head(path, handler)`

Like GET, but no response body. Use it to check if a resource exists.

### CONNECT and TRACE
Usage: `app.Connect(path, handler)`, `app.Trace(path, handler)`

Advanced methods: CONNECT sets up a tunnel (for SSL), and TRACE is for debugging, echoing back the request.

## Working with `req.Request` and `req.Response`
Now that you’ve seen the routes, let’s talk about the Request and Response objects, your go-to helpers for handling incoming requests and sending responses.

### Request
Each request handler gets a Request object loaded with info on the incoming request. Here’s what you can do with it:

* Query Parameters: Get query parameters with .QueryParam("key").
* Headers: Access headers using .Header("key").
* Body as String: Grab the request body with .BodyAsString().

### Example:
```go
app.Get("/greet", func(r *req.Request, w *req.Response) {
	name := r.QueryParam("name")
	if name == "" {
		name = "stranger"
	}
	w.Send("Hello, " + name + "!")
})
```

### Response
The Response object helps you send a reply back to the client. Here's what you can do:

* Send Text: .Send(data string) writes plain text back.
* Set Status: .Status(code) sets the HTTP status.
* Send JSON: .JSON(data) serializes a Go object to JSON and sends it.
* Handle Errors: .JSONError(message) sends a JSON-formatted error response.

### Example:
```go
app.Get("/user", func(r *req.Request, w *req.Response) {
	user := map[string]string{"name": "Gopher", "language": "Go"}
	w.JSON(user)
})
```

Alright, you've got the basics down! Now you’re ready to build APIs that handle all sorts of requests with ease. Dig into those methods, use Request and Response like a pro, and create something awesome. Go forth and code!


Next step: [learn about our middleware](./middleware.md)