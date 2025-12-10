package modinjector

import usecase "github.com/ibanezv/tallertech_interview/internal/domain/useCase"

func GetDomainInjector() []interface{} {
	return []interface{}{
		usecase.NewCreateEventUc,
		usecase.NewEventGetOneUc,
		usecase.NewGetAllEventUc,
	}
}
