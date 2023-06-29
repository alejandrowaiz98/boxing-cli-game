package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func SetLogger() {

	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

}

func GetLogger() zerolog.Logger {

	return logger

}
