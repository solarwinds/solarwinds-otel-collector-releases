package internal

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type Exporter struct {
	logger         *zap.Logger
	exporter       exporter.Metrics
	modifyResource resourceModifier
}

type resourceModifier func(resource pcommon.Resource) error

func NewExporter(ctx context.Context, set extension.Settings, cfg *Config, modifyResource resourceModifier) (*Exporter, error) {
	set.Logger.Debug("Creating Exporter")
	oCfg, err := cfg.OTLPConfig()
	if err != nil {
		return nil, err
	}
	expSet := toExporterSettings(set)

	exp := &Exporter{
		logger:         set.Logger,
		modifyResource: modifyResource,
	}
	exp.exporter, err = otlpexporter.NewFactory().CreateMetrics(ctx, expSet, oCfg)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (e *Exporter) Start(ctx context.Context, host component.Host) error {
	e.logger.Debug("Starting exporter")
	return e.exporter.Start(ctx, host)
}

func (e *Exporter) Shutdown(ctx context.Context) error {
	e.logger.Debug("Shutting down exporter")
	return e.exporter.Shutdown(ctx)
}

func toExporterSettings(set extension.Settings) exporter.Settings {
	return exporter.Settings{
		ID:                set.ID,
		TelemetrySettings: set.TelemetrySettings,
		BuildInfo:         set.BuildInfo,
	}
}

func (e *Exporter) PushMetrics(ctx context.Context, md pmetric.Metrics) error {
	if md.MetricCount() == 0 {
		// For receivers with no direct output, but scrape pipeline (ie. telegraf)
		return nil
	}

	rms := md.ResourceMetrics()

	for i := 0; i < rms.Len(); i++ {
		rm := rms.At(i)
		if err := e.modifyResource(rm.Resource()); err != nil {
			return err
		}
	}

	return e.exporter.ConsumeMetrics(ctx, md)
}
