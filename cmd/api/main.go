package main

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"
	"go-practice/internal/models/migrations"
	"go-practice/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	fmt.Println("Hello, World!")

	// config setup
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitializeDB(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	// migration
	if err := database.Migrate(db, []*gormigrate.Migration{
		migrations.UsersTable(),
	}); err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed")

	// initialize echo
	e := echo.New()
	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
