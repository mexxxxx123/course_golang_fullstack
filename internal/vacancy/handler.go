package vacancyform

import (
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyFormHandler struct {
	// Role        string
	// Company     string
	// Location    string
	// NameCompany string
	// Zp          string
	// Email       string
	router fiber.Router
	logger *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &VacancyFormHandler{
		router: router,
		logger: logger,
	}
	vacGroup := router.Group("/vacancy")
	vacGroup.Post("/", h.createVacancy)
}

func (h *VacancyFormHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.logger.Info().Msg(email)
	component := components.Notification("Вакансия создана")
	return tadapter.Render(c, component)
}
