package main

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"
	"go-practice/internal/router"

	"github.com/labstack/echo/v4"
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

	fmt.Println(db)

	// initialize echo
	e := echo.New()
	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
