package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

var DefaultLevel = "info"

func New(level string) zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse log level, defaulting to info")
		return logger.Level(zerolog.InfoLevel)
	}
	return logger.Level(lvl)
}

func Get(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
