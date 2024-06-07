package routes

import (
	"ramen-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/broths", controllers.GetAllBroths)
	r.Run(":8000")
}
