package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Map(InitDb("dbname=test sslmode=disable"))

	m.Get("/unit", func(rw http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
		units, err := GetUnit(r, db)
		PanicIf(err)
		for _, u := range units {
			fmt.Fprintf(rw, "%s", u.Address)
		}
	})

	m.Run()
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
