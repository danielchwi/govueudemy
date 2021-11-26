package controllers

import (
	"encoding/csv"
	"govue/database"
	"govue/middlewares"
	"govue/models"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "orders"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}

func Export(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "orders"); err != nil {
		return err
	}

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

type Sales struct {
	Date string `json:"date"`
	Sum  int    `json:"sum"`
}

func Chart(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "orders"); err != nil {
		return err
	}

	sales := []Sales{}

	database.DB.Raw(`
		SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d')  as date, SUM(oi.price) as sum
		FROM orders o 
		JOIN order_items oi ON o.id = oi.order_id
		GROUP BY date;
	`).Scan(&sales)

	return c.JSON(sales)
}
