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
	logRequest(r)
	name := r.URL.Query().Get("name")

	var users []User
	if name == "" {
		users = readUsers(server.DB)
	} else {
		users = readUsersWhereName(server.DB, name)
	}

	response, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readUserHandler handles requests for reading details of specific user.
func (server *Server) readUserHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	user, ok := readUser(server.DB, id)
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

// updatePermissionHandler handles requests for updateing a permission.
func (server *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	_, ok := readUser(server.DB, id)
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondHeader(w, http.StatusBadRequest)
		return
	}

	if rowCnt := updateUser(server.DB, user); rowCnt == 1 {
		response, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}
		respondJSON(w, http.StatusOK, response)
		return
	}
}

// readGroupsHandler handles requests for reading all group.
func (server *Server) readGroupsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	name := r.URL.Query().Get("name")

	var groups []Group
	if name == "" {
		groups = readGroups(server.DB)
	} else {
		groups = readGroupsWhereName(server.DB, name)
	}

	response, err := json.Marshal(groups)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// readUserHandler handles requests for reading details of specific group.
func (server *Server) readGroupHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	group, ok := readGroup(server.DB, id)
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
	logRequest(r)

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

// readPermissionHandler handles requests for reading permission for given id.
func (server *Server) readPermissionHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	permission, ok := readPermission(server.DB, id)
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(permission)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, http.StatusOK, response)
}

// insertPermissionHandler handles requests for inserting a permission.
func (server *Server) insertPermissionHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	var permission Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		respondHeader(w, http.StatusBadRequest)
		return
	}

	if rowCnt := insertPermission(server.DB, permission); rowCnt == 1 {
		response, err := json.Marshal(permission)
		if err != nil {
			log.Println(err)
		}
		respondJSON(w, http.StatusCreated, response)
		return
	}
}

// updatePermissionHandler handles requests for updateing a permission.
func (server *Server) updatePermissionHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	oldPermission, ok := readPermission(server.DB, id)
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	var newPermission Permission
	if err := json.NewDecoder(r.Body).Decode(&newPermission); err != nil {
		respondHeader(w, http.StatusBadRequest)
		return
	}

	if rowCnt := updatePermission(server.DB, oldPermission, newPermission); rowCnt == 1 {
		response, err := json.Marshal(newPermission)
		if err != nil {
			log.Println(err)
		}
		respondJSON(w, http.StatusOK, response)
		return
	}
}

// deletePermissionHandler handles requests for deleting a permission.
func (server *Server) deletePermissionHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		respondHeader(w, http.StatusBadRequest)
		return
	}

	_, ok := readPermission(server.DB, id)
	if !ok {
		respondHeader(w, http.StatusNotFound)
		return
	}

	if rowCnt := deletePermission(server.DB, id); rowCnt == 1 {
		respondHeader(w, http.StatusNoContent)
		return
	}
}
