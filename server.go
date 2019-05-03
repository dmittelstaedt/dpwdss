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
	// TODO: json extension to path, e.g. /api/v1/users.json, /api/v1/users/luke.json
	// TODO: at begining 3 Return Codes are enough:
	// successful --> 200
	// client-side error --> 400 bad request
	// server-side error --> 500 internal server
	// TODO: include message in body if failure
	// TODO: get permissions with queries, e.g. /permissions?uname=luke;gname=d1-read
	// TODO: validate user input in request bodies --> implement validator
	subrouter := server.Router.PathPrefix("/api/v1").Subrouter()
	subrouter.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	server.Router.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	server.Router.Methods("GET").Path("/users/{name}").Name("readUser").HandlerFunc(server.readUserHandler)
	server.Router.Methods("GET").Path("/groups").Name("readGroups").HandlerFunc(server.readGroupsHandler)
	server.Router.Methods("GET").Path("/groups/{name}").Name("readGroup").HandlerFunc(server.readGroupHandler)
	server.Router.Methods("GET").Path("/permissions").Name("readPermissions").HandlerFunc(server.readPermissionsHandler)
	server.Router.Methods("GET").Path("/permissions/{id}").Name("readPermission").HandlerFunc(server.readPermissionHandler)

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
