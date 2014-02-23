package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

type IUnitRepository interface {
	GetAll() []*Unit
	Get(id int) *Unit
	Add(u *Unit)
	Update(u *Unit) int
	Delete(u *Unit) int
}

type UnitRepository struct {
}

var UnitRepo IUnitRepository
var dbmap *gorp.DbMap

func init() {
	dbConnection, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicIf(err)

	dbmap = &gorp.DbMap{Db: dbConnection, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Unit{}, "units").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	PanicIf(err)

	UnitRepo = &UnitRepository{}
}

func (_ *UnitRepository) GetAll() (units []*Unit) {
	_, err := dbmap.Select(&units, "select * from units")
	PanicIf(err)
	return
}

func (_ *UnitRepository) Add(u *Unit) {
	err := dbmap.Insert(u)
	PanicIf(err)
}

func (_ *UnitRepository) Get(id int) *Unit {
	obj, err := dbmap.Get(Unit{}, id)
	PanicIf(err)
	return obj.(*Unit)
}

func (_ *UnitRepository) Update(u *Unit) int {
	cnt, err := dbmap.Update(u)
	PanicIf(err)
	return int(cnt)
}

func (_ *UnitRepository) Delete(u *Unit) int {
	cnt, err := dbmap.Delete(u)
	PanicIf(err)
	return int(cnt)
}

type Unit struct {
	Id      int
	Address string
}
