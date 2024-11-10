package main

import (
	"database/sql"
	"fmt"
	"log"
	"mcs_bab_7/database"
	"mcs_bab_7/routers"
	"os"

	_ "github.com/lib/pq"
)

// input
// go get -u "github.com/gin-gonic/gin"
// go get -u "github.com/lib/pq"
// go get -u "github.com/rubenv/sql-migrate"
// go get -u "github.com/joho/godotenv"

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbName   = "praktikum_mcs_bab7"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	// var PORT = ":49000"

	// psqlInfo := fmt.Sprintf(
	// 	`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	host, port, user, password, dbName,
	// )

	psqlInfo := os.Getenv("DATABASE_URL")
	var PORT = ":" + os.Getenv("PORT")

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error open DB", psqlInfo)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run(PORT)
	fmt.Println("DB Success Connected")
	fmt.Printf("Server running on port %v", PORT)
}
