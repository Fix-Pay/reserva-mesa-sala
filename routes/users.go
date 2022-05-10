package routes

import (
	"github.com/Fix-Pay/reserva-mesa-sala/controllers"
	"github.com/gofiber/fiber/v2"
)

func HandleUsers(app *fiber.App) error {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Post("/users/auth", controllers.Auth)
	app.Post("/tables/create", controllers.CreateTables)
	app.Get("/tables/:id", controllers.GetTables)

	return nil
}
