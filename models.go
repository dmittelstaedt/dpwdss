package main

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
