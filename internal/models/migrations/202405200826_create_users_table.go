package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func UsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202405200826",
		Migrate: func(tx *gorm.DB) error {
			type user struct {
				gorm.Model
				Username string `gorm:"size:255;not null;unique"`
				Email    string `gorm:"size:255;not null;unique"`
			}
			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
