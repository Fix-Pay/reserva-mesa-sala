package main

import (
	"fixpay/reserva-mesa-sala/controllers"
	"fmt"
	goutils "github.com/armando-couto/goutils"
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Tables struct {
	gorm.Model
	ID int
    PLACES string
}

func main() {
	app := fiber.New()

	//GET /api/register
	app.Get("/api/*", controllers.HomePageHandler)

	log.Fatal(app.Listen(fmt.Sprint(":", goutils.Godotenv("port_application"))))

	dsn := "host=localhost user=postgres password=postgres dbname=fixpay_mesa_sala port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(err)

	// Read
	var table Tables
	tablea := db.First(&table, 1) // find product with integer primary key
	fmt.Println(tablea)
	//db.First(&table, "code = ?", "D42") // find product with code D42
}