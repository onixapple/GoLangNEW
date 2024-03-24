package router

import (
	"go-app/controller"

	"github.com/gin-gonic/gin"
)

var context *gin.Context

func RunRouter() {
	router := gin.Default()

	router.GET("/books", controller.GetBooks)

	router.Run("localhost:8000")
}
