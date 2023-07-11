package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"github.com/mvzcanhaco/api-users-crud-verifymy/usecase"
	"github.com/stretchr/testify/assert"
)

type mockUserUseCase struct {
	CreateUserFunc       func(user *usecase.CreateUserData) (*entity.User, error)
	GetUserByIDFunc      func(id uint64) (*entity.User, error)
	GetAllUsersFunc      func(page, pageSize int) ([]*entity.User, error)
	UpdateUserFunc       func(user *entity.User) error
	DeleteUserFunc       func(id uint64) error
	CheckEmailExistsFunc func(email string) (bool, error)
	AuthenticateUserFunc func(email, password string) (string, error)
}

func (m *mockUserUseCase) CreateUser(user *usecase.CreateUserData) (*entity.User, error) {
	return m.CreateUserFunc(user)
}

func (m *mockUserUseCase) GetUserByID(id uint64) (*entity.User, error) {
	return m.GetUserByIDFunc(id)
}

func (m *mockUserUseCase) GetAllUsers(page, pageSize int) ([]*entity.User, error) {
	return m.GetAllUsersFunc(page, pageSize)
}

func (m *mockUserUseCase) UpdateUser(user *entity.User) error {
	return m.UpdateUserFunc(user)
}

func (m *mockUserUseCase) DeleteUser(id uint64) error {
	return m.DeleteUserFunc(id)
}

func (m *mockUserUseCase) CheckEmailExists(email string) (bool, error) {
	return m.CheckEmailExistsFunc(email)
}

func (m *mockUserUseCase) AuthenticateUser(email, password string) (string, error) {
	return m.AuthenticateUserFunc(email, password)
}

func TestUserHandler_CreateUser(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		CreateUserFunc: func(user *usecase.CreateUserData) (*entity.User, error) {
			return &entity.User{
				ID:        1,
				Name:      user.Name,
				Email:     user.Email,
				BirthDate: user.BirthDate,
				Age:       0,
				Profile:   "",
				Address:   nil,
			}, nil
		},
		CheckEmailExistsFunc: func(email string) (bool, error) {
			return false, nil
		},
	}

	// Create UserHandler with mock UserUseCase
	handler := NewUserHandler(mock)

	// Create a new Gin router
	router := gin.Default()

	// Define the route for /api/v1/users
	router.POST("/api/v1/users", func(c *gin.Context) {
		handler.CreateUser(c)
	})

	// Create a new HTTP request with JSON payload
	payload := usecase.CreateUserData{
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		BirthDate: "1990-01-01",
		Password:  "password",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	expectedResponse := UserResponse{
		ID:        1,
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		BirthDate: "1990-01-01",
		Age:       0,
		Profile:   "",
		Address:   nil,
	}
	var responseUser UserResponse
	_ = json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, expectedResponse, responseUser)
}

func TestUserHandler_GetUserByID(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		GetUserByIDFunc: func(id uint64) (*entity.User, error) {
			if id == 1 {
				return &entity.User{
					ID:        1,
					Name:      "John Doe",
					Email:     "johndoe@example.com",
					BirthDate: "1990-01-01",
					Age:       0,
					Profile:   "",
					Address:   nil,
				}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	// Create UserHandler with mock UserUseCase
	handler := NewUserHandler(mock)

	// Create a new Gin router
	router := gin.Default()

	// Define the route for /api/v1/users/:id
	router.GET("/api/v1/users/:id", func(c *gin.Context) {
		handler.GetUserByID(c)
	})

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse := UserResponse{
		ID:        1,
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		BirthDate: "1990-01-01",
		Age:       33, // Altere a idade esperada com base no valor correto
		Profile:   "",
		Address:   nil,
	}
	var responseUser UserResponse
	_ = json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, expectedResponse, responseUser)
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		GetAllUsersFunc: func(page, pageSize int) ([]*entity.User, error) {
			if page == 1 && pageSize == 10 {
				return []*entity.User{
					{
						ID:        1,
						Name:      "John Doe",
						Email:     "johndoe@example.com",
						BirthDate: "1990-01-01",
						Age:       0,
						Profile:   "",
						Address:   nil,
					},
					{
						ID:        2,
						Name:      "Jane Smith",
						Email:     "janesmith@example.com",
						BirthDate: "1992-05-15",
						Age:       0,
						Profile:   "",
						Address:   nil,
					},
				}, nil
			}
			return nil, errors.New("failed to retrievedata")
		},
	}

	// Create UserHandler with mock UserUseCase
	handler := NewUserHandler(mock)

	// Create a new Gin router
	router := gin.Default()

	// Define the route for /api/v1/users
	router.GET("/api/v1/users", func(c *gin.Context) {
		handler.GetAllUsers(c)
	})

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/api/v1/users?page=1&pageSize=10", nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse := []UserResponse{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "johndoe@example.com",
			BirthDate: "1990-01-01",
			Age:       33,
			Profile:   "",
			Address:   nil,
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "janesmith@example.com",
			BirthDate: "1992-05-15",
			Age:       31,
			Profile:   "",
			Address:   nil,
		},
	}
	var responseUsers []UserResponse
	_ = json.Unmarshal(w.Body.Bytes(), &responseUsers)
	assert.Equal(t, expectedResponse, responseUsers)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		GetUserByIDFunc: func(id uint64) (*entity.User, error) {
			if id == 1 {
				return &entity.User{
					ID:        1,
					Name:      "John Doe",
					Email:     "johndoe@example.com",
					BirthDate: "1990-01-01",
					Age:       0,
					Profile:   "",
					Address:   nil,
				}, nil
			}
			return nil, errors.New("user not found")
		},
		UpdateUserFunc: func(user *entity.User) error {
			return nil
		},
	}

	// Create UserHandler with mock UserUseCase
	handler := NewUserHandler(mock)

	// Create a new Gin router
	router := gin.Default()

	// Define the route for /api/v1/users/:id
	router.PATCH("/api/v1/users/:id", func(c *gin.Context) {
		handler.UpdateUser(c)
	})

	// Create a new HTTP request with JSON payload
	payload := usecase.UpdateUserData{
		BirthDate: "1992-02-01",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", "/api/v1/users/1", bytes.NewReader(body))

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse := UserResponse{
		ID:        1,
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		BirthDate: "1992-02-01",
		Age:       0,
		Profile:   "",
		Address: &entity.Address{
			Street:  "R Manuel Jacinto",
			City:    "Sao Paulo",
			State:   "SP",
			Country: "Brasil",
		},
	}
	var responseUser UserResponse
	_ = json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, expectedResponse, responseUser)
}

