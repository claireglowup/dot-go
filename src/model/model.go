package model

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type Music struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
	Writer   string `json:"writer"`
	Year     int    `json:"year"`
}
