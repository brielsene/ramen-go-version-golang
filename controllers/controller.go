package controllers

import (
	"net/http"
	"ramen-go/database"
	"ramen-go/models"

	"github.com/gin-gonic/gin"
)

func GetAllBroths(c *gin.Context) {
	var broths []models.Broth
	database.DB.Find(&broths)
	c.JSON(http.StatusAccepted, &broths)
}
