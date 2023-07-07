package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvzcanhaco/api-users-crud-verifymy/delivery/response"
	"github.com/mvzcanhaco/api-users-crud-verifymy/usecase"
)

type AuthHandler struct {
	userUseCase usecase.UserUseCase
}

func NewAuthHandler(userUseCase usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{userUseCase: userUseCase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		response.BadRequest(c, err)
		return
	}

	token, err := h.userUseCase.AuthenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	if token == "" {
		response.StatusUnauthorized(c)
		return
	}

	// Criar sessão logada ou adicionar o token no cabeçalho da resposta
	response.Success(c, http.StatusOK, gin.H{
		"token": token,
	})

}
