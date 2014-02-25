package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"reflect"
	"strconv"
)

var GetAllUnits func(IUnitRepository) (int, string)
var GetUnit func(martini.Params, IUnitRepository) (int, string)

//var AddUnit func(http.ResponseWriter, Unit, IUnitRepository)

func init() {
	makeFunc := func(base func([]reflect.Value) []reflect.Value, fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), base)
		fn.Set(v)
	}

	// getAll(Repository) (int, string)
	getAll := func(in []reflect.Value) []reflect.Value {
		values := in[0].MethodByName("GetAll").Call([]reflect.Value{})
		// values is []reflect.Value returned by reflect.Call.
		// Since GetAll only returns interface{}, we just want the first object in values
		jsonResponse := string(jsonEncode(values[0].Interface()))
		return genericHandlerReturn(http.StatusFound, jsonResponse)
	}

	makeFunc(getAll, &GetAllUnits)

	/*func AddUnit(rw http.ResponseWriter, u Unit, repo IUnitRepository) (int, string) {
		repo.Add(&u)
		rw.Header().Set("Location", fmt.Sprintf("/unit/%d", u.Id))
		return http.StatusCreated, ""
	}
	// add(http.ResponseWriter, entity, Repository) (int, string)
	add := func(in []reflect.Value) []reflect.Value {
		in[2].MethodByName("Add").Call([]reflect.Value{in[1]})
		header := in[0].MethodByName("Header").Call(nil)
		location := reflect.ValueOf("Location")
		locationValue := reflect.ValueOf(fmt.Sprintf("/unit/%d", in[1].FieldByName("Id")))
		reflect.ValueOf(header).MethodByName("Set").Call([]reflect.Value{location, locationValue})
		return genericHandlerReturn(http.StatusCreated, "")
	}

	makeFunc(add, &AddUnit)
	*/

	// get(martini.Params, Repository) (int, string)
	get := func(in []reflect.Value) []reflect.Value {
		params := in[0].Interface().(martini.Params)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			return notFoundGeneric()
		}

		inGet := []reflect.Value{reflect.ValueOf(id)}
		values := in[1].MethodByName("Get").Call(inGet)

		if values[0].IsNil() {
			return notFoundGeneric()
		}

		jsonResponse := string(jsonEncode(values[0].Interface()))
		return []reflect.Value{reflect.ValueOf(http.StatusOK), reflect.ValueOf(jsonResponse)}
	}

	makeFunc(get, &GetUnit)

}

func AddUnit(rw http.ResponseWriter, u Unit, repo IUnitRepository) (int, string) {
	repo.Add(&u)
	rw.Header().Set("Location", fmt.Sprintf("/unit/%d", u.Id))
	return http.StatusCreated, ""
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

func genericHandlerReturn(status int, text string) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(status), reflect.ValueOf(text)}
}

func notFoundGeneric() []reflect.Value {
	return genericHandlerReturn(notFound())
}
