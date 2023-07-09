package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mvzcanhaco/api-users-crud-verifymy/delivery/response"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/utils"
	"github.com/mvzcanhaco/api-users-crud-verifymy/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err)
		return
	}

	// Verificar se o email já está cadastrado
	emailExists, _ := h.userUseCase.CheckEmailExists(user.Email)
	if emailExists {
		response.StatusConflit(c)
		return
	}

	// Calcula a idade com base na data de nascimento
	age, err := utils.CalculateAge(user.BirthDate)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	// Atribui a idade calculada ao usuário
	user.Age = age

	// Criptografar a senha antes de inserir no banco de dados
	hashedPassword, err := usecase.HashPassword(user.Password)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	// Atribuir a senha criptografada ao usuário
	user.Password = hashedPassword

	if err := h.userUseCase.CreateUser(&user); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		response.NotFound(c, err)
		return
	}

	response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "100")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valor inválido para 'page'",
		})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valor inválido para 'pageSize'",
		})
		return
	}

	users, err := h.userUseCase.GetAllUsers(pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar usuários",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err)
		return
	}

	user.ID = id

	if err := h.userUseCase.UpdateUser(user); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := h.userUseCase.DeleteUser(uint(id)); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.NoContent(c)
}
