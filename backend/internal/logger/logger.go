package logger

import (
	"context"
	"fmt"
	"os"
	"qvarkk/gofiber/internal/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultFilename    = "logs/app.log"
	defaultMaxFilesize = 100 // 100 MB
	defaultMaxBackups  = 5
	defaultMaxAge      = 30 // 30 days
	defaultCompress    = true
)

var LoggerModule = fx.Module("logger",
	fx.Provide(NewZapLogger),
)

func NewZapLogger(lc fx.Lifecycle, env config.AppEnv, logLevel config.LogLevel) (*zap.Logger, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, fmt.Errorf("failed to parse log level %q: %w", logLevel, err)
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   defaultFilename,
		MaxSize:    defaultMaxFilesize,
		MaxBackups: defaultMaxBackups,
		MaxAge:     defaultMaxAge,
		Compress:   defaultCompress,
	}

	var encoderConfig zapcore.EncoderConfig
	if env.IsProduction() {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
	}

	encoderConfig.TimeKey = "time"
	encoderConfig.ConsoleSeparator = " | "

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	if !env.IsProduction() {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	stdoutEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberjackLogger), level),
		zapcore.NewCore(stdoutEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = logger.Sync()
			return nil
		},
	})

	return logger, nil
}
