package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/tallertech_interview/internal/domain"
	usecase "github.com/ibanezv/tallertech_interview/internal/domain/useCase"
)

type GetOneEventController struct {
	useCase usecase.EventGetOne
}

func (c *GetOneEventController) Handler(fCtx *fiber.Ctx) error {
	var event domain.Event

	if err := fCtx.BodyParser(&event); err != nil {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	id := ""
	eventResult, err := c.useCase.Do(id)
	if err != nil {
		return fCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	if eventResult == nil {
		return fCtx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "event not found"})
	}

	return fCtx.Status(fiber.StatusOK).JSON(eventResult)
}

func NewGetOneEventController(uc usecase.EventGetOne) *GetOneEventController {
	return &GetOneEventController{useCase: uc}
}
