package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mvzcanhaco/api-users-crud-verifymy/delivery/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verificar o cabeçalho Authorization no formato "Bearer <token>"
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação não fornecido")
			c.Abort()
			return
		}

		// Extrair o token do cabeçalho
		tokenArr := strings.Split(tokenString, " ")
		if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação inválido")
			c.Abort()
			return
		}
		tokenString = tokenArr[1]

		// Verificar a validade do token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Defina a chave secreta usada para assinar o token
			return []byte("VMYCRUDTEST"), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação inválido1")
			c.Abort()
			return
		}

		// Obter os claims do token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação inválido2")
			c.Abort()
			return
		}

		// Obter o valor do claim "id"
		userID, ok := claims["id"].(float64)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação inválido3")
			c.Abort()
			return
		}

		// Definir os dados do usuário no contexto
		c.Set("ID", uint(userID))

		profile, ok := claims["profile"].(string)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "Token de autenticação inválido4")
			c.Abort()
			return
		}

		// Definir o perfil do usuário no contexto
		c.Set("profile", profile)

		// Continuar para o próximo handler
		c.Next()
	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verificar se o perfil do usuário é "admin"
		userProfile := c.GetString("profile")
		if userProfile != "admin" {
			response.Error(c, http.StatusForbidden, "Apenas usuários com perfil de administrador podem realizar esta operação")
			c.Abort()
			return
		}

		// Continuar para o próximo handler
		c.Next()
	}
}
