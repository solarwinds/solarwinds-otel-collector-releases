package scraper

import (
	"github.com/solarwinds-cloud/uams-otel-collector-plugin/test"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

type SchedulerMock struct {
	scheduleCC     test.CallsCounter
	scheduleResult error
	runtime        *Runtime
}

var _ Scheduler = (*SchedulerMock)(nil)

func CreateSchedulerMock(
	scheduleResult error,
	runtime *Runtime,
) *SchedulerMock {
	return &SchedulerMock{
		scheduleCC:     test.CallsCounter{},
		scheduleResult: scheduleResult,
		runtime:        runtime,
	}
}

// Schedule implements Scheduler.
func (s *SchedulerMock) Schedule(*Descriptor, *types.ScraperConfig) (*Runtime, error) {
	s.scheduleCC.IncrementCount()
	return s.runtime, s.scheduleResult
}
