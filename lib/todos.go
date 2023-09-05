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
	{ID: 1, Text: "Learn Go", Done: false},
	{ID: 2, Text: "Learn Gin", Done: false},
	{ID: 3, Text: "Learn Docker", Done: true},
	{ID: 4, Text: "Learn ACI", Done: false},
	{ID: 5, Text: "Learn ACA", Done: false},
	{ID: 9997, Text: "Learn AKS", Done: false},
	{ID: 6, Text: "Deploying to ACR", Done: false},
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
