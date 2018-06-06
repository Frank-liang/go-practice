package main

import (
	"fmt"
	"net/http"
)

type Database struct {
	Url string
}

func NewDatabase(url string) Database {
	return Database{url}
}

func main() {
	db := NewDatabase("localhost:5432")

	http.HandleFunc("/hello", hello(db))
	http.ListenAndServe(":3000", nil)
}

func hello(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}
