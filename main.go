package main //everything eeds to be part of packages
import (
	"go-app/db"
	"go-app/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// need a starting point for GO compiler - main function

func main() {
	godotenv.Load(".env")
	db.CreateDatabase()
	router.RunRouter()
	//after reading the Book object form db , its included in a variable , to be available on indedx.html

	//listening, if any errors , log Fatal will close with code 1
	log.Fatal(http.ListenAndServe(":8000", nil))

}
