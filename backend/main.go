package main

import (
	"net/http"

	database "example.com/example/backend/Database"
	handlers "example.com/example/backend/Handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.Connect()
	defer db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(db, w)

	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
