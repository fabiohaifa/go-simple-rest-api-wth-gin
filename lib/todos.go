// S.O.L.I.D. principles:
// S ==> Single Responsibility Principle - a class should have only one reason to change
// O ==> Objects / Entities should open for extension, but closed for modification
// L ==> Liskov Substitution Principle - a class should be replaceable with another class with the same interface
// I ==> Interface Segregation Principle - a class should not depend on interfaces it does not use
// D ==> Dependency Inversion Principle - a class should depend on abstractions, not concrete classes

package lib

import (
	mod "api-wth-gin/model"
	"sort"
)

var ToDos = []mod.ToDo{
	{ID: 1, Text: "Learning Go", Done: true},
	{ID: 2, Text: "Learning Gin", Done: true},
	{ID: 3, Text: "Learning Docker", Done: true},
	{ID: 4, Text: "Learning Postman", Done: false},
	{ID: 6, Text: "Learning ACI", Done: false},
	{ID: 7, Text: "Learning ACA", Done: false},
	{ID: 9345, Text: "Learning AKS", Done: false},
	{ID: 5, Text: "Learning GitHub Actions", Done: false},
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

	// Sort ToDos by ID
	sort.Slice(retToDo, func(i, j int) bool {
		return retToDo[i].ID < retToDo[j].ID
	})

	return retToDo
}

func (t T) isToDo(item mod.ToDo, Content string) bool {
	return item.Text == Content
}
