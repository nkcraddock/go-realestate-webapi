package main

type Unit struct {
	Id      int    `form:"id" json:"id"`
	Address string `form:"address" json:"address"`
}
