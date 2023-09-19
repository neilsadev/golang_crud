package helper

import (
	"github.com/rs/zerolog/log"
)

func ErrorPanic(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
