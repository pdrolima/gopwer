package middleware

import "C"
import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/webmasterdro/gopwer/config"
)

func Authenticated() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: errorHandler,
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
