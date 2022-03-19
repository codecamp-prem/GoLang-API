package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Numeric : struct
type Numeric struct {
	Num1   float32 `json:"num1"`
	Num2   float32 `json:"num2"`
	Result float32 `json:"result"`
}

func main() {
	r := gin.Default()
	r.POST("/add", add)
	r.POST("/subtract", subtract)
	r.POST("/multiply", multiply)
	r.POST("/divide", divide)

	r.Run("localhost:8080")
	fmt.Printf("Server is running!")
}

func add(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	num.Result = num.Num1 + num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func subtract(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	num.Result = num.Num1 - num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func multiply(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	num.Result = num.Num1 * num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func divide(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	num.Result = num.Num1 / num.Num2
	c.IndentedJSON(http.StatusOK, num)
}
