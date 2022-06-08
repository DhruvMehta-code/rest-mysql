package main

import (
	"rest-go/rest-go-mysql/mesage"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	route := router.Group("/message")

	route.POST("/", mesage.ReadMessage)
	route.GET("/", mesage.GetAll)
	route.GET("/:id", mesage.GetOne)

	router.Run("localhost:8080")
}
