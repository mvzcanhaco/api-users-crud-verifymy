package http

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
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
