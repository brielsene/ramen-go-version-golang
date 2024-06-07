package controllers

import (
	"net/http"
	"ramen-go/client"
	"ramen-go/database"
	"ramen-go/models"

	"github.com/gin-gonic/gin"
)

func GetAllBroths(c *gin.Context) {
	var broths []models.Broth
	database.DB.Find(&broths)
	c.JSON(http.StatusAccepted, &broths)
}

func CreateOrderId(c *gin.Context) {
	c.JSON(http.StatusCreated, client.CreateReqWithExternalAPI())
}

func GetAllProteins(c *gin.Context) {
	var proteins []models.Protein
	database.DB.Find(&proteins)
	c.JSON(http.StatusAccepted, proteins)
}
