package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
)

func GetAllUnits(rw http.ResponseWriter, repo IUnitRepository) {
	units := repo.GetAll()
	writeJson(rw, units)
}

func AddUnit(rw http.ResponseWriter, u Unit, repo IUnitRepository) {
	repo.Add(&u)
	writeJson(rw, u.Id)
}

func GetUnit(rw http.ResponseWriter, parms martini.Params, repo IUnitRepository) {
	id, err := strconv.Atoi(parms["id"])
	PanicIf(err)
	u := repo.Get(id)
	writeJson(rw, u)
}

func UpdateUnit(rw http.ResponseWriter, u Unit, repo IUnitRepository) {
	cnt := repo.Update(&u)
	writeJson(rw, cnt)
}

func DeleteUnit(rw http.ResponseWriter, parms martini.Params, repo IUnitRepository) {
	id, err := strconv.Atoi(parms["id"])
	PanicIf(err)
	u := repo.Get(id)

	if u == nil {
		panic("Unit " + parms["id"] + " not found.")
	}

	cnt := repo.Delete(u)
	writeJson(rw, cnt)
}

func writeJson(rw http.ResponseWriter, o interface{}) {
	jsonResponse, err := json.MarshalIndent(o, "", "    ")
	PanicIf(err)
	rw.Write(jsonResponse)
}
