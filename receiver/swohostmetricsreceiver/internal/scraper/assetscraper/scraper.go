package assetscraper

import (
	"go.opentelemetry.io/collector/component"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper/metrics/installedsoftware"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper/metrics/installedupdates"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scope"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scraper"
)

//nolint:gochecknoglobals // Private, read-only.
var scraperType component.Type = component.MustNewType("asset")

func ScraperType() component.Type {
	return scraperType
}

const (
	scopeMetricsName = "otelcol/swohostmetricsreceiver/asset"
)

type AssetScraper struct {
	scraper.Manager
	config *Config
}

var _ scraper.Scraper = (*AssetScraper)(nil)

func NewAssetScraper(
	config *Config,
) (*AssetScraper, error) {
	descriptor := &scraper.Descriptor{
		Type: ScraperType(),
		ScopeDescriptors: map[string]scope.Descriptor{
			scopeMetricsName: {
				ScopeName: scopeMetricsName,
				MetricDescriptors: map[string]metric.Descriptor{
					installedsoftware.Name: {Create: installedsoftware.NewEmitter},
					installedupdates.Name:  {Create: installedupdates.NewEmitter},
				},
			},
		},
	}

	managerConfig := &scraper.ManagerConfig{
		ScraperConfig:           &config.ScraperConfig,
		DelayedProcessingConfig: &config.DelayedProcessingConfig,
	}

	s := &AssetScraper{
		Manager: scraper.NewScraperManager(),
		config:  config,
	}

	if err := s.Init(descriptor, managerConfig); err != nil {
		return nil, err
	}

	return s, nil
}
