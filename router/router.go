package router

import (
	"go-app/controller"

	"github.com/gin-gonic/gin"
)

func RunRouter() {
	router := gin.Default()
	router.GET("/book/:name", controller.GetBookByName)
	router.GET("/books", controller.GetBooks)
	router.POST("/newBooks", controller.AddBook)

	router.Run("localhost:8000")
}
