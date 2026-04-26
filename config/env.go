package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const defaultUrlInt = 1
const defaultUrlBool = false
const defaultUrlString = "http://localhost:3001/"

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded!")
}

type DatabaseConfig struct {
	Url string
}

func NewDataBaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getEnvString("DATABASE_URL", ""),
		// url: getEnvInt("DATABASE_URL"),
	}
}

type LogConfig struct {
	Level  int
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getEnvInt("LOG_LEVEL", 0),
		Format: getEnvString("LOG_FORMAT", "json"),
	}
}

func getEnvString(key string, def string) string {
	url := os.Getenv(key)
	if url == "" {
		return def
	}
	return url
}

func getEnvInt(key string, def int) int {
	urlInt, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return def
	}

	return urlInt
}

func getEnvBool(key string, def bool) bool {
	urlBool, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return def
	}
	return urlBool
}
