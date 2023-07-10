package db

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	migrator := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230709000001",
			Migrate: func(tx *gorm.DB) error {
				// Executar a migração para criar a tabela
				if err := tx.AutoMigrate(&entity.User{}); err != nil {
					return err
				}

				// Inserir o registro inicial
				initialUser := &entity.User{
					Name:      "Initial User",
					Email:     "admin@example.com",
					Password:  "$2a$10$nig.ESp4fCFRW.5DPtDJZ.S4hIF7g0AE7UC/yODkb8Pl5PSFMdVra",
					BirthDate: "1992-02-01",
					Profile:   "admin",
					Address: &entity.Address{
						Street:  "exemplo",
						City:    "Sao Paulo",
						State:   "SP",
						Country: "Brasil",
					},
				}
				if err := tx.Create(initialUser).Error; err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		// Mais migrações...
	})

	return migrator.Migrate()
}
