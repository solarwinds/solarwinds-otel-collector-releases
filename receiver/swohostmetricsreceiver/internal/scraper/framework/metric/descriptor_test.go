package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TraverseThroughMetricDescriptors_onNoMatchEmptyEmitterMapIsReturned(t *testing.T) {
	descriptors := map[string]Descriptor{
		"testing.metric.1": {
			Create: createArtificialMetricEmitter,
		},
	}
	enabledMetrics := &Enabled{
		Metrics: map[string]*struct{}{
			"testing.metric.2": {},
		},
	}
	mes := TraverseThroughMetricDescriptors(descriptors, enabledMetrics)

	assert.Zero(t, len(mes), "on no match no emitters are created")
}

func Test_TraverseThroughMetricDescriptors_onMatchOnlyMetricEmittersAreCreated(t *testing.T) {
	expectedMetricName := "testing.metric.1"
	descriptors := map[string]Descriptor{
		expectedMetricName: {
			Create: createArtificialMetricEmitter,
		},
	}
	enabledMetrics := &Enabled{
		Metrics: map[string]*struct{}{
			expectedMetricName: {},
		},
	}
	mes := TraverseThroughMetricDescriptors(descriptors, enabledMetrics)

	assert.Equal(t, 1, len(mes), "on match metric emitter is created")
	_, found := mes[expectedMetricName]
	assert.True(t, found, "emitter for matching metric must be created")
}

func createArtificialMetricEmitter() Emitter {
	return CreateMetricEmitterMockV2("testing.metric.1", 0, 0)
}
