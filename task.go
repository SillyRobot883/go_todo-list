package main

type Tasks struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	isCompleted bool   `json:"isCompleted"`
}

func NewTasks(id int, title string, isCompleted bool, description string) *Tasks {
	return &Tasks{ID: id, Title: title, isCompleted: isCompleted, Description: description}
}
