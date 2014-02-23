package main

type IUnitRepository interface {
	GetAll() []*Unit
	Get(id int) *Unit
	Add(u *Unit)
	Update(u *Unit) int
	Delete(u *Unit) int
}

type UnitRepository struct {
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

	if obj == nil {
		return nil
	}

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
