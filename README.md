# GopherLight ![gopher1](./img/typing-furiously.gif)

![GitHub License](https://img.shields.io/github/license/BrunoCiccarino/express-go?style=for-the-badge&color=blue&link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino%2Fexpress-go%2Fblob%2Fmain%2FLICENSE) ![Go Reference](https://img.shields.io/badge/reference-grey?style=for-the-badge&logo=go&link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino%2Fexpress-go) ![pr's welcome](https://img.shields.io/badge/PR'S-WELCOME-green?style=for-the-badge) ![GitHub Repo stars](https://img.shields.io/github/stars/BrunoCiccarino/express-go) ![GitHub followers](https://img.shields.io/github/followers/BrunoCiccarino?link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino) ![GitHub forks](https://img.shields.io/github/forks/BrunoCiccarino/express-go)


![gopher2](./img/gopher.png)

### What is express-go?
Hey there! So, you know how building web applications can sometimes feel like climbing a mountain? Well, express-go is like that cool hiking buddy who helps you navigate the trail, making things way easier and way more fun!

express-go is a micro framework for Go (Golang) that brings a bit of the simplicity and flexibility of the popular Express.js framework from the Node.js world right to your Go projects. It’s perfect for those times when you want to whip up a web server or an API without getting bogged down in all the nitty-gritty details.

Imagine you want to handle HTTP requests and create endpoints to manage users—just like in a classic CRUD (Create, Read, Update, Delete) app. With express-go, you can define your routes and handlers in a snap. No need to wrestle with the standard net/http package; instead, you get a clean and straightforward way to manage your routes and responses.

The cool part? You get to focus on writing your application logic while the framework handles the heavy lifting under the hood. Need to add a new route? Just call a method and pass in your handler. Want to send a JSON response? Easy peasy!

Plus, it’s lightweight, so it won’t weigh down your application. You get all the goodies of a modern web framework while keeping things simple and fast. Whether you're a seasoned pro or just dipping your toes into web development, express-go makes it a breeze to get your ideas off the ground.

So, if you’re looking for a friendly and efficient way to build web apps in Go, express-go is your new best friend. Grab your backpack, and let’s hit the trail!

## Examples:

### Configuring a simple route and responding with a string

```go
package main

import (
	"github.com/BrunoCiccarino/express-go/req"
	"github.com/BrunoCiccarino/express-go/router"
	"fmt"
)

// main sets up the application, defines a route, and starts the server.
// It listens on port 3333 and responds to the "/hello" path with a plain text message.
func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/hello".
	app.Route("/hello", func(r *req.Request, w *req.Response) {
		w.Send("Hello, World!")
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
```

### Returning JSON data

```go
package main

import (
	"github.com/BrunoCiccarino/express-go/req"
	"github.com/BrunoCiccarino/express-go/router"
	"fmt"
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
```

### Using HTTP headers

```go
package main

import (
	"github.com/BrunoCiccarino/express-go/req"
	"github.com/BrunoCiccarino/express-go/router"
	"fmt"
)

// main sets up the application, defines a route that checks for an Authorization header,
// and responds with an appropriate message based on the header's presence.
// It listens on port 3333 and responds to the "/auth" path.
func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/auth".
	app.Route("/auth", func(r *req.Request, w *req.Response) {
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
```

### Crud example using express-go

```go
package main

import (
	"encoding/json"
	"github.com/BrunoCiccarino/express-go/req"
	"github.com/BrunoCiccarino/express-go/router"
	"log"
	"strconv"
    "fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]User)
var nextID = 1

// CreateUser adds a new user to the system.
//
// This function decodes user data from the request,
// assigns a unique ID to the user and stores it in the in-memory "database".
//
// Parameters:
//
// req: The received request, containing user data in the request body.
// res: The response to be sent, containing the status of the operation and the created user.
//
// Returns:
//
// Sends a JSON response with the created user data or an error if the input is invalid.
func CreateUser(req *req.Request, res *req.Response) {
	var user User
	err := json.Unmarshal([]byte(req.BodyAsString()), &user)
	if err != nil {
		res.Status(400).Send("Invalid input")
		log.Println("Error decoding JSON:", err)
		return
	}
	user.ID = nextID
	nextID++
	users[user.ID] = user
	res.Status(201).JSON(user)
}

// GetUser returns a user by their ID.
//
// Params:
//
// req: The received request (containing the user ID).
// res: The response to be sent (containing the user or an error message).
func GetUser(req *req.Request, res *req.Response) {
	idParam := req.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		res.Status(400).Send("Invalid user ID")
		return
	}

	user, exists := users[id]
	if !exists {
		res.Status(404).Send("User not found")
		return
	}

	res.Status(200).JSON(user)
}

// UpdateUser updates a user's data.
//
// Params:
//
// req: The received request (containing the new data).
// res: The response to be sent (containing the updated status and user).
func UpdateUser(req *req.Request, res *req.Response) {
	idParam := req.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		res.Status(400).Send("Invalid user ID")
		return
	}

	var updatedUser User
	err = json.Unmarshal([]byte(req.BodyAsString()), &updatedUser)
	if err != nil {
		res.Status(400).Send("Invalid input")
		log.Println("Error decoding JSON:", err)
		return
	}

	user, exists := users[id]
	if !exists {
		res.Status(404).Send("User not found")
		return
	}

	user.Name = updatedUser.Name
	user.Age = updatedUser.Age
	users[id] = user

	res.Status(200).JSON(user)
}

// DeleteUser removes a user by ID.
//
// Parameters:
//
// req: The received request (containing the user ID).
// res: The response to be sent (success or error status).
func DeleteUser(req *req.Request, res *req.Response) {
	idParam := req.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		res.Status(400).Send("Invalid user ID")
		return
	}

	_, exists := users[id]
	if !exists {
		res.Status(404).Send("User not found")
		return
	}

	delete(users, id)
	res.Status(200).Send(fmt.Sprintf("User %d deleted", id))
}

func main() {
	app := router.NewApp()

	app.Route("/users/create", CreateUser)
	app.Route("/users/get", GetUser)
	app.Route("/users/update", UpdateUser)
	app.Route("/users/delete", DeleteUser)

	app.Listen(":3333")
}
```
