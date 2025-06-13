package controller

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name": "Robert",
	})
}
