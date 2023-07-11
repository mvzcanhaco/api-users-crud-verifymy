package usecase

import (
	"errors"
	"testing"

	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/utils"
)

// Implementação de um repositório fictício para fins de teste
type MockUserRepository struct {
	users []*entity.User
}

func (repo *MockUserRepository) Create(user *entity.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	repo.users = append(repo.users, user)
	return nil
}

func (repo *MockUserRepository) FindByID(id uint64) (*entity.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *MockUserRepository) FindAll(page, pageSize int) ([]*entity.User, error) {
	startIndex := (page - 1) * pageSize
	if startIndex < 0 || startIndex >= len(repo.users) {
		return nil, errors.New("invalid page")
	}
	endIndex := startIndex + pageSize
	if endIndex > len(repo.users) {
		endIndex = len(repo.users)
	}
	return repo.users[startIndex:endIndex], nil
}

func (repo *MockUserRepository) Update(user *entity.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	for i, u := range repo.users {
		if u.ID == user.ID {
			repo.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (repo *MockUserRepository) Delete(id uint64) error {
	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (repo *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func TestCreateUser(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	createUserData := &CreateUserData{
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}

	user, err := uc.CreateUser(createUserData)
	if err != nil {
		t.Errorf("Error creating user: %s", err.Error())
	}

	if user.Name != createUserData.Name {
		t.Errorf("Expected user name to be %s, got %s", createUserData.Name, user.Name)
	}

	// Test other user properties

	// Test error case: nil user
	_, err = uc.CreateUser(nil)
	if err == nil {
		t.Error("Expected error for nil user, got nil")
	}
}

func TestGetUserByID(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	// Test existing user
	existingUser := &entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	uc.userRepo.Create(existingUser)

	user, err := uc.GetUserByID(existingUser.ID)
	if err != nil {
		t.Errorf("Error getting user by ID: %s", err.Error())
	}

	if user != existingUser {
		t.Error("Expected user to be the same as the existing user")
	}

	// Test non-existing user
	_, err = uc.GetUserByID(123)
	if err == nil {
		t.Error("Expected error for non-existing user, got nil")
	}
}

func TestGetAllUsers(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	// Test with some users in the repository
	users := []*entity.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "password",
			BirthDate: "1992-02-01",
			Address: &entity.Address{
				Street:  "R Manuel Jacinto",
				City:    "Sao Paulo",
				State:   "SP",
				Country: "Brasil",
			},
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Password:  "password",
			BirthDate: "1992-02-01",
			Address: &entity.Address{
				Street:  "R Manuel Jacinto",
				City:    "Sao Paulo",
				State:   "SP",
				Country: "Brasil",
			},
		},
	}
	for _, user := range users {
		uc.userRepo.Create(user)
	}

	// Test case 1: Valid page and page size
	users, err := uc.GetAllUsers(1, 10)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, but got %d", len(users))
	}

	// Test case 2: Invalid page
	_, err = uc.GetAllUsers(0, 10)
	if err == nil {
		t.Errorf("Expected error for invalid page, but got nil")
	} else if err.Error() != "page and pageSize must be greater than 0" {
		t.Errorf("Expected error message 'page and pageSize must be greater than 0', but got '%s'", err.Error())
	}

	// Test case 3: Invalid page size
	_, err = uc.GetAllUsers(2, 0)
	if err == nil {
		t.Errorf("Expected error for invalid pageSize, but got nil")
	} else if err.Error() != "page and pageSize must be greater than 0" {
		t.Errorf("Expected error message 'page and pageSize must be greater than 0', but got '%s'", err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	// Test existing user
	existingUser := &entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	uc.userRepo.Create(existingUser)

	// Update user
	existingUser.Name = "Updated Name"
	err := uc.UpdateUser(existingUser)
	if err != nil {
		t.Errorf("Error updating user: %s", err.Error())
	}

	// Check if user was updated
	updatedUser, err := uc.GetUserByID(existingUser.ID)
	if err != nil {
		t.Errorf("Error getting user by ID: %s", err.Error())
	}

	if updatedUser.Name != existingUser.Name {
		t.Errorf("Expected user name to be %s, got %s", existingUser.Name, updatedUser.Name)
	}

	// Test error case: nil user
	err = uc.UpdateUser(nil)
	if err == nil {
		t.Error("Expected error for nil user, got nil")
	}
}

func TestDeleteUser(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	// Test existing user
	existingUser := &entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	uc.userRepo.Create(existingUser)

	// Delete user
	err := uc.DeleteUser(existingUser.ID)
	if err != nil {
		t.Errorf("Error deleting user: %s", err.Error())
	}

	// Check if user was deleted
	_, err = uc.GetUserByID(existingUser.ID)
	if err == nil {
		t.Error("Expected error for non-existing user, got nil")
	}

	// Test error case: non-existing user
	err = uc.DeleteUser(123)
	if err == nil {
		t.Error("Expected error for non-existing user, got nil")
	}
}

func TestCheckEmailExists(t *testing.T) {
	uc := &UserUseCaseImpl{
		userRepo: &MockUserRepository{},
	}

	// Test existing email
	existingUser := &entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	err := uc.userRepo.Create(existingUser)
	if err != nil {
		t.Errorf("Erro ao criar usuario: %s", err.Error())
	}
	exists, err := uc.CheckEmailExists(existingUser.Email)
	if err != nil {
		t.Errorf("Error checking email exists: %s", err.Error())
	}

	if !exists {
		t.Error("Expected email to exist, got false")
	}

	// Test non-existing email
	exists, err = uc.CheckEmailExists("nonexisting@example.com")

	if exists {
		t.Error("Expected email to not exist, got true")
	}
}

func TestCalculateAge(t *testing.T) {
	// Test with a known birth date and current date
	birthDate := "1992-02-01"
	expectedAge := 31 // Assuming current date is 2023-07-11

	age, err := utils.CalculateAge(birthDate)
	if err != nil {
		t.Errorf("Error calculating age: %s", err.Error())
	}

	if age != expectedAge {
		t.Errorf("Expected age to be %d, got %d", expectedAge, age)
	}

	// Test with future birth date
	futureDate := "2050-02-01"

	age, err = utils.CalculateAge(futureDate)
	if err == nil {
		t.Error("Expected error for future birth date, got nil")
	}

}
