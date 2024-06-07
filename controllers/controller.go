package controllers

import (
	"log"
	"net/http"
	"ramen-go/client"
	"ramen-go/database"
	"ramen-go/dto"
	"ramen-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBroths(c *gin.Context) {
	var broths []models.Broth
	database.DB.Find(&broths)
	c.JSON(http.StatusAccepted, &broths)
}

// func CreateOrderId(c *gin.Context) {
// 	c.JSON(http.StatusCreated, client.CreateReqWithExternalAPI())
// }

func GetAllProteins(c *gin.Context) {
	var proteins []models.Protein
	database.DB.Find(&proteins)
	c.JSON(http.StatusAccepted, proteins)
}

func CreateOrder(c *gin.Context) {
	var dto dto.OrderRequestDto
	c.ShouldBindJSON(&dto)
	var brothSearch models.Broth
	var proteinSearch models.Protein
	brothId, _ := strconv.Atoi(dto.BrothId)
	proteinId, _ := strconv.Atoi(dto.ProteinId)
	database.DB.First(&brothSearch, brothId)
	if err := database.DB.First(&brothSearch, dto.BrothId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Broth not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	if err := database.DB.First(&proteinSearch, proteinId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Protein not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return

	}
	var clientOrderId = client.CreateReqWithExternalAPI()
	idOrder, err := strconv.Atoi(clientOrderId.OrderId)
	if err != nil {
		log.Fatal("Erro de conversão: ", err)
	}
	order := models.Order{Id: idOrder, Description: proteinSearch.Name + " " + brothSearch.Name, Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png"}
	database.DB.Save(&order)

}
