package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// ErrorHandlingMiddleware é um middleware para tratar erros
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se houve algum erro durante o processamento da requisição
		errs := c.Errors
		if len(errs) > 0 {
			err := errs[0].Err
			var statusCode int

			// Aqui você pode adicionar lógica personalizada para diferentes tipos de erros
			switch err {
			case gorm.ErrRecordNotFound:
				statusCode = http.StatusBadRequest
				c.JSON(statusCode, gin.H{
					"error": "both brothId and proteinId are required",
				})

			default:
				statusCode = http.StatusInternalServerError
				c.JSON(statusCode, gin.H{
					"error": "could not place order",
				})

			}

			// // Retorna a resposta de erro com o status code apropriado
			// c.JSON(statusCode, ErrorResponse{Message: errorMessage})
		}
	}
}
