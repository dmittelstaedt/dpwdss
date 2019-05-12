package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

const userHeader = "ID\tFIRSTNAME\tLASTNAME\tNAME\tROLE"
const groupHeader = "ID\tNAME"
const permissionHeader = "ID\tUSER\tGROUP"

func printUser(user User) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, userHeader)
	fmt.Fprintln(w, strconv.Itoa(user.ID)+"\t"+user.FirstName+"\t"+user.LastName+"\t"+user.Name+"\t"+user.Role)
	fmt.Fprint(w)
	w.Flush()
}

func printUsers(users []User) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, userHeader)
	for _, user := range users {
		fmt.Fprintln(w, strconv.Itoa(user.ID)+"\t"+user.FirstName+"\t"+user.LastName+"\t"+user.Name+"\t"+user.Role)
	}
	fmt.Fprint(w)
	w.Flush()
}

func printGroup(group Group) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, groupHeader)
	fmt.Fprintln(w, strconv.Itoa(group.ID)+"\t"+group.Name)
	fmt.Fprint(w)
	w.Flush()
}

func printGroups(groups []Group) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, groupHeader)
	for _, group := range groups {
		fmt.Fprintln(w, strconv.Itoa(group.ID)+"\t"+group.Name)
	}
	fmt.Fprint(w)
	w.Flush()
}

func printPermission(permission PermissionOut) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, permissionHeader)
	fmt.Fprintln(w, strconv.Itoa(permission.ID)+"\t"+permission.UserName+"\t"+permission.GroupName)
	fmt.Fprint(w)
	w.Flush()
}

func printPermissions(permissions []PermissionOut) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, permissionHeader)
	for _, permission := range permissions {
		fmt.Fprintln(w, strconv.Itoa(permission.ID)+"\t"+permission.UserName+"\t"+permission.GroupName)
	}
	fmt.Fprint(w)
	w.Flush()
}
