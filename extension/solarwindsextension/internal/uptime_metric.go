package internal

import (
	"context"
	"github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"time"
)

func NewUptimeMetric(logger *zap.Logger) *UptimeMetric {
	logger.Debug("Creating UptimeMetric")
	return &UptimeMetric{logger: logger, uptime: newUptimeCounter()}
}

type UptimeMetric struct {
	logger *zap.Logger
	uptime *uptimeCounter
}

func (um *UptimeMetric) AddUptimeMetric(_ context.Context, md pmetric.Metrics) error {
	um.logger.Debug("Adding uptime metric")
	res := md.ResourceMetrics().AppendEmpty()
	scopeMetrics := res.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName(metadata.ScopeName)
	scopeMetrics.Scope().SetVersion("0.0.1")
	m := scopeMetrics.Metrics().AppendEmpty()
	m.SetName("otelcol.uptime")
	dataPoint := m.SetEmptyGauge().DataPoints().AppendEmpty()
	dataPoint.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dataPoint.SetDoubleValue(um.uptime.Get())
	return nil
}
