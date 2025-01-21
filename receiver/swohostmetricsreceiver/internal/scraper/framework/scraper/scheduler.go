package scraper

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

// Schedule prescribes how transformation of
// configuration oriented structure to runtime
// configuration should look like.
type Scheduler interface {
	// Schedule method process the transformation mentioned above.
	// Provider descriptor pointer together with scraper configuration
	// is "translated" into Runtime object. When fails error is returned,
	// otherwise nil is return in place of error.
	Schedule(
		*Descriptor,
		*types.ScraperConfig,
	) (*Runtime, error)
}

type scheduler struct{}

var _ Scheduler = (*scheduler)(nil)

func NewScraperScheduler() Scheduler {
	return new(scheduler)
}

// Schedule implements ScraperScheduler.
func (*scheduler) Schedule(
	descriptor *Descriptor,
	config *types.ScraperConfig,
) (*Runtime, error) {
	sn := descriptor.Type

	// Obtains enabled metrics for scheduled scraper.
	enabledMetrics, err := metric.GetEnabledMetrics(sn.String(), config)
	if err != nil {
		m := fmt.Sprintf("failed to get enabled metrics for scraper '%s'", sn)
		zap.L().Error(m, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", m, err)
	}

	// Assembly Scraper runtime based on enabled metrics.
	scraperRuntime, err := createScraperRuntime(descriptor, enabledMetrics)
	if err != nil {
		m := fmt.Sprintf("failed to create scraper runtime for scraper '%s'", sn)
		zap.L().Error(m, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", m, err)
	}

	zap.L().Sugar().Debugf(
		"scheduling of scraper '%s' finished successfully",
		sn)
	return scraperRuntime, nil
}
