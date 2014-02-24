package main

type IRepository interface {
	GetAll() interface{}
	Get(id int) interface{}
	Add(u interface{})
	Update(u interface{}) int
	Delete(u interface{}) int
}

type Repository struct {
	Type  interface{}
	Table string
}

func (repo *Repository) GetAll() interface{} {
	units, err := dbmap.Select(repo.Type, "select * from "+repo.Table)
	PanicIf(err)
	return units
}

func (repo *Repository) Get(id int) interface{} {
	obj, err := dbmap.Get(repo.Type, id)
	PanicIf(err)

	if obj == nil {
		return nil
	}

	return obj
}

func (_ *Repository) Add(u interface{}) {
	err := dbmap.Insert(u)
	PanicIf(err)
}

func (_ *Repository) Update(u interface{}) int {
	cnt, err := dbmap.Update(u)
	PanicIf(err)
	return int(cnt)
}

func (_ *Repository) Delete(u interface{}) int {
	cnt, err := dbmap.Delete(u)
	PanicIf(err)
	return int(cnt)
}
