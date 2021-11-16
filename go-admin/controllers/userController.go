package controllers

import (
	"govue/database"
	"govue/models"
	"govue/util"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := ((page - 1) * limit)

	var users []models.User
	var total int64

	database.DB.Preload("Role").Offset(int(offset)).Limit(limit).Find(&users)
	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Password = user.SetPassword(user.Password)

	database.DB.Create(&user)
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}
	database.DB.Delete(&user)

	return nil
}

func UpdateInfo(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Updates(&user)
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)

}

func UpdatePassword(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
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
		Id: uint(userId),
	}
	user.SetPassword([]byte(data["password"]))

	database.DB.Updates(&user)
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}
