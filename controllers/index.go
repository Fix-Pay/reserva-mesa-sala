package controllers

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
)

func HomePageHandler(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", c.Params("api"))
	return c.SendString(msg) // => ✋ register
}
