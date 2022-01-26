package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webmasterdro/gopwer/controllers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

}
