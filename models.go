package main

// User represents an user
type User struct {
	ID        int
	FirstName string
	LastName  string
	Name      string
	Realm     string
	Role      string
	Password  string
}

// Permission represents a permission of a user
type Permission struct {
	Name string
}
