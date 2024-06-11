package routes

import (
	"ramen-go/controllers"
	"ramen-go/middleware"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	// Grupo de rotas protegidas pela chave de API
	protected := r.Group("/", middleware.APIKeyMiddleware())
	protected.GET("/broths", controllers.GetAllBroths)
	protected.POST("/create-order", controllers.CreateOrder)
	protected.GET("/proteins", controllers.GetAllProteins)
	protected.POST("/order", controllers.CreateOrder)
	r.Run(":8000")
}
