package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1/hotels")
	{
		v1.POST("/", controllers.createHotel)
		v1.GET("/", controllers.getAllHotels)
		v1.GET("/:id", controllers.getHotel)
		v1.PUT("/:id", controllers.updateHotel)
		v1.DELETE("/:id", controllers.deleteHotel)

	}
}
