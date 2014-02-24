package main

type IUnitRepository interface {
	GetAll() interface{}
	Get(id int) interface{}
	Add(u interface{})
	Update(u interface{}) int
	Delete(u interface{}) int
}

type UnitRepository struct {
	Repository
}
