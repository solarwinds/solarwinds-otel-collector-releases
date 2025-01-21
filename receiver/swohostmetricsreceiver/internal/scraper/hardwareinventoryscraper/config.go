package hardwareinventoryscraper

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
	"go.opentelemetry.io/collector/component"
)

// Config for Hardware Inventory scraper.
type Config struct {
	types.DelayedProcessingConfig `mapstructure:",squash"`
	types.ScraperConfig           `mapstructure:",squash"`
}

// Config implements component.Config interface.
var _ component.Config = (*Config)(nil)
