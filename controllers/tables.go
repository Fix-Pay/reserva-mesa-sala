package controllers

import (
	"github.com/Fix-Pay/reserva-mesa-sala/db"
	"github.com/Fix-Pay/reserva-mesa-sala/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func CreateTables(c *fiber.Ctx) error {
	var table models.Table
	if err := c.BodyParser(&table); err != nil {
		log.Fatalln(err)
	}
	if err := table.Create(); err != nil {
		log.Fatalln(err)
	}
	return c.JSON(table)
}

func GetTables(c *fiber.Ctx) error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}
	var tables []models.Table
	engine.Find(&tables)
	return c.JSON(tables)
}
func CreateReserva(c *fiber.Ctx) error {
	var enviado models.Enviado
	err := c.BodyParser(&enviado)
	if err != nil {
		log.Fatalln(err)
	}

	inicioint, err := strconv.Atoi(enviado.HoraInicio)
	if err != nil {
		log.Fatalln(err)
	}
	finalint, err := strconv.Atoi(enviado.HoraFinal)
	if err != nil {
		log.Fatalln(err)
	}

	for i := inicioint; i <= finalint; i++ {
		repete := strconv.Itoa(i)
		reserva := models.Reserva{enviado.TableId, enviado.Data, repete}
		if err = reserva.Create(); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
func GetReserva(c *fiber.Ctx) error {
	engine, err := db.CreateDB()
	if err != nil {
		return err
	}
	var reserva []models.Reserva
	engine.Find(&reserva)
	return c.JSON(reserva)
}
