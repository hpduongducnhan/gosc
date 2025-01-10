package envvars

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().
		Timestamp().
		Str("logger_name", "gosc-envvars").
		Logger().
		Level(zerolog.WarnLevel) // Set a specific level for the special logger
	logger.Debug().Msg("Init elk client")
}
