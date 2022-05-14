package models

import "github.com/Fix-Pay/reserva-mesa-sala/db"

type User struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Id       uint   `json:"id" gorm:"primaryKey"`
}

func (s *User) Get() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}

	result := engine.Where(s).First(s)
	return result.Error
}

// Create User in DB //
func (s *User) Create() error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}
	result := engine.Create(s)
	return result.Error
}
