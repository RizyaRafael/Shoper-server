package router

import (
	"server/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", controller.GetUser)

}
