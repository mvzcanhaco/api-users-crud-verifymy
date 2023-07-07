package usecase

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
)

// Implemente as funções da interface UserUseCase
// Exemplo:
func (uc *UserUseCaseImpl) CreateUser(user *entity.User) error {
	// Valide os dados, aplique regras de negócio e chame o repositório para criar o usuário
	return uc.userRepo.Create(user)
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
