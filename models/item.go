package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ItemId      int    `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     int    `json:"orderId"`
}
