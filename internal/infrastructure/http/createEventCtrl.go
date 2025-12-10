package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/tallertech_interview/internal/domain"
	usecase "github.com/ibanezv/tallertech_interview/internal/domain/useCase"
)

type CreateEventController struct {
	useCase usecase.EventCreation
}

func (c *CreateEventController) Handler(fCtx *fiber.Ctx) error {
	var event domain.Event

	if err := fCtx.BodyParser(&event); err != nil {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	eventResult, err := c.useCase.Do(&event)
	if err != nil {
		return fCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	return fCtx.Status(fiber.StatusOK).JSON(eventResult)
}

func (c *CreateEventController) Validate(fCtx *fiber.Ctx) error {
	var event domain.Event

	if err := fCtx.BodyParser(&event); err != nil {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	if len(event.Description) == 0 || len(event.Description) > 100 {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	if !event.EndTime.After(event.StartTime) {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	return fCtx.Next()
}

func NewCreateEventController(uc usecase.EventCreation) *CreateEventController {
	return &CreateEventController{useCase: uc}
}
