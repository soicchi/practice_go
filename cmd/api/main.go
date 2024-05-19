package main

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"
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
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
