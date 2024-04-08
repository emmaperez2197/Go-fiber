package handlers

import (
	"github.com/emmaperez2197/go-fiber/database"
	"github.com/emmaperez2197/go-fiber/jwt"
	"github.com/emmaperez2197/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	error := c.BodyParser(&user)

	if error != nil {
		return c.Status(400).JSON(error)
	}

	email := user.Email

	err := db.Where("email = ?", email).First(&user).Error

	if err == nil {
		return c.Status(400).JSON("This email is already registered")
	}

	passwordPlane := []byte(user.Password)

	hash, err := bcrypt.GenerateFromPassword(passwordPlane, bcrypt.DefaultCost)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	user.Password = string(hash)

	db.Create(&user)

	return c.Status(200).JSON(user)
}

func LogIn(c *fiber.Ctx) error {
	db := database.DB

	var user models.User

	error := c.BodyParser(&user)

	if error != nil {
		return c.Status(400).JSON(error)
	}
	email := user.Email
	passwordHash := user.Password

	err := db.Where("email =?", email).First(&user).Error

	if err != nil {
		return c.Status(400).JSON("No exist user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordHash))

	if err != nil {
		return c.Status(400).JSON("contraseña incorrecta")
	}

	t, err := jwt.GeneratetToken(user)

	if err != nil {
		c.Status(400).JSON("Error generating token")
	}

	respose := fiber.Map{
		"token": t,
	}
	return c.Status(fiber.StatusOK).JSON(respose)
}
