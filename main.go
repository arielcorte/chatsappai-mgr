package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	CreateURLMappings()

	if err := Engine.Run(); err != nil {
		panic(err)
	}
}
