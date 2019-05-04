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
	subrouter := server.Router.PathPrefix("/api/v1").Subrouter()
	subrouter.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	subrouter.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	subrouter.Methods("GET").Path("/users/{name}").Name("readUser").HandlerFunc(server.readUserHandler)
	subrouter.Methods("GET").Path("/groups").Name("readGroups").HandlerFunc(server.readGroupsHandler)
	subrouter.Methods("GET").Path("/groups/{name}").Name("readGroup").HandlerFunc(server.readGroupHandler)
	subrouter.Methods("GET").Path("/permissions").Name("readPermissions").HandlerFunc(server.readPermissionsHandler)
	subrouter.Methods("GET").Path("/permissions/{id}").Name("readPermission").HandlerFunc(server.readPermissionHandler)

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
