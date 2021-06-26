package migrations

import (
	"articles-go/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1624350474AddCategoryIDToArticles() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1624350474",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Migrator().AddColumn(&models.Article{}, "category_id"); err != nil {
				return err
			}

			var articles []models.Article
			tx.Unscoped().Find(&articles)
			for _, article := range articles {
				article.CategoryID = 2
				tx.Save(&article)
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.Article{}, "category_id")
		},
	}
}
