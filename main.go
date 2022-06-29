package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var Todos = []Todo{
	{ID: 1, Text: "Learn Go", Done: false},
	{ID: 2, Text: "Learn Gin", Done: false},
	{ID: 3, Text: "Learn Docker", Done: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
