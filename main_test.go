package main

import (
	"bytes"
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

const pathPrefix = "/api/v1"

var testServer Server
var luke User
var han User
var d1Read Group
var d1Write Group
var p1 Permission
var p2 Permission

func TestMain(m *testing.M) {
	testServer = NewServer("root", "david", "pwdss_test")
	testServer.SetRoutes()

	initializeStructs()

	ensureTablesExist()

	clearTables()

	code := m.Run()

	os.Exit(code)
}

func TestReadEmptyUsers(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/users", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var users []User
	decodeResponse(t, rr.Body, &users)

	if len(users) != 0 {
		t.Errorf("Expected number of users %v, Got %v", 0, len(users))
	}
}

func TestReadUsers(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addUserT(t, han)

	rr := executeRequest(t, "GET", "/users", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var users []User
	decodeResponse(t, rr.Body, &users)

	if len(users) != 2 {
		t.Errorf("Expected number of users %v, Got %v", 2, len(users))
	}
}

func TestReadUsersWhereName(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addUserT(t, han)

	rr := executeRequest(t, "GET", "/users?name=luke", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var users []User
	decodeResponse(t, rr.Body, &users)

	if len(users) != 1 {
		t.Errorf("Expected number of users %v, Got %v", 1, len(users))
	}
}

func TestReadNonIDUser(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/users/test", nil)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestReadNonExistentUser(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/users/36", nil)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestReadUser(t *testing.T) {
	clearTables()
	addUserT(t, luke)

	rr := executeRequest(t, "GET", "/users/1", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var user User
	decodeResponse(t, rr.Body, &user)

	if user != luke {
		t.Errorf("Expected user %v. Got %v", luke, user)
	}
}

func TestUpdateNonIDUser(t *testing.T) {
	clearTables()

	body, err := json.Marshal(&luke)
	if err != nil {
		t.Logf("Error marhaling request body")
	}

	rr := executeRequest(t, "PUT", "/users/test", body)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateNonExistentUser(t *testing.T) {
	clearTables()

	body, err := json.Marshal(&luke)
	if err != nil {
		t.Logf("Error marshaling request body")
	}

	rr := executeRequest(t, "PUT", "/users/36", body)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestUpdateUser(t *testing.T) {
	clearTables()
	addUserT(t, luke)

	lukeNew := User{
		ID:        luke.ID,
		FirstName: luke.FirstName,
		LastName:  luke.LastName,
		Name:      luke.Name,
		Realm:     luke.Realm,
		Role:      luke.Role,
		Password:  "skywalker",
	}

	body, err := json.Marshal(&lukeNew)
	if err != nil {
		t.Logf("Error marshaling request body")
	}

	rr := executeRequest(t, "PUT", "/users/1", body)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var user User
	decodeResponse(t, rr.Body, &user)

	if user != lukeNew {
		t.Errorf("Expected user %v, Got %v", lukeNew, user)
	}
}

func TestReadEmptyGroups(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/groups", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var groups []Group
	decodeResponse(t, rr.Body, &groups)

	if len(groups) != 0 {
		t.Errorf("Expected number of groups %v, Got %v", 0, len(groups))
	}
}

func TestReadGroups(t *testing.T) {
	clearTables()
	addGroupT(t, d1Read)
	addGroupT(t, d1Write)

	rr := executeRequest(t, "GET", "/groups", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var groups []Group
	decodeResponse(t, rr.Body, &groups)

	if len(groups) != 2 {
		t.Errorf("Expected number of groups %v, Got %v", 2, len(groups))
	}
}

func TestReadGroupssWhereName(t *testing.T) {
	clearTables()
	addGroupT(t, d1Read)
	addGroupT(t, d1Write)

	rr := executeRequest(t, "GET", "/groups?name=d1-read", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var groups []Group
	decodeResponse(t, rr.Body, &groups)

	if len(groups) != 1 {
		t.Errorf("Expected number of groups %v, Got %v", 1, len(groups))
	}
}

func TestReadNonIDGroup(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/groups/test", nil)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestReadNonExistentGroup(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/groups/36", nil)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestReadGroup(t *testing.T) {
	clearTables()
	addGroupT(t, d1Read)

	rr := executeRequest(t, "GET", "/groups/1", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var group Group
	decodeResponse(t, rr.Body, &group)

	if group != d1Read {
		t.Errorf("Expected group %v. Got %v", d1Read, group)
	}
}

func TestReadEmptyPermissions(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/permissions", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var permissions []Permission
	decodeResponse(t, rr.Body, &permissions)

	if len(permissions) != 0 {
		t.Errorf("Expected number of groups %v, Got %v", 0, len(permissions))
	}
}

func TestReadPermissions(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addGroupT(t, d1Read)
	addPermissionT(t, p1)

	rr := executeRequest(t, "GET", "/permissions", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var permissions []Permission
	decodeResponse(t, rr.Body, &permissions)

	if len(permissions) != 1 {
		t.Errorf("Expected number of permissions %v, Got %v", 1, len(permissions))
	}
}

func TestReadNonExistentPermission(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/permissions/36", nil)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestReadNonIDPermission(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "GET", "/permissions/text", nil)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestReadPermission(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addGroupT(t, d1Read)
	addPermissionT(t, p1)

	rr := executeRequest(t, "GET", "/permissions/1", nil)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var permission Permission
	decodeResponse(t, rr.Body, &permission)

	if permission != p1 {
		t.Errorf("Expected permission %v, Got %v", p1, permission)
	}
}

func TestInsertPermission(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addGroupT(t, d1Read)

	body, err := json.Marshal(p1)
	if err != nil {
		t.Logf("Error parsing request")
	}

	rr := executeRequest(t, "POST", "/permissions", body)

	checkResponseCode(t, http.StatusCreated, rr.Code)

	var permission Permission
	decodeResponse(t, rr.Body, &permission)

	if permission != p1 {
		t.Errorf("Expected permission %v, Got %v", p1, permission)
	}
}

func TestUpdateNonExistentPermission(t *testing.T) {
	clearTables()

	body, err := json.Marshal(&p1)
	if err != nil {
		t.Logf("Error marshaling request body")
	}
	rr := executeRequest(t, "PUT", "/permissions/1", body)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestUpdateNonIDPermission(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "PUT", "/permissions/text", nil)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestUpdatePermission(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addUserT(t, han)
	addGroupT(t, d1Read)
	addGroupT(t, d1Write)
	addPermissionT(t, p1)
	addPermissionT(t, p2)

	p1New := Permission{
		ID:      p1.ID,
		UserID:  p1.UserID,
		GroupID: d1Write.ID,
	}

	body, err := json.Marshal(&p1New)
	if err != nil {
		t.Logf("Error marshaling request body")
	}
	rr := executeRequest(t, "PUT", "/permissions/1", body)

	checkResponseCode(t, http.StatusOK, rr.Code)

	var permission Permission
	decodeResponse(t, rr.Body, &permission)

	if permission != p1New {
		t.Errorf("Expected permission %v, Got %v", p1New, permission)
	}
}

func TestDeleteNonExistentPermission(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "DELETE", "/permissions/text", nil)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteNonIDPermission(t *testing.T) {
	clearTables()

	rr := executeRequest(t, "DELETE", "/permissions/1", nil)

	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestDeletePermission(t *testing.T) {
	clearTables()
	addUserT(t, luke)
	addUserT(t, han)
	addGroupT(t, d1Read)
	addGroupT(t, d1Write)
	addPermissionT(t, p1)
	addPermissionT(t, p2)

	rr := executeRequest(t, "DELETE", "/permissions/1", nil)

	checkResponseCode(t, http.StatusNoContent, rr.Code)
}

func initializeStructs() {
	luke = User{
		ID:        1,
		FirstName: "Luke",
		LastName:  "Skywalker",
		Name:      "luke",
		Realm:     "share",
		Role:      "app_user",
		Password:  "luke",
	}

	han = User{
		ID:        2,
		FirstName: "Han",
		LastName:  "Solo",
		Name:      "han",
		Realm:     "share",
		Role:      "app_user",
		Password:  "han",
	}

	d1Read = Group{
		ID:   1,
		Name: "d1-read",
	}

	d1Write = Group{
		ID:   2,
		Name: "d1-write",
	}

	p1 = Permission{
		ID:      1,
		UserID:  1,
		GroupID: 1,
	}

	p2 = Permission{
		ID:      2,
		UserID:  2,
		GroupID: 2,
	}
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
	testServer.DB.Exec("DELETE FROM user_to_group")
	testServer.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
	testServer.DB.Exec("ALTER TABLE groups AUTO_INCREMENT = 1")
	testServer.DB.Exec("ALTER TABLE user_to_group AUTO_INCREMENT = 1")
}

func addUserT(t *testing.T, user User) {
	if _, err := testServer.DB.Exec("insert into users (firstname, lastname, name, realm, role, password) VALUES(?, ?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Name, user.Realm, user.Role, user.Password); err != nil {
		t.Logf("Error inserting user %v", err)
	}
}

func addGroupT(t *testing.T, group Group) {
	if _, err := testServer.DB.Exec("insert into groups (name) values (?)", group.Name); err != nil {
		t.Logf("Error inserting group %v", err)
	}
}

func addPermissionT(t *testing.T, permission Permission) {
	if _, err := testServer.DB.Exec("insert into user_to_group (user_id, group_id) values (?,?)", permission.UserID, permission.GroupID); err != nil {
		t.Logf("Error inserting permission %v", err)
	}
}

func executeRequest(t *testing.T, method, route string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, pathPrefix+route, bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	testServer.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expectedCode, actualCode int) {
	if actualCode != expectedCode {
		t.Errorf("Expected response code %v. Got %v", expectedCode, actualCode)
	}
}

func decodeResponse(t *testing.T, body *bytes.Buffer, v interface{}) {
	if err := json.NewDecoder(body).Decode(v); err != nil {
		t.Log("Error during parsing respone body")
	}
}
