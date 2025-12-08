package usecase

import "github.com/ibanezv/tallertech_interview/internal/domain"

type EventCreation interface {
	Do(*domain.Event) (domain.Event, error)
}

type CreateEventUc struct {
}

func (c *CreateEventUc) Do(event *domain.Event) (domain.Event, error) {

	return domain.Event{}, nil
}
