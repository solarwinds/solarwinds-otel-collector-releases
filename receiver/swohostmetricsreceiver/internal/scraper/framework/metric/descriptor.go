package metric

import (
	"go.uber.org/zap"
)

// MetricEmitterCreateFunc is a functor for creation of
// metric emitter instance.
type EmitterCreateFunc func() Emitter

// MetricDescriptor represent description of one metric.
type Descriptor struct {
	// Creator function for creation of specific metric emitter for
	// metric represented by this descriptor.
	Create EmitterCreateFunc
}

func TraverseThroughMetricDescriptors(
	metricDescriptors map[string]Descriptor,
	enabledMetrics *Enabled,
) map[string]Emitter {
	mes := make(map[string]Emitter, 0)

	for mName, mDescriptor := range metricDescriptors {
		// Metric is not among enabled ones.
		if _, found := enabledMetrics.Metrics[mName]; !found {
			continue
		}

		// Metric is enabled by config, let's use it.
		zap.L().Sugar().Debugf("creating metric emitter for '%s", mName)
		me := mDescriptor.Create()
		mes[me.Name()] = me
	}

	return mes
}
