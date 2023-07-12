package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func Success(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	}
}

func NotFound(c *gin.Context, err error) {
	errorMessage := "Not found"
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
}

func InternalServerError(c *gin.Context, err error) {
	errorMessage := "Internal server error"
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
}

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
