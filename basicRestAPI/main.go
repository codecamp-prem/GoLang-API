package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTTP GET, POST, PUT, DELETE
	r.GET("/", getHello)
	r.POST("/", postHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)
	r.Run() //Default port 8080
	fmt.Println("Server is running!!")
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
