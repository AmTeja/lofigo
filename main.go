package main

import (
	"github.com/amteja/lofig/database"
	"github.com/amteja/lofig/env"
	"github.com/amteja/lofig/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	env.Setup()

	database.Connect()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.PostRoutes(app)
	routes.UserRoutes(app)

	app.Listen(":3000")
}
