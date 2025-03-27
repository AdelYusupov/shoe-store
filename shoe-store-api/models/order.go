// models/order.go
package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	FullName    string      `json:"full_name"`
	Phone       string      `json:"phone"`
	City        string      `json:"city"`
	Address     string      `json:"address"`
	Delivery    string      `json:"delivery"`
	PickupPoint string      `json:"pickup_point"`
	Comment     string      `json:"comment"`
	Contact     string      `json:"contact"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Total       float64     `json:"total"`
	Status      string      `json:"status" gorm:"default:'new'"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Size      string  `json:"size"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
