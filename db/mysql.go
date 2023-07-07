package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(mysql:3306)/vmyCrud?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func SetupDatabase() *gorm.DB {
	time.Sleep(5 * time.Second) // Atraso de 5 segundos

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
