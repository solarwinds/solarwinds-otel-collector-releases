package scope

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func CreateCustomScopeEmitter(
	name string,
	mes map[string]metric.Emitter,
) Emitter {
	return CreateDefaultScopeEmitter(name, mes)
}

type emitterMock struct {
	name string
}

var _ Emitter = (*emitterMock)(nil)

// Emit implements ScopeEmitter.
func (s *emitterMock) Emit() *Result {
	return &Result{
		Data:  pmetric.ScopeMetricsSlice{},
		Error: nil,
	}
}

// Init implements ScopeEmitter.
func (s *emitterMock) Init() error {
	return nil
}

// Name implements ScopeEmitter.
func (s *emitterMock) Name() string {
	return s.name
}
