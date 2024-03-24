package main //everything eeds to be part of packages
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// need a starting point for GO compiler - main function
var db *sql.DB

type Book struct {
	Name   string
	Author string
	Stock  uint
}

type Film struct {
	Title    string
	Director string
}

func main() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	var name, author string
	var quantity uint

	row, err := db.Query("SELECT * from Books")
	if err != nil {
		log.Fatal(err)
	}
	row.Next()
	err = row.Scan(&name, &author, &quantity)
	if err != nil {
		log.Fatal(err)
	}
	//after reading the Book object form db , its included in a variable , to be available on indedx.html

	films := map[string][]Film{
		"Films": {
			{Title: "Pirates 4", Director: "Someone De Niro"},
			{Title: "Toy Story", Director: "Tony Stark"},
			{Title: "Grand Casino", Director: "Mellstroy"},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	}
	//handles the function and assigns an endpoint
	http.HandleFunc("/", h1)

	//listening, if any errors , log Fatal will close with code 1
	log.Fatal(http.ListenAndServe(":8000", nil))

}
