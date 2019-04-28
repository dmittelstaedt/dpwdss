package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func readPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	// var permissions string
	params := mux.Vars(r)

	log.Println("user-name: " + params["user-name"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}

// Experimentell hanlders
func readAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	testUsers := []User{
		User{ID: 1, Name: "David Mittelstaedt"},
		User{ID: 2, Name: "Luke Skywalker"},
	}

	output, _ := json.Marshal(testUsers)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func readUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)
	testUser := User{
		ID:   3,
		Name: "Han Solo",
	}

	log.Println(params["id"])

	response, _ := json.Marshal(testUser)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	var testUser User
	if err := json.NewDecoder(r.Body).Decode(&testUser); err != nil {
		log.Fatal(err)
	}
	testUser.ID = 3
	log.Println("User updated: " + testUser.Name)

	response, _ := json.Marshal(testUser)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
