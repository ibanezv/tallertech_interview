package usecase

import (
	"github.com/ibanezv/tallertech_interview/internal/domain"
	"github.com/ibanezv/tallertech_interview/internal/infrastructure/repo"
)

type EventCreation interface {
	Do(*domain.Event) (*domain.Event, error)
}

type CreateEventUc struct {
	db repo.DataBase
}

func NewCreateEventUc(db repo.DataBase) *CreateEventUc {
	return &CreateEventUc{db: db}
}

func (c *CreateEventUc) Do(event *domain.Event) (*domain.Event, error) {
	eventDb, err := c.db.Save(domain.ToDbEvent(event))
	if err != nil {
		return nil, err
	}
	return domain.ToEvent(eventDb), nil
}
