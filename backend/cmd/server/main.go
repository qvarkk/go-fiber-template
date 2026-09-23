package main

import (
	"qvarkk/gofiber/internal/config"
	"qvarkk/gofiber/internal/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(config.Load),

		logger.LoggerModule,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx.Invoke(func(log *zap.Logger) {
			log.Info("Application initialized successfully.")
		}),
	).Run()
}
