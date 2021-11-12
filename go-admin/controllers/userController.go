package controllers

import (
	"govue/database"
	"govue/models"

	"github.com/gofiber/fiber/v2"
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

	user.Password = user.SetPassword(user.Password)

	database.DB.Create(&user)

	return c.JSON(user)
}
