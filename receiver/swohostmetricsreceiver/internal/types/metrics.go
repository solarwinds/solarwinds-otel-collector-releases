package types

import "go.opentelemetry.io/collector/pdata/pmetric"

// Functor capable of providing metric slice.
type MetricsEmittingFunc func() (pmetric.MetricSlice, error)

// Interface prescribing what metric emitter (actual metric producer)
// needs to implement.
type MetricEmitterInterface interface {
	// callback for initializing metrics internals.
	Initialize() error

	// emitter call back capable of providing metric slice.
	GetEmittingFunction() MetricsEmittingFunc
}
