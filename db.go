package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

type IUnitRepository interface {
	GetAll() []*Unit
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

func (ur *UnitRepository) GetAll() (units []*Unit) {
	_, err := dbmap.Select(&units, "select * from units")
	PanicIf(err)
	return
}

type Unit struct {
	Id      int64
	Address string
	Created int64
}
