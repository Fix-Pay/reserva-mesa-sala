package main

import (
	"github.com/Fix-Pay/reserva-mesa-sala/db"
	"github.com/Fix-Pay/reserva-mesa-sala/models"
	"github.com/Fix-Pay/reserva-mesa-sala/routes"
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
	routes.HandleUsers(app)

	// Declaring port //
	if err := app.Listen(":8000"); err != nil {
		log.Fatalln("Erro listen and server", err)
	}
}
