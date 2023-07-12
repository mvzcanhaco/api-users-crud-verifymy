package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mvzcanhaco/api-users-crud-verifymy/delivery/middleware"
	"github.com/mvzcanhaco/api-users-crud-verifymy/usecase"
)

type Router struct {
	authHandler *AuthHandler
	userHandler *UserHandler
}

func NewRouter(userUseCase usecase.UserUseCase) *Router {
	authHandler := NewAuthHandler(userUseCase)
	userHandler := NewUserHandler(userUseCase)

	return &Router{
		authHandler: authHandler,
		userHandler: userHandler,
	}
}

func (r *Router) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// Anotações do Swagger para a rota de login
		// @Summary Fazer login
		// @Description Realiza a autenticação do usuário usando e-mail e senha
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Param input body LoginInput true "Credenciais de login"
		// @Success 200 {object} TokenResponse
		// @Router /api/v1/login [post]
		v1.POST("/login", r.authHandler.Login)

		// Anotações do Swagger para a rota de criação de usuário
		// @Summary Criar usuário
		// @Description Cria um novo usuário
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Param input body CreateUserInput true "Dados do usuário a ser criado"
		// @Success 201 {object} UserResponse
		// @Router /api/v1/users [post]
		v1.POST("/users", r.userHandler.CreateUser)

		// Rotas protegidas pelo middleware
		v1.Use(middleware.AuthMiddleware())

		// Anotações do Swagger para a rota de busca de usuário por ID
		// @Summary Obter usuário por ID
		// @Description Retorna um usuário com base no ID fornecido
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Param id path int true "ID do usuário"
		// @Success 200 {object} UserResponse
		// @Router /api/v1/users/{id} [get]
		v1.GET("/users/:id", r.userHandler.GetUserByID)

		// Anotações do Swagger para a rota de busca de todos os usuários
		// @Summary Obter todos os usuários
		// @Description Retorna uma lista de todos os usuários
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Success 200 {array} UserResponse
		// @Router /api/v1/users [get]
		v1.GET("/users", r.userHandler.GetAllUsers)

		// Anotações do Swagger para a rota de atualização de usuário
		// @Summary Atualizar usuário
		// @Description Atualiza as informações de um usuário existente
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Param id path int true "ID do usuário"
		// @Param input body UpdateUserInput true "Novos dados do usuário"
		// @Success 200 {object} UserResponse
		// @Router /api/v1/users/{id} [put]
		v1.PATCH("/users/:id", middleware.AdminOnlyMiddleware(), r.userHandler.UpdateUser)

		// Anotações do Swagger para a rota de exclusão de usuário
		// @Summary Excluir usuário
		// @Description Exclui um usuário existente
		// @Tags Users
		// @Accept json
		// @Produce json
		// @Param id path int true "ID do usuário"
		// @Success 204 "No Content"
		// @Router /api/v1/users/{id} [delete]
		v1.DELETE("/users/:id", middleware.AdminOnlyMiddleware(), r.userHandler.DeleteUser)

	}

	return router
}

func SetupRoutes(userUseCase usecase.UserUseCase) *gin.Engine {
	router := NewRouter(userUseCase)
	r := router.RegisterRoutes()

	return r
}
