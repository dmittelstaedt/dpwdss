package main

import "fmt"

const apiEndpoint = "http://localhost:8080/api/v1/"

func main() {
	// users := readUsers()
	// printUsers(users)

	// fmt.Println("")

	// user := readUser(1)
	// printUser(user)

	// fmt.Println("")

	groups := readGroups()
	printGroups(groups)

	fmt.Println("")
	group := readGroup(1)
	printGroup(group)

	fmt.Println("")
}

// Create substrings --> Move to Client logic
// substrs := strings.Split(params["permission-name"], "-")
// dir := strings.Join(substrs[:len(substrs)-1], "-")
// rw := strings.Join(substrs[len(substrs)-1:], "")
