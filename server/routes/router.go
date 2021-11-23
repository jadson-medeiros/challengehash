package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/challengehash/server/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.ShowProducts)
		}
	}

	return router
}
