package hardwareinventoryscraper

import (
	"go.opentelemetry.io/collector/component"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scope"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/scraper"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/hardwareinventoryscraper/metrics/cpu"
)

const (
	cpuScopeName = "otelcol/swohostmetricsreceiver/hardwareinventory/cpu"
)

//nolint:gochecknoglobals // Private, read-only.
var scraperType component.Type = component.MustNewType("hardwareinventory")

func ScraperType() component.Type {
	return scraperType
}

type Scraper struct {
	scraper.Manager
	config *Config
}

var _ scraper.Scraper = (*Scraper)(nil)

func NewHardwareInventoryScraper(
	config *Config,
) (*Scraper, error) {
	descriptor := &scraper.Descriptor{
		Type: ScraperType(),
		ScopeDescriptors: map[string]scope.Descriptor{
			cpuScopeName: {
				ScopeName: cpuScopeName,
				MetricDescriptors: map[string]metric.Descriptor{
					cpu.Name: {Create: cpu.NewEmitter},
				},
			},
		},
	}

	managerConfig := &scraper.ManagerConfig{
		ScraperConfig:           &config.ScraperConfig,
		DelayedProcessingConfig: &config.DelayedProcessingConfig,
	}

	s := &Scraper{
		Manager: scraper.NewScraperManager(),
		config:  config,
	}

	if err := s.Init(descriptor, managerConfig); err != nil {
		return nil, err
	}

	return s, nil
}
