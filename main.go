package main

import (
	"fixpay/reserva-mesa-sala/controllers"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//=============================== Routes ===============================
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.SignIn)
	app.Get("/private", controllers.Private)
	app.Get("/public", controllers.Public)

	//================================ Run =================================
	fmt.Println("Carregando servidor...")
	app.Listen(":8000")
}
