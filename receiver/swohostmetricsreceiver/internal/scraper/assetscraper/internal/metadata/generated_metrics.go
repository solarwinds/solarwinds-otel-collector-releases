// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

type metricSwoAssetInstalledsoftware struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills swo.asset.installedsoftware metric with initial data.
func (m *metricSwoAssetInstalledsoftware) init() {
	m.data.SetName("swo.asset.installedsoftware")
	m.data.SetDescription("Currently installed software.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSwoAssetInstalledsoftware) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, installedsoftwareDateAttributeValue string, installedsoftwareNameAttributeValue string, installedsoftwarePublisherAttributeValue string, installedsoftwareVersionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("installedsoftware.date", installedsoftwareDateAttributeValue)
	dp.Attributes().PutStr("installedsoftware.name", installedsoftwareNameAttributeValue)
	dp.Attributes().PutStr("installedsoftware.publisher", installedsoftwarePublisherAttributeValue)
	dp.Attributes().PutStr("installedsoftware.version", installedsoftwareVersionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSwoAssetInstalledsoftware) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSwoAssetInstalledsoftware) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSwoAssetInstalledsoftware(cfg MetricConfig) metricSwoAssetInstalledsoftware {
	m := metricSwoAssetInstalledsoftware{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSwoAssetInstalledupdates struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills swo.asset.installedupdates metric with initial data.
func (m *metricSwoAssetInstalledupdates) init() {
	m.data.SetName("swo.asset.installedupdates")
	m.data.SetDescription("Currently installed OS updates.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSwoAssetInstalledupdates) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, installedupdateCaptionAttributeValue string, installedupdateDescriptionAttributeValue string, installedupdateHotfixidAttributeValue string, installedupdateInstalledbyAttributeValue string, installedupdateInstalledonAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("installedupdate.caption", installedupdateCaptionAttributeValue)
	dp.Attributes().PutStr("installedupdate.description", installedupdateDescriptionAttributeValue)
	dp.Attributes().PutStr("installedupdate.hotfixid", installedupdateHotfixidAttributeValue)
	dp.Attributes().PutStr("installedupdate.installedby", installedupdateInstalledbyAttributeValue)
	dp.Attributes().PutStr("installedupdate.installedon", installedupdateInstalledonAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSwoAssetInstalledupdates) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSwoAssetInstalledupdates) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSwoAssetInstalledupdates(cfg MetricConfig) metricSwoAssetInstalledupdates {
	m := metricSwoAssetInstalledupdates{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                          MetricsBuilderConfig // config of the metrics builder.
	startTime                       pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                 int                  // maximum observed number of metrics per resource.
	metricsBuffer                   pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                       component.BuildInfo  // contains version information.
	metricSwoAssetInstalledsoftware metricSwoAssetInstalledsoftware
	metricSwoAssetInstalledupdates  metricSwoAssetInstalledupdates
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                          mbc,
		startTime:                       pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                   pmetric.NewMetrics(),
		buildInfo:                       settings.BuildInfo,
		metricSwoAssetInstalledsoftware: newMetricSwoAssetInstalledsoftware(mbc.Metrics.SwoAssetInstalledsoftware),
		metricSwoAssetInstalledupdates:  newMetricSwoAssetInstalledupdates(mbc.Metrics.SwoAssetInstalledupdates),
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSwoAssetInstalledsoftware.emit(ils.Metrics())
	mb.metricSwoAssetInstalledupdates.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSwoAssetInstalledsoftwareDataPoint adds a data point to swo.asset.installedsoftware metric.
func (mb *MetricsBuilder) RecordSwoAssetInstalledsoftwareDataPoint(ts pcommon.Timestamp, val int64, installedsoftwareDateAttributeValue string, installedsoftwareNameAttributeValue string, installedsoftwarePublisherAttributeValue string, installedsoftwareVersionAttributeValue string) {
	mb.metricSwoAssetInstalledsoftware.recordDataPoint(mb.startTime, ts, val, installedsoftwareDateAttributeValue, installedsoftwareNameAttributeValue, installedsoftwarePublisherAttributeValue, installedsoftwareVersionAttributeValue)
}

// RecordSwoAssetInstalledupdatesDataPoint adds a data point to swo.asset.installedupdates metric.
func (mb *MetricsBuilder) RecordSwoAssetInstalledupdatesDataPoint(ts pcommon.Timestamp, val int64, installedupdateCaptionAttributeValue string, installedupdateDescriptionAttributeValue string, installedupdateHotfixidAttributeValue string, installedupdateInstalledbyAttributeValue string, installedupdateInstalledonAttributeValue string) {
	mb.metricSwoAssetInstalledupdates.recordDataPoint(mb.startTime, ts, val, installedupdateCaptionAttributeValue, installedupdateDescriptionAttributeValue, installedupdateHotfixidAttributeValue, installedupdateInstalledbyAttributeValue, installedupdateInstalledonAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
