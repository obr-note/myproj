package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type Comment struct {
	Id      int64     `db:"id,primarykey,autoincrement"`
	Name    string    `db:"name,notnull,default:'名無し',size:200"`
	Text    string    `db:"text,notnull,size:400"`
	Created time.Time `db:"created,notnull"`
	Updated time.Time `db:"updated,notnull"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, _ := sql.Open("postgres", dsn)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")
	err := dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
	err = dbmap.Insert(&Comment{Text: "こんにちわ"})
	if err != nil {
		log.Fatal(err)
	}
}
