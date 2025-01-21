package metric

import "go.opentelemetry.io/collector/pdata/pmetric"

// Result is structure representing result from
// metric emitter.
type Result struct {
	// Data contains metric slice on success.
	Data pmetric.MetricSlice
	// Error is filled on failure, otherwise nil is returned.
	Error error
}

// Emitter is prescription for metric emitter in
// scraping framework.
type Emitter interface {
	// Init initializes metric emitter. Returns error
	// when fail, otherwise nil is returned.
	Init() error
	// Emit produces pointer to emitted metric result.
	Emit() *Result
	// Name returns name of metric emitted by metric emitter.
	Name() string
}
