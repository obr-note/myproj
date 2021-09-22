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

	row := db.QueryRow(`SELECT name, age FROM users WHERE id=$1`, 1)
	var name string
	var age int64
	err = row.Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, age)
}
