package main

import (
	"server/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Shopy",
	})
	
	router.UserRouter(app)
	router.ItemRouter(app)

	app.Listen(":3001")
}
