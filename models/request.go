package models

import (
	"log"
	"time"
)

type OrderRequest struct {
	OrderedAt    time.Time
	CustomerName string
	Items        []ItemRequest
}

type OrderRequestUpdate struct {
	OrderedAt    time.Time
	CustomerName string
	Items        []ItemRequestUpdate
}

type ItemRequestUpdate struct {
	ItemId      uint `json:"lineItemId"`
	ItemCode    string
	Description string
	Quantity    uint
}

type ItemRequest struct {
	ItemCode    string
	Description string
	Quantity    uint
}

func (orderRequest OrderRequest) ConvertToOrder() Order {
	var order Order
	var items []Item

	order.CustomerName = orderRequest.CustomerName
	order.OrderedAt = orderRequest.OrderedAt

	for _, or := range orderRequest.Items {
		var item Item
		item.ItemCode = or.ItemCode
		item.Description = or.Description
		item.Quantity = or.Quantity

		items = append(items, item)
	}
	order.Items = items
	log.Println(order)
	return order
}
