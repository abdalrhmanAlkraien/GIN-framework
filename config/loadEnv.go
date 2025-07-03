package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("call file")
	DbConfig()

}
