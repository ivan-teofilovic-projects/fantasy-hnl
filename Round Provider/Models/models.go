package models

type Round struct {
	Round int `json:"round"`
}

type Rounds struct {
	CurrentRound Round `json:"currentRound"`
}

type Events struct {
	Events []Event
}

type Event struct {
	HomeTeam       Team `json:"homeTeam"`
	AwayTeam       Team `json:"awayTeam"`
	StartTimestamp int  `json:"startTimestamp"`
	Id             int  `json:"id"`
}

type Team struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Player struct {
	Player     General    `json:"player"`
	Team       Team       `json:"team"`
	Statistics Statistics `json:"statistics"`
}

type General struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Position string `json:"position"`
}

type Statistics struct {
	Goals         int `json:"goals"`
	MinutesPlayed int `json:"minutesPlayed"`
	Assists       int `json:"goalAssist"`
	Saves         int `json:"saves"`
}
