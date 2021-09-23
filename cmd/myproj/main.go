package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/labstack/echo"

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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment
		_, err := dbmap.Select(
			&comments,
			"SELECT * FROM comments ORDER BY created desc LIMIT 10",
		)
		if err != nil {
			c.Logger().Error("Select: ", err)
			return c.String(http.StatusBadRequest, "Select: "+err.Error())
		}
		return c.JSON(http.StatusOK, comments)
	})
	// e.Static("/", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
