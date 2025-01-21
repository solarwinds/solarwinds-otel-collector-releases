package internal

import "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"

const (
	ArtificialMetricName = "swo.artifial.metric"
)

type ArtificialMetricEmitter struct {
	Emitter types.MetricsEmittingFunc
}

func (emitter *ArtificialMetricEmitter) Initialize() error {
	return nil
}

func (emitter *ArtificialMetricEmitter) GetEmittingFunction() types.MetricsEmittingFunc {
	return emitter.Emitter
}
