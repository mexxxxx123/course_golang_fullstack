package main

import (
	"log"
	"mexxx1/golang-fullstack/config"
	"mexxx1/golang-fullstack/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDataBaseConfig()
	log.Println(dbConf)

	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":3001")
}
