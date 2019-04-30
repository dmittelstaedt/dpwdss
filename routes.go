package main

import "github.com/gorilla/mux"

// newRouter returns router with all defined routes.
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/users").Name("readUsers").HandlerFunc(readUsersHandler)
	router.Methods("GET").Path("/user/{name}").Name("readUser").HandlerFunc(readUserHandler)
	router.Methods("GET").Path("/user/{name}/permissions").Name("readUserPermissions").HandlerFunc(readUserPermissionsHandler)
	router.Methods("GET").Path("/user/{name}/permission/{permission-name}").Name("readUserPermission").HandlerFunc(readUserPermissionHandler)
	router.Methods("PUT").Path("/user/{name}/permission/{permission-name}").Name("updateUserPermission").HandlerFunc(updateUserPermissionHandler)
	router.Methods("GET").Path("/permissions").Name("readPermissions").HandlerFunc(readPermissionsHandler)
	router.Methods("GET").Path("/permission/{name}").Name("readPermission").HandlerFunc(readPermissionHandler)

	// Experimentell routes
	router.Methods("POST").Path("/update").Name("update").HandlerFunc(updateUserHandler)

	return router
}
