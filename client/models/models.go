package models

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

// Group represents a group.
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Permission represents a permission.
type Permission struct {
	ID      int `json:"id"`
	UserID  int `json:"userid"`
	GroupID int `json:"groupid"`
}

// PermissionOut represents a permission for printing to console.
type PermissionOut struct {
	ID        int
	UserName  string
	GroupName string
}
