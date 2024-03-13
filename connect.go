package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "root:password@tcp(127.0.0.1:5432)/postgres")
	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	log.Println("OK ... Not Error")
}
