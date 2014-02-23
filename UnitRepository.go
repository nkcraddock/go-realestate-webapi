package main

type IUnitRepository interface {
	GetAll() interface{}
	Get(id int) interface{}
	Add(u interface{})
	Update(u interface{}) int
	Delete(u interface{}) int
}

type UnitRepository struct {
	Type interface{}
}

func (_ *UnitRepository) GetAll() interface{} {
	var units []Unit // figure out how to make this from Type
	_, err := dbmap.Select(&units, "select * from units")
	PanicIf(err)
	return units
}

func (_ *UnitRepository) Add(u interface{}) {
	err := dbmap.Insert(u)
	PanicIf(err)
}

func (repo *UnitRepository) Get(id int) interface{} {
	obj, err := dbmap.Get(repo.Type, id)
	PanicIf(err)

	if obj == nil {
		return nil
	}

	return obj
}

func (_ *UnitRepository) Update(u interface{}) int {
	cnt, err := dbmap.Update(u)
	PanicIf(err)
	return int(cnt)
}

func (_ *UnitRepository) Delete(u interface{}) int {
	cnt, err := dbmap.Delete(u)
	PanicIf(err)
	return int(cnt)
}
