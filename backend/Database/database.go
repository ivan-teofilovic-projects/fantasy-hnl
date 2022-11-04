package database

import (
	"database/sql"
	"log"

	models "example.com/example/backend/Models"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(localhost:33060)/Players")

	if err != nil {
		log.Fatal(err)
	}
	pong := db.Ping()
	if pong != nil {
		log.Fatal(pong)
	}
	return db
}

func GetPlayerById(db *sql.DB) []models.Player {
	rows, err := db.Query(`SELECT * FROM players;`)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var Players []models.Player
	for rows.Next() {
		var u models.Player
		err := rows.Scan(&u.Id, &u.Name, &u.Team, &u.Position, &u.Value)
		if err != nil {
			log.Fatal(err)
		}
		Players = append(Players, u)
	}

	return Players
}
