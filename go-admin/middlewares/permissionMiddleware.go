package middlewares

import (
	"errors"
	"govue/database"
	"govue/models"
	"govue/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorize(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")
	id, err := util.ParseJwt(cookie)
	if err != nil {
		return err
	}

	user_id, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(user_id),
	}

	database.DB.Preload("Role").Find(&user)
	role := models.Role{
		Id: user.Role.Id,
	}
	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("Unauthorize - You can't access " + page + " page")
}
