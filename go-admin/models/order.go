package models

import "gorm.io/gorm"

type Order struct {
	Id        uint        `json:"id"`
	FirstName string      `json:"-"`
	LastName  string      `json:"-"`
	Name      string      `json:"name" gorm:"-"`
	Email     string      `json:"email"`
	Total     float32     `json:"total" gorm"-"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated-at"`
	OrderItem []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (*Order) Take(db *gorm.DB, offset int, limit int) interface{} {
	var orders []Order

	db.Preload("OrderItem").Offset(offset).Limit(limit).Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FirstName + " " + order.LastName
		var total float32 = 0
		for _, orderItem := range order.OrderItem {
			total += orderItem.Price
		}
		orders[i].Total = total
	}

	return orders
}

func (o Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&o).Count(&total)
	return total
}
