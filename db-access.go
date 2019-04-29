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

// readUsers returns a slice with all users.
func readUsers(db *sql.DB) ([]User, error) {
	stmt, err := db.Prepare("select id, firstname, lastname, name, role from users")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Role); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// readUser returns user struct with all information.
func readUser(db *sql.DB, userName string) (User, error) {
	stmt, err := db.Prepare("select id, firstname, lastname, name, role from users where name=? limit 1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userName)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Role); err != nil {
			log.Println(err)
		}
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

// readPermissions returns all possible permissions.
func readPermissions(db *sql.DB) ([]Permission, error) {
	stmt, err := db.Prepare("select id, name from groups")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println()
	}
	defer rows.Close()

	var permissions []Permission
	for rows.Next() {
		var permission Permission
		if err = rows.Scan(&permission.ID, &permission.Name); err != nil {
			log.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// ReadPermissions returns all permissions of given user.
func readPermission(db *sql.DB, userName string) ([]Permission, error) {
	stmt, err := db.Prepare("select name from groups where id in (select group_id from user_to_group where user_id=(select id from users where name=?))")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userName)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var permissions []Permission
	for rows.Next() {
		var permission Permission
		if err = rows.Scan(&permission.Name); err != nil {
			log.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
