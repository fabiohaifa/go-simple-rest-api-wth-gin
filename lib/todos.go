// S.O.L.I.D. principles:
// Single Responsibility Principle >> 1 job only!

package lib

import (
	mod "example/hello/model"
)

var ToDos = []mod.ToDo{
	{ID: 1, Text: "Learn Go", Done: true},
	{ID: 2, Text: "Learn Gin", Done: false},
	{ID: 3, Text: "Learn Docker", Done: true},
}

func GetAllToDos(Text string) []mod.ToDo {
	var retToDo []mod.ToDo
	if Text == "" {
		retToDo = ToDos
	} else {
		for i, _ := range ToDos {
			if isToDo(ToDos[i], Text) {
				retToDo = append(retToDo, ToDos[i])
			}
		}
	}
	return retToDo
}

func isToDo(item mod.ToDo, Content string) bool {
	return item.Text == Content
}
