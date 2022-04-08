package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codecamp-prem/productapi/models"
	"github.com/gin-gonic/gin"
)

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

func (repo *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product

	c.BindJSON(&product)
	err := models.CreateProduct(repo.Products, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (repo *ProductRepo) ReadProducts(c *gin.Context) {
	c.JSON(http.StatusOK, repo.Products)
}

func (repo *ProductRepo) ReadProductByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	product := models.ReadProductByID(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}

func (repo *ProductRepo) UpdateProductByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request <=0"})
		return
	}
	var product models.Product
	c.BindJSON(&product)
	fmt.Println(product.ID)
	fmt.Println(id)
	if product.ID != int(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Homi"})
		return
	}

	updatedProduct := models.UpdateProductByID(repo.Products, &product)
	c.JSON(http.StatusOK, &updatedProduct)
}

func (repo *ProductRepo) DeleteProductByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	product := models.DeleteProductByID(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}
