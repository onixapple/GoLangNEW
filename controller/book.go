package controller

import (
	"go-app/model"
	"go-app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(context *gin.Context) {
	books := model.GetBooks()
	context.IndentedJSON(http.StatusOK, books)
}

func AddBook(context *gin.Context) {
	var newBook types.Book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	id, err := model.AddBook(newBook)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
	context.IndentedJSON(http.StatusCreated, id)
}
