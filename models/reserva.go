package models

import (
	"github.com/Fix-Pay/reserva-mesa-sala/db"
)

type Reserva struct {
	TableId      string `json:"table_id"`
	DiaReservado string `json:"dia_reservado"`
	HoraCoberta  string `json:"hora_coberta"`
}

func (r *Reserva) Create() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}

	result := engine.Create(r)
	return result.Error
}
