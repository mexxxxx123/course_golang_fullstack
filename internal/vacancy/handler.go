package vacancy

import (
	"fmt"
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyFormHandler struct {
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или неверный",
		},
	)
	h.logger.Info().Msg(form.Email)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification("Ошибки", components.NotificationFail)
		return tadapter.Render(c, component)
	}
	for key, value := range errors.Errors {

		h.logger.Info().Msg(form.Email)
		fmt.Println(key, value)
	}
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
