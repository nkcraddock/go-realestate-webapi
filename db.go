package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

var dbmap *gorp.DbMap

func init() {
	dbConnection, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicIf(err)

	dbmap = &gorp.DbMap{Db: dbConnection, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Unit{}, "units").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	PanicIf(err)
}
