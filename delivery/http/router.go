package http

import (
	"github.com/gin-gonic/gin"
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
		v1.POST("/login", r.authHandler.Login)
		v1.POST("/users", r.userHandler.CreateUser)
		v1.GET("/users/:id", r.userHandler.GetUserByID)
		v1.GET("/users", r.userHandler.GetAllUsers)
		v1.PUT("/users/:id", r.userHandler.UpdateUser)
		v1.DELETE("/users/:id", r.userHandler.DeleteUser)
	}

	return router
}

func SetupRoutes(userUseCase usecase.UserUseCase) *gin.Engine {
	router := NewRouter(userUseCase)
	r := router.RegisterRoutes()

	return r
}
