package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	database "example.com/example/backend/Database"
)

func Home(db *sql.DB, w http.ResponseWriter) {
	players := database.GetPlayerById(db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(players)
}
