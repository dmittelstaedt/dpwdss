package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func setDB(dbf *sql.DB) {
	db = dbf
}

// readUsersHandler handles requests for reading all users.
// TODO: show only name and id?
func readUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	users, err := readUsers(db)
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readUserHandler handles requests for reading details of specific user.
func readUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	user, err := readUser(db, params["name"])
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println()
	}

	respondJSON(w, http.StatusOK, response)
}

// readPermissionsHandler hanldes requests for reading all existing permissions.
func readPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	permissions, err := readPermissions(db)
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readPermissionHandler handles requests for reading permissions for specific user.
func readPermissionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	permissions, err := readPermission(db, params["name"])
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// Experimentell hanlders

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
