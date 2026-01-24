package gutils

import (
	"github.com/skypbc/goutils/gerrors"
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Close закрывает ресурс и если задан уровень логирования, то логирует ошибку закрытия при ее наличии.
func Close(closer io.Closer, level ...zerolog.Level) {
	if closer == nil {
		return
	}
	if len(level) == 0 {
		_ = closer.Close()
		return
	}
	var logger *zerolog.Event
	switch level[0] {
	case zerolog.DebugLevel:
		logger = log.Debug()
	case zerolog.InfoLevel:
		logger = log.Info()
	case zerolog.WarnLevel:
		logger = log.Warn()
	case zerolog.ErrorLevel:
		logger = log.Error()
	case zerolog.FatalLevel:
		logger = log.Fatal()
	case zerolog.PanicLevel:
		logger = log.Panic()
	case zerolog.TraceLevel:
		logger = log.Trace()
	case zerolog.NoLevel:
		logger = log.WithLevel(zerolog.NoLevel)
	case zerolog.Disabled:
		logger = log.WithLevel(zerolog.Disabled)
	default:
		logger = log.Error()
	}
	if err := closer.Close(); err != nil {
		err = gerrors.WrapWithSkip(2, err)
		logger.Err(err).Msgf(`Error closing resource: %s`, err.Error())
		return
	}
}
