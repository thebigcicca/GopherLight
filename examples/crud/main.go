package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BrunoCiccarino/GopherLight/req"
	"github.com/BrunoCiccarino/GopherLight/router"
)

// Middleware example: Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

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
		res.Status(http.StatusBadRequest).JSONError("Invalid input")
		log.Println("Error decoding JSON:", err)
		return
	}
	user.ID = nextID
	nextID++
	users[user.ID] = user
	res.Status(http.StatusCreated).JSON(user)
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
		res.Status(http.StatusBadRequest).JSONError("Invalid user ID")
		return
	}

	user, exists := users[id]
	if !exists {
		res.Status(http.StatusNotFound).JSONError("User not found")
		return
	}

	res.JSON(user)
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
		res.Status(http.StatusBadRequest).JSONError("Invalid user ID")
		return
	}

	var updatedUser User
	err = json.Unmarshal([]byte(req.BodyAsString()), &updatedUser)
	if err != nil {
		res.Status(http.StatusBadRequest).JSONError("Invalid input")
		log.Println("Error decoding JSON:", err)
		return
	}

	user, exists := users[id]
	if !exists {
		res.Status(http.StatusNotFound).JSONError("User not found")
		return
	}

	user.Name = updatedUser.Name
	user.Age = updatedUser.Age
	users[id] = user

	res.JSON(user)
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
		res.Status(http.StatusBadRequest).JSONError("Invalid user ID")
		return
	}

	_, exists := users[id]
	if !exists {
		res.Status(http.StatusNotFound).JSONError("User not found")
		return
	}

	delete(users, id)
	res.JSON(map[string]string{"message": fmt.Sprintf("User %d deleted", id)})
}

// Define the "/hello" route handler
func HelloHandler(req *req.Request, res *req.Response) {
	res.Send("Hello, World!")
}

func main() {
	app := router.NewApp()

	// Use the logging middleware
	app.Use(loggingMiddleware)

	// Register routes
	app.Get("/hello", HelloHandler)
	app.Post("/users/create", CreateUser)
	app.Get("/users/get", GetUser)
	app.Put("/users/update", UpdateUser)
	app.Delete("/users/delete", DeleteUser)

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
