package vacancy

import (
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/validator"
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
	repo   *VacancyRepository
}

func NewHandler(router fiber.Router, logger *zerolog.Logger, repo *VacancyRepository) {
	h := &VacancyFormHandler{
		router: router,
		logger: logger,
		repo:   repo,
	}
	vacGroup := router.Group("/vacancy")
	vacGroup.Post("/", h.createVacancy)
}

func (h *VacancyFormHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Role:        c.FormValue("role"),
		CompanyType: c.FormValue("company-type"),
		Location:    c.FormValue("location"),
		CompanyName: c.FormValue("company-name"),
		Salary:      c.FormValue("salary"),
		Email:       c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{
			Name:    "Role",
			Field:   form.Role,
			Message: "Не задана Должность",
		},
		&validators.StringIsPresent{
			Name:    "CompanyType",
			Field:   form.CompanyType,
			Message: "Не задана Сфера компании",
		},
		&validators.StringIsPresent{
			Name:    "Location",
			Field:   form.Location,
			Message: "Не задано Местоположение",
		},
		&validators.StringIsPresent{
			Name:    "CompanyName",
			Field:   form.CompanyName,
			Message: "Не задано Название компании",
		},
		&validators.StringIsPresent{
			Name:    "Salary",
			Field:   form.Salary,
			Message: "Не задана Заработная плата",
		},
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Не задан или введен неверно Email",
		},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification((validator.PrintErrors(*errors)), components.NotificationFail)
		return tadapter.Render(c, component)
	}

	err := h.repo.AddVacancy(form)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		component = components.Notification("Ошибка на сервере, попробуйте позднее", components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
