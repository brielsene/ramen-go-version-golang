package routes

import (
	"ramen-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/broths", controllers.GetAllBroths)
	r.POST("/test", controllers.TesteClientHttp)
	r.Run(":8000")
}
