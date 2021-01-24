package config

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

// SentryConfig ...
type SentryConfig struct {
	SentryDSN string `yaml:"dsn" env:"SENTRY_DSN"`
}

func (s *SentryConfig) init() {
	if s.SentryDSN == "" {
		return
	}

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   s.SentryDSN,
		Debug: true,
	}); err != nil {
		log.Errorf("Sentry initialization failed: %v", err)
	}
}
