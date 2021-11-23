package controllers

import "github.com/gin-gonic/gin"

func ShowProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": "ok",
	})
}
