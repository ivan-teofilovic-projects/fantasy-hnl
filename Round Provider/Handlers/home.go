package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "example.com/round/Models"
	utility "example.com/round/Utility"
)

func Home(w http.ResponseWriter, r *http.Request, events *models.Events) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var list []int
	json.NewDecoder(r.Body).Decode(&list)
	score := 0.0
	for _, player_id := range list {
		for _, event := range events.Events {
			player_url := "https://api.sofascore.com/api/v1/event/" + strconv.Itoa(event.Id) + "/player/" + strconv.Itoa(player_id) + "/statistics"
			player := new(models.Player)
			code := utility.GetResponse(player_url, player)
			if code != 404 {
				player_score := utility.CalculateScore(player)
				score += player_score
			}
		}
	}
	json.NewEncoder(w).Encode(score)
}
