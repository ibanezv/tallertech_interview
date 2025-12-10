package usecase

import (
	"github.com/ibanezv/tallertech_interview/internal/domain"
	"github.com/ibanezv/tallertech_interview/internal/infrastructure/repo"
)

type EventGetAll interface {
	Do() ([]domain.Event, error)
}

type GetAllEventUc struct {
	db repo.DataBase
}

func NewGetAllEventUc(db repo.DataBase) *GetAllEventUc {
	return &GetAllEventUc{db: db}
}

func (g *GetAllEventUc) Do() ([]domain.Event, error) {
	events, err := g.db.GetAll()
	if err != nil {
		return nil, err
	}

	return domain.ToEventList(events), nil
}
