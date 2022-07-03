package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Lib "example/hello/lib"
)

func getToDos(context *gin.Context) {
	name := context.Param("text")
	context.IndentedJSON(http.StatusOK, Lib.GetAllToDos(name))
}

func main() {
	router := gin.Default()
	// Retrive ToDos filtering by text
	router.GET("/todo-list/:text", getToDos)
	// Retrive all ToDos
	router.GET("/todo-list", getToDos)
	// Start the server at 9090 port
	router.Run("localhost:9090")
}
