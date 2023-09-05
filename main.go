package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Lib "api-wth-gin/lib"
)

func getToDos(context *gin.Context) {
	name := context.Param("text")
	r := Lib.T.GetAllToDos(0, name)
	if len(r) > 0 {
		context.IndentedJSON(http.StatusOK, r)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
}

func main() {

	router := gin.Default()
	// Retrive ToDos filtering by text
	router.GET("/todo-list/:text", getToDos)
	// Retrive all ToDos
	router.GET("/todo-list", getToDos)
	// Start the server at 9090 port
	router.Run("0.0.0.0:80")

}
