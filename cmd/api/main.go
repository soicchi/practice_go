package main

import (
	"fmt"
	"log"

	"go-practice/config"
)

func main() {
	fmt.Println("Hello, World!")

	// config setup
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
