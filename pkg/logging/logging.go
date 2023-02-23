// Package logging
// shortcuts for logging using https://github.com/rs/zerolog
package logging

import (
	"github.com/rs/zerolog/log"
)

func LogInfo(message string) {
	log.Info().Msg(message)
}

func LogWarning(message string) {
	log.Warn().Msg(message)
}

func LogError(message string, err error) {
	log.Error().Err(err).Msg(message)
}

func LogPanic(message string, err error) {
	log.Panic().Err(err).Msg(message)
}
