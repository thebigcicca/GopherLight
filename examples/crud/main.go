

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
    app.Route("GET", "/hello", HelloHandler)
    app.Route("POST", "/users/create", CreateUser)
    app.Route("GET", "/users/get", GetUser)
    app.Route("PUT", "/users/update", UpdateUser)
    app.Route("DELETE", "/users/delete", DeleteUser)

    fmt.Println("Server listening on port 3333")
    app.Listen(":3333")
}
