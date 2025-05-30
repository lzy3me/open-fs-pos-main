package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lzy3me/open-fs-pos-main/controllers"
)

func AuthRoute(app fiber.Router) {
	api := app.Group("/auth")

	api.Post("/signin", controllers.LoginController)
}