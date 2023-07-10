package main

import (
	"log"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar o servidor HTTP na porta especificada
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
