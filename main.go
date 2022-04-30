package main

import (
	"fixpay/reserva-mesa-sala/controllers"
	"fmt"
	goutils "github.com/armando-couto/goutils"
	fiber "github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/api/*", controllers.HomePageHandler)

	log.Fatal(app.Listen(fmt.Sprint(":", goutils.Godotenv("port_application"))))
}
