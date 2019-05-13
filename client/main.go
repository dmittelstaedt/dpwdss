package main

import (
	"flag"
	"log"
	"strconv"
)

const apiEndpoint = "http://localhost:8080/api/v1/"

var userName string
var groupName string
var setPermission bool
var resetPassword bool
var version bool

// TODO: Check if flags are visit
// TODO: Version, Git Commit and Build Date as flag
// TODO: all users for specific group
// TODO: read Groups by name --> add to API

// Check if flags are correctly set, e.g. resetPassword in combination with groupName is not allowed
func checkCombination() bool {
	return false
}

func main() {
	flag.StringVar(&userName, "user", "", "Name of the user. Printing all users with value all.")
	flag.StringVar(&groupName, "group", "", "Name of the group. Printing all groups with value all.")
	flag.Parse()

	if userName == "all" && groupName == "" {
		users := readUsers()
		printUsers(users)
	}

	if userName != "" && userName != "all" && groupName == "" {
		user, ok := readUserByName(userName)
		if ok {
			printUser(user)
		}
	}

	if userName == "" && groupName == "all" {
		groups := readGroups()
		printGroups(groups)
	}

	if userName == "" && groupName != "" && groupName != "all" {
		groupID, err := strconv.Atoi(groupName)
		if err != nil {
			log.Println(err)
		}
		group := readGroup(groupID)
		printGroup(group)
	}

	if userName != "" && userName != "all" && groupName == "all" {
		permissionsOut := readPermissionsByUser(userName)
		printPermissions(permissionsOut)
	}
}

// Create substrings --> Move to Client logic
// substrs := strings.Split(params["permission-name"], "-")
// dir := strings.Join(substrs[:len(substrs)-1], "-")
// rw := strings.Join(substrs[len(substrs)-1:], "")
