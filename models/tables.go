package models

import "github.com/Fix-Pay/reserva-mesa-sala/db"

type Table struct {
	Local  string `json:"local"`
	Status bool   `json:"status"`
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
