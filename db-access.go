package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

// Read all permissions of given user.
func readPermissions(db *sql.DB, userName string) ([]string, error) {
	var permissions []string
	stmt, err := db.Prepare("select name from groups where id in (select group_id from user_to_group where user_id=(select id from users where name=?))")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var permission string
		if err = rows.Scan(&permission); err != nil {
			log.Fatal(err)
		}
		permissions = append(permissions, permission)
		log.Println("Group: " + permission)
	}

	return permissions, nil
}
