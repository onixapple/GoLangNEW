package db

import (
	"database/sql"
	"fmt"
	"go-app/types"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

var DB *sql.DB

func CreateDatabase() {

	cfg := mysql.Config{
		User:                 goDotEnvVariable("USER"),
		Passwd:               goDotEnvVariable("PASS"),
		Net:                  "tcp",
		Addr:                 goDotEnvVariable("ADDR"),
		DBName:               goDotEnvVariable("DATABASE"),
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
