package controllers

import (
	"govue/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "daniel",
	}
	user.Lastname = "widodo"
	return c.JSON(user)
	//return c.SendString("Hello, World 👋!")
}
