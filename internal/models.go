package models

type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Error struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}
