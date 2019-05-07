package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
)

const userHeader = "ID\tFIRSTNAME\tLASTNAME\tNAME\tROLE"

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

func executeRequest(method, url string) *http.Response {
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

func readUsers() []User {
	resp := executeRequest("GET", "http://localhost:8080/api/v1/users")
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Println(err)
	}

	return users
}

func readUser(id int) User {
	resp := executeRequest("GET", "http://localhost:8080/api/v1/users/"+strconv.Itoa(id))
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Println(err)
	}

	return user
}

func main() {
	users := readUsers()
	printUsers(users)

	user := readUser(1)
	printUser(user)
}

// Create substrings --> Move to Client logic
// substrs := strings.Split(params["permission-name"], "-")
// dir := strings.Join(substrs[:len(substrs)-1], "-")
// rw := strings.Join(substrs[len(substrs)-1:], "")
