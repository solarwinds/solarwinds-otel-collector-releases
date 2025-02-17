// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper/scrapertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := scrapertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			allMetricsCount++
			mb.RecordSwoHostinfoFirewallDataPoint(ts, 1, "firewall.profile.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSwoHostinfoUptimeDataPoint(ts, 1, "hostdetails.domain-val", "hostdetails.domain.fqdn-val", 23, "hostdetails.model.serialnumber-val", "hostdetails.model.manufacturer-val", "hostdetails.model.name-val", 25, "hostdetails.timezone.caption-val", "hostdetails.timezone.standardname-val", "osdetails.hostname-val", 18, "osdetails.os-val", "osdetails.platform-val", "osdetails.platform.family-val", "osdetails.platform.version-val", "osdetails.kernel.version-val", "osdetails.kernel.architecture-val", "osdetails.virtualization.system-val", "osdetails.virtualization.role-val", "osdetails.host.id-val", 23, "osdetails.language.name-val", "osdetails.language.displayname-val")

			allMetricsCount++
			mb.RecordSwoHostinfoUserLastLoggedDataPoint(ts, 1, "user.name-val", "user.displayname-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "swo.hostinfo.firewall":
					assert.False(t, validatedMetrics["swo.hostinfo.firewall"], "Found a duplicate in the metrics slice: swo.hostinfo.firewall")
					validatedMetrics["swo.hostinfo.firewall"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Metric provides firewall profiles statuses. This metric is supported only on Windows.", ms.At(i).Description())
					assert.Equal(t, "status", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("firewall.profile.name")
					assert.True(t, ok)
					assert.EqualValues(t, "firewall.profile.name-val", attrVal.Str())
				case "swo.hostinfo.uptime":
					assert.False(t, validatedMetrics["swo.hostinfo.uptime"], "Found a duplicate in the metrics slice: swo.hostinfo.uptime")
					validatedMetrics["swo.hostinfo.uptime"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Host uptime in seconds.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("hostdetails.domain")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.domain-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.domain.fqdn")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.domain.fqdn-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.domain.role")
					assert.True(t, ok)
					assert.EqualValues(t, 23, attrVal.Int())
					attrVal, ok = dp.Attributes().Get("hostdetails.model.serialnumber")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.model.serialnumber-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.model.manufacturer")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.model.manufacturer-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.model.name")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.model.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.timezone.bias")
					assert.True(t, ok)
					assert.EqualValues(t, 25, attrVal.Int())
					attrVal, ok = dp.Attributes().Get("hostdetails.timezone.caption")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.timezone.caption-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("hostdetails.timezone.standardname")
					assert.True(t, ok)
					assert.EqualValues(t, "hostdetails.timezone.standardname-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.hostname")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.hostname-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.boottime")
					assert.True(t, ok)
					assert.EqualValues(t, 18, attrVal.Int())
					attrVal, ok = dp.Attributes().Get("osdetails.os")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.os-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.platform")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.platform-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.platform.family")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.platform.family-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.platform.version")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.platform.version-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.kernel.version")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.kernel.version-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.kernel.architecture")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.kernel.architecture-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.virtualization.system")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.virtualization.system-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.virtualization.role")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.virtualization.role-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.host.id")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.host.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.language.lcid")
					assert.True(t, ok)
					assert.EqualValues(t, 23, attrVal.Int())
					attrVal, ok = dp.Attributes().Get("osdetails.language.name")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.language.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("osdetails.language.displayname")
					assert.True(t, ok)
					assert.EqualValues(t, "osdetails.language.displayname-val", attrVal.Str())
				case "swo.hostinfo.user.lastLogged":
					assert.False(t, validatedMetrics["swo.hostinfo.user.lastLogged"], "Found a duplicate in the metrics slice: swo.hostinfo.user.lastLogged")
					validatedMetrics["swo.hostinfo.user.lastLogged"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Host last logged-in user. Supported for Windows and Linux.", ms.At(i).Description())
					assert.Equal(t, "user", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("user.name")
					assert.True(t, ok)
					assert.EqualValues(t, "user.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("user.displayname")
					assert.True(t, ok)
					assert.EqualValues(t, "user.displayname-val", attrVal.Str())
				}
			}
		})
	}
}
