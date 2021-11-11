package controllers

import (
	"govue/database"
	"govue/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetAll(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Password, _ = bcrypt.GenerateFromPassword([]byte("1234"), 14)

	database.DB.Create(&user)

	return c.JSON(user)
}
