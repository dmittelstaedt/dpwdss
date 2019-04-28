package main

func testDBAccess() {
	db := connection()
	// getUser(db, "david", "app_admin")
	readPermissions(db, "d1r")
	defer db.Close()
}

func main() {
	// router := newRouter()
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatal(err)
	// }

	testDBAccess()
}
