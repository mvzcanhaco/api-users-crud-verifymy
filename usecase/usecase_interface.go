package usecase

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/repository"
)

type UserUseCase interface {
	CreateUser(user *CreateUserData) error
	GetUserByID(id uint) (*entity.User, error)
	GetAllUsers(page, pageSize int) ([]*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	CheckEmailExists(email string) (bool, error)
	AuthenticateUser(email, password string) (string, error)
}

type UserUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUseCaseImpl(userRepo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

type CreateUserData struct {
	Name      string          `json:"name" validate:"required"`
	Email     string          `json:"email" validate:"required,email"`
	Password  string          `json:"password" validate:"required,min=6"`
	BirthDate string          `json:"birthDate" validate:"required"`
	Adress    *entity.Address `json:"address" validate:"required"`
}

type UpdateUserData struct {
	Email     string          `json:"email" validate:"required,email"`
	BirthDate string          `json:"birthDate" validate:"required"`
	Adress    *entity.Address `json:"address" validate:"required"`
}
