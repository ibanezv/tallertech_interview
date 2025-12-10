package modinjector

import (
	"context"

	"github.com/ibanezv/tallertech_interview/internal/infrastructure/http"
)

func GetInfraInjector() []interface{} {
	return []interface{}{
		context.Background,
		http.NewCreateEventController,
		http.NewGetAllEventController,
		http.NewGetOneEventController,
	}
}
