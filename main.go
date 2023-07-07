package main

import (
	"log"

	"github.com/mvzcanhaco/api-users-crud-verifymy/db"
	"github.com/mvzcanhaco/api-users-crud-verifymy/delivery/http"
	"github.com/mvzcanhaco/api-users-crud-verifymy/repository"
	"github.com/mvzcanhaco/api-users-crud-verifymy/usecase"
)

func main() {

	db := db.SetupDatabase()

	// Inicializar as dependências
	userRepo := repository.NewUserRepositoryImpl(db)
	userUseCase := usecase.NewUserUseCaseImpl(userRepo)
	r := http.SetupRoutes(userUseCase)
	// Iniciar o servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
