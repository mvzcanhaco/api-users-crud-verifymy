package db

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	migrator := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Defina suas migrações aqui
		// Exemplo:
		{
			ID: "20230709000001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entity.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		// Mais migrações...
	})

	return migrator.Migrate()
}
