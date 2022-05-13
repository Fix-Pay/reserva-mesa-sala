package controllers

import "github.com/gofiber/fiber/v2"

func Private(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "path": "private"})
}

func Public(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "path": "public"})
}
