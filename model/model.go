package model

type ToDoHandler interface {
	GetAllToDos(text string) []ToDo
	isToDo(item ToDo, Content string) bool
}

type ToDo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
