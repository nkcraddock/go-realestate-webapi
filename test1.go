package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.MapTo(UnitRepo, (*IUnitRepository)(nil))

	m.Get("/unit", func(rw http.ResponseWriter, r *http.Request, repo IUnitRepository) {
		units := repo.GetAll()

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
