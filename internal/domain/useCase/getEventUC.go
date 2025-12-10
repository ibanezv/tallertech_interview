package usecase

import (
	"github.com/ibanezv/tallertech_interview/internal/domain"
	"github.com/ibanezv/tallertech_interview/internal/infrastructure/repo"
)

type EventGetOne interface {
	Do(string) (*domain.Event, error)
}

type EventGetOneUc struct {
	db repo.DataBase
}

func NewEventGetOneUc(db repo.DataBase) *EventGetOneUc {
	return &EventGetOneUc{db: db}
}

func (e *EventGetOneUc) Do(id string) (*domain.Event, error) {
	event, err := e.db.GetById(id)
	if err != nil {
		return nil, err
	}
	return domain.ToEvent(event), nil
}
