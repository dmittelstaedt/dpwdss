package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/permissions/{user-name}").Name("permissions").HandlerFunc(readPermissionsHandler)

	// Experimentell routes
	router.Methods("GET").Path("/users").Name("users").HandlerFunc(readAllUsersHandler)
	router.Methods("GET").Path("/users/{id}").Name("getUser").HandlerFunc(readUserHandler)
	router.Methods("POST").Path("/update").Name("update").HandlerFunc(updateUserHandler)
	return router
}
