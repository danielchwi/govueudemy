package controllers

import (
	"govue/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messages": "Password do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
		Password:  data["password"],
	}

	return c.JSON(user)
	//return c.SendString("Hello, World 👋!")
}
