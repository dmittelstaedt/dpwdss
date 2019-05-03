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

	rr := executeRequest(t, "GET", "/users")

	checkResponseCode(t, http.StatusOK, rr.Code)

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

	rr := executeRequest(t, "GET", "/users")

	checkResponseCode(t, http.StatusOK, rr.Code)

	var users []User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Log("Error during parsing respone body")
	}

	if len(users) != 2 {
		t.Errorf("Expected number of users %v, Got %v", 2, len(users))
	}
}

func TestReadNonExistentUser(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/users/test")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
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

	rr := executeRequest(t, "GET", "/users/luke")

	checkResponseCode(t, http.StatusOK, rr.Code)

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
	clearTables()

	rr := executeRequest(t, "GET", "/groups")

	checkResponseCode(t, http.StatusOK, rr.Code)

	var groups []Group
	if err := json.NewDecoder(rr.Body).Decode(&groups); err != nil {
		t.Log("Error during parsing respone body")
	}

	if len(groups) != 0 {
		t.Errorf("Expected number of groups %v, Got %v", 0, len(groups))
	}
}

func TestReadGroups(t *testing.T) {
	clearTables()
	d1Read := Group{
		ID:   1,
		Name: "d1-read",
	}
	d1Write := Group{
		ID:   2,
		Name: "d1-write",
	}
	addGroupT(t, d1Read)
	addGroupT(t, d1Write)

	rr := executeRequest(t, "GET", "/groups")

	checkResponseCode(t, http.StatusOK, rr.Code)

	var groups []Group
	if err := json.NewDecoder(rr.Body).Decode(&groups); err != nil {
		t.Log("Error during parsing respone body")
	}

	if len(groups) != 2 {
		t.Errorf("Expected number of groups %v, Got %v", 2, len(groups))
	}
}

func TestReadNonExistentGroup(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/groups/test")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestReadGroup(t *testing.T) {
	clearTables()
	d1Read := Group{
		ID:   1,
		Name: "d1-read",
	}
	addGroupT(t, d1Read)

	rr := executeRequest(t, "GET", "/groups/d1-read")

	checkResponseCode(t, http.StatusOK, rr.Code)

	var group Group
	if err := json.NewDecoder(rr.Body).Decode(&group); err != nil {
		t.Log("Error during parsing respone body")
	}

	if group != d1Read {
		t.Errorf("Expected user: %v. Got %v", group, d1Read)
	}
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
		t.Logf("Error inserting user: %v", err)
	}
}

func addGroupT(t *testing.T, group Group) {
	if _, err := testServer.DB.Exec("insert into groups (name) values (?)", group.Name); err != nil {
		t.Logf("Error inserting group: %v", err)
	}
}

func executeRequest(t *testing.T, method, route string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, route, nil)
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expectedCode, actualCode int) {
	if actualCode != expectedCode {
		t.Errorf("Expected response code %v. Got %v", expectedCode, actualCode)
	}
}
