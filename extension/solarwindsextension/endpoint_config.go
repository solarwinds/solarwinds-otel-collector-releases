package solarwindsextension

import (
	"github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension/internal"
	"go.opentelemetry.io/collector/config/configopaque"
)

type EndpointConfig interface {
	Url() (string, error)
	Token() configopaque.String
}

type endpointConfig struct{ cfg *internal.Config }

var _ EndpointConfig = (*endpointConfig)(nil)

func newEndpointConfig(cfg *internal.Config) *endpointConfig {
	return &endpointConfig{cfg: cfg}
}

func (c *endpointConfig) Url() (string, error) {
	return c.cfg.EndpointUrl()
}

func (c *endpointConfig) Token() configopaque.String {
	return c.cfg.IngestionToken
}
