package models

import (
	"github.com/Fix-Pay/reserva-mesa-sala/db"
)

type Table struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	Id     uint   `json:"id" gorm:"primaryKey"`
}

type Reserva struct {
	TableId      string `json:"table_id"`
	DiaReservado string `json:"dia_reservado"`
	HoraCoberta  string `json:"hora_coberta"`
}
type Enviado struct {
	TableId    string `json:"table_id"`
	Data       string `json:"data"`
	HoraInicio string `json:"hora_inicio"`
	HoraFinal  string `json:"hora_final"`
}

func (t *Table) Create() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}
	result := engine.Create(t)
	return result.Error
}

func (r *Reserva) Create() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}

	result := engine.Create(r)
	return result.Error
}
