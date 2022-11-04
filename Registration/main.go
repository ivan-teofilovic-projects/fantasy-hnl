package main

import (
	"log"
	"net/http"

	database "example.com/greetings/Database"
	handlers "example.com/greetings/Handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.ConnectDatabase()
	defer db.Close()
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(r, db, w)
	})
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignIn(r, db, w)
	})

	log.Fatal(http.ListenAndServeTLS("127.0.0.1:9000", "localhost.crt", "localhost.key", nil))
}
