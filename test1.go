package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()

	m.MapTo(UnitRepo, (*IUnitRepository)(nil))

	m.Get("/unit/:id", GetUnit)
	m.Get("/unit", GetAllUnits)
	m.Post("/unit", binding.Bind(Unit{}), AddUnit)
	m.Put("/unit", binding.Bind(Unit{}), UpdateUnit)
	m.Delete("/unit/:id", DeleteUnit)

	m.Run()
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
