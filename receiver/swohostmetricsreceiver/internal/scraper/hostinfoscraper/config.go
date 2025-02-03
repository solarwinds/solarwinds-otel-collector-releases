package hostinfoscraper

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/hostinfoscraper/metrics/uptime"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
	"go.opentelemetry.io/collector/component"
)

func CreateDefaultConfig() component.Config {
	return &types.ScraperConfig{
		Metrics: map[string]types.MetricSettingsConfig{
			uptime.MetricName: {
				Enabled: true,
			},
		},
	}
}
