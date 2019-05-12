package main

import (
	"encoding/json"
	"log"
	"strconv"
)

func readUsers() []User {
	resp := sendRequest("GET", apiEndpoint+"users")
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Println(err)
	}

	return users
}

func readUser(id int) User {
	resp := sendRequest("GET", apiEndpoint+"users/"+strconv.Itoa(id))
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Println(err)
	}

	return user
}

func readGroups() []Group {
	resp := sendRequest("GET", apiEndpoint+"groups")
	defer resp.Body.Close()

	var groups []Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		log.Println(err)
	}
	return groups
}

func readGroup(id int) Group {
	resp := sendRequest("GET", apiEndpoint+"groups/"+strconv.Itoa(id))
	defer resp.Body.Close()

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		log.Println(err)
	}
	return group
}

func readPermissions() []Permission {
	resp := sendRequest("GET", apiEndpoint+"permissions")
	defer resp.Body.Close()

	var permissions []Permission
	if err := json.NewDecoder(resp.Body).Decode(&permissions); err != nil {
		log.Println(err)
	}

	return permissions
}

func readPermission(id int) Permission {
	resp := sendRequest("GET", apiEndpoint+"permissions"+strconv.Itoa(id))
	defer resp.Body.Close()

	var permission Permission
	if err := json.NewDecoder(resp.Body).Decode(&permission); err != nil {
		log.Println(err)
	}
	return permission
}

func updatePermission(id int) Permission {
	return Permission{}
}

func createPermission() Permission {
	return Permission{}
}
