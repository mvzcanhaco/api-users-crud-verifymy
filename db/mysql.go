package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construir a DSN usando as variáveis de ambiente
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func SetupDatabase() *gorm.DB {
	time.Sleep(10 * time.Second) // Atraso de 5 segundos

	// Configurar conexão com o banco de dados
	dbCon, err := NewDB()
	if err != nil {
		panic(err)
	}

	// Executar migrações
	err = RunMigrations(dbCon)
	if err != nil {
		panic(err)
	}

	return dbCon
}
