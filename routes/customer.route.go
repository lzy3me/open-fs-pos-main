package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lzy3me/open-fs-pos-main/controllers"
)

func CustomerRoute(app fiber.Router) {
	api := app.Group("/customer")

	api.Get("/", controllers.CustomerController)
}