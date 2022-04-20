package main

import (
	"assignment2/configs"
	"assignment2/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.NewPostgres()
	if db == nil {
		fmt.Println("running db is fail")
	}

	orderController := controllers.NewControllerOrder(db)

	router := gin.Default()

	router.GET("/orders", orderController.GetOrders)
	router.POST("/orders", orderController.CreateOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrderByID)
	router.DELETE("/orders/:orderId", orderController.DeleteOrderById)

	router.Run(":8080")
}
