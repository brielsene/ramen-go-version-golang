package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func APIKeyMiddleware() gin.HandlerFunc {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY_PERSONAL")
	fmt.Println("API-KEY_PERSONAL: ", apiKey)
	return func(c *gin.Context) {
		requestApiKey := c.GetHeader("x-api-key")
		if requestApiKey == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "x-api-key header missing"})
			c.Abort()
			return
		}

		if requestApiKey != apiKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "x-api-key header missing"})
			c.Abort()
			return
		}

		c.Next()
	}
}
