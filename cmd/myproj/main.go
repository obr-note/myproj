package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(`INSERT INTO users(name, age) VALUES($1, $2)`, "Bob", 18)
	if err != nil {
		log.Fatal(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(affected)
}
