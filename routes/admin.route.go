package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lzy3me/open-fs-pos-main/controllers"
)

func AdminRoute(app fiber.Router) {
	api := app.Group("/admin")
	
	api.Get("/", controllers.AdminController)
}