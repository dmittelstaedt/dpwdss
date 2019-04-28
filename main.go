package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TestUser realizes a testUser
type TestUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// db := connection()
	// getUser(db, "david", "app_admin")
	// defer db.Close()

	r := mux.NewRouter().StrictSlash(true)
	r.Methods("GET").Path("/users").Name("users").HandlerFunc(getAllUsersHandler)
	r.Methods("GET").Path("/users/{id}").Name("getUser").HandlerFunc(getUserHandler)
	r.Methods("POST").Path("/update").Name("update").HandlerFunc(updateUserHandler)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	testUsers := []TestUser{
		TestUser{ID: 1, Name: "David Mittelstaedt"},
		TestUser{ID: 2, Name: "Luke Skywalker"},
	}

	output, _ := json.Marshal(testUsers)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)
	testUser := TestUser{
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
	var testUser TestUser
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
