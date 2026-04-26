package home

import (
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
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("info")
	return fiber.NewError(fiber.StatusBadRequest, "Unauthorized")

}

type User struct {
	Id   int
	Name string
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	// tmpl := template.Must(template.ParseFiles("./html/page.html"))
	// var tpl bytes.Buffer
	// if err := tmpl.Execute(&tpl, data); err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	// }
	// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	// return c.Send(tpl.Bytes())

	// data := struct {
	// 	Count   int
	// 	IsAdmin bool
	// }{
	// 	Count:   4,
	// 	IsAdmin: true,
	// }

	users := []User{
		{Id: 1, Name: "Anton"},
		{Id: 2, Name: "Petr"},
	}

	return c.Render("page", users)
}
