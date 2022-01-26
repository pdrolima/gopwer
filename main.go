package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/webmasterdro/gopwer/database"
	"github.com/webmasterdro/gopwer/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.Connect()

	// Setup routes
	router.Routes(app)

	app.Listen(":3000")
}
