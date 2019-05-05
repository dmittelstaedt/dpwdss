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

	// users
	subrouter.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	subrouter.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(server.readUsersHandler)
	subrouter.Methods("GET").Path("/users/{name}").Name("readUser").HandlerFunc(server.readUserHandler)
	subrouter.Methods("PUT").Path("/users/{name}").Name("updateUser").HandlerFunc(server.updateUserHandler)

	// groups
	subrouter.Methods("GET").Path("/groups").Name("readGroups").HandlerFunc(server.readGroupsHandler)
	subrouter.Methods("GET").Path("/groups/{name}").Name("readGroup").HandlerFunc(server.readGroupHandler)

	// permissions
	subrouter.Methods("GET").Path("/permissions").Name("readPermissions").HandlerFunc(server.readPermissionsHandler)
	subrouter.Methods("GET").Path("/permissions/{id}").Name("readPermission").HandlerFunc(server.readPermissionHandler)
	subrouter.Methods("POST").Path("/permissions").Name("insertPermission").HandlerFunc(server.insertPermissionHandler)
	subrouter.Methods("PUT").Path("/permissions/{id}").Name("updatePermission").HandlerFunc(server.updatePermissionHandler)
	subrouter.Methods("DELETE").Path("/permissions/{id}").Name("deletePermission").HandlerFunc(server.deletePermissionHandler)
}

// Run starts the server
func (server *Server) Run() {
	if err := http.ListenAndServe(":8080", server.Router); err != nil {
		log.Fatal(err)
	}
	defer server.DB.Close()
}
