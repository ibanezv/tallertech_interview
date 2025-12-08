package server

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ibanezv/tallertech_interview/internal/infrastructure/http"
)

type AppServer struct {
	app              *fiber.App
	createController *http.CreateEventController
}

func (s *AppServer) Listen() error {
	app := fiber.New()

	app.Use(recover.New())

	api := app.Group("api/v1")
	eventsGroup := api.Group("/event")
	eventsGroup.Get("/", nil)
	eventsGroup.Post("/", s.createController.Handler)
	eventsGroup.Get("/:eventId", nil)

	return nil
}
