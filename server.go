package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
)

var m *martini.Martini

func main() {
	m.Run()
}

func init() {
	m = martini.New()

	// middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	// injection
	m.MapTo(&UnitRepository{Type: Unit{}}, (*IUnitRepository)(nil))

	// routes
	r := martini.NewRouter()
	r.Get("/unit/:id", GetUnit)
	r.Get("/unit", GetAllUnits)
	r.Post("/unit", binding.Bind(Unit{}), AddUnit)
	r.Put("/unit", binding.Bind(Unit{}), UpdateUnit)
	r.Delete("/unit/:id", DeleteUnit)

	// wireup router action
	m.Action(r.Handle)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
