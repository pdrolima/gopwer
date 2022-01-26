package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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

	hash, err := hashing.HashPassword("teste", "teste", "md5")

	if err != nil {
		fmt.Println(err)
	}

	db := database.DB

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}

	user.Passwd = "0x" + hash
	user.Birthday = time.Now()
	user.Creatime = time.Now()

	// Save user to database
	db.Create(&user)

	return nil
}
