package migrations

import (
	"articles-go/config"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1624269650CreateArticlesTable(),
			m1624347947CreateCategoriesTable(),
			m1624350474AddCategoryIDToArticles(),
			m1624355844CreateUsersTable(),
			m1624363890AddUserIDToArticles(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
