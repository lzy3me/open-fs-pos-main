package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/lzy3me/open-fs-pos-main/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == fiber.StatusInternalServerError {
				log.Println(err.Error())
			}

			return c.Status(code).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	api := app.Group("/api")
	routes.RootRoute(api)
	fmt.Println("Server is running on port", port)
	if err := app.Listen(port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
