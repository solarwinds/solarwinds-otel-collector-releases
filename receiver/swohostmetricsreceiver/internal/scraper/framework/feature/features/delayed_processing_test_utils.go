package features

import (
	"time"

	"github.com/solarwinds-cloud/uams-otel-collector-plugin/test"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
)

type DelayedProcessingMock struct {
	InitCC        test.CallsCounter
	IsReadyCC     test.CallsCounter
	UpdateCC      test.CallsCounter
	initResult    error
	isReadyResult bool
}

var _ DelayedProcessing = (*DelayedProcessingMock)(nil)

func CreateDelayedProcessingMock(
	initResult error,
	isReadyResult bool,
) *DelayedProcessingMock {
	return &DelayedProcessingMock{
		InitCC:        test.CallsCounter{},
		IsReadyCC:     test.CallsCounter{},
		UpdateCC:      test.CallsCounter{},
		initResult:    initResult,
		isReadyResult: isReadyResult,
	}
}

// InitDelayedProcessing implements DelayedProcessing.
func (d *DelayedProcessingMock) InitDelayedProcessing(
	*types.DelayedProcessingConfig,
	time.Time,
) error {
	d.InitCC.IncrementCount()
	return d.initResult
}

// IsReady implements DelayedProcessing.
func (d *DelayedProcessingMock) IsReady(time.Time) bool {
	d.IsReadyCC.IncrementCount()
	return d.isReadyResult
}

// UpdateLastProcessedTime implements DelayedProcessing.
func (d *DelayedProcessingMock) UpdateLastProcessedTime(time.Time) {
	d.UpdateCC.IncrementCount()
}
