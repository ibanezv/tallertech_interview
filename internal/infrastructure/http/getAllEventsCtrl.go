package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/tallertech_interview/internal/domain"
	usecase "github.com/ibanezv/tallertech_interview/internal/domain/useCase"
)

type GetAllEventController struct {
	useCase usecase.EventGetAll
}

func (c *GetAllEventController) Handler(fCtx *fiber.Ctx) error {
	var event domain.Event

	if err := fCtx.BodyParser(&event); err != nil {
		return fCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	eventResult, err := c.useCase.Do()
	if err != nil {
		return fCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	return fCtx.Status(fiber.StatusOK).JSON(eventResult)
}

func NewGetAllEventController(uc usecase.EventGetAll) *GetAllEventController {
	return &GetAllEventController{useCase: uc}
}
