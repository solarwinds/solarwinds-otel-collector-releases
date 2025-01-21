package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func Test_Example_HowToFillScopeMetrics(t *testing.T) {
	// product of metric emitter
	ms := pmetric.NewMetricSlice()
	m := ms.AppendEmpty()
	m.SetName("kokoha.metric")
	m.SetDescription("This is mighty kokoha metric")
	s := m.SetEmptySum()
	s.DataPoints().EnsureCapacity(2)
	s.DataPoints().AppendEmpty().SetIntValue(1701)
	s.DataPoints().AppendEmpty().SetIntValue(1702)

	// scope metric emitter product
	rm := pmetric.NewResourceMetrics()
	sm := rm.ScopeMetrics().AppendEmpty()
	ms.MoveAndAppendTo(sm.Metrics())

	assert.Equal(t, 1, rm.ScopeMetrics().At(0).Metrics().Len(), "There must be exactly one metric")
	assert.Equal(t, "kokoha.metric", rm.ScopeMetrics().At(0).Metrics().At(0).Name(), "Metric name must be the same")
	assert.Equal(t, 2, rm.ScopeMetrics().At(0).Metrics().At(0).Sum().DataPoints().Len(), "Number of data points must fit")
	assert.Equal(t, int64(1701), rm.ScopeMetrics().At(0).Metrics().At(0).Sum().DataPoints().At(0).IntValue(), "Value must be the same")
}
