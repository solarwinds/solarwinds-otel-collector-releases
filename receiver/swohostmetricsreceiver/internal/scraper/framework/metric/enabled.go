package metric

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

type Name = string

type void = struct{}

// Enabled is representation of enabled metrics.
type Enabled struct {
	// Metrics itself indexed by metric name.
	// When metric is contained in map, it is supposed to be
	// enable by configuration.
	Metrics map[Name]*void
}

// GetEnabledMetrics returns enabled metrics from whole scraperConfig.
// Only metrics, which are enabled, will be contained in returned struct.
// On failure error is returned, on success nil is returned.
func GetEnabledMetrics(
	scraperName string,
	scraperConfig *types.ScraperConfig,
) (*Enabled, error) {
	// Check if there are at least some metrics configured.
	if len(scraperConfig.Metrics) == 0 {
		message := fmt.Sprintf(
			"no configured metrics for scraper '%s'",
			scraperName,
		)
		zap.L().Error(message)
		return nil, errors.New(message)
	}

	// Traverse scraper config and setup only enabled metrics.
	em := new(Enabled)
	em.Metrics = make(map[Name]*void, 0)
	for mn, c := range scraperConfig.Metrics {
		if c.Enabled {
			em.Metrics[mn] = new(void)
		}
	}

	if len(em.Metrics) == 0 {
		message := fmt.Sprintf(
			"no enabled metrics available for scpraper '%s'",
			scraperName,
		)
		zap.L().Error(message)
		return nil, errors.New(message)
	}

	return em, nil
}
