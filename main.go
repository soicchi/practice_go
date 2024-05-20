package main

import (
	"log"
	"os"

	"go-practice/cmd"
	"go-practice/internal/config"
	"go-practice/internal/database"
	"go-practice/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// command line
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	// config setup
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
	log.Println(db)

	// initialize echo
	e := echo.New()
	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
