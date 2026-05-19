package logger

import (
	"mexxx1/golang-fullstack/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(conf *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(conf.Level))
	var logger zerolog.Logger
	if conf.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return &logger
}
