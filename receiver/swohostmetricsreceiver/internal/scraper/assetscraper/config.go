package assetscraper

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
	"go.opentelemetry.io/collector/component"
)

// Asset scraper config.
type Config struct {
	types.DelayedProcessingConfig `mapstructure:",squash"`
	types.ScraperConfig           `mapstructure:",squash"`
}

// implements compnent.Config interface.
var _ component.Config = (*Config)(nil)
