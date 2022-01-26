package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/webmasterdro/gopwer/config"
	"github.com/webmasterdro/gopwer/database"
	"github.com/webmasterdro/gopwer/hashing"
	"github.com/webmasterdro/gopwer/models"
	"time"
)

func Login(c *fiber.Ctx) error {
	fmt.Println("Login")

	return nil
}

func Register(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}

	errors := models.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	hash, _ := hashing.HashPassword(user.Name, user.Passwd, config.Config("PASSWORD_HASH"))

	user.Passwd = hash
	user.Birthday = time.Now()
	user.Creatime = time.Now()

	// Save user to database
	db := database.DB
	db.Create(&user)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "User created",
		"user":    user,
	})
}
