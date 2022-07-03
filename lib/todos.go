// S.O.L.I.D. principles:
// S ==> Single Responsibility Principle >> 1 job only!
//

package lib

import (
	mod "example/hello/model"
)

var ToDos = []mod.ToDo{
	{ID: 1, Text: "Learn Go", Done: true},
	{ID: 2, Text: "Learn Gin", Done: false},
	{ID: 3, Text: "Learn Docker", Done: true},
}

type T int

func (t T) GetAllToDos(text string) []mod.ToDo {
	var retToDo []mod.ToDo
	if text == "" {
		retToDo = ToDos
	} else {
		for i, _ := range ToDos {
			if t.isToDo(ToDos[i], text) {
				retToDo = append(retToDo, ToDos[i])
			}
		}
	}
	return retToDo
}

func (t T) isToDo(item mod.ToDo, Content string) bool {
	return item.Text == Content
}
