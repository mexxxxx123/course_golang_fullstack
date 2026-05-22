package main

import (
	"mexxx1/golang-fullstack/config"
	"mexxx1/golang-fullstack/internal/home"
	"mexxx1/golang-fullstack/internal/vacancy"
	"mexxx1/golang-fullstack/pkg/logger"
	"mexxx1/golang-fullstack/pkg/logger/database"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Config`s

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

	//DbPool

	dbpool := database.CreateDbPool(dbConf, logger)
	defer dbpool.Close()

	//Repo`s

	VacancyRepo := vacancy.NewVacancyRepository(dbpool, logger)

	//Handler`s

	home.NewHandler(app, logger, VacancyRepo)
	vacancy.NewHandler(app, logger, VacancyRepo)

	app.Listen(":3001")
}
