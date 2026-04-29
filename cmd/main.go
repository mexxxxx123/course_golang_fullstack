package main

import (
	"mexxx1/golang-fullstack/config"
	"mexxx1/golang-fullstack/internal/home"
	"mexxx1/golang-fullstack/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Configs

	config.Init()
	logConf := config.NewLogConfig()
	logger := logger.NewLogger(logConf)

	// App

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())
	app.Static("/static", "./public")

	home.NewHandler(app, logger)

	app.Listen(":3001")
}
