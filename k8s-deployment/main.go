package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", print)
	router.Run(":8081")
}

func print(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hey dude, welcome to the demo application!")
}
