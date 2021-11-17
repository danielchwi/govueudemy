package models

import "gorm.io/gorm"

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (*Product) Take(db *gorm.DB, offset int, limit int) interface{} {
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}

func (p Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&p).Count(&total)
	return total
}
