package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, e Entity, page int) fiber.Map {

	limit := 5
	offset := ((page - 1) * limit)

	total := e.Count(db)
	data := e.Take(db, offset, limit)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	}
}
