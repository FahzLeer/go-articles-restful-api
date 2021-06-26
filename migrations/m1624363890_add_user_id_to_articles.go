package migrations

import (
	"articles-go/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1624363890AddUserIDToArticles() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1624363890",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AddColumn(&models.Article{}, "user_id")
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.Article{}, "user_id")
		},
	}
}
