package database

import (
	"database/sql"
	"log"
)

func ConnectDatabase() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@(0.0.0.0:3307)/path")
	if err != nil {
		log.Fatal(err)
		log.Println("err")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		log.Println("err")
	}
	return db
}
