package main

import (
	"fmt"

	"github.com/codecamp-prem/productapi/controllers"
	"github.com/codecamp-prem/productapi/models"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{ID: 1, Name: "Product 1", Stock: 100, Price: 2.20},
	{ID: 2, Name: "Product 2", Stock: 123, Price: 1.98},
	{ID: 3, Name: "Product 3", Stock: 176, Price: 1.06},
}

func main() {
	r := gin.Default()
	controller := controllers.Init(&products)

	r.GET("/product", controller.ReadProducts)
	r.GET("/product/:id", controller.ReadProductByID)
	r.POST("/product", controller.CreateProduct)
	r.PUT("/product/:id", controller.UpdateProductByID)
	r.DELETE("/product/:id", controller.DeleteProductByID)

	r.Run("localhost:8080")
	fmt.Println("Server is running")
}
