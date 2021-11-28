package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/challengehash/database"
	"github.com/jadson-medeiros/challengehash/models"
)

func ShowProduct(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": "ID need to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var product models.Product

	err = db.First(&product, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": "Cannot find a product: " + err.Error(),
		})

		return
	}

	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDatabase()

	var product models.Product

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": "Cannot create a product: " + err.Error(),
		})

		return
	}

	c.JSON(200, product)
}

func ShowProducts(c *gin.Context) {
	db := database.GetDatabase()

	var products []models.Product

	err := db.Find(&products).Error

	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": "Cannot list products: " + err.Error(),
		})

		return
	}

	c.JSON(200, products)
}
