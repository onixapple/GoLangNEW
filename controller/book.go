package controller

import (
	"go-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(context *gin.Context) {
	books := model.GetBooks()
	context.IndentedJSON(http.StatusOK, books)
}
