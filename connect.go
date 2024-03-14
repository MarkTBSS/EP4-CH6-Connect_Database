package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	database, err := sql.Open("postgres", "user=root password=Pass1234 dbname=postgres host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	log.Println("OK ... Not Error")
}
