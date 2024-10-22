package main

import (
	"os"
	"strings"

	"github.com/thmelodev/ddd-events-api/src/modules/events"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
	"github.com/thmelodev/ddd-events-api/src/utils/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	loggerEnv := logger.Environment(appEnv)
	logger.Init(loggerEnv)

	var loggerOption fx.Option
	if strings.TrimSpace(appEnv) == "production" {
		loggerOption = fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		})
	} else {
		loggerOption = fx.Logger(logger.Get())
	}

	app := fx.New(
		loggerOption,
		fx.Provide(config.Init),
		httpServer.Module(),
		events.Module(),
	)

	app.Run()
}
