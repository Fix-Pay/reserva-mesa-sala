package models

import "github.com/Fix-Pay/reserva-mesa-sala/db"

type Table struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	Place  string `json:"place"`
	Id     uint   `json:"id" gorm:"primaryKey"`
}

func (t *Table) Create() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}
	result := engine.Create(t)
	return result.Error
}
func (t *Table) Get() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}

	result := engine.Where(t).First(t)
	return result.Error
}
