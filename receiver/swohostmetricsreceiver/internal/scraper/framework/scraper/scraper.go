package scraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Scraper represents general prescription for scraping.
// It mimics functions used in OTEL collector in scraper helper.
// Start function mimics component.StartFunc by signature.
// Shutdown function mimics component.ShutdownFunc.
// Scraper function mimics scraperhelper.ScrapeFunc.
type Scraper interface {
	Start(context.Context, component.Host) error
	Shutdown(context.Context) error
	Scrape(context.Context) (pmetric.Metrics, error)
}
