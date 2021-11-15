package controllers

import (
	"govue/database"
	"govue/models"
	"govue/util"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.Password = user.SetPassword([]byte(data["password"]))

	database.DB.Create(&user)
	database.DB.Preload("Role").Find(&user)

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

	if err := user.ComparePassword([]byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messages": "Incorect password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(token)
}

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Preload("Role").Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Messages": "Success logout",
	})

}
