package example

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"

// ScraperConfig represents typical scraper config, which is tend
// to be used in this example package.
type ScraperConfig struct {
	types.DelayedProcessingConfig `mapstructure:",squash"`
	types.ScraperConfig           `mapstructure:",squash"`
}
