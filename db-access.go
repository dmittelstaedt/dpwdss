package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// readUsers returns a slice with all users.
func readUsers(db *sql.DB) []User {
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
	return users
}

// readUser returns user struct with all information.
func readUser(db *sql.DB, userName string) (User, bool) {
	stmt, err := db.Prepare("select id, firstname, lastname, name, role from users where name=? limit 1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	var user User
	err = stmt.QueryRow(userName).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Name, &user.Role)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	ok := false
	if err == nil {
		ok = true
	}

	return user, ok
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

func readGroup(db *sql.DB, groupName string) (Group, bool) {
	stmt, err := db.Prepare("select id, name from groups where name=?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	var group Group
	err = stmt.QueryRow(groupName).Scan(&group.ID, &group.Name)
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

// func updateUser(firstName, lastName, userName, permission string) {
// 	// update users set firstname='$firstName', lastname='$lastName', permission='$permission' where name='$username'
// }

// Remove existing user from table users.
// func removeUser(userName string) {
// 	// delete from users where name='$userName'
// }

// Update password of given user in table users.
// func updatePassword(userName, currentPassword, newPassword string) {
// 	// update users set password='$newPasswordHashed' where name='$username' and password='$currentPasswordHashed'
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

// readUserPermissions returns all permissions of given user.
// func readUserPermissions(db *sql.DB, userName string) ([]Permission, error) {
// 	stmt, err := db.Prepare("select id, name from groups where id in (select group_id from user_to_group where user_id=(select id from users where name=?))")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query(userName)
// 	if err != nil {
// 		log.Println()
// 	}
// 	defer rows.Close()

// 	var permissions []Permission
// 	for rows.Next() {
// 		var permission Permission
// 		if err = rows.Scan(&permission.ID, &permission.Name); err != nil {
// 			log.Println(err)
// 		}
// 		permissions = append(permissions, permission)
// 	}

// 	return permissions, nil
// }

// readUserPermission returns permission of given user and permission
// func readUserPermission(db *sql.DB, userName, permissionName string) (Permission, bool) {
// 	stmt, err := db.Prepare("select id, name from groups where id in (select group_id from user_to_group where user_id=(select id from users where name=?)) and name=?")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer stmt.Close()

// 	var permission Permission
// 	err = stmt.QueryRow(userName, permissionName).Scan(&permission.ID, &permission.Name)
// 	if err != nil && err != sql.ErrNoRows {
// 		log.Println(err)
// 	}

// 	ok := false
// 	if err == nil {
// 		ok = true
// 	}

// 	return permission, ok
// }

// insertUserPermission inserts a new permission for the given user.
func insertUserPermission(db *sql.DB, userName, permissionName string) int64 {
	stmt, err := db.Prepare("insert into user_to_group (user_id, group_id) values ((select id from users where name=?), (select id from groups where name=?))")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(userName, permissionName)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}

// updateUserPermission updates the old permission with tne new permission. Returns number of affecetd rows.
func updateUserPermission(db *sql.DB, userName, oldPermissionName, newPermissionName string) int64 {
	stmt, err := db.Prepare("update user_to_group set group_id=(select id from groups where name=?) where user_id=(select id from users where name=?) and group_id=(select id from groups where name=?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newPermissionName, userName, oldPermissionName)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	return rowCnt
}

// deleteUserPermission deletes the given permission for the given user.
func deleteUserPermission(db *sql.DB, userName, permissionName string) int64 {
	stmt, err := db.Prepare("delete from user_to_group where user_id=(select id from users where name=?) and group_id=(select id from groups where name=?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(userName, permissionName)
	if err != nil {
		log.Println(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rowCnt
}
