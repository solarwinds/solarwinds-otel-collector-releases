package hostinfoscraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scraper"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

//nolint:gochecknoglobals // Private, read-only.
var scraperType component.Type = component.MustNewType("hostinfo")

func ScraperType() component.Type {
	return scraperType
}

type Factory struct{}

var _ types.ScraperFactory = (*Factory)(nil)

// CreateDefaultConfig implements types.ScraperFactory.
func (f *Factory) CreateDefaultConfig() component.Config {
	return types.CreateDefaultScraperConfig()
}

// CreateScraper implements types.ScraperFactory.
func (f *Factory) CreateScraper(
	_ context.Context,
	_ receiver.Settings,
	cfg component.Config,
) (scraperhelper.Scraper, error) {
	return scraper.CreateScraper[types.ScraperConfig, Scraper](
		ScraperType(),
		cfg,
		NewHostInfoScraper,
	)
}
