package usecase

import (
	"errors"

	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/utils"
)

func (uc *UserUseCaseImpl) CreateUser(user *CreateUserData) (*entity.User, error) {

	if user == nil {
		return nil, errors.New("user is nil")
	}
	// Crie a entidade User com base nos dados da estrutura CreateUser
	newUser := &entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		Address:   user.Address,
	}

	// Calcula a idade com base na data de nascimento
	age, err := utils.CalculateAge(newUser.BirthDate)
	if err != nil {
		age = 0
	}

	// Atribui a idade calculada ao usuário
	newUser.Age = age
	newUser.Profile = "user"

	// Chame a função uc.userRepo.Create com a entidade User
	return newUser, uc.userRepo.Create(newUser)
}

func (uc *UserUseCaseImpl) GetUserByID(id uint64) (*entity.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *UserUseCaseImpl) GetAllUsers(page, pageSize int) ([]*entity.User, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("page and pageSize must be greater than 0")
	}
	// Chamar o método FindAll do repositório passando os índices
	users, err := uc.userRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCaseImpl) UpdateUser(user *entity.User) error {

	if user == nil {
		return errors.New("user is nil")
	}
	// Calcula a idade com base na data de nascimento
	age, err := utils.CalculateAge(user.BirthDate)
	if err != nil {
		age = 0
	}

	// Atribui a idade calculada ao usuário
	user.Age = age
	return uc.userRepo.Update(user)
}

func (uc *UserUseCaseImpl) DeleteUser(id uint64) error {
	return uc.userRepo.Delete(id)
}

func (u *UserUseCaseImpl) CheckEmailExists(email string) (bool, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
