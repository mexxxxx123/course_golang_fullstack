package main

import (
	"mexxx1/golang-fullstack/config"
	"mexxx1/golang-fullstack/internal/home"
	vacancyform "mexxx1/golang-fullstack/internal/vacancy"
	"mexxx1/golang-fullstack/pkg/logger"
	"mexxx1/golang-fullstack/pkg/logger/database"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Configs

	config.Init()
	logConf := config.NewLogConfig()
	dbConf := config.NewDataBaseConfig()
	logger := logger.NewLogger(logConf)

	// App
	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	dbpool := database.CreateDbPool(dbConf, logger)
	defer dbpool.Close()

	home.NewHandler(app, logger)
	vacancyform.NewHandler(app, logger)

	app.Listen(":3001")
}
