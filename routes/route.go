package routes

import (
	"ramen-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/broths", controllers.GetAllBroths)
	r.POST("/create-order", controllers.CreateOrder)
	r.GET("/proteins", controllers.GetAllProteins)
	r.POST("/order", controllers.CreateOrder)
	r.Run(":8000")
}
