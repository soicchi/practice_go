package database

import (
	"fmt"

	"go-practice/internal/config"
	"go-practice/internal/models/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cfg *config.Database) (*gorm.DB, error) {
	dsn := createDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsList())
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}

func RollbackLast(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsList())
	if err := m.RollbackLast(); err != nil {
		return err
	}

	return nil
}

func createDSN(cfg *config.Database) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)
}

func migrationsList() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		migrations.UsersTable(),
	}
}