func TestUserHandler_DeleteUser(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		DeleteUserFunc: func(id uint64) error {
			return nil
		},
	}

	// Create UserHandler with mock UserUseCase
	handler := NewUserHandler(mock)

	// Create a new Gin router
	router := gin.Default()

	// Define the route for /api/v1/users/:id
	router.DELETE("/api/v1/users/:id", func(c *gin.Context) {
		handler.DeleteUser(c)
	})

	// Create a new HTTP request
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAuthHandler_Login(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		AuthenticateUserFunc: func(email, password string) (string, error) {
			if email == "johndoe@example.com" && password == "password" {
				return "token123", nil
			}
			return "", errors.New("authentication failed")
		},
	}

	// Create AuthHandler with mock UserUseCase
	handler := NewAuthHandler(mock)

	// Create a new Gin router and register the Login route
	router := gin.Default()
	router.POST("/login", handler.Login)

	// Create a new HTTP request with JSON payload
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "johndoe@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))

	// Perform the request through the router
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse := gin.H{
		"token": "token123",
	}
	var responseJSON gin.H
	_ = json.Unmarshal(w.Body.Bytes(), &responseJSON)
	assert.Equal(t, expectedResponse, responseJSON)
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{
		AuthenticateUserFunc: func(email, password string) (string, error) {
			return "", errors.New("authentication failed")
		},
	}

	// Create AuthHandler with mock UserUseCase
	handler := NewAuthHandler(mock)

	// Create a new Gin router and register the Login route
	router := gin.Default()
	router.POST("/login", handler.Login)

	// Create a new HTTP request with JSON payload
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "johndoe@example.com",
		Password: "wrongpassword",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))

	// Perform the request through the router
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check the response body
	expectedResponse := gin.H{"error": "Credenciais inválidas"}
	var responseJSON gin.H
	_ = json.Unmarshal(w.Body.Bytes(), &responseJSON)
	assert.Equal(t, expectedResponse, responseJSON)
}

func TestAuthHandler_Login_InvalidRequest(t *testing.T) {
	// Mock UserUseCase
	mock := &mockUserUseCase{}

	// Create AuthHandler with mock UserUseCase
	handler := NewAuthHandler(mock)

	// Create a new Gin router and register the Login route
	router := gin.Default()
	router.POST("/login", handler.Login)

	// Create a new HTTP request with invalid JSON payload
	invalidPayload := []byte(`{"email": "johndoe@example.com"}`) // Missing "password" field
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(invalidPayload))

	// Perform the request through the router
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check the response body
	expectedResponse := gin.H{"error": "email and password are required"}
	var responseJSON gin.H
	_ = json.Unmarshal(w.Body.Bytes(), &responseJSON)
	assert.Equal(t, expectedResponse, responseJSON)
}

func TestRegisterRoutes(t *testing.T) {
	userUseCase := &mockUserUseCase{}

	router := NewRouter(userUseCase)
	r := router.RegisterRoutes()

	assert.NotNil(t, r)

}

func TestSetupRoutes(t *testing.T) {
	userUseCase := &mockUserUseCase{}

	r := SetupRoutes(userUseCase)

	assert.NotNil(t, r)

}

// You can add more test functions for each individual route handler if needed.
