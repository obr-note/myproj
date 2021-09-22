package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows := db.QueryRow(`SELECT id, name, age FROM users`)
	var id int64
	var name string
	var age myType
	err = rows.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, name, age)
}

type myType string

func (mt *myType) Scan(src interface{}) error {
	switch x := src.(type) {
	case int64:
		*mt = myType(fmt.Sprint(src))
	case float64:
		*mt = myType(fmt.Sprintf("%.2f", src))
	case bool:
		*mt = myType(fmt.Sprint(src))
	case []byte:
		if len(x) < 10 {
			*mt = myType(fmt.Sprintf("[% 02X]", x))
		} else {
			*mt = myType(fmt.Sprintf("[% 02X...]", x))
		}
	case string:
		*mt = myType(x)
	case time.Time:
		*mt = myType(x.Format("2006/01/02 15:04:05"))
	case nil:
		*mt = "nil"
	}
	return nil
}
