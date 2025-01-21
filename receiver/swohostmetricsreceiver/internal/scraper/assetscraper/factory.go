package assetscraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scraper"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

type Factory struct{}

var _ types.ScraperFactory = (*Factory)(nil)

// CreateDefaultConfig implements types.ScraperFactory.
func (*Factory) CreateDefaultConfig() component.Config {
	// in fact returns asset's scraper configuration covered by component.Config
	// type
	return &Config{
		DelayedProcessingConfig: types.DelayedProcessingConfig{},
		ScraperConfig:           *types.CreateDefaultScraperConfig(),
	}
}

// CreateScraper implements types.ScraperFactory.
func (*Factory) CreateScraper(
	_ context.Context,
	_ receiver.Settings,
	cfg component.Config,
) (scraperhelper.Scraper, error) {
	return scraper.CreateScraper[Config, AssetScraper](
		ScraperType(),
		cfg,
		NewAssetScraper,
	)
}
