package db

import (
	"database/sql"
	"fmt"
	"go-app/goenv"
	"go-app/types"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func CreateDatabase() {

	cfg := mysql.Config{
		User:                 goenv.GoDotEnvVariable("USER"),
		Passwd:               goenv.GoDotEnvVariable("PASS"),
		Net:                  "tcp",
		Addr:                 goenv.GoDotEnvVariable("ADDR"),
		DBName:               goenv.GoDotEnvVariable("DATABASE"),
		AllowNativePasswords: true,
	}

	var err error

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

}

func GetBooks(DB *sql.DB) ([]types.Book, error) {

	rows, err := DB.Query("Select * from Books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []types.Book

	for rows.Next() {
		var book types.Book
		if err := rows.Scan(&book.Author, &book.Name, &book.Stock); err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}
	return books, nil
}
