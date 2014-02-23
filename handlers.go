package main

import (
	"github.com/coopernurse/gorp"
	"net/http"
)

func GetUnit(r *http.Request, db *gorp.DbMap) (units []Unit, err error) {
	_, err = db.Select(&units, "select * from units")
	return
}
