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

type UserResponse struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	BirthDate string          `json:"birthDate"`
	Age       int             `json:"age"`
	Profile   string          `json:"profile"`
	Address   *entity.Address `json:"address"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var createUser usecase.CreateUserData
	if err := c.ShouldBindJSON(&createUser); err != nil {
		response.BadRequest(c, err)
		return
	}

	// Verificar se o email já está cadastrado
	emailExists, _ := h.userUseCase.CheckEmailExists(createUser.Email)
	if emailExists {
		response.StatusConflit(c)
		return
	}

	// Criptografar a senha antes de inserir no banco de dados
	hashedPassword, err := usecase.HashPassword(createUser.Password)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	// Atribuir a senha criptografada ao usuário
	createUser.Password = hashedPassword
	user, err := h.userUseCase.CreateUser(&createUser)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, http.StatusCreated, mapUserToResponse(user))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		response.NotFound(c, err)
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

	response.Success(c, http.StatusOK, mapUserToResponse(user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "100")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		response.BadRequest(c, err)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		response.BadRequest(c, err)
		return
	}

	users, err := h.userUseCase.GetAllUsers(pageInt, pageSizeInt)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	// Mapear e calcular a idade de cada usuário
	var responseUsers []*UserResponse
	for _, user := range users {
		// Calcula a idade com base na data de nascimento
		age, err := utils.CalculateAge(user.BirthDate)
		if err != nil {
			// Lida com o erro, se necessário
			continue // Ou retorne um erro ou pule para o próximo usuário
		}

		// Cria um novo objeto ResponseUser
		responseUser := mapUserToResponse(user)

		// Atribui a idade ao usuário
		responseUser.Age = age

		// Adiciona o usuário mapeado ao array de responseUsers
		responseUsers = append(responseUsers, responseUser)
	}

	response.Success(c, http.StatusOK, responseUsers)

}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		response.NotFound(c, err)
		return
	}

	var updateUser usecase.UpdateUserData
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		response.BadRequest(c, err)
		return
	}

	// Verificar se houve alguma alteração nos campos
	if updateUser.BirthDate != "" {
		user.BirthDate = updateUser.BirthDate
	}

	if updateUser.Address != nil {
		user.Address = updateUser.Address
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	if err := h.userUseCase.UpdateUser(user); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, http.StatusOK, mapUserToResponse(user))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := h.userUseCase.DeleteUser(id); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.NoContent(c)
}

func mapUserToResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		Age:       user.Age,
		Profile:   user.Profile,
		Address:   user.Address,
	}
}
