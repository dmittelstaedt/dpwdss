package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User represents an user
type User struct {
	ID                 int
	FirstName          string
	LastName           string
	Name               string
	Realm              string
	Role               string
	Password           string
	LastChange         time.Time
	LastChangePassword time.Time
	LastChangeRecord   time.Time
}

func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:david@tcp(127.0.0.1:3306)/pwdss")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

// Check if user is valid. Returns name and role of user.
func isValidUser(userName, password string) {
	// select name, role from users where username='$username' and password='$passwordHashed' limit 1
}

// Get name of all users from table users.
func getAllUsers() {
	// select name from users where role = '$role' order by username
}

// GetUser returns user struct with all information.
func getUser(db *sql.DB, userName, role string) (User, error) {
	var user User

	stmt, err := db.Prepare("select firstname, lastname, name from users where name=? and role=? limit 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userName, role)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.FirstName, &user.LastName, &user.Name); err != nil {
			log.Fatal(err)
		}
		log.Println("Firstname: " + user.FirstName + " Lastname: " + user.LastName)
	}

	return user, nil
}

// Add new user to table users.
func addUser(firstName, lastName, userName, permission string) {
	// insert into users (firstname, lastname, name, realm, role, permission, password, last_change_password) VALUES ('$firstName', '$lastName', '$username', '$realm', '$role', '$permission', '$passwordHashed', NOW())
}

func updateUser(firstName, lastName, userName, permission string) {
	// update users set firstname='$firstName', lastname='$lastName', permission='$permission' where name='$username'
}

// Remove existing user from table users.
func removeUser(userName string) {
	// delete from users where name='$userName'
}

// Update password of given user in table users.
func updatePassword(userName, currentPassword, newPassword string) {
	// update users set password='$newPasswordHashed' where name='$username' and password='$currentPasswordHashed'
}
