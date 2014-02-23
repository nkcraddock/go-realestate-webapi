package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.MapTo(UnitRepo, (*IUnitRepository)(nil))

	m.Get("/unit/:id", GetUnit)
	m.Get("/unit", GetAllUnits)
	m.Post("/unit", AddUnit)
	m.Put("/unit", UpdateUnit)
	m.Delete("/unit/:id", DeleteUnit)

	m.Run()
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
