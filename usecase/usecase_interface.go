package usecase

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/repository"
)

type UserUseCase interface {
	CreateUser(user *entity.User) error
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
