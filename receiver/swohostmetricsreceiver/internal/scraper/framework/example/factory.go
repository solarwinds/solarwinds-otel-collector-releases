package example

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scraper"
)

// CreateScraperExplicitely represents the way how to create scraper with
// a possibility to run some additional code.
func CreateScraperExplicitly(
	_ context.Context,
	_ receiver.Settings,
	cfg component.Config,
) (scraperhelper.Scraper, error) {
	// Create scraper directly through allocating callback.
	exampleScraper, err := NewExemplaryScraper(cfg.(*ScraperConfig))
	if err != nil {
		m := fmt.Sprintf("scraper '%s	' creation failed", ScraperType())
		zap.L().Error(m, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", m, err)
	}

	// In case there is a need to add some more code to be run.
	//
	// Create an OTEL scraper.
	otelScraper, err := scraperhelper.NewScraper(
		ScraperType(),
		exampleScraper.Scrape,
		scraperhelper.WithStart(exampleScraper.Start),
		scraperhelper.WithShutdown(exampleScraper.Shutdown),
	)

	return otelScraper, err
}

func CreateScraperImplicitly(
	_ context.Context,
	_ receiver.Settings,
	cfg component.Config,
) (scraperhelper.Scraper, error) {
	return scraper.CreateScraper[ScraperConfig, ExemplaryScraper](
		ScraperType(),
		cfg,
		NewExemplaryScraper,
	)
}
