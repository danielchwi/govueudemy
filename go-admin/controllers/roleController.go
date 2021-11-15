package controllers

import (
	"govue/database"
	"govue/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]*models.Permission, len(list))

	for i, permissionId := range list {
		var permissionIdFloat float64
		s, ok := permissionId.(string)
		if ok {
			permissionIdFloat, _ = strconv.ParseFloat(s, 64)
		}
		f, ok := permissionId.(float64)
		if ok {
			permissionIdFloat = f
		}

		permissions[i] = &models.Permission{
			Id: uint(permissionIdFloat),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{Id: uint(id)}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	//return c.JSON(list)

	permissions := make([]*models.Permission, len(list))

	for i, permissionId := range list {
		var permissionIdFloat float64
		s, ok := permissionId.(string)
		if ok {
			permissionIdFloat, _ = strconv.ParseFloat(s, 64)
		}
		f, ok := permissionId.(float64)
		if ok {
			permissionIdFloat = f
		}

		permissions[i] = &models.Permission{
			Id: uint(permissionIdFloat),
		}
	}

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Association("Permissions").Replace(role.Permissions)
	database.DB.Updates(&role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{Id: uint(id)}

	database.DB.Model(&role).Association("Permissions").Clear()
	database.DB.Delete(&role)

	return nil
}
