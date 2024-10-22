package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/BrunoCiccarino/GopherLight/req"
	"github.com/BrunoCiccarino/GopherLight/router"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]User)
var nextID = 1

// CreateUser adds a new user to the system.
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

	app.Route("POST", "/users/create", CreateUser)
	app.Route("GET", "/users/get", GetUser)
	app.Route("PUT", "/users/update", UpdateUser)
	app.Route("DELETE", "/users/delete", DeleteUser)

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
