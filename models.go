package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Server holds database and router for the application
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

// User represents an user
type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firsname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Name      string `json:"name,omitempty"`
	Realm     string `json:"realm,omitempty"`
	Role      string `json:"role,omitempty"`
	Password  string `json:"password,omitempty"`
}

// Permission represents a permission
type Permission struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// TODO: strcut for holding groups
