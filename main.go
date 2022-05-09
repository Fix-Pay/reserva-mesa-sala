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

	//=============================== Routes ===============================
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.SignIn)
	app.Get("/private", controllers.Private)
	app.Get("/public", controllers.Public)

	//================================ Run =================================
	fmt.Println("Carregando servidor...")
	log.Fatal(app.Listen(fmt.Sprint(":", goutils.Godotenv("port_application"))))
}
