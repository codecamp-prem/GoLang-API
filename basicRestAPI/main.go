package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Product : struct
type Product struct {
	ID    int     `json:"ID" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

func main() {
	r := gin.Default()

	// HTTP GET, POST, PUT, DELETE
	r.GET("/", getHello)
	r.POST("/", postHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)

	// groupin
	r1 := r.Group("/api")
	{
		r1.GET("/hello", getHello)
		r1.POST("/hello", postHello)
		r1.PUT("/hello", putHello)
		r1.DELETE("/hello", deleteHello)
	}

	// handling path params
	r.GET("/product/:id", getProductByID)
	r.GET("/profile/:username", showProfile)
	r.GET("/compute/:num1/add/:num2", compute)

	// Handling query params
	// /employee?firstname=Jane&lastname=Doe&id=2
	r.GET("/employee", showEmployee)

	//binding post data
	r.POST("/product", performProduct)
	r.POST("/products", performProducts)

	r.Run(":8080") //Default port 8080
	fmt.Println("Server is running!!")
}

func performProduct(c *gin.Context) {
	var product Product

	err := c.BindJSON(&product)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}

func performProducts(c *gin.Context) {
	var products []Product

	err := c.BindJSON(&products)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func showEmployee(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "")
	lastName := c.DefaultQuery("lastname", "")
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 0)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":         id,
		"First Name": firstName,
		"Last Name":  lastName,
	})
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")                   //string
	idn, _ := strconv.ParseInt(id, 10, 0) //convert sting to integer

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":   idn,
		"name": "Product A",
	})
}

func showProfile(c *gin.Context) {
	username := c.Param("username") //string
	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
		"name":     "Doe Family",
	})
}

func compute(c *gin.Context) {
	num1 := c.Param("num1") //string

	num1Int, _ := strconv.ParseInt(num1, 10, 0)            //convert sting to integer
	num2Int, _ := strconv.ParseInt(c.Param("num2"), 10, 0) //convert sting to integer
	result := num1Int + num2Int

	c.IndentedJSON(http.StatusOK, gin.H{
		"num1":   num1Int,
		"num2":   num2Int,
		"result": result,
	})
}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "API- HTTP GET",
	})
}
func postHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "API- HTTP POST",
	})
}
func putHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "API- HTTP PUT",
	})
}
func deleteHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "API- HTTP DELTE",
	})
}
