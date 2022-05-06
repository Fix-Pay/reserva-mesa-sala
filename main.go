package main

import (
	"fixpay/reserva-mesa-sala/controllers"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}
type LoginRequest struct {
	Email    string
	Password string
}

func main() {
	app := fiber.New()

	fmt.Println("Carregando servidor...")

	app.Post("/signup", controllers.HandleSignup)
	app.Post("/login", controllers.HandleLogin)
	app.Get("/private", controllers.HandlePrivate)
	app.Get("/public", controllers.HandlePublic)
	app.Listen(":8000")
}
