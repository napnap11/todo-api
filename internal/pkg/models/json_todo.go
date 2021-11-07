package models

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Status string `json:"status"`
	Image  string `json:"image"`
}
