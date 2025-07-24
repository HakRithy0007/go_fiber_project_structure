package loggers

import (
	"github.com/rs/zerolog"
)

func NewLog(loglevel string) {
	switch loglevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "trancing":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
}
