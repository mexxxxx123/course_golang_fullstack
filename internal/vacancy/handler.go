package vacancy

import (
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/validator"
	"mexxx1/golang-fullstack/views/components"
	"net/http"
	"time"

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
	vacGroup.Get("/", h.getAll)
}
func (h *VacancyFormHandler) getAll(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	vacancies, err := h.repo.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.logger.Error().Msg(err.Error())
	}

	return c.JSON(vacancies)

}

func (h *VacancyFormHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Role:        c.FormValue("role"),
		CompanyType: c.FormValue("company-type"),
		Location:    c.FormValue("location"),
		CompanyName: c.FormValue("company-name"),
		Salary:      c.FormValue("salary"),
		Email:       c.FormValue("email"),
		Createdat:   time.Now(),
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
		for key, value := range errors.Errors {
			for value1 := range value {
				h.logger.Error().Msg(errors.Errors[key][value1])
			}
		}
		return tadapter.Render(c, component, http.StatusBadRequest)
	}

	err := h.repo.AddVacancy(form)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		component = components.Notification("Ошибка на сервере, попробуйте позднее", components.NotificationFail)
		return tadapter.Render(c, component, http.StatusInternalServerError)

	}
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}
