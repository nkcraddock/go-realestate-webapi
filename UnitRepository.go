package main

type IUnitRepository interface {
	IRepository
}

type UnitRepository struct {
	Repository
}
