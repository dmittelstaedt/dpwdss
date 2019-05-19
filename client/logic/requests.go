package logic

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dmittelstaedt/dpwdss/client/models"
)

const apiEndpoint = "http://localhost:8080/api/v1/"

// ReadUsers returns all users.
func ReadUsers() []models.User {
	resp := sendRequest("GET", apiEndpoint+"users")
	defer resp.Body.Close()

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Println(err)
	}

	return users
}

// ReadUser returns user for given ID.
func ReadUser(id int) models.User {
	resp := sendRequest("GET", apiEndpoint+"users/"+strconv.Itoa(id))
	defer resp.Body.Close()

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Println(err)
	}

	return user
}

// ReadUserByName returns user for given name.
func ReadUserByName(name string) (models.User, bool) {
	resp := sendRequest("GET", apiEndpoint+"users?name="+name)
	defer resp.Body.Close()

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Println(err)
	}

	if len(users) == 1 {
		return users[0], true
	}

	return models.User{}, false
}

// ReadGroups returns all groups.
func ReadGroups() []models.Group {
	resp := sendRequest("GET", apiEndpoint+"groups")
	defer resp.Body.Close()

	var groups []models.Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		log.Println(err)
	}
	return groups
}

// ReadGroup returns group for given ID.
func readGroup(id int) models.Group {
	resp := sendRequest("GET", apiEndpoint+"groups/"+strconv.Itoa(id))
	defer resp.Body.Close()

	var group models.Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		log.Println(err)
	}
	return group
}

// ReadGroupByName returns group for given name.
func ReadGroupByName(name string) (models.Group, bool) {
	resp := sendRequest("GET", apiEndpoint+"groups?name="+name)
	defer resp.Body.Close()

	var groups []models.Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		log.Println(err)
	}

	if len(groups) == 1 {
		return groups[0], true
	}

	return models.Group{}, false
}

// ReadPermissions returns all permissions.
func ReadPermissions() []models.Permission {
	resp := sendRequest("GET", apiEndpoint+"permissions")
	defer resp.Body.Close()

	var permissions []models.Permission
	if err := json.NewDecoder(resp.Body).Decode(&permissions); err != nil {
		log.Println(err)
	}

	return permissions
}

// ReadPermissionsByUser returns all permissions for user with given name.
func ReadPermissionsByUser(name string) []models.PermissionOut {
	user, _ := ReadUserByName(name)
	permissions := ReadPermissions()

	var userPermissions []models.Permission

	for _, permission := range permissions {
		if permission.UserID == user.ID {
			userPermissions = append(userPermissions, permission)
		}
	}

	var permissionsOut []models.PermissionOut
	for _, permission := range userPermissions {
		group := readGroup(permission.GroupID)
		permissionsOut = append(permissionsOut, models.PermissionOut{
			ID:        permission.ID,
			UserName:  user.Name,
			GroupName: group.Name,
		})
	}
	return permissionsOut
}

// ReadPermission returns permission for given ID.
func ReadPermission(id int) models.Permission {
	resp := sendRequest("GET", apiEndpoint+"permissions"+strconv.Itoa(id))
	defer resp.Body.Close()

	var permission models.Permission
	if err := json.NewDecoder(resp.Body).Decode(&permission); err != nil {
		log.Println(err)
	}
	return permission
}

// UpdatePermission updates permission for given ID.
func UpdatePermission(id int) models.Permission {
	return models.Permission{}
}

// CreatePermission creates new permission.
func CreatePermission() models.Permission {
	return models.Permission{}
}

// sendRequest sends request for given method and URL.
func sendRequest(method, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	return resp
}
