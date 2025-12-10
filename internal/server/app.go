package server

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/ibanezv/tallertech_interview/internal/server/infrastructure"
	"github.com/ibanezv/tallertech_interview/internal/server/infrastructure/modinjector"
	"go.uber.org/fx"
)

const (
	startMsg           = "starting application"
	shutdownMsg        = "shutdown application"
	serverListenErrMsg = "failure at web server startup"
)

func initialize(ctx context.Context, app *infrastructure.AppServer) {
	log.Info(ctx, startMsg)
	if err := app.Listen(); err != nil {
		log.Fatal(ctx, serverListenErrMsg)
	}

	log.Info(ctx, shutdownMsg)
}

// InitServer initializes the FX with the Application.
func InitServer() {
	fx.New(
		NewServer(),
		fx.Invoke(initialize),
	).Run()
}

// NewServer returns the Application's http server with all its dependencies.
func NewServer() fx.Option {
	injections := []interface{}{
		infrastructure.NewAppServer,
	}

	injections = append(injections, modinjector.GetAppInjector()...)
	injections = append(injections, modinjector.GetDomainInjector()...)
	injections = append(injections, modinjector.GetInfraInjector()...)

	return fx.Provide(injections...)
}
