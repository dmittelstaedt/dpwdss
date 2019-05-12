package main

const apiEndpoint = "http://localhost:8080/api/v1/"

func getUserPermissions() {
	user := readUser(4)
	permissions := readPermissions()

	var userPermissions []Permission

	for _, permission := range permissions {
		if permission.UserID == user.ID {
			userPermissions = append(userPermissions, permission)
		}
	}

	var permissionsOut []PermissionOut
	for _, permission := range userPermissions {
		group := readGroup(permission.GroupID)
		permissionsOut = append(permissionsOut, PermissionOut{
			ID:        permission.ID,
			UserName:  user.Name,
			GroupName: group.Name,
		})
	}
	printPermissions(permissionsOut)
}

func main() {
	// users := readUsers()
	// printUsers(users)

	// fmt.Println("")

	// user := readUser(1)
	// printUser(user)

	// fmt.Println("")

	// groups := readGroups()
	// printGroups(groups)

	// fmt.Println("")
	// group := readGroup(1)
	// printGroup(group)

	// fmt.Println("")

	getUserPermissions()
}

// Create substrings --> Move to Client logic
// substrs := strings.Split(params["permission-name"], "-")
// dir := strings.Join(substrs[:len(substrs)-1], "-")
// rw := strings.Join(substrs[len(substrs)-1:], "")
