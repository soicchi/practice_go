package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func PostsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202405200953",
		Migrate: func(tx *gorm.DB) error {
			type post struct {
				gorm.Model
				Title  string `gorm:"size:255;not null"`
				UserID uint
			}
			if err := tx.Migrator().CreateTable(&post{}); err != nil {
				return err
			}

			return tx.Exec("ALTER TABLE posts ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE").Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("posts")
		},
	}
}
