package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var db *sql.DB

func setDB(dbf *sql.DB) {
	db = dbf
}

// readUsersHandler handles requests for reading all users.
func readUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	users := readUsers(db)

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

	permissions, _ := readPermission(db, params["name"])

	response, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

func readUserPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	permissions, err := readUserPermissions(db, params["name"])
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

func readUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	permission, _ := readUserPermission(db, params["name"], params["permission-name"])

	response, err := json.Marshal(permission)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// updateUserPermissionHandler handles requests for updating a permission of a user.
func updateUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	// Check if user exists --> return 404

	// Check if new permission is valid --> return 422
	_, ok := readPermission(db, params["permission-name"])
	if !ok {
		respondHeader(w, http.StatusUnprocessableEntity)
		return
	}

	// Check if new permission already exists --> retrun 409
	_, ok = readUserPermission(db, params["name"], params["permission-name"])
	if ok {
		respondHeader(w, http.StatusConflict)
		return
	}

	// Create substrings
	substrs := strings.Split(params["permission-name"], "-")
	dir := strings.Join(substrs[:len(substrs)-1], "-")
	rw := strings.Join(substrs[len(substrs)-1:], "")

	// Check if old permission already exists --> 404
	var oldPermissionName string
	if rw == "read" {
		oldPermissionName = dir + "-write"
	} else {
		oldPermissionName = dir + "-read"
	}
	log.Println(oldPermissionName)

	_, ok = readUserPermission(db, params["name"], oldPermissionName)

	if ok {
		if rowCnt := updateUserPermission(db, params["name"], oldPermissionName, params["permission-name"]); rowCnt == 1 {
			log.Println("Permission updated")
			respondHeader(w, http.StatusOK)
		}
	} else { // TODO: move to separate function
		if rowCnt := insertUserPermission(db, params["name"], params["permission-name"]); rowCnt == 1 {
			log.Println("Permission updated")
			respondHeader(w, http.StatusCreated)
		}
	}
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
