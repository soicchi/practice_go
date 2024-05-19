package main

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"
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
}
