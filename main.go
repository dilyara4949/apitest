package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type userStorage struct {
	mu      sync.Mutex
	storage map[string]string
}

var users = userStorage{
	storage: make(map[string]string),
}

func GetUser(email string) string {
	users.mu.Lock()
	defer users.mu.Unlock()

	return users.storage[email]
}

func AddUser(email, name string) {
	users.mu.Lock()
	defer users.mu.Unlock()

	users.storage[email] = name
}

func DeleteUser(email string) {
	users.mu.Lock()
	defer users.mu.Unlock()

	delete(users.storage, email)
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
*
Test HTTP service.
GET endpoint should return user by it's email. URL Params are: email (string).
POST endpoint should creat new user. URL Params are: email (string) and name(string).
DELETE endpoint should remove user by it's email. URL params are: email(string).
*/
func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		params := r.URL.Query()
		email := params.Get("email")

		name := GetUser(email)
		if len(name) == 0 {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		response := UserResponse{
			Name:  name,
			Email: email,
		}

		json, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "500 not found.", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, string(json))

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		params := r.URL.Query()
		name := params.Get("name")
		email := params.Get("email")

		AddUser(email, name)
		fmt.Fprintf(w, "%s has been added with next email:  %s\n", name, email)
	case "DELETE":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		params := r.URL.Query()
		email := params.Get("email")

		DeleteUser(email)
		fmt.Fprintf(w, "%s has been deleted", email)

	default:
		http.Error(w, "400 not found.", http.StatusBadRequest)
		fmt.Fprintf(w, "Sorry, only GET, POST and DELETE methods are supported.")
	}
}

func main() {

	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
