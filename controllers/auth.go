package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/webmasterdro/gopwer/config"
	"github.com/webmasterdro/gopwer/database"
	"github.com/webmasterdro/gopwer/hashing"
	"github.com/webmasterdro/gopwer/models"
	"time"
)

func Login(c *fiber.Ctx) error {

	type Login struct {
		Name     string `json:"name" validate:"required,min=3,max=32"`
		Password string `json:"password" validate:"required,min=6,max=32"`
	}

	type UserData struct {
		Name     string `json:"name"`
		Truename string `json:"truename"`
	}

	input := Login{}
	ud := UserData{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid input",
		})
	}

	errors := models.ValidateStruct(input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	user, err := models.GetUserByUserName(input.Name)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Wrong username",
		})
	}

	hash, _ := hashing.CheckPassword(input.Name, input.Password, user.Passwd)

	if !hash {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Wrong username or password",
		})
	}

	ud = UserData{
		Name:     user.Name,
		Truename: user.Truename,
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = ud.Name
	claims["truename"] = ud.Truename
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"token":  tokenString,
		"user":   ud,
	})
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

	hash, _ := hashing.HashPassword(user.Name, user.Passwd)

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
