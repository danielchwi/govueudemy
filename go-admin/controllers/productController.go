package controllers

import (
	"govue/database"
	"govue/middlewares"
	"govue/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)

}

func DeleteProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}
	database.DB.Delete(&product)

	return nil
}
