package scraper

import (
	"github.com/solarwinds/solarwinds-otel-collector/pkg/testutil"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

type SchedulerMock struct {
	scheduleCC     testutil.CallsCounter
	scheduleResult error
	runtime        *Runtime
}

var _ Scheduler = (*SchedulerMock)(nil)

func CreateSchedulerMock(
	scheduleResult error,
	runtime *Runtime,
) *SchedulerMock {
	return &SchedulerMock{
		scheduleCC:     testutil.CallsCounter{},
		scheduleResult: scheduleResult,
		runtime:        runtime,
	}
}

// Schedule implements Scheduler.
func (s *SchedulerMock) Schedule(*Descriptor, *types.ScraperConfig) (*Runtime, error) {
	s.scheduleCC.IncrementCount()
	return s.runtime, s.scheduleResult
}
