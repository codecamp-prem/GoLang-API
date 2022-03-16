package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	r.Run() //Default port 8080
	fmt.Println("Server is running!!")
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
