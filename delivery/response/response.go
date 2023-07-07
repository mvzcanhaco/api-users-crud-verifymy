package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success envia uma resposta de sucesso com o status e o payload fornecidos.
func Success(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// BadRequest envia uma resposta de erro com o status de BadRequest (400) e a mensagem de erro fornecida.
func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

// NotFound envia uma resposta de erro com o status de NotFound (404) e a mensagem de erro fornecida.
func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

// InternalServerError envia uma resposta de erro com o status de InternalServerError (500) e a mensagem de erro fornecida.
func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

// NoContent envia uma resposta vazia com o status de NoContent (204).
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func StatusConflit(c *gin.Context) {
	c.Status(http.StatusConflict)
}

func StatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Credenciais inválidas",
	})
}
