package controllers

import (
	"assignment2/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type controllerOrder struct {
	db *gorm.DB
}

func NewControllerOrder(db *gorm.DB) *controllerOrder {
	return &controllerOrder{
		db: db,
	}
}

func (in *controllerOrder) GetOrders(c *gin.Context) {
	var (
		order []models.Order
	)

	err := in.db.Preload("Items").Find(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &order,
	})
}

func (in *controllerOrder) CreateOrder(c *gin.Context) {
	var (
		orderRequest models.OrderRequest
		order        models.Order
	)

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order = orderRequest.ConvertToOrder()
	err := in.db.Create(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "save data success!",
		"data": order,
	})
}

func (in *controllerOrder) UpdateOrderByID(c *gin.Context) {
	var (
		orderRequestUpdate models.OrderRequestUpdate
		orderRequest       models.OrderRequest
		order              models.Order
		item               models.Item
	)

	if err := c.ShouldBindJSON(&orderRequestUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("orderId")
	err := in.db.First(&order, id).Error

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		}
		return
	}

	err = in.db.Model(&order).Updates(orderRequest).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	for _, itemOrderRequestUpdate := range orderRequestUpdate.Items {
		err := in.db.First(&item, itemOrderRequestUpdate.ItemId).Error

		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

			}
			return
		}

		err = in.db.Model(&item).Updates(itemOrderRequestUpdate).Error
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		order.Items = append(order.Items, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "update data success!",
		"data": order,
	})
}

func (in *controllerOrder) DeleteOrderById(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("orderId")
	log.Println(id)
	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success !",
	})
}
