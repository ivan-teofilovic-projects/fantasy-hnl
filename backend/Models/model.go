package models

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Team     string `json:"team"`
	Position string `json:"position"`
	Value    string `json:"value"`
}
