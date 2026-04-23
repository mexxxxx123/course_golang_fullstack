package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded!")
}

type DatabaseConfig struct {
	url string
}

func NewDataBaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}
