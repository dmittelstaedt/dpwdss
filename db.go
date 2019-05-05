package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// readUsers returns a slice with all users.
func readUsers(db *sql.DB) []User {
	stmt, err := db.Prepare("select id, firstname, lastname, name, realm, role, password from users")
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
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Realm, &user.Role, &user.Password); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users
}

// readUsers returns a slice with all users based on the given name.
func readUsersWhereName(db *sql.DB, name string) []User {
	stmt, err := db.Prepare("select id, firstname, lastname, name, realm, role, password from users where name=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Realm, &user.Role, &user.Password); err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users
}

// readUser returns user struct with all information.
func readUser(db *sql.DB, id int) (User, bool) {
	stmt, err := db.Prepare("select id, firstname, lastname, name, realm, role, password from users where id=? limit 1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	var user User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Realm, &user.Role, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	ok := false
	if err == nil {
		ok = true
	}

	return user, ok
}

// updateUser updates the given user.
func updateUser(db *sql.DB, user User) int64 {
	stmt, err := db.Prepare("update users set firstname=?, lastname=?, realm=?, role=?, password=? where id=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.FirstName, user.LastName, user.Realm, user.Role, user.Password, user.ID)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}

// readUsers returns a slice with all users.
func readGroups(db *sql.DB) []Group {
	stmt, err := db.Prepare("select id, name from groups")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		if err = rows.Scan(&group.ID, &group.Name); err != nil {
			log.Println(err)
		}
		groups = append(groups, group)
	}
	return groups
}

func readGroup(db *sql.DB, id int) (Group, bool) {
	stmt, err := db.Prepare("select id, name from groups where id=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	var group Group
	err = stmt.QueryRow(id).Scan(&group.ID, &group.Name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	ok := false
	if err == nil {
		ok = true
	}

	return group, ok
}

// Add new user to table users.
// func addUser(firstName, lastName, userName, permission string) {
// 	// insert into users (firstname, lastname, name, realm, role, permission, password, last_change_password) VALUES ('$firstName', '$lastName', '$username', '$realm', '$role', '$permission', '$passwordHashed', NOW())
// }

// Remove existing user from table users.
// func removeUser(userName string) {
// 	// delete from users where name='$userName'
// }

// readPermissions returns all possible permissions.
func readPermissions(db *sql.DB) ([]Permission, error) {
	// select user_to_group.id, users.name, groups.name from user_to_group inner join users on user_to_group.user_id = users.id inner join groups on user_to_group.group_id = groups.id;
	stmt, err := db.Prepare("select id, user_id, group_id from user_to_group")
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
		if err = rows.Scan(&permission.ID, &permission.UserID, &permission.GroupID); err != nil {
			log.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// readPermissions returns Permission for given name of a Permission.
func readPermission(db *sql.DB, permissionID int) (Permission, bool) {
	stmt, err := db.Prepare("select id, user_id, group_id from user_to_group where id=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	var permission Permission
	err = stmt.QueryRow(permissionID).Scan(&permission.ID, &permission.UserID, &permission.GroupID)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	ok := false
	if err == nil {
		ok = true
	}

	return permission, ok
}

// insertPermission inserts a new permission with the given user_id and group_id.
func insertPermission(db *sql.DB, permmission Permission) int64 {
	stmt, err := db.Prepare("insert into user_to_group (user_id, group_id) values (?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(permmission.UserID, permmission.GroupID)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}

// updatePermission updates the old permission with the new permission. Returns number of affecetd rows.
func updatePermission(db *sql.DB, oldPermission, newPermission Permission) int64 {
	stmt, err := db.Prepare("update user_to_group set user_id=?, group_id=? where id=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newPermission.UserID, newPermission.GroupID, oldPermission.ID)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}

// deletePermission deletes the given permission.
func deletePermission(db *sql.DB, id int) int64 {
	stmt, err := db.Prepare("delete from user_to_group where id=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}
