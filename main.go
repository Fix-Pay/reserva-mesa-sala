package main

import (
	"github.com/Fix-Pay/reserva-mesa-sala/db"
	"github.com/Fix-Pay/reserva-mesa-sala/models"
	"github.com/Fix-Pay/reserva-mesa-sala/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// Migration with GORM //
	engine, err := db.CreateDB()
	if err != nil {
		log.Fatalln(err)
	}
	if err := engine.AutoMigrate(&models.User{}); err != nil {
		log.Fatalln(err)
	}
	// Declaring Routes //
	if err := routes.HandleUsers(app); err != nil {
		log.Fatalln(err)
	}

	// Declaring port //
	if err := app.Listen(":8000"); err != nil {
		log.Fatalln("Deu erro doidao", err)
	}
}
