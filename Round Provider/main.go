package main

import (
	"net/http"

	handlers "example.com/round/Handlers"
	utility "example.com/round/Utility"
)

func main() {
	//get current round number
	rounds := utility.GetCurrentRound()

	//use previously obtained current round number to get all of the events of the round
	events := utility.GetEvents(rounds)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(w, r, events)
	})

	http.ListenAndServe("127.0.0.1:8083", nil)
}
