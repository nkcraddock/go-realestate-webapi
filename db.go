package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

type UnitRepository interface {
	GetAll() []*Unit
}

func InitDb(connection string) *gorp.DbMap {
	db, err := sql.Open("postgres", connection)
	PanicIf(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Unit{}, "units").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()

	PanicIf(err)

	return dbmap
}

type Unit struct {
	Id      int64
	Address string
	Created int64
}
