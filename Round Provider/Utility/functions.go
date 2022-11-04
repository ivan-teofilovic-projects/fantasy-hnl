package utility

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	models "example.com/round/Models"
)

func CalculateScore(player *models.Player) float64 {
	score := math.Min(float64(player.Statistics.MinutesPlayed), 90.0)/90.0*2 + float64(player.Statistics.Assists) + float64(player.Statistics.Goals) + float64(player.Statistics.Saves)/3*1
	return score
}

func GetResponse(url string, class interface{}) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&class)
	return resp.StatusCode
}

func GetCurrentRound() *models.Rounds {
	rounds_url := "https://api.sofascore.com/api/v1/unique-tournament/170/season/42138/rounds"
	rounds := new(models.Rounds)
	GetResponse(rounds_url, rounds)
	rounds.CurrentRound.Round -= 1
	return rounds
}

func GetEvents(rounds *models.Rounds) *models.Events {
	round_url := "https://api.sofascore.com/api/v1/unique-tournament/170/season/42138/events/round/" + strconv.Itoa(rounds.CurrentRound.Round)
	events := new(models.Events)
	GetResponse(round_url, events)
	return events
}
