package routes

import (
	"ramen-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/broths", controllers.GetAllBroths)
	r.POST("/test", controllers.CreateOrderId)
	r.GET("/proteins", controllers.GetAllProteins)
	r.Run(":8000")
}
