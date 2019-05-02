package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const usersCreationQuery = `
create table if not exists users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	firstname VARCHAR(64) NOT NULL,
	lastname VARCHAR(64) NOT NULL,
	name VARCHAR(64) NOT NULL,
	realm VARCHAR(64) NOT NULL,
	role enum('app_user', 'app_admin') NOT NULL,
	permission enum('read', 'read_write') NOT NULL,
	password VARCHAR(64) NOT NULL,
	last_change TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	last_change_password DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_change_record TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	CONSTRAINT uc_user UNIQUE (id,name)
  )
`

const groupsCreationQuery = `
create table if not exists groups (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(32) NOT NULL,
	CONSTRAINT gc_group UNIQUE (id,name)
  )
`

const permissionsCreationQuery = `
create table if not exists user_to_group (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	user_id INT(6) UNSIGNED,
	group_id INT(6) UNSIGNED,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (group_id) REFERENCES groups(id),
	CONSTRAINT c_utg UNIQUE (id)
  )
`

var testServer Server

func TestMain(m *testing.M) {
	testServer = NewServer("root", "david", "pwdss_test")
	testServer.SetRoutes()

	ensureTablesExist()

	code := m.Run()

	clearTables()

	os.Exit(code)
}

func TestReadEmptyUsers(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %v. Got %v", http.StatusOK, rr.Code)
	}

	var users []User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Log("Error during parsing respone body")
	}

	if len(users) != 0 {
		t.Errorf("Expected number of users %v, Got %v", 0, len(users))
	}
}

func TestReadUsers(t *testing.T) {
	clearTables()
	luke := User{
		ID:        1,
		FirstName: "Luke",
		LastName:  "Skywalker",
		Name:      "luke",
		Role:      "app_user",
	}
	han := User{
		ID:        2,
		FirstName: "Han",
		LastName:  "Solo",
		Name:      "han",
		Role:      "app_user",
	}
	addUserT(t, luke)
	addUserT(t, han)

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %v. Got %v", http.StatusOK, rr.Code)
	}

	var users []User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Log("Error during parsing respone body")
	}

	if len(users) != 2 {
		t.Errorf("Expected number of users %v, Got %v", 0, len(users))
	}
}

func TestReadNonExistentUser(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/user/test", nil)
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected response code %v. Got %v", http.StatusNotFound, rr.Code)
	}
}

func TestReadUser(t *testing.T) {
	clearTables()
	luke := User{
		ID:        1,
		FirstName: "Luke",
		LastName:  "Skywalker",
		Name:      "luke",
		Role:      "app_user",
	}
	addUserT(t, luke)

	req, _ := http.NewRequest("GET", "/user/luke", nil)
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %v. Got %v", http.StatusNotFound, rr.Code)
	}

	var user User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Log("Error during parsing respone body")
	}

	if user != luke {
		t.Errorf("Expected user: %v. Got %v", luke, user)
	}

}

func TestUpdateNonExistentUser(t *testing.T) {
	// TODO: Implement
}

func TestUpdateUser(t *testing.T) {
	// TODO: Implement
}

func TestReadEmptyGroups(t *testing.T) {
	// TODO: Implement
}

func TestReadGroups(t *testing.T) {
	// TODO: Implement
}

func TestReadNonExistentGroup(t *testing.T) {
	// TODO: Implement
}

func TestReadGroup(t *testing.T) {
	// TODO: Implement
}

func TestReadEmptyPermissions(t *testing.T) {
	// TODO: Implement
}

func TestReadPermissions(t *testing.T) {
	// TODO: Implement
}

func TestReadNonExistentPermissions(t *testing.T) {
	// TODO: Implement
}

func TestReadPermission(t *testing.T) {
	// TODO: Implement
}

func TestUpdateNonExistentPermission(t *testing.T) {
	// TODO: Implement
}

func TestUpdatePermission(t *testing.T) {
	// TODO: Implement
}

func TestInsertPermission(t *testing.T) {
	// TODO: Implement
}

func TestDeletePermission(t *testing.T) {
	// TODO: Implement
}

func ensureTablesExist() {
	if _, err := testServer.DB.Exec(usersCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := testServer.DB.Exec(groupsCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := testServer.DB.Exec(permissionsCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTables() {
	testServer.DB.Exec("DELETE FROM users")
	testServer.DB.Exec("DELETE FROM groups")
	testServer.DB.Exec("DELETE FROM user_to_groups")
	testServer.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
	testServer.DB.Exec("ALTER TABLE groups AUTO_INCREMENT = 1")
	testServer.DB.Exec("ALTER TABLE user_to_groups AUTO_INCREMENT = 1")
}

func addUserT(t *testing.T, user User) {
	if _, err := testServer.DB.Exec("insert into users (firstname, lastname, name, role) VALUES(?, ?, ?, ?)", user.FirstName, user.LastName, user.Name, user.Role); err != nil {
		t.Logf("Error inserting user")
	}
}
