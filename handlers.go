package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// readUsersHandler handles requests for reading all users.
func (server *Server) readUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	users := readUsers(server.DB)

	response, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readUserHandler handles requests for reading details of specific user.
func (server *Server) readUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	user, ok := readUser(server.DB, params["name"])
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println()
	}

	respondJSON(w, http.StatusOK, response)
}

// readGroupsHandler handles requests for reading all group.
func (server *Server) readGroupsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	groups := readGroups(server.DB)

	response, err := json.Marshal(groups)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readUserHandler handles requests for reading details of specific group.
func (server *Server) readGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	group, ok := readGroup(server.DB, params["name"])
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(group)
	if err != nil {
		log.Println()
	}

	respondJSON(w, http.StatusOK, response)
}

// readPermissionsHandler hanldes requests for reading all existing permissions.
func (server *Server) readPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)

	permissions, err := readPermissions(server.DB)
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
func (server *Server) readPermissionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: " + r.URL.Path)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	permissions, ok := readPermission(server.DB, id)
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// func (server *Server) readUserPermissionsHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request: " + r.URL.Path)
// 	params := mux.Vars(r)

// 	permissions, err := readUserPermissions(server.DB, params["name"])
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	response, err := json.Marshal(permissions)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// TODO: Return 404 if no permission is set for user
// 	respondJSON(w, http.StatusOK, response)
// }

// func (server *Server) readUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request: " + r.URL.Path)
// 	params := mux.Vars(r)

// 	permission, _ := readUserPermission(server.DB, params["name"], params["permission-name"])

// 	response, err := json.Marshal(permission)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	respondJSON(w, http.StatusOK, response)
// }

// insertUserPermissionHandler handles requests for inserting permission of a user.
// func (server *Server) insertUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request: " + r.URL.Path)
// 	params := mux.Vars(r)

// 	// Check if user resource exists --> 404
// 	_, ok := readUser(server.DB, params["name"])
// 	if !ok {
// 		respondHeader(w, http.StatusNotFound)
// 		return
// 	}

// 	// Check if permission resource exists --> 409
// 	_, ok = readUserPermission(server.DB, params["name"], params["permission-name"])
// 	if ok {
// 		respondHeader(w, http.StatusConflict)
// 		return
// 	}

// 	// Decode Body --> 500
// 	var newPermission Permission
// 	if err := json.NewDecoder(r.Body).Decode(&newPermission); err != nil {
// 		respondHeader(w, http.StatusInternalServerError)
// 		return
// 	}

// 	// Check if new permission is valid --> 422
// 	_, ok = readPermission(server.DB, newPermission.Name)
// 	if !ok || newPermission.Name != params["permission-name"] {
// 		respondHeader(w, http.StatusUnprocessableEntity)
// 		return
// 	}

// 	if rowCnt := insertUserPermission(server.DB, params["name"], newPermission.Name); rowCnt == 1 {
// 		log.Println("Permission inserted")
// 		respondHeader(w, http.StatusCreated)
// 	}
// }

// updateUserPermissionHandler handles requests for updating a permission of a user.
// func (server *Server) updateUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request: " + r.URL.Path)
// 	params := mux.Vars(r)

// 	// Check if user resource exists --> 404
// 	_, ok := readUser(server.DB, params["name"])
// 	if !ok {
// 		respondHeader(w, http.StatusNotFound)
// 		return
// 	}

// 	// Check if permission resource exists --> 404
// 	_, ok = readUserPermission(server.DB, params["name"], params["permission-name"])
// 	if !ok {
// 		respondHeader(w, http.StatusNotFound)
// 		return
// 	}

// 	// Decode Body --> 500
// 	var newPermission Permission
// 	if err := json.NewDecoder(r.Body).Decode(&newPermission); err != nil {
// 		respondHeader(w, http.StatusInternalServerError)
// 		return
// 	}

// 	// Check if new permission is valid --> 422
// 	_, ok = readPermission(server.DB, newPermission.Name)
// 	if !ok {
// 		respondHeader(w, http.StatusUnprocessableEntity)
// 		return
// 	}

// 	// Check if new permission is already set --> 409
// 	_, ok = readUserPermission(server.DB, params["name"], newPermission.Name)
// 	if ok {
// 		respondHeader(w, http.StatusConflict)
// 		return
// 	}

// 	// Update resource --> 200
// 	if rowCnt := updateUserPermission(server.DB, params["name"], params["permission-name"], newPermission.Name); rowCnt == 1 {
// 		log.Println("Permission updated")
// 		respondHeader(w, http.StatusOK)
// 	}

// 	// Create substrings --> Move to Client logic
// 	// substrs := strings.Split(params["permission-name"], "-")
// 	// dir := strings.Join(substrs[:len(substrs)-1], "-")
// 	// rw := strings.Join(substrs[len(substrs)-1:], "")

// 	// Check if old permission already exists --> 404
// 	// var oldPermissionName string
// 	// if rw == "read" {
// 	// 	oldPermissionName = dir + "-write"
// 	// } else {
// 	// 	oldPermissionName = dir + "-read"
// 	// }
// 	// log.Println(oldPermissionName)

// 	// _, ok = readUserPermission(db, params["name"], oldPermissionName)
// }

// func (server *Server) deleteUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request: " + r.URL.Path)
// 	params := mux.Vars(r)

// 	// Check if user resource exists --> 404
// 	_, ok := readUser(server.DB, params["name"])
// 	if !ok {
// 		respondHeader(w, http.StatusNotFound)
// 		return
// 	}

// 	// Check if permission resource exists --> 404
// 	_, ok = readUserPermission(server.DB, params["name"], params["permission-name"])
// 	if !ok {
// 		respondHeader(w, http.StatusNotFound)
// 		return
// 	}

// 	if rowCnt := deleteUserPermission(server.DB, params["name"], params["permission-name"]); rowCnt == 1 {
// 		respondHeader(w, http.StatusNoContent)
// 	}
// }

// Experimentell hanlders
func (server *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
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
