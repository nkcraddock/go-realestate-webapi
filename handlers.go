package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
)

func GetAllUnits(rw http.ResponseWriter, r *http.Request, repo IUnitRepository) {
	units := repo.GetAll()
	jsonResponse, err := json.Marshal(units)
	PanicIf(err)
	rw.Write(jsonResponse)
}

func AddUnit(rw http.ResponseWriter, r *http.Request, repo IUnitRepository) {
	u := getUnitFromRequest(r)
	repo.Add(u)
	jsonResponse, err := json.Marshal(u.Id)
	PanicIf(err)
	rw.Write(jsonResponse)
}

func GetUnit(rw http.ResponseWriter, parms martini.Params, repo IUnitRepository) {
	id, err := strconv.Atoi(parms["id"])
	PanicIf(err)
	u := repo.Get(id)
	jsonResponse, err := json.MarshalIndent(u, "", "    ")
	PanicIf(err)
	rw.Write(jsonResponse)
}

func UpdateUnit(rw http.ResponseWriter, r *http.Request, repo IUnitRepository) {
	u := getUnitFromRequest(r)
	cnt := repo.Update(u)
	writeJson(rw, cnt)
}

func DeleteUnit(rw http.ResponseWriter, parms martini.Params, repo IUnitRepository) {
	id, err := strconv.Atoi(parms["id"])
	PanicIf(err)
	u := repo.Get(id)
	cnt := repo.Delete(u)
	writeJson(rw, cnt)
}

func writeJson(rw http.ResponseWriter, o interface{}) {
	jsonResponse, err := json.MarshalIndent(o, "", "    ")
	PanicIf(err)
	rw.Write(jsonResponse)
}

func getUnitFromRequest(r *http.Request) *Unit {
	idstr, address := r.FormValue("id"), r.FormValue("address")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		id = 0
	}

	return &Unit{
		Id:      id,
		Address: address,
	}
}
