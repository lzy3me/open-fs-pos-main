package routes

import "github.com/gofiber/fiber/v2"

func RootRoute(app fiber.Router) {
	AuthRoute(app)
	AdminRoute(app)
	CustomerRoute(app)
}