package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Rest API")
	})
	r.GET("/hello2", hello)
	r.Run()
	fmt.Println("Server is running!!")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Rest API from functionHandler")
}
