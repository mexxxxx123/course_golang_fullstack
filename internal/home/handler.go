package home

import (
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	logger *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &HomeHandler{
		router: router,
		logger: logger,
	}
	router.Get("/", h.home)
	router.Get("/error", h.error)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("info")
	return fiber.NewError(fiber.StatusBadRequest, "Unauthorized")

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}
