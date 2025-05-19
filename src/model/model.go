package model

type User struct {
	Name     string
	Email    string
	Password string
}

type Music struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
	Writer   string `json:"writer"`
	Year     int    `json:"year"`
}
