package main

import (
	"log"
	"net/http"
)

func main() {
	db := connection()
	defer db.Close()
	setDB(db)

	router := newRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
