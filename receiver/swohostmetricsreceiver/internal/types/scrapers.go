package types

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Interface prescibing what scraper factory needs to implement.
type ScraperFactory interface {
	// Creates default scraper configuration.
	CreateDefaultConfig() component.Config

	// Creates scraper object, in case of failure error is returned.
	CreateScraper(ctx context.Context, settings receiver.Settings, cfg component.Config) (scraperhelper.Scraper, error)
}
