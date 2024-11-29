package models

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Competed bool   `json:"completed"`
}
