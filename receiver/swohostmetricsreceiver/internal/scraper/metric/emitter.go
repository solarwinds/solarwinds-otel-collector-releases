package metric

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Result contains metrics as data and possible error
// from emit of one metric emitter.
type Result struct {
	MetricSlice pmetric.MetricSlice
	Error       error
}
