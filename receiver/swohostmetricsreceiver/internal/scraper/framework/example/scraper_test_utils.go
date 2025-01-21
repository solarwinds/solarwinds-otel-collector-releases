package example

import (
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
)

func NewMetricEmitterS1M1() metric.Emitter {
	return metric.CreateMetricEmitterMockV2(scope1metric1, 1, 101)
}

func NewMetricEmitterS1M2() metric.Emitter {
	return metric.CreateMetricEmitterMockV2(scope1metric2, 2, 102)
}

func NewMetricEmitterS2M1() metric.Emitter {
	return metric.CreateMetricEmitterMockV2(scope2metric1, 1, 201)
}
