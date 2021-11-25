package controllers

import (
	"encoding/csv"
	"govue/database"
	"govue/models"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}

func Export(c *fiber.Ctx) error {

	filepath := "./csv/order.csv"

	if err := CreateFile(filepath); err != nil {
		return err
	}

	return c.Download(filepath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	var Orders []models.Order

	database.DB.Preload("OrderItem").Find(&Orders)

	for _, order := range Orders {
		writer.Write([]string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		})
		for _, orderItem := range order.OrderItem {
			writer.Write([]string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			})
		}
	}

	return nil
}
