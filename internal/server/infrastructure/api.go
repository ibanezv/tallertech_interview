package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ibanezv/tallertech_interview/internal/infrastructure/http"
)

type AppServer struct {
	app              *fiber.App
	createController *http.CreateEventController
	getAllController *http.GetAllEventController
	getOneController *http.GetOneEventController
}

func (s *AppServer) Listen() error {
	app := fiber.New()

	app.Use(recover.New())

	api := s.app.Group("api/v1")
	eventsGroup := api.Group("/event")
	eventsGroup.Get("/", s.getAllController.Handler)
	eventsGroup.Post("/", s.createController.Validate, s.createController.Handler)
	eventsGroup.Get("/:id", s.getOneController.Handler)

	return app.Listen(":8080")
}

func NewAppServer(createCtrl *http.CreateEventController,
	getCtrl *http.GetOneEventController,
	getAllCtrl *http.GetAllEventController) *AppServer {
	return &AppServer{
		app:              fiber.New(),
		createController: createCtrl,
		getAllController: getAllCtrl,
		getOneController: getCtrl,
	}
}
