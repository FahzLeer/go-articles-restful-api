package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1624347947CreateCategoriesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1624347947",
		Migrate: func(tx *gorm.DB) error {
			type category struct {
				gorm.Model
				Name string `gorm:"unique;not null"`
				Desc string `gorm:"not null"`
			}

			return tx.Migrator().CreateTable(&category{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("categories")
		},
	}
}
