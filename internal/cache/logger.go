package cache

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/rs/zerolog"
)

var _ badger.Logger = (*logger)(nil)

type logger struct {
	log zerolog.Logger
}

func (l *logger) Errorf(s string, i ...any) {
	l.log.Error().Msgf(s, i...)
}

func (l *logger) Warningf(s string, i ...any) {
	l.log.Warn().Msgf(s, i...)
}

func (l *logger) Infof(s string, i ...any) {
	l.log.Info().Msgf(s, i...)
}

func (l *logger) Debugf(s string, i ...any) {
	l.log.Debug().Msgf(s, i...)
}
