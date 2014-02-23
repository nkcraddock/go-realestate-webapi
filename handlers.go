package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
)

func GetAllUnits(repo IUnitRepository) (int, string) {
	units := repo.GetAll()
	return http.StatusOK, string(jsonEncode(units))
}

func AddUnit(rw http.ResponseWriter, u Unit, repo IUnitRepository) (int, string) {
	repo.Add(&u)
	rw.Header().Set("Location", fmt.Sprintf("/unit/%d", u.Id))
	return http.StatusCreated, ""
}

func GetUnit(parms martini.Params, repo IUnitRepository) (int, string) {
	id, err := strconv.Atoi(parms["id"])

	if err != nil {
		return notFound()
	}

	u := repo.Get(id)

	if u == nil {
		return notFound()
	}

	return http.StatusOK, string(jsonEncode(u))
}

func UpdateUnit(u Unit, repo IUnitRepository) (int, string) {
	cnt := repo.Update(&u)
	if cnt < 1 {
		return notFound()
	}
	return http.StatusOK, string(cnt)
}

func DeleteUnit(parms martini.Params, repo IUnitRepository) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	if err != nil {
		return notFound()
	}

	u := repo.Get(id)
	if u == nil {
		return notFound()
	}

	cnt := repo.Delete(u)
	return http.StatusOK, string(cnt)
}

func jsonEncode(o interface{}) []byte {
	jsonResponse, err := json.MarshalIndent(o, "", "    ")
	PanicIf(err)
	return jsonResponse
}

func notFound() (int, string) {
	return http.StatusNotFound, "Not found"
}
