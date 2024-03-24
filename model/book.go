package model

import (
	"errors"
	"go-app/db"
	"go-app/types"
	"log"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Stock  uint   `json:"stock"`
}

func GetBooks() []types.Book {
	books, err := db.GetBooks(db.DB)

	if err != nil {
		log.Fatal(err)
	}

	return books

}

func AddBook(book types.Book) (uint, error) {
	id, err := db.AddBook(book)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetBookByName(name string) (types.Book, error) {
	var book types.Book

	book, err := db.GetBookByName(name)

	if err != nil {
		return book, errors.New("book not found")
	}
	return book, nil
}
