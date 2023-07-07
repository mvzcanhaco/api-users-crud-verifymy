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
			ID: "20230707000001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entity.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "20230707000003",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					gorm.Model
					ID        uint            `gorm:"primaryKey" json:"id,omitempty"`
					Name      string          `gorm:"not null" json:"name,omitempty"`
					Email     string          `gorm:"not null;unique" json:"email,omitempty"`
					Password  string          `gorm:"not null" json:"password,omitempty"`
					BirthDate string          `gorm:"not null" json:"birthDate,omitempty"`
					Age       int             `gorm:"not null" json:"age,omitempty"`
					Address   *entity.Address `json:"address,omitempty"`
					Profile   string          `json:"profile,omitempty" gorm:"type:varchar(255);default:'user'"`
				}
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		// Mais migrações...
	})

	return migrator.Migrate()
}
