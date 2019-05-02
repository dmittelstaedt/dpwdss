package main

func main() {

}

// import (
// 	"encoding/json"
// 	"net/http"
// 	"testing"
// )

// func TestReadUsers(t *testing.T) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", "http://localhost:8080/users", nil)
// 	if err != nil {
// 		t.Log(err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Log(err)
// 	}
// 	defer resp.Body.Close()

// 	statusCode := resp.StatusCode
// 	if statusCode != 200 {
// 		t.Errorf("readUsers was incorrect, got %v, want: %v", statusCode, 200)
// 	}

// 	var users []User
// 	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
// 		t.Log("Error during parsing respone body")
// 	}
// 	if len(users) != 7 {
// 		t.Errorf("readUsers was incorrect, got %v, want %v", len(users), 7)
// 	}

// 	dummy := User{
// 		ID:        1,
// 		FirstName: "David",
// 		LastName:  "Mittelstaedt",
// 		Name:      "david",
// 		Role:      "app_admin",
// 	}

// 	found := false
// 	for _, user := range users {
// 		if user == dummy {
// 			found = true
// 		}
// 	}
// 	if !found {
// 		t.Errorf("readUsers was incorrect, got %v, want %v", "", dummy)
// 	}
// }
