package main

// User represents a user.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firsname"`
	LastName  string `json:"lastname"`
	Name      string `json:"name"`
	Realm     string `json:"realm"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}
