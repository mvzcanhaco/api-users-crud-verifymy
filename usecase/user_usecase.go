package usecase

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/utils"
)

// Implemente as funções da interface UserUseCase
// Exemplo:
func (uc *UserUseCaseImpl) CreateUser(user *CreateUserData) error {
	// Crie a entidade User com base nos dados da estrutura CreateUser
	newUser := &entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		Address:   user.Adress,
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
	return uc.userRepo.Create(newUser)
}

func (uc *UserUseCaseImpl) GetUserByID(id uint) (*entity.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *UserUseCaseImpl) GetAllUsers(page, pageSize int) ([]*entity.User, error) {

	// Chamar o método FindAll do repositório passando os índices
	users, err := uc.userRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCaseImpl) UpdateUser(user *entity.User) error {

	// Calcula a idade com base na data de nascimento
	age, err := utils.CalculateAge(user.BirthDate)
	if err != nil {
		age = 0
	}

	// Atribui a idade calculada ao usuário
	user.Age = age
	return uc.userRepo.Update(user)
}

func (uc *UserUseCaseImpl) DeleteUser(id uint) error {
	return uc.userRepo.Delete(id)
}

func (u *UserUseCaseImpl) CheckEmailExists(email string) (bool, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
