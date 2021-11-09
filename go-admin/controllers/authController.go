package controllers

import (
	"govue/database"
	"govue/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
	//return c.SendString("Hello, World 👋!")
}

func Login(c *fiber.Ctx) error {
	user := models.User{}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Where("email= ?", data["email"]).First(&user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messages": "Incorect password",
		})
	}

	return c.JSON(user)
}
