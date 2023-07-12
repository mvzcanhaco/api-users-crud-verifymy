package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	router := gin.Default()
	router.Use(AuthMiddleware())

	validToken := generateValidToken()

	router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Access granted", w.Body.String())
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := gin.Default()
	router.Use(AuthMiddleware())

	invalidToken := "invalid-token"

	router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Token de autenticação inválido"}`, w.Body.String())
}

func TestAdminOnlyMiddleware_AdminUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(AdminOnlyMiddleware())

	router.GET("/admin-only", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	req.Header.Set("Authorization", "Bearer "+generateNonAdminToken())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"error": "Apenas usuários com perfil de administrador podem realizar esta operação"}`, w.Body.String())
}

func TestAdminOnlyMiddleware_NonAdminUser(t *testing.T) {
	router := gin.Default()
	router.Use(AdminOnlyMiddleware())

	router.GET("/admin-only", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	req.Header.Set("Authorization", "Bearer "+generateNonAdminToken())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"error": "Apenas usuários com perfil de administrador podem realizar esta operação"}`, w.Body.String())
}

func generateValidToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = 1
	claims["profile"] = "admin"
	tokenString, _ := token.SignedString([]byte("VMYCRUDTEST"))
	return tokenString
}

func generateNonAdminToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = 2
	claims["profile"] = "user"
	tokenString, _ := token.SignedString([]byte("VMYCRUDTEST"))
	return tokenString
}
