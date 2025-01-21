package scraper

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

// CreateScraper creates scraper in implicit way. Function packs
// all required checks and allocation into single call to be minimalistic
// in usage.
func CreateScraper[TConfig component.Config, TScraper Scraper](
	scraperName component.Type,
	config component.Config,
	sAllocator func(*TConfig) (*TScraper, error),
) (scraperhelper.Scraper, error) {
	scraper, err := sAllocator(config.(*TConfig))
	if err != nil {
		m := fmt.Sprintf("scraper '%s' creation failed", scraperName)
		zap.L().Error(m, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", m, err)
	}

	otelScraper, err := scraperhelper.NewScraper(
		scraperName,
		(*scraper).Scrape,
		scraperhelper.WithStart((*scraper).Start),
		scraperhelper.WithShutdown((*scraper).Shutdown),
	)
	if err != nil {
		m := fmt.Sprintf("scraperhelper scraper '%s' creation failed", scraperName)
		zap.L().Error(m, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", m, err)
	}

	return otelScraper, nil
}
