package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NewServer creates Database connection and sets routes. Returns Server struct.
func NewServer(user, password, dbName string) Server {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	return Server{router, db}
}

// SetRoutes sets the routes for the server
func (server *Server) SetRoutes() {
	// TODO: add version and json extension to path, e.g. /api/v1/users.json, /api/v1/users/luke.json
	// TODO: no singular --> always plural resources
	// TODO: at begining 3 Return Codes are enough:
	// successful --> 200
	// client-side error --> 400 bad request
	// server-side error --> 500 internal server
	// TODO: include message in body if failure
	server.Router.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	server.Router.Methods("GET").Path("/user/{name}").Name("readUser").HandlerFunc(server.readUserHandler)
	server.Router.Methods("GET").Path("/user/{name}/permissions").Name("readUserPermissions").HandlerFunc(server.readUserPermissionsHandler)

	server.Router.Methods("GET").Path("/user/{name}/permission/{permission-name}").Name("readUserPermission").HandlerFunc(server.readUserPermissionHandler)
	server.Router.Methods("POST").Path("/user/{name}/permission/{permission-name}").Name("insertUserPermission").HandlerFunc(server.insertUserPermissionHandler)
	server.Router.Methods("PUT").Path("/user/{name}/permission/{permission-name}").Name("updateUserPermission").HandlerFunc(server.updateUserPermissionHandler)
	server.Router.Methods("DELETE").Path("/user/{name}/permission/{permission-name}").Name("deleteUserPermission").HandlerFunc(server.deleteUserPermissionHandler)

	server.Router.Methods("GET").Path("/permissions").Name("readPermissions").HandlerFunc(server.readPermissionsHandler)
	server.Router.Methods("GET").Path("/permission/{name}").Name("readPermission").HandlerFunc(server.readPermissionHandler)

	// Experimentell routes
	server.Router.Methods("POST").Path("/update").Name("update").HandlerFunc(server.updateUserHandler)
}

// Run starts the server
func (server *Server) Run() {
	if err := http.ListenAndServe(":8080", server.Router); err != nil {
		log.Fatal(err)
	}
	defer server.DB.Close()
}
