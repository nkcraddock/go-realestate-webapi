package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"net/http"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lesson4 sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	m.Get("/", func(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		rows, err := db.Query("select title, author, description from books")
		PanicIf(err)
		defer rows.Close()

		var title, author, description string
		for rows.Next() {
			err := rows.Scan(&title, &author, &description)
			PanicIf(err)
			fmt.Fprintf(rw, "%s, %s, %s", title, author, description)
		}
	})

	m.Run()
}
