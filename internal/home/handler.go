package home

import (
	"math"
	"mexxx1/golang-fullstack/internal/vacancy"
	"mexxx1/golang-fullstack/pkg/logger/tadapter"
	"mexxx1/golang-fullstack/views"
	"mexxx1/golang-fullstack/views/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *vacancy.VacancyRepository
	store      *session.Store
}

func NewHandler(router fiber.Router, logger *zerolog.Logger, repo *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:     router,
		logger:     logger,
		repository: repo,
		store:      store,
	}
	router.Get("/", h.home)
	router.Get("/error", h.error)
	router.Get("/loginPage", h.loginPage)
	router.Post("/login", h.login)
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
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	sses, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	if name, ok := sses.Get("name").(string); ok {
		h.logger.Info().Msg(name)
	}
	count := h.repository.CountAll()

	component := views.Main(vacancies, int(math.Ceil(float64(count)/float64(PAGE_ITEMS))), page)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) loginPage(c *fiber.Ctx) error {
	component := views.Login()
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("name", "Anton")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return tadapter.Render(c, component, http.StatusOK)

}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	form := SessionInfo{
		Email:    c.FormValue("LoginEmail"),
		Password: c.FormValue("LoginPassword"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "LoginEmail",
			Field:   form.Email,
			Message: "Не задан или введен неверно Email",
		},
		&validators.StringIsPresent{
			Name:    "LoginPassword",
			Field:   form.Password,
			Message: "Не задан пароль",
		},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		for key, value := range errors.Errors {
			for value1 := range value {
				h.logger.Error().Msg(errors.Errors[key][value1])
			}
		}
		component = components.Notification("Ошибка на сервере, попробуйте позднее", components.NotificationFail)
		h.logger.Error().Msg("1123")
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Set("email", form.Email)
	sess.Set("password", form.Password)

	if name, ok := sess.Get("password").(string); ok {
		h.logger.Info().Msg(name)
	}
	if name, ok := sess.Get("email").(string); ok {
		h.logger.Info().Msg(name)
	}

	if err := sess.Save(); err != nil {
		panic(err)
	}
	component = components.Notification("Логин выполнен", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)
}
