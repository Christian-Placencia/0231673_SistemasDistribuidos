package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var users []User
var nextID = 1

// Handles POST requests to add a new user. It decodes the incoming JSON, assigns a new ID, and adds the user to the users slice.
func addUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    newUser.ID = nextID
    nextID++
    users = append(users, newUser)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

// Handles GET requests to retrieve all users.
func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Sets up the HTTP server with two endpoints:
func main() {
	// To add a new user.
    http.HandleFunc("/adduser", addUser)
    http.HandleFunc("/getusers", getUsers)

	// To get all users.
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
